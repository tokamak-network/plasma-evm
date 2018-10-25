package pls

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
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

	keys  []*ecdsa.PrivateKey
	addrs []common.Address

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
	testPlsConfig = &DefaultConfig
	testEthClient *ethclient.Client

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
	maxTxFee        *big.Int

	err error
)

func init() {
	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))

	testTxPoolConfig.Journal = ""
	testPlsConfig.TxPool = *testTxPoolConfig
	testPlsConfig.Operator = accounts.Account{Address: params.Operator}

	testPlsConfig.RootChainURL = rootchainUrl

	testEthClient, err = ethclient.Dial(testPlsConfig.RootChainURL)
	if err != nil {
		log.Error("Failed to connect rootchian provider", err)
	}

	keys = []*ecdsa.PrivateKey{key1, key2, key3, key4}
	addrs = []common.Address{addr1, addr2, addr3, addr4}

	maxTxFee = new(big.Int).Mul(defaultGasPrice, big.NewInt(int64(defaultGasLimit)))
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

	startETHDeposit(t, rcm.rootchainContract, key1, ether(1))
	startETHDeposit(t, rcm.rootchainContract, key2, ether(1))
	startETHDeposit(t, rcm.rootchainContract, key3, ether(1))
	startETHDeposit(t, rcm.rootchainContract, key4, ether(1))

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

	startETHDeposit(t, rcm.rootchainContract, key1, ether(1))
	startETHDeposit(t, rcm.rootchainContract, key2, ether(1))
	startETHDeposit(t, rcm.rootchainContract, key3, ether(1))
	startETHDeposit(t, rcm.rootchainContract, key4, ether(1))

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

	startETHWithdraw(t, rcm.rootchainContract, key1, ether(1), rcm.contractParams.costERO)
	startETHWithdraw(t, rcm.rootchainContract, key2, ether(1), rcm.contractParams.costERO)
	startETHWithdraw(t, rcm.rootchainContract, key3, ether(1), rcm.contractParams.costERO)
	startETHWithdraw(t, rcm.rootchainContract, key4, ether(1), rcm.contractParams.costERO)

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

	// NRBEpoch#1 / Block#1
	tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t, testEthClient, plsClient)

	wait(4)
	checkBlockNumber(t, pls, 1)

	opt := makeTxOpt(operatorKey, 0, nil, nil)

	_, err = tokenInRootChain.Mint(opt, addr1, ether(100))
	if err != nil {
		t.Fatalf("Failed to mint token: %v", err)
	}
	wait(2)

	ts1, err := tokenInRootChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from root chain: %v", err)
	}

	ts2, err := tokenInChildChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from child chain: %v", err)
	}

	log.Info("Token total supply", "rootchain", ts1, "childchain", ts2)

	wait(3)

	_, err = pls.rootchainManager.rootchainContract.MapRequestableContractByOperator(opt, tokenAddrInRootChain, tokenAddrInChildChain)
	if err != nil {
		t.Fatalf("Failed to map token addresses to RootChain contract: %v", err)
	}
	wait(2)

	tokenAddr, err := pls.rootchainManager.rootchainContract.RequestableContracts(baseCallOpt, tokenAddrInRootChain)
	wait(2)
	if err != nil {
		t.Fatalf("Failed to fetch token address from RootChain contract: %v", err)
	} else if tokenAddr != tokenAddrInChildChain {
		t.Fatalf("RootChain doesn't know requestable contract address in child chain: %v != %v", tokenAddrInChildChain, tokenAddr)
	}

	// deposit 1 ether for each account
	ETHBalances1 := getETHBalances(addrs, testEthClient)
	PETHBalances1 := getPETHBalances(addrs, plsClient)

	depositAmount := ether(10)
	depositAmountNeg := new(big.Int).Neg(depositAmount)

	for _, key := range keys {
		startETHDeposit(t, pls.rootchainManager.rootchainContract, key, depositAmount)
	}

	// NRBEpoch#1 / Block#2
	makeSampleTx(pls.rootchainManager)

	// ORBEpoch#2 / Block#3
	wait(4)

	// NRBEpoch#3 / Block#4
	makeSampleTx(pls.rootchainManager)

	wait(2)

	checkBlockNumber(t, pls, 4)

	ETHBalances2 := getETHBalances(addrs, testEthClient)
	PETHBalances2 := getPETHBalances(addrs, plsClient)

	// check eth balance in root chain
	for i := range keys {
		checkBalance(t, ETHBalances1[i], ETHBalances2[i], depositAmountNeg, maxTxFee, "Failed to check ETH balance(1)")
	}

	// check peth balance in child chain
	for i := range keys {
		checkBalance(t, PETHBalances1[i], PETHBalances2[i], depositAmount, nil, "Failed to check ETH balance(1)")
	}

	// deposit 1 token from addr1
	tokenAmount := ether(1)
	tokenAmountNeg := new(big.Int).Neg(tokenAmount)

	TokenBalances1 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances1 := getTokenBalances(addrs, tokenInChildChain)

	startTokenDeposit(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, testEthClient, key1, tokenAmount)

	wait(2)

	checkBlockNumber(t, pls, 4)

	// NRBEpoch#3 / Block#5
	makeSampleTx(pls.rootchainManager)

	// ORBEpoch#4 / Block#6
	wait(6)

	checkBlockNumber(t, pls, 6)

	TokenBalances2 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances2 := getTokenBalances(addrs, tokenInChildChain)

	// check Token balance
	checkBalance(t, TokenBalances1[0], TokenBalances2[0], tokenAmountNeg, nil, "Failed to check Token Balance (1)")

	// check PToken balance
	checkBalance(t, PTokenBalances1[0], PTokenBalances2[0], tokenAmount, nil, "Failed to check PToken Balance (1)")

	// transfer token in child chain
	tokenAmountToTransfer := new(big.Int).Div(ether(1), big.NewInt(2))
	tokenAmountToTransferNeg := new(big.Int).Neg(tokenAmountToTransfer)

	// NRBEpoch#5 / Block#7
	transferToken(t, tokenInChildChain, key1, addr2, tokenAmountToTransfer)
	wait(3)

	checkBlockNumber(t, pls, 7)

	PTokenBalances3 := getTokenBalances(addrs, tokenInChildChain)
	checkBalance(t, PTokenBalances2[1], PTokenBalances3[1], tokenAmountToTransferNeg, nil, "Failed to check PToken Balance (2-1)")
	checkBalance(t, PTokenBalances2[2], PTokenBalances3[2], tokenAmountToTransfer, nil, "Failed to check PToken Balance (2-2)")

	// NRBEpoch#5 / Block#8
	transferToken(t, tokenInChildChain, key1, addr2, tokenAmountToTransfer)
	wait(3)

	checkBlockNumber(t, pls, 8)

	PTokenBalances4 := getTokenBalances(addrs, tokenInChildChain)
	checkBalance(t, PTokenBalances3[1], PTokenBalances4[1], tokenAmountToTransferNeg, nil, "Failed to check PToken Balance (3-1)")
	checkBalance(t, PTokenBalances3[2], PTokenBalances4[2], tokenAmountToTransfer, nil, "Failed to check PToken Balance (3-2)")

	t.Log("Test finished")
}

func startETHDeposit(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey, value *big.Int) {
	if value.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot deposit 0 ETH")
	}

	opt := makeTxOpt(key, 0, nil, value)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	isTransfer := true

	if _, err := rootchainContract.StartEnter(opt, isTransfer, addr, empty32Bytes, empty32Bytes); err != nil {
		t.Fatalf("Failed to make an ETH deposit request: %v", err)
	}
}

func startTokenDeposit(t *testing.T, rootchainContract *rootchain.RootChain, tokenContract *token.RequestableSimpleToken, tokenAddress common.Address, ethClient *ethclient.Client, key *ecdsa.PrivateKey, amount *big.Int) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot deposit 0 Token")
	}

	opt := makeTxOpt(key, 0, nil, nil)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	isTransfer := false

	trieKey, err := tokenContract.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := amount.Bytes()
	trieValue = common.LeftPadBytes(trieValue, 32)
	trieValue2 := common.BytesToHash(trieValue)

	tx, err := rootchainContract.StartEnter(opt, isTransfer, tokenAddress, trieKey, trieValue2)

	if err != nil {
		t.Fatalf("Failed to make an token deposit request: %v", err)
	}

	wait(2)

	receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("Failed to get receipt: %v", err)
	}
	if receipt == nil {
		t.Fatal("Failed to send transaction: %v", tx.Hash().Hex())
	}

	log.Info("Token Deposit receipt", "deposit.tx", tx.Hash().Hex(), "receipt.status", receipt.Status, "trieKey", common.Bytes2Hex(trieKey[:]), "trieValue", common.Bytes2Hex(trieValue2[:]))
}

func startETHWithdraw(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey, value, cost *big.Int) {
	opt := makeTxOpt(key, 0, nil, cost)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	isTransfer := true

	if _, err := rootchainContract.StartExit(opt, isTransfer, addr, value, empty32Bytes, empty32Bytes); err != nil {
		t.Fatalf("Failed to make an exit request: %v", err)
	}
}

func transferToken(t *testing.T, tokenContract *token.RequestableSimpleToken, key *ecdsa.PrivateKey, to common.Address, amount *big.Int) {
	opt := makeTxOpt(key, 0, nil, nil)
	_, err := tokenContract.Transfer(opt, to, amount)
	if err != nil {
		t.Fatalf("Failed to transfer toekn: %v", err)
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
		testEthClient,
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
		testEthClient,
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
		testEthClient,
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

func checkBlockNumber(t *testing.T, pls *Plasma, targetBlockNumber uint64) {
	if pls.blockchain.CurrentBlock().NumberU64() != targetBlockNumber {
		t.Fatalf("Expected block number: %d, actual block %d", targetBlockNumber, pls.blockchain.CurrentBlock().NumberU64())
	}
}

// checkBalance check after = before + diff if offset is nil.
// Otherwise, check -offset < after - (before + diff) < offset
func checkBalance(t *testing.T, before, after *big.Int, diff, offset *big.Int, caption string) {
	// truncate up to 100 ether
	truncate := func(v *big.Int) (*big.Int, *big.Int) {
		v0 := v

		if v.Sign() < 0 {
			v0.Abs(v0)
		}

		_, v1 := new(big.Int).DivMod(v0, ether(100), new(big.Int))

		v1, v2 := new(big.Int).DivMod(v1, ether(1), new(big.Int))
		v2.Div(v2, big.NewInt(1e9)) // remove 9 digits from wei

		if v.Sign() < 0 {
			v1.Neg(v1)
		}
		return v1, v2
	}

	toString := func(q, r *big.Int) string {

		return fmt.Sprintf("%s.%s", q.String(), r.String())
	}

	b := before
	a := after

	if offset == nil {
		if a.Cmp(new(big.Int).Add(b, diff)) != 0 {
			bt := toString(truncate(b))
			at := toString(truncate(a))
			dt := toString(truncate(diff))

			wait(1)
			t.Fatalf(caption+"\t : Expected %s != %s + %s, but it isn't", bt, at, dt)
		}
	} else {
		target := new(big.Int).Sub(new(big.Int).Sub(a, b), diff)
		s := new(big.Int).Neg(offset)
		e := offset

		// out of range
		if s.Cmp(target) > 0 || target.Cmp(e) > 0 {
			st := toString(truncate(s))
			et := toString(truncate(e))
			targett := toString(truncate(target))

			wait(1)
			t.Fatalf(caption+"\t : Expected %s < %s < %s, but it isn't", st, targett, et)
		}
	}
}

func getETHBalances(addrs []common.Address, ethClient *ethclient.Client) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = ethClient.BalanceAt(context.Background(), addr, nil)
	}

	return balances
}

func getPETHBalances(addrs []common.Address, plsClient *plsclient.Client) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = plsClient.BalanceAt(context.Background(), addr, nil)
	}

	return balances
}

func getTokenBalances(addrs []common.Address, tokenContract *token.RequestableSimpleToken) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = tokenContract.Balances(baseCallOpt, addr)
	}

	return balances
}
