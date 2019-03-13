package tx

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"testing"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/epochhandler"
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

	keysHex = []string{
		"b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291",
		"78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a115",
		"bfaa65473b85b3c33b2f5ddb511f0f4ef8459213ada2920765aaac25b4fe38c5",
		"067394195895a82e685b000e592f771f7899d77e87cc8c79110e53a2f0b0b8fc",
		"ae03e057a5b117295db86079ba4c8505df6074cdc54eec62f2050e677e5d4e66",
	}

	keys  []*ecdsa.PrivateKey
	addrs []common.Address
	accs  []accounts.Account
	opts  []*bind.TransactOpts

	testConfig = DefaultConfig
	backend    *ethclient.Client

	defaultGasLimit uint64 = 7000000
	defaultResubmit        = 500 * time.Millisecond
	maxTxFee        *big.Int

	err error
)

func init() {
	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))

	testConfig.Interval = defaultResubmit

	backend, err = ethclient.Dial(rootchainUrl)
	if err != nil {
		log.Error("Failed to connect rootchian provider", "err", err)
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
	db := ethdb.NewMemDatabase()
	tm := makeTestManager(db)

	tm.Start()

	// addrs[0] send 1 ETH to addrs[1] n1 times
	n1 := 10
	nonce1, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	for i := 0; i < n1; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(int64(1e18+i)), []byte{}, false, fmt.Sprintf("raw tx %d", i))
		if err := tm.Add(accs[0], rawTx); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
		log.Debug(fmt.Sprintf("raw tx %d added", i))
	}

	if numAddrs := ReadNumAddr(tm.db); numAddrs != 1 {
		t.Errorf("Number of account is expected %d, but actual is %d", 1, numAddrs)
	}

	timer := time.NewTimer(40 * time.Second)

	<-timer.C
	nonce2, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	if nonce2-nonce1 != uint64(n1) {
		t.Fatalf("Nonce doesn't match. expected %d + %d == %d", nonce1, n1, nonce2)
	}

	tm.Stop()
}

func TestRestart(t *testing.T) {
	db := ethdb.NewMemDatabase()
	tm := makeTestManager(db)
	tm.Start()

	// addrs[0] send 1 ETH to addrs[1] n1 times
	n1 := 10
	nonce1, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	for i := 0; i < n1; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(int64(1e18+i)), []byte{}, false, fmt.Sprintf("raw tx %d", i))
		if err := tm.Add(accs[0], rawTx); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
		log.Debug(fmt.Sprintf("raw tx %d added", i))
	}

	timer := time.NewTimer(5 * time.Second)

	<-timer.C

	// restart TransactionManager
	tm.Stop()

	<-time.NewTimer(5 * time.Second).C
	tm = makeTestManager(db)
	tm.Start()

	// addrs[0] send 1 ETH to addrs[1] n2 times
	n2 := 10

	for i := 0; i < n2; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(int64(1e18+i+n1)), []byte{}, false, fmt.Sprintf("raw tx %d", n1+i))
		if err := tm.Add(accs[0], rawTx); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
		log.Debug(fmt.Sprintf("raw tx %d added", n1+i))
	}

	timer2 := time.NewTimer(45 * time.Second)

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
	db := ethdb.NewMemDatabase()
	tm := makeTestManager(db)

	tm.Start()

	// addrs[0] send 1 ETH to addrs[1] n1 times
	n1 := 10
	nonce1, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	// addrs[1] send lots of tx to congest network
	quit := make(chan struct{})
	defer func() {
		close(quit)
	}()

	go func() {
		addr := addrs[2]
		opt := opts[2]

		nonce, _ := backend.NonceAt(context.Background(), addr, nil)
		opt.GasPrice = big.NewInt(2 * params.GWei)
		opt.GasLimit = 6500000
		for {
			select {
			case <-quit:
				return
			default:
				opt.Nonce = big.NewInt(int64(nonce))
				_, _, _, err := epochhandler.DeployEpochHandler(opt, backend)
				if err != nil {
					nonce++
				}
				nonce++
			}
		}
	}()

	for i := 0; i < n1; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(int64(1e18+i)), []byte{}, false, fmt.Sprintf("raw tx %d", i))
		if err := tm.Add(accs[0], rawTx); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
		log.Debug(fmt.Sprintf("raw tx %d added", i))
	}

	if numAddrs := ReadNumAddr(tm.db); numAddrs != 1 {
		t.Errorf("Number of account is expected %d, but actual is %d", 1, numAddrs)
	}

	timer := time.NewTimer(40 * time.Second)

	<-timer.C
	nonce2, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	if nonce2-nonce1 != uint64(n1) {
		t.Fatalf("Nonce doesn't match. expected %d + %d == %d", nonce1, n1, nonce2)
	}

	tm.Stop()
}
