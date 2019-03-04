package tx

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/mattn/go-colorable"
	"io/ioutil"
	"math/big"
	"testing"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/log"
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

	testConfig = DefaultConfig
	backend    *ethclient.Client

	defaultGasLimit uint64 = 7000000
	maxTxFee        *big.Int

	err error
)

func init() {
	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))

	testConfig.Resubmit = 500 * time.Millisecond

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
	}
}

func makeTestManager() *TransactionManager {
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

	db := ethdb.NewMemDatabase()
	tm, _ := NewTransactionManager(ks, backend, db, testConfig)

	return tm
}

func TestTransactionManager(t *testing.T) {
	tm := makeTestManager()

	go tm.Start()

	// addrs[0] send 1 ETH to addrs[1] 20 times
	n := 20
	nonce1, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	for i := 0; i < n; i++ {
		rawTx := NewRawTransaction(addrs[0], 21000, &addrs[1], big.NewInt(1e18), []byte{}, false, fmt.Sprintf("raw tx %d", i))
		if err := tm.Add(accs[0], rawTx); err != nil {
			t.Fatalf("Failed to add rawTx: %v", err)
		}
	}

	timer := time.NewTimer(30 * time.Second)

	<-timer.C
	nonce2, _ := backend.NonceAt(context.Background(), addrs[0], nil)

	if nonce2-nonce1 != uint64(n) {
		t.Fatalf("Nonce doesn't match. expected %d + %d == %d", nonce1, n, nonce2)
	}

	if len(tm.addresses) != 1 {
		t.Errorf("Number of account is expected %d, but actual is %d", 1, len(tm.addresses))
	}

	defer tm.Stop()
}
