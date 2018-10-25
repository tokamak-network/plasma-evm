package pls

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"math/big"
	"runtime"
	"testing"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/consensus"
	"github.com/Onther-Tech/plasma-evm/consensus/ethash"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/token"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/bloombits"
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
	"github.com/Onther-Tech/plasma-evm/pls/gasprice"
	"github.com/Onther-Tech/plasma-evm/plsclient"
	"github.com/Onther-Tech/plasma-evm/rpc"
	"github.com/mattn/go-colorable"
)

var (
	loglevel = flag.Int("loglevel", 4, "verbosity of logs")

	rootchainUrl   = "ws://localhost:8546"
	plasmachainUrl = "http://localhost:8547"

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

	operatorNonce uint64 = 0
	addr1Nonce    uint64 = 0
	addr2Nonce    uint64 = 0
	addr3Nonce    uint64 = 0
	addr4Nonce    uint64 = 0

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

	// pls ~ rootchain
	testPlsConfig        = &DefaultConfig
	testEthBackendClient *ethclient.Client

	// pls ~ plasmachain
	testPlsBackendClient *plsclient.Client

	testTxPoolConfig = &core.DefaultTxPoolConfig

	// rootchain contract
	NRBEpochLength = big.NewInt(2)
	development    = false

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
	testPlsConfig.TxPool = *testTxPoolConfig
	testPlsConfig.Operator = accounts.Account{Address: params.Operator}

	testPlsConfig.RootChainURL = rootchainUrl

	testEthBackendClient, err = ethclient.Dial(testPlsConfig.RootChainURL)
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

	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
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
		makeSampleTx(rcm)
		i++
		ev := <-events.Chan()

		blockInfo := ev.Data.(core.NewMinedBlockEvent)

		if rcm.env.IsRequest {
			t.Fatal("Block should not be request block, but it is not. blockNumber:", blockInfo.Block.NumberU64())
		}
	}

	ev := <-events.Chan()
	blockInfo := ev.Data.(core.NewMinedBlockEvent)
	if !rcm.env.IsRequest {
		t.Fatal("Block should be request block", "blockNumber", blockInfo.Block.NumberU64())
	}

	for i = 0; i < NRBEpochLength.Uint64()*2; {
		makeSampleTx(rcm)
		i++
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)
		makeSampleTx(rcm)

		if rcm.env.IsRequest {
			t.Fatal("Block should not be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
	}

	log.Info("test finished")
	return
}

// TestScenario2 tests enter and exit between root chain & plasma chain
func TestScenario2(t *testing.T) {
	rcm, stopFn, err := makeManager()
	defer stopFn()

	ctx := context.Background()

	if err != nil {
		t.Fatalf("Failed to make rootchian manager: %v", err)
	}

	NRBEpochLength, err := rcm.NRBEpochLength()

	if err != nil {
		t.Fatalf("Failed to get NRBEpochLength: %v", err)
	}

	// balance check in root chain before enter
	bal1be, _ := rcm.backend.BalanceAt(ctx, addr1, nil)
	bal2be, _ := rcm.backend.BalanceAt(ctx, addr2, nil)
	bal3be, _ := rcm.backend.BalanceAt(ctx, addr3, nil)
	bal4be, _ := rcm.backend.BalanceAt(ctx, addr4, nil)

	log.Info("balance of addr1 in root chain before enter", "balance", bal1be)
	log.Info("balance of addr2 in root chain before enter", "balance", bal2be)
	log.Info("balance of addr3 in root chain before enter", "balance", bal3be)
	log.Info("balance of addr4 in root chain before enter", "balance", bal4be)

	// balance check in plasma chain before enter
	db, _ := rcm.blockchain.State()
	log.Info("balance of addr1 in plasma chain before enter", "balance", db.GetBalance(addr1))
	log.Info("balance of addr2 in plasma chain before enter", "balance", db.GetBalance(addr2))
	log.Info("balance of addr3 in plasma chain before enter", "balance", db.GetBalance(addr3))
	log.Info("balance of addr4 in plasma chain before enter", "balance", db.GetBalance(addr4))

	startDepositEnter(t, rcm.rootchainContract, key1, ether(1))
	startDepositEnter(t, rcm.rootchainContract, key2, ether(1))
	startDepositEnter(t, rcm.rootchainContract, key3, ether(1))
	startDepositEnter(t, rcm.rootchainContract, key4, ether(1))

	wait(3)

	numEROs, _ := rcm.rootchainContract.GetNumEROs(baseCallOpt)

	if numEROs.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("numEROs should not be 0")
	}

	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
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

	// #1 NRB epoch check
	for i = 0; i < NRBEpochLength.Uint64(); {
		makeSampleTx(rcm)
		i++
		ev := <-events.Chan()

		blockInfo := ev.Data.(core.NewMinedBlockEvent)

		if rcm.env.IsRequest {
			t.Fatal("Block should not be request block, but it is not. blockNumber:", blockInfo.Block.NumberU64())
		}
	}

	// #2 ORB epoch check
	for i = 0; i < 1; {
		i++
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)
		if !rcm.env.IsRequest {
			t.Fatal("Block should be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
		// balance check in plasma chain after enter
		db, _ := rcm.blockchain.State()
		log.Info("balance of addr1 in plasma chain after enter", "balance", db.GetBalance(addr1))
		log.Info("balance of addr2 in plasma chain after enter", "balance", db.GetBalance(addr2))
		log.Info("balance of addr3 in plasma chain after enter", "balance", db.GetBalance(addr3))
		log.Info("balance of addr4 in plasma chain after enter", "balance", db.GetBalance(addr4))
	}

	// balance check in root chain after enter
	bal1ae, _ := rcm.backend.BalanceAt(ctx, addr1, nil)
	bal2ae, _ := rcm.backend.BalanceAt(ctx, addr2, nil)
	bal3ae, _ := rcm.backend.BalanceAt(ctx, addr3, nil)
	bal4ae, _ := rcm.backend.BalanceAt(ctx, addr4, nil)

	log.Info("balance of addr1 in root chain after enter", "balance", bal1ae)
	log.Info("balance of addr2 in root chain after enter", "balance", bal2ae)
	log.Info("balance of addr3 in root chain after enter", "balance", bal3ae)
	log.Info("balance of addr4 in root chain after enter", "balance", bal4ae)

	// #3 NRB epoch check
	for i = 0; i < NRBEpochLength.Uint64(); {
		makeSampleTx(rcm)
		i++
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)
		makeSampleTx(rcm)

		if rcm.env.IsRequest {
			t.Fatal("Block should not be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
	}

	startExit(t, rcm.rootchainContract, key1, ether(1), rcm.contractParams.costERO)
	startExit(t, rcm.rootchainContract, key2, ether(1), rcm.contractParams.costERO)
	startExit(t, rcm.rootchainContract, key3, ether(1), rcm.contractParams.costERO)
	startExit(t, rcm.rootchainContract, key4, ether(1), rcm.contractParams.costERO)

	wait(3)

	// #4 ORB epoch check
	for i = 0; i < 1; {
		i++
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)
		if !rcm.env.IsRequest {
			t.Fatal("Block should be request block", "blockNumber", blockInfo.Block.NumberU64())
		}

		db, _ := rcm.blockchain.State()
		log.Info("balance of addr1 in plasma chain after exit", "balance", db.GetBalance(addr1))
		log.Info("balance of addr2 in plasma chain after exit", "balance", db.GetBalance(addr2))
		log.Info("balance of addr3 in plasma chain after exit", "balance", db.GetBalance(addr3))
		log.Info("balance of addr4 in plasma chain after exit", "balance", db.GetBalance(addr4))
	}

	// #5+ NRB epoch progress
	for i = 0; i < NRBEpochLength.Uint64()*6; {
		makeSampleTx(rcm)
		i++
		ev := <-events.Chan()
		blockInfo := ev.Data.(core.NewMinedBlockEvent)
		makeSampleTx(rcm)

		if rcm.env.IsRequest {
			t.Fatal("Block should not be request block", "blockNumber", blockInfo.Block.NumberU64())
		}
	}

	applyRequests(t, rcm.rootchainContract, operatorKey)

	// balance check in root chain after enter
	bal1aex, _ := rcm.backend.BalanceAt(ctx, addr1, nil)
	bal2aex, _ := rcm.backend.BalanceAt(ctx, addr2, nil)
	bal3aex, _ := rcm.backend.BalanceAt(ctx, addr3, nil)
	bal4aex, _ := rcm.backend.BalanceAt(ctx, addr4, nil)

	log.Info("balance of addr1 in root chain after exit", "balance", bal1aex)
	log.Info("balance of addr2 in root chain after exit", "balance", bal2aex)
	log.Info("balance of addr3 in root chain after exit", "balance", bal3aex)
	log.Info("balance of addr4 in root chain after exit", "balance", bal4aex)

	log.Info("test finished")
	return
}

// TestScenario3 tests RootChainManager with Plasma service providing RPC endpoint to
// plasma network
func TestScenario3(t *testing.T) {
	pls, rpcServer, err := makePls()
	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	// pls.Start()
	pls.protocolManager.Start(1)

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}

	pls.StartMining(runtime.NumCPU())

	rpcClient := rpc.DialInProc(rpcServer)
	plsClient := plsclient.NewClient(rpcClient)

	wait(3)

	log.Info("All backend is set up")

	tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t, testEthBackendClient, plsClient)

	wait(3)

	ts1, err := tokenInRootChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from root chain", "err", err)
	}

	ts2, err := tokenInChildChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from child chain", "err", err)
	}

	log.Info("Token total supply", "rootchain", ts1, "childchain", ts2)

	wait(3)

	opt := makeTxOpt(operatorKey, 0, nil, nil)

	pls.rootchainManager.rootchainContract.MapRequestableContractByOperator(opt, tokenAddrInRootChain, tokenAddrInChildChain)

	log.Info("Test finished")
}

func startDepositEnter(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey, value *big.Int) {
	opt := makeTxOpt(key, 0, nil, ether(10))
	addr := crypto.PubkeyToAddress(key.PublicKey)
	isTransfer := true

	if _, err := rootchainContract.StartEnter(opt, isTransfer, addr, empty32Bytes, empty32Bytes); err != nil {
		t.Fatalf("Failed to make an enter (deposit) request: %v", err)
	}
}

func startExit(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey, value, cost *big.Int) {
	opt := makeTxOpt(key, 0, nil, cost)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	isTransfer := true

	if _, err := rootchainContract.StartExit(opt, isTransfer, addr, value, empty32Bytes, empty32Bytes); err != nil {
		t.Fatalf("Failed to make an exit request: %v", err)
	}
}

func applyRequests(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey) {
	opt := makeTxOpt(key, 0, nil, nil)

	if _, err := rootchainContract.ApplyRequest(opt); err != nil {
		t.Fatalf("Failed to apply request: %v", err)
	}
}

func deployRootChain(genesis *types.Block) (address common.Address, rootchainContract *rootchain.RootChain, err error) {
	opt := bind.NewKeyedTransactor(operatorKey)

	address, _, rootchainContract, err = rootchain.DeployRootChain(
		opt,
		testEthBackendClient,
		development,
		NRBEpochLength,
		genesis.Header().Root,
		genesis.Header().TxHash,
		genesis.Header().IntermediateStateHash,
	)

	testPlsConfig.RootChainContract = address

	wait(2)

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
	blocks, _ := core.GenerateChain(params.PlasmaChainConfig, parent, engine, db, n, func(i int, b *core.BlockGen) {
		b.SetCoinbase(common.Address{0: byte(seed), 19: byte(i)})
	})
	return blocks
}

func newTxPool(blockchain *core.BlockChain) *core.TxPool {
	pool := core.NewTxPool(*testTxPoolConfig, params.PlasmaChainConfig, blockchain)

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

func makePls() (*Plasma, *rpc.Server, error) {
	var err error

	db, blockchain, err := newCanonical(0, true)

	if err != nil {
		log.Error("Failed to creaet blockchain", "err", err)
		return nil, nil, err
	}

	config := testPlsConfig
	chainConfig := params.PlasmaChainConfig

	rootchainAddress, rootchainContract, err := deployRootChain(blockchain.Genesis())

	if err != nil {
		log.Error("Failed to deploy rootchain contract", "err", err)
		return nil, nil, err
	}

	config.RootChainContract = rootchainAddress

	// configure account manager with empty keystore backend
	backends := []accounts.Backend{}
	accManager := accounts.NewManager(backends...)

	pls := &Plasma{
		config:         config,
		chainDb:        db,
		chainConfig:    chainConfig,
		eventMux:       new(event.TypeMux),
		accountManager: accManager,
		blockchain:     blockchain,
		engine:         engine,
		shutdownChan:   make(chan bool),
		networkID:      config.NetworkId,
		gasPrice:       config.MinerGasPrice,
		etherbase:      config.Etherbase,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(db, params.BloomBitsBlocks, params.BloomConfirms),
	}
	pls.bloomIndexer.Start(pls.blockchain)
	pls.txPool = core.NewTxPool(config.TxPool, pls.chainConfig, pls.blockchain)

	if pls.protocolManager, err = NewProtocolManager(pls.chainConfig, config.SyncMode, config.NetworkId, pls.eventMux, pls.txPool, pls.engine, pls.blockchain, db); err != nil {
		return nil, nil, err
	}

	epochEnv := miner.NewEpochEnvironment()
	pls.miner = miner.New(pls, pls.chainConfig, pls.EventMux(), pls.engine, epochEnv, config.MinerRecommit, config.MinerGasFloor, config.MinerGasCeil)
	pls.miner.SetExtra(makeExtraData(config.MinerExtraData))
	pls.APIBackend = &PlsAPIBackend{pls, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.MinerGasPrice
	}
	pls.APIBackend.gpo = gasprice.NewOracle(pls.APIBackend, gpoParams)
	// Dial rootchain provider
	rootchainBackend, err := ethclient.Dial(config.RootChainURL)
	if err != nil {
		return nil, nil, err
	}
	log.Info("Rootchain provider connected", "url", config.RootChainURL)

	if err != nil {
		return nil, nil, err
	}

	stopFn := func() { pls.Stop() }

	if pls.rootchainManager, err = NewRootChainManager(
		config,
		stopFn,
		pls.txPool,
		pls.blockchain,
		rootchainBackend,
		rootchainContract,
		pls.eventMux,
		pls.accountManager,
		pls.miner,
		epochEnv,
	); err != nil {
		return nil, nil, err
	}

	handler := rpc.NewServer()
	apis := pls.APIs()

	for _, api := range apis {
		if api.Service == nil || api.Namespace != "eth" {
			log.Debug("InProc skipped to register service", "service", api.Service, "namespace", api.Namespace)
			continue
		}

		if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
			return nil, nil, err
		}
		log.Debug("InProc registered", "service", api.Service, "namespace", api.Namespace)
	}

	return pls, handler, nil
}

func deployTokenContracts(t *testing.T, ethClient *ethclient.Client, plsClient *plsclient.Client) (*token.RequestableSimpleToken, *token.RequestableSimpleToken, common.Address, common.Address) {
	opt := makeTxOpt(operatorKey, 0, nil, nil)

	tokenAddrInRootChain, _, tokenInRootChain, err := token.DeployRequestableSimpleToken(
		opt,
		testEthBackendClient,
	)
	if err != nil {
		t.Fatalf("Failed to deploy token contract in root chain", "err", err)
	}
	log.Info("Token deployed in root chain", "address", tokenAddrInRootChain)

	tokenAddrInChildChain, _, tokenInChildChain, err := token.DeployRequestableSimpleToken(
		opt,
		plsClient,
	)
	if err != nil {
		t.Fatalf("Failed to deploy token contract in child chain", "err", err)
	}
	log.Info("Token deployed in child chain", "address", tokenAddrInChildChain)
	operatorNonce += 1

	return tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain
}

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
	epochEnv := miner.NewEpochEnvironment()
	miner := miner.New(minerBackend, params.PlasmaChainConfig, mux, engine, epochEnv, testPlsConfig.MinerRecommit, testPlsConfig.MinerGasFloor, testPlsConfig.MinerGasCeil)

	var rcm *RootChainManager

	stopFn := func() {
		blockchain.Stop()
		txPool.Stop()
		miner.Stop()
		mux.Stop()
		rcm.Stop()
	}
	rcm, err = NewRootChainManager(
		testPlsConfig,
		stopFn,
		txPool,
		blockchain,
		testEthBackendClient,
		rootchainContract,
		mux,
		nil,
		miner,
		epochEnv,
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

// TODO: any user sends tx
func makeSampleTx(rcm *RootChainManager) error {
	pool := rcm.txPool
	nonce := operatorNonce
	operatorNonce += 1

	// self transfer
	var err error

	tx := types.NewTransaction(nonce, operator, nil, 21000, nil, []byte{})

	signer := types.NewEIP155Signer(params.PlasmaChainConfig.ChainID)

	tx, err = types.SignTx(tx, signer, operatorKey)
	if err != nil {
		log.Error("Failed to sign sample tx", "err", err)
		return err
	}

	if err = pool.AddLocal(tx); err != nil {
		log.Error("Failed to insert sample tx to tx pool", "err", err)

		return err
	}

	return nil
}
