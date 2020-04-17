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

	ADDR0 = "0xb79749F25Ef64F9AC277A4705887101D3311A0F4"
	ADDR1 = "0x5E3230019fEd7aB462e3AC277E7709B9b2716b4F"
	ADDR2 = "0x515B385bDc89bCc29077f2B00a88622883bfb498"
	ADDR3 = "0xC927A0CF2d4a1B59775B5D0A35ec76d099e1FaD4"
	ADDR4 = "0x48aFf0622a866d77651eAaA462Ea77b5F39D0ae1"
	ADDR5 = "0xb715125A08140AEA83588a4b569599cde4a0a336"
	ADDR6 = "0x499De281cd965781F1422b7cB73367C15DC416D2"
	ADDR7 = "0xaA60af9BD19dc7438fd19457955C52982D070D27"
	ADDR8 = "0x37da08b6Cd15c3aE905A25Df57B6841A5D80aC93"
	ADDR9 = "0xec4A610a07e81264e8f7F1CAeAe522fEdD7e59c1"

	KEY0 = "2628ca66087c6bc7f9eff7d70db7413d435e170040e8342e67b3db4e55ce752f"
	KEY1 = "86e60281da515184c825c3f46c7ec490b075af1e74607e2e9a66e3df0fa22122"
	KEY2 = "b77b291fab2b0a9e03b5ee0fb0f1140ff41780e93a39e534d54a05ccfad3eead"
	KEY3 = "54a93b74538a7ab51062c7314ea9838519acae6b4ea3d47a7f367e866010364d"
	KEY4 = "434e494f59f6228481256c0c88a375eef2c57be70e612576f302337f48a4634b"
	KEY5 = "c85ab6a568ce788082664c0c17f86e332793895750455090f30f4578e4d20f9a"
	KEY6 = "83d58f7a18e85b728bf5b00ce92d0d8491ae51a962331c8626e51ac32ba8b5f7"
	KEY7 = "85a7751420007fba52e23eca493ac40c770b63c7a16f27ffec39fa01061bc435"
	KEY8 = "5c148c5ba69b7b5c4e53d222e74e6edbbea72f3744fe2ab770320ae70b8d42c0"
	KEY9 = "65d2ecce5d466cb3f9e0ca9acdf53575047ca71527f7c2ed2ab0de620918b2e7"

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
	var wg sync.WaitGroup
	db := rawdb.NewMemoryDatabase()
	tm := makeTestManager(db)

	tm.Start(&wg)

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
	var wg sync.WaitGroup

	db := rawdb.NewMemoryDatabase()
	tm := makeTestManager(db)
	tm.Start(&wg)

	// addrs[0] sends n1 transactions
	n1 := 10
	nonce1, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	for i := 0; i < n1; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(int64(1e18+i)), []byte{}, false, fmt.Sprintf("raw tx %d", i))
		if err := tm.Add(accs[0], rawTx,false); err != nil {
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
	tm.Start(&wg)
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
	var wg sync.WaitGroup
	db := rawdb.NewMemoryDatabase()
	tm := makeTestManager(db)

	tm.Start(&wg)

	// addrs[0] send n1 transactions
	n1 := 3
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
