package pls

import (
	"context"
	"crypto/ecdsa"
	"errors"
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
	ethClient     *ethclient.Client

	// pls ~ plasmachain
	plsClient *plsclient.Client

	testTxPoolConfig = &core.DefaultTxPoolConfig

	// rootchain contract
	NRBEpochLength = big.NewInt(2)
	development    = true

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

	ethClient, err = ethclient.Dial(testPlsConfig.RootChainURL)
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
		t.Fatalf("Failed to start rootchain manager: %v", err)
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
		t.Fatalf("Failed to start rootchain manager: %v", err)
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

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	events := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	wait(3)

	log.Info("All backends are set up")

	// NRBEpoch#1 / Block#1 (1/2)
	tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t)

	wait(4)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

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
	ETHBalances1 := getETHBalances(addrs)
	PETHBalances1 := getPETHBalances(addrs)

	depositAmount := ether(10)
	depositAmountNeg := new(big.Int).Neg(depositAmount)

	for _, key := range keys {
		startETHDeposit(t, pls.rootchainManager.rootchainContract, key, depositAmount)
	}

	// NRBEpoch#1 / Block#2 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#2 / Block#3 (1/1): 4 ETH deposits
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	ETHBalances2 := getETHBalances(addrs)
	PETHBalances2 := getPETHBalances(addrs)

	// check eth balance in root chain
	for i := range keys {
		if err := checkBalance(ETHBalances1[i], ETHBalances2[i], depositAmountNeg, maxTxFee, "Failed to check ETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// check peth balance in child chain
	for i := range keys {
		if err := checkBalance(PETHBalances1[i], PETHBalances2[i], depositAmount, nil, "Failed to check ETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// deposit 1 token from addr1
	tokenAmount := ether(1)
	tokenAmountNeg := new(big.Int).Neg(tokenAmount)

	TokenBalances1 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances1 := getTokenBalances(addrs, tokenInChildChain)

	startTokenDeposit(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key1, tokenAmount)
	wait(2)

	// NRBEpoch#3 / Block#4 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#3 / Block#5 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#4 / Block#6 (1/1): 1 Token deposit

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	TokenBalances2 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances2 := getTokenBalances(addrs, tokenInChildChain)

	// check Token balance
	if err := checkBalance(TokenBalances1[0], TokenBalances2[0], tokenAmountNeg, nil, "Failed to check Token Balance (1)"); err != nil {
		t.Fatal(err)
	}

	// check PToken balance
	if err := checkBalance(PTokenBalances1[0], PTokenBalances2[0], tokenAmount, nil, "Failed to check PToken Balance (1)"); err != nil {
		t.Fatal(err)
	}

	// transfer token from addr1 to addr2, in child chain
	tokenAmountToTransfer := new(big.Int).Div(ether(1), big.NewInt(2))
	tokenAmountToTransferNeg := new(big.Int).Neg(tokenAmountToTransfer)

	// NRBEpoch#5 / Block#7 (1/2)
	transferToken(t, tokenInChildChain, key1, addr2, tokenAmountToTransfer)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	PTokenBalances3 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances2[0], PTokenBalances3[0], tokenAmountToTransferNeg, nil, "Failed to check PToken Balance (2-1)"); err != nil {
		t.Fatal(err)
	}
	if err := checkBalance(PTokenBalances2[1], PTokenBalances3[1], tokenAmountToTransfer, nil, "Failed to check PToken Balance (2-2)"); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#5 / Block#8 (2/2)
	transferToken(t, tokenInChildChain, key1, addr2, tokenAmountToTransfer)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	PTokenBalances4 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances3[0], PTokenBalances4[0], tokenAmountToTransferNeg, nil, "Failed to check PToken Balance (3-1)"); err != nil {
		t.Fatal(err)
	}
	if err := checkBalance(PTokenBalances3[1], PTokenBalances4[1], tokenAmountToTransfer, nil, "Failed to check PToken Balance (3-2)"); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#6 / Block#9 (1/2)
	tokenAmountToWithdraw := new(big.Int).Div(tokenAmount, big.NewInt(4)) // 4 witndrawal requests
	tokenAmountToWithdrawNeg := new(big.Int).Neg(tokenAmountToWithdraw)

	PTokenBalances5 := getTokenBalances(addrs, tokenInChildChain)

	// (1/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, pls.rootchainManager.contractParams.costERO)

	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#6 / Block#10 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#7 / Block#11 (1/1)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	PTokenBalances6 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances5[1], PTokenBalances6[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance (4)"); err != nil {
		t.Fatal(err)
	}

	// (2/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, pls.rootchainManager.contractParams.costERO)

	// NRBEpoch#8 / Block#12 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#8 / Block#13 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#9 / Block#14 (1/1)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	PTokenBalances7 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances6[1], PTokenBalances7[1], tokenAmountToWithdrawNeg, nil, "Failed to check Token Balance (5)"); err != nil {
		t.Fatal(err)
	}

	// (3/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, pls.rootchainManager.contractParams.costERO)

	// NRBEpoch#10 / Block#15 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#10 / Block#16 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#11 / Block#17 (1/1)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	PTokenBalances8 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances7[1], PTokenBalances8[1], tokenAmountToWithdrawNeg, nil, "Failed to check Token Balance (6)"); err != nil {
		t.Fatal(err)
	}

	// (4/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, pls.rootchainManager.contractParams.costERO)

	// NRBEpoch#12 / Block#18 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#12 / Block#19 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#13 / Block#20 (1/1)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	PTokenBalances9 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances8[1], PTokenBalances9[1], tokenAmountToWithdrawNeg, nil, "Failed to check Token Balance (7)"); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#14 / Block#21 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#14 / Block#22 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// apply requests (4 ETH deposits, 1 Token deposits, 4 Token withdrawals)
	for i := 0; i < 4+1; i++ {
		applyRequests(t, pls.rootchainManager.rootchainContract, operatorKey)
	}
	for i := 0; i < 4; i++ {
		tokenBalanceBefore, _ := tokenInRootChain.Balances(baseCallOpt, addr2)
		applyRequests(t, pls.rootchainManager.rootchainContract, operatorKey)
		tokenBalanceAfter, _ := tokenInRootChain.Balances(baseCallOpt, addr2)

		if tokenBalanceAfter.Cmp(new(big.Int).Add(tokenBalanceBefore, tokenAmountToWithdraw)) != 0 {
			t.Fatalf("applyRequest() does not increase token balance")
		}
	}

	t.Log("Test finished")
}

// test challenge invalid exit
func TestScenario4(t *testing.T) {
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

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	events := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	wait(3)

	log.Info("All backends are set up")

	// NRBEpoch#1 / Block#1 (1/2)
	tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t)

	wait(4)
	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

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
	ETHBalances1 := getETHBalances(addrs)
	PETHBalances1 := getPETHBalances(addrs)

	depositAmount := ether(10)
	depositAmountNeg := new(big.Int).Neg(depositAmount)

	for _, key := range keys {
		startETHDeposit(t, pls.rootchainManager.rootchainContract, key, depositAmount)
	}

	// NRBEpoch#1 / Block#2 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#2 / Block#3 (1/1): 4 ETH deposits
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	ETHBalances2 := getETHBalances(addrs)
	PETHBalances2 := getPETHBalances(addrs)

	// check eth balance in root chain
	for i := range keys {
		if err := checkBalance(ETHBalances1[i], ETHBalances2[i], depositAmountNeg, maxTxFee, "Failed to check ETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// check peth balance in child chain
	for i := range keys {
		if err := checkBalance(PETHBalances1[i], PETHBalances2[i], depositAmount, nil, "Failed to check ETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// deposit 1 token from addr1
	tokenAmount := ether(1)
	tokenAmountNeg := new(big.Int).Neg(tokenAmount)

	TokenBalances1 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances1 := getTokenBalances(addrs, tokenInChildChain)

	startTokenDeposit(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key1, tokenAmount)
	wait(2)

	// NRBEpoch#3 / Block#4 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#3 / Block#5 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#4 / Block#6 (1/1): 1 Token deposit

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	TokenBalances2 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances2 := getTokenBalances(addrs, tokenInChildChain)

	// check Token balance
	if err := checkBalance(TokenBalances1[0], TokenBalances2[0], tokenAmountNeg, nil, "Failed to check Token Balance (1)"); err != nil {
		t.Fatal(err)
	}

	// check PToken balance
	if err := checkBalance(PTokenBalances1[0], PTokenBalances2[0], tokenAmount, nil, "Failed to check PToken Balance (1)"); err != nil {
		t.Fatal(err)
	}

	// invalid withdrawal
	startTokenWithdraw(t, pls.rootchainManager.rootchainContract, tokenInRootChain, tokenAddrInRootChain, key2, ether(100), pls.rootchainManager.contractParams.costERO)

	wait(2)

	// NRBEpoch#5 / Block#7 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#5 / Block#8 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// ORBEpoch#6 / Block#9 (1/1): invalid withdrawal

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#7 / Block#10 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#7 / Block#11 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#8 / Block#12 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	// NRBEpoch#8 / Block#12 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkBlockNumber(pls, events); err != nil {
		t.Fatal(err)
	}

	wait(10)

	ERO, err := pls.rootchainManager.rootchainContract.EROs(baseCallOpt, big.NewInt(5))
	if err != nil {
		t.Fatal("failed to get ERO")
	}

	if !ERO.Challenged {
		t.Fatal("ERO is not challenged successfully")
	}

	applyRequests(t, pls.rootchainManager.rootchainContract, operatorKey)
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

func startTokenDeposit(t *testing.T, rootchainContract *rootchain.RootChain, tokenContract *token.RequestableSimpleToken, tokenAddress common.Address, key *ecdsa.PrivateKey, amount *big.Int) {
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
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rootchainContract.StartEnter(opt, isTransfer, tokenAddress, trieKey, trieValue32Bytes)

	if err != nil {
		t.Fatalf("Failed to make an token deposit request: %v", err)
	}

	wait(2)

	receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("Failed to get receipt: %v", err)
	}
	if receipt == nil {
		t.Fatalf("Failed to send transaction: %v", tx.Hash().Hex())
	}
	if receipt.Status == 0 {
		t.Fatalf("Transaction reverted: %v", tx.Hash().Hex())
	}
}

func startETHWithdraw(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey, value, cost *big.Int) {
	opt := makeTxOpt(key, 0, nil, cost)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	if _, err := rootchainContract.StartExit(opt, addr, value, empty32Bytes, empty32Bytes); err != nil {
		t.Fatalf("Failed to make an exit request: %v", err)
	}
}

func startTokenWithdraw(t *testing.T, rootchainContract *rootchain.RootChain, tokenContract *token.RequestableSimpleToken, tokenAddress common.Address, key *ecdsa.PrivateKey, amount, cost *big.Int) {
	opt := makeTxOpt(key, 0, nil, cost)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	trieKey, err := tokenContract.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := amount.Bytes()
	trieValue = common.LeftPadBytes(trieValue, 32)
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rootchainContract.StartExit(opt, tokenAddress, big.NewInt(0), trieKey, trieValue32Bytes)

	if err != nil {
		t.Fatalf("Failed to make an token withdrawal request: %v", err)
	}

	wait(3)

	receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("Failed to get receipt: %v", err)
	}
	if receipt == nil {
		t.Fatalf("Failed to send transaction: %v", tx.Hash().Hex())
	}
	if receipt.Status == 0 {
		t.Fatalf("Transaction reverted: %v", tx.Hash().Hex())
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

	wait(2)

	tx, err := rootchainContract.ApplyRequest(opt)

	if err != nil {
		t.Fatalf("Failed to apply request: %v", err)
	}

	wait(4)

	receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
	log.Error("applyRequest receipt", "receipt", receipt)
	if err != nil {
		t.Fatalf("Failed to get receipt: %v", err)
	}
	if receipt == nil {
		t.Fatalf("Failed to send transaction: %v", tx.Hash().Hex())
	}
	if receipt.Status == 0 {
		t.Fatalf("Transaction reverted: %v", tx.Hash().Hex())
	}
}

func deployRootChain(genesis *types.Block) (address common.Address, rootchainContract *rootchain.RootChain, err error) {
	opt := bind.NewKeyedTransactor(operatorKey)

	address, _, rootchainContract, err = rootchain.DeployRootChain(
		opt,
		ethClient,
		development,
		NRBEpochLength,
		genesis.Header().Root,
		genesis.Header().TxHash,
		genesis.Header().ReceiptHash,
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

func deployTokenContracts(t *testing.T) (*token.RequestableSimpleToken, *token.RequestableSimpleToken, common.Address, common.Address) {
	opt := makeTxOpt(operatorKey, 0, nil, nil)

	tokenAddrInRootChain, _, tokenInRootChain, err := token.DeployRequestableSimpleToken(
		opt,
		ethClient,
	)
	if err != nil {
		t.Fatal("Failed to deploy token contract in root chain", "err", err)
	}
	log.Info("Token deployed in root chain", "address", tokenAddrInRootChain)

	tokenAddrInChildChain, _, tokenInChildChain, err := token.DeployRequestableSimpleToken(
		opt,
		plsClient,
	)
	if err != nil {
		t.Fatal("Failed to deploy token contract in child chain", "err", err)
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
		ethClient,
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

func checkBlockNumber(pls *Plasma, events *event.TypeMuxSubscription) error {
	ev := <-events.Chan()
	blockInfo := ev.Data.(core.NewMinedBlockEvent)

	log.Info("Check Block Number", "current block number", pls.blockchain.CurrentBlock().NumberU64(), "event block number", blockInfo.Block.NumberU64())

	if pls.blockchain.CurrentBlock().NumberU64() != blockInfo.Block.NumberU64() {
		return errors.New(fmt.Sprintf("Expected block number: %d, actual block %d", blockInfo.Block.NumberU64(), pls.blockchain.CurrentBlock().NumberU64()))
	}

	return nil
}

// checkBalance check after = before + diff if offset is nil.
// Otherwise, check -offset < after - (before + diff) < offset
func checkBalance(before, after *big.Int, diff, offset *big.Int, caption string) error {
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
			return errors.New(fmt.Sprintf(caption+"\t : Expected %s != %s + %s, but it isn't", bt, at, dt))
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
			return errors.New(fmt.Sprintf(caption+"\t : Expected %s < %s < %s, but it isn't", st, targett, et))
		}
	}

	return nil
}

func getETHBalances(addrs []common.Address) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = ethClient.BalanceAt(context.Background(), addr, nil)
	}

	return balances
}

func getPETHBalances(addrs []common.Address) []*big.Int {
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
