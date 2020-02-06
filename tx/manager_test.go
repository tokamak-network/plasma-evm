package tx

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/epochhandler"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/mattn/go-colorable"
)

var (
	loglevel = flag.Int("loglevel", 4, "verbosity of logs")

	rootchainUrl = "ws://localhost:8546"

	ADDR0 = "0x8aFD9CCC0EC274961BC22a7d4fC93ba813d6a061"
	ADDR1 = "0x0F4d6130AF2db37F14d4C576c4bE03a6BC0537e1"
	ADDR2 = "0xF7681f3bcd42742520F7441decfe17FF3B4af4a0"
	ADDR3 = "0x3746221Fd27164Aaf0aFe7432046F0176bd1cBa7"
	ADDR4 = "0x1CE66b805281ADa64c11Aac743038e05EEaCDacb"
	ADDR5 = "0x98Bca9B1cdC6F96F7c99fA57e6dD04c3250d9E45"
	ADDR6 = "0x31ff53cBED624d079b935AdfE8D8452A17512Bb1"
	ADDR7 = "0x02f3F30C52CEa117fFF723B82db757559C868F29"
	ADDR8 = "0x4Da60C9fEe1367f7a5488Eb03163887ae1E35549"
	ADDR9 = "0x75f9A88FCfdEcfcDddC0c388465C9B1622e44F40"

	KEY0 = "2106171baa6d881a96e1926352aef3ad4aa8c4f27ac8367f6696d95d88f997f2"
	KEY1 = "6999eb7868035fc0ff2328b545cc5336e99900683e86b80d9752ab3de18fe193"
	KEY2 = "8ec15a59afe2dfad6502fed2284c969d519b394be91a6e9eadd6f68a58b09bee"
	KEY3 = "6a6e1af88a12e68012c2ba584b5aa5d1990e64f1b6851322500dbe3a8fe36134"
	KEY4 = "26c553dffe59894183c397bfb9008482f20cbc4526eed062c3e802ea843497a6"
	KEY5 = "02f45c0281bc76ab38a8e10f283c197d0a32b218c850ead61d7c65f38e028c48"
	KEY6 = "66a9443aa3c5bdee0783a2fc3bda4a6d12a412cbd075f15bfb45a1da4810840e"
	KEY7 = "820eaf6b0269d6040a17443b2c994596caa3c30c381145989de549a341c74d19"
	KEY8 = "440e25a6b8d5cd41b19823f85fa7690c7f091cd62394a4eaf99278d45428269b"
	KEY9 = "54a9d136007d301469cfd4d5e142170302f0bbdb2336e53dedfe9ed767b7bd3d"

	keysHex = []string{
		KEY0,
		KEY1,
		KEY2,
		KEY3,
		KEY4,
		KEY5,
		KEY6,
		KEY7,
		KEY8,
		KEY9,
	}

	keys  []*ecdsa.PrivateKey
	addrs []common.Address
	accs  []accounts.Account
	opts  []*bind.TransactOpts

	testConfig = DefaultConfig
	backend    *ethclient.Client

	defaultGasLimit uint64 = 7000000
	defaultResubmit        = 3 * time.Second
	maxTxFee        *big.Int

	err error
)

func init() {
	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))

	testConfig.Interval = defaultResubmit

	backend, err = ethclient.Dial(rootchainUrl)
	if err != nil {
		log.Error("Failed to connect rootchain provider", "err", err)
	}

	networkId, err := backend.NetworkID(context.Background())
	testConfig.ChainId = new(big.Int).Set(networkId)

	if err != nil {
		log.Error("Failed to read rootchain network id", "err", err)
	}

	log.Info("rootchain connected", "network id", networkId)

	for _, hex := range keysHex {
		key, err := crypto.HexToECDSA(hex)
		if err != nil {
			log.Error("Failed to convert hex string to private key", "err", err)
		}
		addr := crypto.PubkeyToAddress(key.PublicKey)

		keys = append(keys, key)
		addrs = append(addrs, addr)
		accs = append(accs, accounts.Account{Address: addr})
		opts = append(opts, bind.NewKeyedTransactor(key))
	}
}

func makeTestManager(db ethdb.Database) *TransactionManager {
	d, err := ioutil.TempDir("", "pls-transaction-manager-test")
	if err != nil {
		log.Error("Failed to set temporary keystore directory", "err", err)
	}

	ks := keystore.NewKeyStore(d, 2, 1)

	for _, key := range keys {
		acc, err := ks.ImportECDSA(key, "")
		if err != nil {
			log.Error("Failed to import private key", "err", err)
		}

		if err := ks.Unlock(acc, ""); err != nil {
			log.Error("Failed to unlock account", "err", err)
		}
	}

	tm, _ := NewTransactionManager(ks, backend, db, testConfig)

	return tm
}

func TestBasic(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	tm := makeTestManager(db)

	tm.Start()

	// addrs[0] send 1 ETH to addrs[1] n1 times
	n1 := 10
	nonce1, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	for i := 0; i < n1; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(int64(1e18+i)), []byte{}, false, fmt.Sprintf("raw tx %d", i))
		if err := tm.Add(accs[0], rawTx, false); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
		log.Debug(fmt.Sprintf("raw tx %d added", i))
	}

	if numAddrs := ReadNumAddr(tm.db); numAddrs != 1 {
		t.Errorf("Number of account is expected %d, but actual is %d", 1, numAddrs)
	}

	timer := time.NewTimer(10 * time.Minute)

	<-timer.C
	nonce2, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	if nonce2-nonce1 != uint64(n1) {
		t.Fatalf("Nonce doesn't match. expected %d + %d == %d", nonce1, n1, nonce2)
	}

	tm.Stop()
}

func TestRestart(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	tm := makeTestManager(db)
	tm.Start()

	// addrs[0] sends n1 transactions
	n1 := 10
	nonce1, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	for i := 0; i < n1; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(int64(1e18+i)), []byte{}, false, fmt.Sprintf("raw tx %d", i))
		if err := tm.Add(accs[0], rawTx, false); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
		log.Debug(fmt.Sprintf("raw tx %d added", i))
	}

	timer := time.NewTimer(25 * time.Second)

	<-timer.C

	// restart TransactionManager
	tm.Stop()

	<-time.NewTimer(5 * time.Second).C
	tm = makeTestManager(db)
	tm.Start()
	log.Info("TranasctionManager restarted")

	// addrs[0] sends n2 transactions
	n2 := 10

	for i := 0; i < n2; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(int64(1e18+i+n1)), []byte{}, false, fmt.Sprintf("raw tx %d", n1+i))
		if err := tm.Add(accs[0], rawTx, false); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
		log.Debug(fmt.Sprintf("raw tx %d added", n1+i))
	}

	timer2 := time.NewTimer(2 * time.Minute)

	<-timer2.C
	nonce2, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	if nonce2-nonce1 != uint64(n1+n2) {
		t.Fatalf("Nonce doesn't match. expected %d + %d == %d", nonce1, n1+n2, nonce2)
	}

	if ReadNumAddr(tm.db) != 1 {
		t.Errorf("Number of account is expected %d, but actual is %d", 1, len(tm.addresses))
	}

	tm.Stop()
}

func TestCongestedNetwork(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	tm := makeTestManager(db)

	tm.Start()

	// addrs[0] send n1 transactions
	n1 := 10
	nonce1, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	// addrs[1, ..., 8] send lots of tx to congest network
	quit := make(chan struct{})
	defer func() {
		close(quit)
	}()

	go func() {

		senders := map[common.Address]*bind.TransactOpts{
			addrs[1]: opts[1],
			addrs[2]: opts[2],
			addrs[3]: opts[3],
			addrs[4]: opts[4],
			addrs[5]: opts[5],
			addrs[6]: opts[6],
			addrs[7]: opts[7],
			addrs[8]: opts[8],
		}

		for addr, opt := range senders {
			opt.GasPrice = new(big.Int).Sub(testConfig.MaxGasPrice, big.NewInt(params.GWei))
			opt.GasLimit = 3000000

			nonce, err := tm.backend.PendingNonceAt(context.Background(), addr)

			if err != nil {
				t.Fatalf("Faile to fetch nonce: %v", err)
			}
			opt.Nonce = big.NewInt(int64(nonce))
		}

		i := 0
		for {
			interval := 200 * time.Millisecond
			ticker := time.NewTicker(interval)
			select {
			case <-quit:
				return
			case <-ticker.C:
				i++
				if i%10 == 0 {
					log.Info("Congestion Transactions", "num", i)
				}

				wg := sync.WaitGroup{}
				for addr, opt := range senders {
					go func(addr common.Address, opt *bind.TransactOpts) {
						wg.Add(1)
						_, _, _, err := epochhandler.DeployEpochHandler(opt, backend)
						if err != nil {
							log.Warn("Faile to send congest transaction", "err", err, "addr", addr, "nonce", opt.Nonce)
						}

						opt.Nonce = new(big.Int).Add(opt.Nonce, big.NewInt(1))
						wg.Done()
					}(addr, opt)
				}
				wg.Wait()
			}
		}
	}()

	<-time.NewTimer(3 * time.Second).C

	var txs RawTransactions

	data := common.FromHex(epochhandler.EpochHandlerBin)

	for i := 0; i < n1; i++ {
		rawTx := NewRawTransaction(addrs[0], 3000000, &addrs[1], big.NewInt(0), append(data, byte(i)), false, fmt.Sprintf("raw tx %d", i))
		if err := tm.Add(accs[0], rawTx, false); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
		log.Debug(fmt.Sprintf("raw tx %d added", i))

		txs = append(txs, rawTx)
	}

	if numAddrs := ReadNumAddr(tm.db); numAddrs != 1 {
		t.Errorf("Number of account is expected %d, but actual is %d", 1, numAddrs)
	}

	timer := time.NewTimer(300 * time.Second)

	<-timer.C
	nonce2, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	if nonce2-nonce1 != uint64(n1) {
		t.Fatalf("Nonce doesn't match. expected %d + %d == %d", nonce1, n1, nonce2)
	}

	if len(ReadPendingTxs(tm.db, addrs[0])) != 0 {
		t.Fatalf("Rwa transactions are not mined")
	}

	tm.Stop()
}
