package pls

import (
	"crypto/ecdsa"
	"flag"
	"math/big"
	"testing"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/consensus"
	"github.com/Onther-Tech/plasma-evm/consensus/ethash"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/contract"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/core/vm"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/miner"
	"github.com/Onther-Tech/plasma-evm/node"
	"github.com/Onther-Tech/plasma-evm/p2p"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/mattn/go-colorable"
)

var (
	loglevel = flag.Int("loglevel", 4, "verbosity of logs")

	operator       = params.Operator
	operatorKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	opt0           = bind.NewKeyedTransactor(operatorKey)

	addr1 = common.HexToAddress("0x5df7107c960320b90a3d7ed9a83203d1f98a811d")
	addr2 = common.HexToAddress("0x3cd9f729c8d882b851f8c70fb36d22b391a288cd")
	addr3 = common.HexToAddress("0x57ab89f4eabdffce316809d790d5c93a49908510")
	addr4 = common.HexToAddress("0x6c278df36922fea54cf6f65f725267e271f60dd9")

	key1, _ = crypto.HexToECDSA("78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a115")
	key2, _ = crypto.HexToECDSA("bfaa65473b85b3c33b2f5ddb511f0f4ef8459213ada2920765aaac25b4fe38c5")
	key3, _ = crypto.HexToECDSA("067394195895a82e685b000e592f771f7899d77e87cc8c79110e53a2f0b0b8fc")
	key4, _ = crypto.HexToECDSA("ae03e057a5b117295db86079ba4c8505df6074cdc54eec62f2050e677e5d4e66")

	opt1 = bind.NewKeyedTransactor(key1)
	opt2 = bind.NewKeyedTransactor(key2)
	opt3 = bind.NewKeyedTransactor(key3)
	opt4 = bind.NewKeyedTransactor(key4)

	empty32Bytes = common.Hash{}

	// blockchain
	canonicalSeed = 1
	engine        = ethash.NewFaker()

	// node
	testNodeKey, _ = crypto.GenerateKey()
	testNodeConfig = &node.Config{
		Name: "test node",
		P2P:  p2p.Config{PrivateKey: testNodeKey},
	}

	// pls
	testPlsConfig     = DefaultConfig
	testClientBackend *ethclient.Client

	testTxPoolConfig = core.DefaultTxPoolConfig

	// rootchain contract
	NRBEpochLength = big.NewInt(2)

	// transaction
	defaultGasPrice        = big.NewInt(1e9) // 1 Gwei
	defaultValue           = big.NewInt(0)
	defaultGasLimit uint64 = 4000000

	err error
)

func init() {
	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))

	testTxPoolConfig.Journal = ""
	testPlsConfig.TxPool = testTxPoolConfig
	testPlsConfig.Operator = accounts.Account{Address: params.Operator}

	testPlsConfig.RootChainURL = "ws://localhost:8546"

	testClientBackend, err = ethclient.Dial(testPlsConfig.RootChainURL)
	if err != nil {
		log.Error("Failed to connect rootchian provider", err)
	}
}

func TestScenario1(t *testing.T) {
	rcm, stopFn, err := makeManager()
	defer stopFn()

	if err != nil {
		t.Fatalf("Failed to make rootchian manager: %v", err)
	}

	NRBEpochLength, err := rcm.NRBEpochLength()

	if err != nil {
		t.Fatalf("Failed to get NRBEpochLength: %v", err)
	}

	startDepositEnter(t, rcm.rootchainContract, key1, ether(1))
	startDepositEnter(t, rcm.rootchainContract, key2, ether(1))
	startDepositEnter(t, rcm.rootchainContract, key3, ether(1))
	startDepositEnter(t, rcm.rootchainContract, key4, ether(1))

	wait(3)

	numEROs, _ := rcm.rootchainContract.GetNumEROs(baseCallOpt)

	if numEROs.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("numEROs should not be 0")
	}

	events := rcm.eventMux.Subscribe(miner.BlockMined{})
	defer events.Unsubscribe()

	if err = rcm.Start(); err != nil {
		t.Fatal("Failed to start rootchain manager: %v", err)
	}

	timer := time.NewTimer(1 * time.Minute)
	go func() {
		<-timer.C
		t.Fatal("Out of time")
	}()

	var i uint64

	for i = 0; i < NRBEpochLength.Uint64(); {
		i++
		ev:= <-events.Chan()

		blockInfo := ev.Data.(miner.BlockMined)


		if blockInfo.IsRequest {
			t.Fatal("Block should not be request block, but it is not. blockNumber:", blockInfo.BlockNumber.Uint64())
		}
	}

	ev := <-events.Chan()
	blockInfo := ev.Data.(miner.BlockMined)
	if !blockInfo.IsRequest {
		t.Fatal("Block should be request block", "blockNumber", blockInfo.BlockNumber.Uint64())
	}

	for i = 0; i < NRBEpochLength.Uint64(); {
		i++
		ev := <-events.Chan()
		blockInfo := ev.Data.(miner.BlockMined)

		if blockInfo.IsRequest {
			t.Fatal("Block should not be request block", "blockNumber", blockInfo.BlockNumber.Uint64())
		}
	}

	log.Info("test finished")
	return
}

func startDepositEnter(t *testing.T, rootchainContract *contract.RootChain, key *ecdsa.PrivateKey, value *big.Int) {
	opt := makeTxOpt(key, 0, nil, ether(1))
	addr := crypto.PubkeyToAddress(key.PublicKey)
	isTransfer := true

	if _, err := rootchainContract.StartEnter(opt, isTransfer, addr, empty32Bytes, empty32Bytes); err != nil {
		t.Fatalf("Failed to make an enter (deposit) request: %v", err)
	}

}

func deployRootChain(genesis *types.Block) (address common.Address, rootchainContract *contract.RootChain, err error) {
	opt := bind.NewKeyedTransactor(operatorKey)
	development := false

	address, _, rootchainContract, err = contract.DeployRootChain(
		opt,
		testClientBackend,
		development,
		NRBEpochLength,
		genesis.Header().Root,
		genesis.Header().TxHash,
		genesis.Header().IntermediateStateHash,
	)

	testPlsConfig.RootChainContract = address

	return address, rootchainContract, err
}

func newCanonical(n int, full bool) (ethdb.Database, *core.BlockChain, error) {
	// gspec = core.DefaultGenesisBlock()
	gspec := core.DefaultGenesisBlock()
	// Initialize a fresh chain with only a genesis block
	db := ethdb.NewMemDatabase()
	genesis := gspec.MustCommit(db)

	blockchain, _ := core.NewBlockChain(db, nil, gspec.Config, engine, vm.Config{})
	// Create and inject the requested chain
	if n == 0 {
		return db, blockchain, nil
	}
	if full {
		// Full block-chain requested
		blocks := makeBlockChain(genesis, n, engine, db, canonicalSeed)
		_, err := blockchain.InsertChain(blocks)
		return db, blockchain, err
	}
	// Header-only chain requested
	headers := makeHeaderChain(genesis.Header(), n, engine, db, canonicalSeed)
	_, err := blockchain.InsertHeaderChain(headers, 1)
	return db, blockchain, err
}

func makeHeaderChain(parent *types.Header, n int, engine consensus.Engine, db ethdb.Database, seed int) []*types.Header {
	blocks := makeBlockChain(types.NewBlockWithHeader(parent), n, engine, db, seed)
	headers := make([]*types.Header, len(blocks))
	for i, block := range blocks {
		headers[i] = block.Header()
	}
	return headers
}

func makeBlockChain(parent *types.Block, n int, engine consensus.Engine, db ethdb.Database, seed int) []*types.Block {
	blocks, _ := core.GenerateChain(params.TestChainConfig, parent, engine, db, n, func(i int, b *core.BlockGen) {
		b.SetCoinbase(common.Address{0: byte(seed), 19: byte(i)})
	})
	return blocks
}

func newTxPool(blockchain *core.BlockChain) *core.TxPool {
	pool := core.NewTxPool(testTxPoolConfig, params.TestChainConfig, blockchain)

	return pool
}

type testPlsBackend struct {
	acm        *accounts.Manager
	blockchain *core.BlockChain
	txPool     *core.TxPool
	db         ethdb.Database
}

func (b *testPlsBackend) AccountManager() *accounts.Manager { return b.acm }
func (b *testPlsBackend) BlockChain() *core.BlockChain      { return b.blockchain }
func (b *testPlsBackend) TxPool() *core.TxPool              { return b.txPool }
func (b *testPlsBackend) ChainDb() ethdb.Database           { return b.db }

func makeManager() (*RootChainManager, func(), error) {
	db, blockchain, _ := newCanonical(0, true)
	contractAddress, rootchainContract, err := deployRootChain(blockchain.Genesis())
	if err != nil {
		return nil, func() {}, err
	}
	wait(3)
	log.Info("Contract deployed at", "address", contractAddress)

	testPlsConfig.RootChainContract = contractAddress

	txPool := newTxPool(blockchain)
	minerBackend := &testPlsBackend{
		acm:        nil,
		blockchain: blockchain,
		txPool:     txPool,
		db:         db,
	}

	mux := new(event.TypeMux)
	miner := miner.New(minerBackend, params.TestChainConfig, mux, engine, testPlsConfig.MinerRecommit, testPlsConfig.MinerGasFloor, testPlsConfig.MinerGasCeil)

	var rcm *RootChainManager

	stopFn := func() {
		blockchain.Stop()
		txPool.Stop()
		miner.Stop()
		mux.Stop()
		rcm.Stop()
	}

	rcm, err = NewRootChainManager(
		&testPlsConfig,
		stopFn,
		txPool,
		blockchain,
		testClientBackend,
		rootchainContract,
		mux,
		nil,
		miner,
	)

	if err != nil {
		return nil, func() {}, err
	}

	go miner.Start(operator)
	return rcm, stopFn, nil
}

func makeTxOpt(key *ecdsa.PrivateKey, gasLimit uint64, gasPrice, value *big.Int) *bind.TransactOpts {
	opt := bind.NewKeyedTransactor(key)
	opt.GasLimit = defaultGasLimit
	opt.GasPrice = defaultGasPrice
	opt.Value = defaultValue

	if gasLimit != 0 {
		opt.GasLimit = gasLimit
	}

	if gasPrice != nil {
		opt.GasPrice = gasPrice
	}

	if value != nil {
		opt.Value = value
	}

	return opt
}

func ether(v float64) *big.Int {
	f := new(big.Float).Mul(big.NewFloat(v), big.NewFloat(1e18))
	out, _ := f.Int(nil)
	return out
}

func wait(t time.Duration) {
	timer := time.NewTimer(t * time.Second)
	<-timer.C
}
