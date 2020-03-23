package pls

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/consensus"
	"github.com/Onther-Tech/plasma-evm/consensus/ethash"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/epochhandler"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/ethertoken"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/mintabletoken"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/submithandler"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/token"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/bloombits"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/core/vm"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/miner"
	"github.com/Onther-Tech/plasma-evm/miner/epoch"
	"github.com/Onther-Tech/plasma-evm/node"
	"github.com/Onther-Tech/plasma-evm/p2p"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/pls/gasprice"
	"github.com/Onther-Tech/plasma-evm/plsclient"
	"github.com/Onther-Tech/plasma-evm/rpc"
	"github.com/Onther-Tech/plasma-evm/tx"
	"github.com/mattn/go-colorable"
)

var (
	loglevel = flag.Int("loglevel", 4, "verbosity of logs")

	rootchainUrl   = "ws://localhost:8546"
	plasmachainUrl = "http://localhost:8547"

	operatorKey, _   = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	operator         = crypto.PubkeyToAddress(operatorKey.PublicKey)
	challengerKey, _ = crypto.HexToECDSA("78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a114")
	challenger       = crypto.PubkeyToAddress(challengerKey.PublicKey)
	operatorOpt      = bind.NewKeyedTransactor(operatorKey)

	addr1 = common.HexToAddress("0x5df7107c960320b90a3d7ed9a83203d1f98a811d")
	addr2 = common.HexToAddress("0x3cd9f729c8d882b851f8c70fb36d22b391a288cd")
	addr3 = common.HexToAddress("0x57ab89f4eabdffce316809d790d5c93a49908510")
	addr4 = common.HexToAddress("0x6c278df36922fea54cf6f65f725267e271f60dd9")
	addrs = []common.Address{addr1, addr2, addr3, addr4}

	key1, _ = crypto.HexToECDSA("78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a115")
	key2, _ = crypto.HexToECDSA("bfaa65473b85b3c33b2f5ddb511f0f4ef8459213ada2920765aaac25b4fe38c5")
	key3, _ = crypto.HexToECDSA("067394195895a82e685b000e592f771f7899d77e87cc8c79110e53a2f0b0b8fc")
	key4, _ = crypto.HexToECDSA("ae03e057a5b117295db86079ba4c8505df6074cdc54eec62f2050e677e5d4e66")
	keys    = []*ecdsa.PrivateKey{key1, key2, key3, key4}

	locks = map[common.Address]*sync.Mutex{
		operator: &sync.Mutex{},
		addr1:    &sync.Mutex{},
		addr2:    &sync.Mutex{},
		addr3:    &sync.Mutex{},
		addr4:    &sync.Mutex{},
	}

	operatorNonceRootChain uint64 = 0
	addr1NonceRootChain    uint64 = 0
	addr2NonceRootChain    uint64 = 0
	addr3NonceRootChain    uint64 = 0
	addr4NonceRootChain    uint64 = 0
	noncesRootChain               = map[common.Address]*uint64{
		operator: &operatorNonceRootChain,
		addr1:    &addr1NonceRootChain,
		addr2:    &addr2NonceRootChain,
		addr3:    &addr3NonceRootChain,
		addr4:    &addr4NonceRootChain,
	}

	operatorNonceChildChain uint64 = 0
	addr1NonceChildChain    uint64 = 0
	addr2NonceChildChain    uint64 = 0
	addr3NonceChildChain    uint64 = 0
	addr4NonceChildChain    uint64 = 0
	noncesChildChain               = map[common.Address]*uint64{
		operator: &operatorNonceChildChain,
		addr1:    &addr1NonceChildChain,
		addr2:    &addr2NonceChildChain,
		addr3:    &addr3NonceChildChain,
		addr4:    &addr4NonceChildChain,
	}

	opt1 = bind.NewKeyedTransactor(key1)
	opt2 = bind.NewKeyedTransactor(key2)
	opt3 = bind.NewKeyedTransactor(key3)
	opt4 = bind.NewKeyedTransactor(key4)

	opts = map[common.Address]*bind.TransactOpts{
		operator: operatorOpt,
		addr1:    opt1,
		addr2:    opt2,
		addr3:    opt3,
		addr4:    opt4,
	}

	empty32Bytes = common.Hash{}

	// contracts
	mintableToken     *mintabletoken.ERC20Mintable
	mintableTokenAddr common.Address

	etherToken     *ethertoken.EtherToken
	etherTokenAddr common.Address

	etherTokenInChildChain     *ethertoken.EtherToken
	etherTokenAddrInChildChain common.Address

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
	testVmConfg   = vm.Config{EnablePreimageRecording: true}
	testPlsConfig = &DefaultConfig
	ethClient     *ethclient.Client

	// pls ~ plasmachain
	plsClient *plsclient.Client

	testTxPoolConfig = &core.DefaultTxPoolConfig

	// rootchain contract
	NRELength               = big.NewInt(2)
	development             = false
	swapEnabledInRootChain  = false
	swapEnabledInChildChain = true

	// transaction
	defaultGasPrice        = big.NewInt(1) // 1 Gwei
	defaultValue           = big.NewInt(0)
	defaultGasLimit uint64 = 7000000
	maxTxFee        *big.Int

	err error
)

func init() {
	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))

	testTxPoolConfig.Journal = ""
	testPlsConfig.TxPool = *testTxPoolConfig
	testPlsConfig.Operator = accounts.Account{Address: operator}
	testPlsConfig.Challenger = accounts.Account{Address: challenger}
	testPlsConfig.NodeMode = ModeOperator

	testPlsConfig.RootChainURL = rootchainUrl

	testPlsConfig.TxConfig.Interval = 2 * time.Second
	testPlsConfig.Miner.Recommit = 10 * time.Second

	ethClient, err = ethclient.Dial(testPlsConfig.RootChainURL)
	if err != nil {
		log.Error("Failed to connect rootchain provider", "err", err)
	}

	networkId, err := ethClient.NetworkID(context.Background())
	if err != nil {
		log.Error("Failed to get network id", "err", err)
	}
	testPlsConfig.RootChainNetworkID = networkId.Uint64()
	testPlsConfig.TxConfig.ChainId = networkId

	keys = []*ecdsa.PrivateKey{key1, key2, key3, key4}
	addrs = []common.Address{addr1, addr2, addr3, addr4}

	operatorNonceRootChain, err = ethClient.NonceAt(context.Background(), operator, nil)
	addr1NonceRootChain, _ = ethClient.NonceAt(context.Background(), addr1, nil)
	addr2NonceRootChain, _ = ethClient.NonceAt(context.Background(), addr2, nil)
	addr3NonceRootChain, _ = ethClient.NonceAt(context.Background(), addr3, nil)
	addr4NonceRootChain, _ = ethClient.NonceAt(context.Background(), addr4, nil)

	for _, opt := range opts {
		opt.GasLimit = defaultGasLimit
	}

	maxTxFee = new(big.Int).Mul(defaultGasPrice, big.NewInt(int64(defaultGasLimit)))
}

// TestBasic tests enter & exit with token transfer in child chain.
func TestBasic(t *testing.T) {
	pls, rpcServer, dir, err := makePls()
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}
	pls.protocolManager.Start(1)

	rpcClient := rpc.DialInProc(rpcServer)

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
	blockSubmitWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
	defer blockFilterer.Unsubscribe()

	epochPreparedEvents := make(chan *rootchain.RootChainEpochPrepared)
	epochPreparedWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	epochPreparedFilterer, _ := pls.rootchainManager.rootchainContract.WatchEpochPrepared(epochPreparedWatchOpts, epochPreparedEvents)
	defer epochPreparedFilterer.Unsubscribe()

	var tx *types.Transaction

	wait(3)

	log.Info("All backends are set up")

	// NRE#1 / Block#1 (1/2)
	tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 1); err != nil {
		t.Fatal(err)
	}

	setNonce(operatorOpt, &operatorNonceRootChain)
	tx, err = tokenInRootChain.Mint(operatorOpt, addr1, ether(100))
	if err != nil {
		t.Fatalf("Failed to mint token: %v", err)
	}

	waitEthTx(tx.Hash(), "")

	ts1, err := tokenInRootChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from root chain: %v", err)
	}

	ts2, err := tokenInChildChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from child chain: %v", err)
	}

	log.Info("Token total supply", "rootchain", ts1, "childchain", ts2)

	setNonce(operatorOpt, &operatorNonceRootChain)
	tx, err = pls.rootchainManager.rootchainContract.MapRequestableContractByOperator(operatorOpt, tokenAddrInRootChain, tokenAddrInChildChain)
	if err != nil {
		t.Fatalf("Failed to map token addresses to RootChain contract: %v", err)
	}

	waitEthTx(tx.Hash(), "")

	tokenAddr, err := pls.rootchainManager.rootchainContract.RequestableContracts(baseCallOpt, tokenAddrInRootChain)
	if err != nil {
		t.Fatalf("Failed to fetch token address from RootChain contract: %v", err)
	} else if tokenAddr != tokenAddrInChildChain {
		t.Fatalf("RootChain doesn't know requestable contract address in child chain: %v != %v", tokenAddrInChildChain.Hex(), tokenAddr.Hex())
	}

	// NRE#1 -> ORE#4 -- deposit 1 ether for each account
	ETHBalances1 := getETHBalances(addrs)
	PETHBalances1 := getPETHBalances(addrs)

	depositAmount := ether(1)
	depositAmountNeg := new(big.Int).Neg(depositAmount)

	for _, key := range keys {
		startETHDeposit(t, pls.rootchainManager, key, depositAmount)
	}

	// NRE#1 / Block#2 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 2); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 1); err != nil {
		t.Fatal(err)
	}

	// ORE#2 is empty
	if err := checkEpochAfter(pls, epochPreparedEvents, true, true, 0, 2); err != nil {
		t.Fatal(err)
	}

	// NRE#3 -> ORE#6 -- deposit 1 token from addr1
	tokenAmount := ether(1)
	tokenAmountNeg := new(big.Int).Neg(tokenAmount)

	TokenBalances1 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances1 := getTokenBalances(addrs, tokenInChildChain)

	startTokenDeposit(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key1, tokenAmount)

	// NRE#3 / Block#3 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 3); err != nil {
		t.Fatal(err)
	}

	// NRE#3 / Block#4 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 4); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 3); err != nil {
		t.Fatal(err)
	}

	// ORE#4 / Block#5 (1/1): deposit 1 ether for each account
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 5); err != nil {
		t.Fatal(err)
	}

	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 5); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 4); err != nil {
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
		if err := checkBalance(PETHBalances1[i], PETHBalances2[i], depositAmount, nil, "Failed to check PETH balance(1)"); err != nil {
			t.Fatal(err)
		}
	}

	// NRE#5 / Block#6 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 6); err != nil {
		t.Fatal(err)
	}

	// NRE#5 / Block#7 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 7); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 5); err != nil {
		t.Fatal(err)
	}

	// ORE#6 / Block#8 (1/1) -- deposit 1 token from addr1
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 8); err != nil {
		t.Fatal(err)
	}

	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 8); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 6); err != nil {
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

	// NRE#7 / Block#9 (1/2)
	transferToken(t, tokenInChildChain, key1, addr2, tokenAmountToTransfer, false)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 9); err != nil {
		t.Fatal(err)
	}

	PTokenBalances3 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances2[0], PTokenBalances3[0], tokenAmountToTransferNeg, nil, "Failed to check PToken Balance (2-1)"); err != nil {
		t.Fatal(err)
	}
	if err := checkBalance(PTokenBalances2[1], PTokenBalances3[1], tokenAmountToTransfer, nil, "Failed to check PToken Balance (2-2)"); err != nil {
		t.Fatal(err)
	}

	// NRE#7 / Block#10 (2/2)
	transferToken(t, tokenInChildChain, key1, addr2, tokenAmountToTransfer, false)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 10); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 7); err != nil {
		t.Fatal(err)
	}

	PTokenBalances4 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances3[0], PTokenBalances4[0], tokenAmountToTransferNeg, nil, "Failed to check PToken Balance (3-1)"); err != nil {
		t.Fatal(err)
	}
	if err := checkBalance(PTokenBalances3[1], PTokenBalances4[1], tokenAmountToTransfer, nil, "Failed to check PToken Balance (3-2)"); err != nil {
		t.Fatal(err)
	}

	// ORE#8 is empty
	if err := checkEpochAfter(pls, epochPreparedEvents, true, true, 0, 8); err != nil {
		t.Fatal(err)
	}

	// NRE#9 / Block#11 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 11); err != nil {
		t.Fatal(err)
	}

	// NRE#9 -> ORE#12 -- (1/4) withdraw addr2's token to root chain
	tokenAmountToWithdraw := new(big.Int).Div(tokenAmount, big.NewInt(4)) // 4 witndrawal requests
	tokenAmountToWithdrawNeg := new(big.Int).Neg(tokenAmountToWithdraw)

	PTokenBalances5 := getTokenBalances(addrs, tokenInChildChain)
	startTokenWithdraw(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#9 / Block#12 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 12); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 9); err != nil {
		t.Fatal(err)
	}

	// ORE#10 is empty
	if err := checkEpochAfter(pls, epochPreparedEvents, true, true, 0, 10); err != nil {
		t.Fatal(err)
	}

	// NRE#11 / Block#13 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 13); err != nil {
		t.Fatal(err)
	}

	// NRE#11 / Block#14 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 14); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 11); err != nil {
		t.Fatal(err)
	}

	// ORE#12 / Block#15 (1/1) -- (1/4) withdraw addr2's token to root chain
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 15); err != nil {
		t.Fatal(err)
	}

	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 15); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 12); err != nil {
		t.Fatal(err)
	}

	PTokenBalances6 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances5[1], PTokenBalances6[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance - token exit (1/4)"); err != nil {
		t.Fatal(err)
	}

	// NRE#13/ Block#16 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 16); err != nil {
		t.Fatal(err)
	}

	// NRE#13 -> ORE#16 -- (2/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#13/ Block#17 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 17); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 13); err != nil {
		t.Fatal(err)
	}

	// ORE#14 is empty
	if err := checkEpochAfter(pls, epochPreparedEvents, true, true, 0, 14); err != nil {
		t.Fatal(err)
	}

	// NRE#15/ Block#18 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 18); err != nil {
		t.Fatal(err)
	}

	// NRE#15 -> ORE#18 -- (3/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#15/ Block#19 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 19); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 15); err != nil {
		t.Fatal(err)
	}

	// ORE#16 / Block#20 -- (2/4) withdraw addr2's token to root chain
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 20); err != nil {
		t.Fatal(err)
	}
	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 20); err != nil {
		t.Fatal(err)
	}
	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 16); err != nil {
		t.Fatal(err)
	}

	PTokenBalances7 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances6[1], PTokenBalances7[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance - token exit (2/4)"); err != nil {
		t.Fatal(err)
	}

	// NRE#17/ Block#21 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 21); err != nil {
		t.Fatal(err)
	}

	// NRE#17 -> ORE#20 -- (4/4) withdraw addr2's token to root chain
	startTokenWithdraw(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key2, tokenAmountToWithdraw, big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#17/ Block#22 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 22); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 17); err != nil {
		t.Fatal(err)
	}

	// ORE#18 / Block#23 -- (3/4) withdraw addr2's token to root chain
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 23); err != nil {
		t.Fatal(err)
	}
	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 23); err != nil {
		t.Fatal(err)
	}
	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 18); err != nil {
		t.Fatal(err)
	}

	PTokenBalances8 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances7[1], PTokenBalances8[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance - token exit (3/4)"); err != nil {
		t.Fatal(err)
	}

	// NRE#19/ Block#24 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 24); err != nil {
		t.Fatal(err)
	}

	// NRE#19/ Block#25 (2/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 25); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 19); err != nil {
		t.Fatal(err)
	}

	// ORE#20 / Block#26 -- (4/4) withdraw addr2's token to root chain
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 26); err != nil {
		t.Fatal(err)
	}
	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 26); err != nil {
		t.Fatal(err)
	}
	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 20); err != nil {
		t.Fatal(err)
	}

	PTokenBalances9 := getTokenBalances(addrs, tokenInChildChain)
	if err := checkBalance(PTokenBalances8[1], PTokenBalances9[1], tokenAmountToWithdrawNeg, nil, "Failed to check PToken Balance - token exit (4/4)"); err != nil {
		t.Fatal(err)
	}

	// finalize until block#26
	finalizeBlocks(t, pls.rootchainManager.rootchainContract, 26)

	// wait request challenge period ends
	wait(10)

	// apply requests (4 ETH deposits, 1 Token deposits, 4 Token withdrawals)
	for i := 0; i < 4; i++ {
		if err := applyRequest(pls.rootchainManager.rootchainContract, key2); err != nil {
			t.Fatalf("Failed to finalize ether deposit request: %v", err)
		}
	}

	wait(10)

	if err := applyRequest(pls.rootchainManager.rootchainContract, key2); err != nil {
		t.Fatalf("Failed to finalize token deposit request: %v", err)
	}

	wait(10)

	for i := 0; i < 4; i++ {
		tokenBalanceBefore, _ := tokenInRootChain.Balances(baseCallOpt, addr2)
		if err := applyRequest(pls.rootchainManager.rootchainContract, key2); err != nil {
			t.Fatalf("Failed to finalize %dth token withdrawal request: %v", i, err)
		}

		tokenBalanceAfter, _ := tokenInRootChain.Balances(baseCallOpt, addr2)

		if tokenBalanceAfter.Cmp(new(big.Int).Add(tokenBalanceBefore, tokenAmountToWithdraw)) != 0 {
			t.Fatalf("applyRequest() does not increase token balance")
		}

		log.Info("Exit request applied")
	}

	t.Log("Test finished")
}

func TestInvalidExit(t *testing.T) {
	pls, rpcServer, dir, err := makePls()
	defer os.RemoveAll(dir)

	if err != nil {
		t.Fatalf("Failed to make pls service: %v", err)
	}
	defer pls.Stop()
	defer rpcServer.Stop()

	if err := pls.rootchainManager.Start(); err != nil {
		t.Fatalf("Failed to start RootChainManager: %v", err)
	}
	pls.protocolManager.Start(1)

	rpcClient := rpc.DialInProc(rpcServer)

	// assign to global variable
	plsClient = plsclient.NewClient(rpcClient)

	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
	blockSubmitWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
	defer blockFilterer.Unsubscribe()

	epochPreparedEvents := make(chan *rootchain.RootChainEpochPrepared)
	epochPreparedWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	epochPreparedFilterer, _ := pls.rootchainManager.rootchainContract.WatchEpochPrepared(epochPreparedWatchOpts, epochPreparedEvents)
	defer epochPreparedFilterer.Unsubscribe()

	var tx *types.Transaction

	wait(3)

	log.Info("All backends are set up")

	// NRE#1 / Block#1 (1/2)
	tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 1); err != nil {
		t.Fatal(err)
	}

	setNonce(operatorOpt, &operatorNonceRootChain)
	tx, err = tokenInRootChain.Mint(operatorOpt, addr1, ether(100))
	if err != nil {
		t.Fatalf("Failed to mint token: %v", err)
	}

	waitEthTx(tx.Hash(), "")

	ts1, err := tokenInRootChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from root chain: %v", err)
	}

	ts2, err := tokenInChildChain.TotalSupply(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get total supply from child chain: %v", err)
	}

	log.Info("Token total supply", "rootchain", ts1, "childchain", ts2)

	setNonce(operatorOpt, &operatorNonceRootChain)
	tx, err = pls.rootchainManager.rootchainContract.MapRequestableContractByOperator(operatorOpt, tokenAddrInRootChain, tokenAddrInChildChain)
	if err != nil {
		t.Fatalf("Failed to map token addresses to RootChain contract: %v", err)
	}

	waitEthTx(tx.Hash(), "")

	tokenAddr, err := pls.rootchainManager.rootchainContract.RequestableContracts(baseCallOpt, tokenAddrInRootChain)
	if err != nil {
		t.Fatalf("Failed to fetch token address from RootChain contract: %v", err)
	} else if tokenAddr != tokenAddrInChildChain {
		t.Fatalf("RootChain doesn't know requestable contract address in child chain: %v != %v", tokenAddrInChildChain.Hex(), tokenAddr.Hex())
	}

	// NRE#1 -> ORE#4 -- deposit 1 ether for each account
	ETHBalances1 := getETHBalances(addrs)
	PETHBalances1 := getPETHBalances(addrs)

	depositAmount := ether(10)
	depositAmountNeg := new(big.Int).Neg(depositAmount)

	for _, key := range keys {
		startETHDeposit(t, pls.rootchainManager, key, depositAmount)
	}

	// NRE#1 / Block#2 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 2); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 1); err != nil {
		t.Fatal(err)
	}

	// ORB#2 is empty
	if err := checkEpochAfter(pls, epochPreparedEvents, true, true, 0, 2); err != nil {
		t.Fatal(err)
	}

	// NRE#3 -> ORE#6 -- deposit 1 token from addr1
	tokenAmount := ether(1)
	tokenAmountNeg := new(big.Int).Neg(tokenAmount)

	TokenBalances1 := getTokenBalances(addrs, tokenInRootChain)
	PTokenBalances1 := getTokenBalances(addrs, tokenInChildChain)

	startTokenDeposit(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key1, tokenAmount)

	// NRE#3 / Block#3 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 3); err != nil {
		t.Fatal(err)
	}

	// NRE#3 / Block#4 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 4); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 3); err != nil {
		t.Fatal(err)
	}

	// ORE#4 / Block#5 (1/1) -- deposit 1 ether for each account
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 5); err != nil {
		t.Fatal(err)
	}

	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 5); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 4); err != nil {
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

	// NRE#5 / Block#6 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 6); err != nil {
		t.Fatal(err)
	}

	// NRE#5 / Block#7 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 7); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 5); err != nil {
		t.Fatal(err)
	}

	// ORE#6 / Block#8 (1/1) -- deposit 1 token from addr1
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 8); err != nil {
		t.Fatal(err)
	}

	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 8); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 6); err != nil {
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

	// NRE#7 -> ORB#9 -- invalid withdrawal
	startTokenWithdraw(t, pls.rootchainManager, tokenInRootChain, tokenAddrInRootChain, key2, ether(100), big.NewInt(int64(pls.rootchainManager.state.costERO)))

	// NRE#7 / Block#9 (1/2)
	makeSampleTx(pls.rootchainManager)
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 9); err != nil {
		t.Fatal(err)
	}

	// NRE#7 / Block#10 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 10); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 7); err != nil {
		t.Fatal(err)
	}

	// ORB#8 is empty
	if err := checkEpochAfter(pls, epochPreparedEvents, true, true, 0, 8); err != nil {
		t.Fatal(err)
	}

	// NRE#9 / Block#11 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 11); err != nil {
		t.Fatal(err)
	}

	// NRE#9 / Block#12 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 12); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 9); err != nil {
		t.Fatal(err)
	}

	// ORE#10 / Block#13 (1/1) -- invalid withdrawal
	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, true, 0, 13); err != nil {
		t.Fatal(err)
	}

	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 13); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, true, 0, 10); err != nil {
		t.Fatal(err)
	}

	// NRE#11 / Block#14 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 14); err != nil {
		t.Fatal(err)
	}

	// NRE#11 / Block#15 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 15); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 11); err != nil {
		t.Fatal(err)
	}

	// ORB#12 is empty
	if err := checkEpochAfter(pls, epochPreparedEvents, true, true, 0, 12); err != nil {
		t.Fatal(err)
	}

	// NRE#13 / Block#16 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 16); err != nil {
		t.Fatal(err)
	}

	// NRE#13 / Block#17 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 17); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 13); err != nil {
		t.Fatal(err)
	}

	// ORB#14 is empty
	if err := checkEpochAfter(pls, epochPreparedEvents, true, true, 0, 14); err != nil {
		t.Fatal(err)
	}

	// NRE#15 / Block#18 (1/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 18); err != nil {
		t.Fatal(err)
	}

	// NRE#15 / Block#19 (2/2)
	makeSampleTx(pls.rootchainManager)

	if err := checkMinedBlock(pls, plasmaBlockMinedEvents, false, 0, 19); err != nil {
		t.Fatal(err)
	}

	if err := checkEpochAfter(pls, epochPreparedEvents, false, false, 0, 15); err != nil {
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
}

//func TestStress(t *testing.T) {
//	// override test parameters
//	NRELength = big.NewInt(1024)
//
//	testPlsConfig.Miner.Recommit = 10 * time.Second
//	testPlsConfig.TxConfig.Interval = 10 * time.Second
//	timeout := testPlsConfig.Miner.Recommit / 100
//
//	targetBlockNumber := 10
//
//	pls, rpcServer, dir, err := makePls()
//	defer os.RemoveAll(dir)
//
//	if err != nil {
//		t.Fatalf("Failed to make pls service: %v", err)
//	}
//	defer pls.Stop()
//	defer rpcServer.Stop()
//
//	if err := pls.rootchainManager.Start(); err != nil {
//		t.Fatalf("Failed to start RootChainManager: %v", err)
//	}
//	pls.protocolManager.Start(1)
//
//	rpcClient := rpc.DialInProc(rpcServer)
//
//	// assign to global variable
//	plsClient = plsclient.NewClient(rpcClient)
//
//	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
//	defer plasmaBlockMinedEvents.Unsubscribe()
//
//	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
//	blockSubmitWatchOpts := &bind.WatchOpts{
//		Start:   nil,
//		Context: context.Background(),
//	}
//	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
//	defer blockFilterer.Unsubscribe()
//
//	blockNumber := 0
//
//	timer := time.NewTimer(5 * time.Minute)
//	go func() {
//		<-timer.C
//		t.Fatal("Out of time")
//	}()
//
//	txs := types.Transactions{}
//	nTxsInBlocks := 0
//
//	blockNumber++
//
//	for _, addr := range addrs {
//		var tx *types.Transaction
//		if tx, err = transferETH(operatorKey, addr, ether(1), false); err != nil {
//			t.Fatalf("Failed to transfer PETH: %v", err)
//		}
//		txs = append(txs, tx)
//	}
//
//	if err := checkSubmittedBlock(pls, blockSubmitEvents, 0, 13); err != nil {
//		t.Fatal(err)
//	}
//
//	for blockNumber < targetBlockNumber {
//		blockNumber++
//
//		var tx *types.Transaction
//		for _, addr := range addrs {
//			if tx, err = transferETH(operatorKey, addr, ether(1), false); err != nil {
//				t.Fatalf("Failed to transfer PETH: %v", err)
//			}
//		}
//
//		txs = append(txs, tx)
//
//		done := make(chan struct{})
//		wg := sync.WaitGroup{}
//
//		wg.Add(1)
//		go func(t *testing.T) {
//			for {
//				timer := time.NewTimer(timeout)
//				select {
//				case <-timer.C:
//					timer.Reset(timeout)
//
//					for i, addr := range addrs {
//						go func(t *testing.T, i int, addr common.Address) {
//							var tx *types.Transaction
//
//							if tx, err = transferETH(operatorKey, addr, ether(1), false); err != nil {
//								t.Fatalf("Failed to transfer PETH: %v", err)
//							}
//							txs = append(txs, tx)
//
//							key := keys[i]
//
//							if tx, err = transferETH(key, addr, ether(0.0001), false); err != nil {
//								t.Fatalf("Failed to transfer PETH: %v", err)
//							}
//							txs = append(txs, tx)
//						}(t, i, addr)
//					}
//
//				case <-done:
//					wg.Done()
//					return
//				}
//			}
//		}(t)
//
//		// check new block is mined
//		wg.Add(1)
//		go func() {
//			if err := checkSubmittedBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, int64(blockNumber)); err != nil {
//				t.Fatal(err)
//			}
//			close(done)
//			b := pls.blockchain.GetBlockByNumber(uint64(blockNumber))
//			nTxsInBlocks += len(b.Transactions())
//			wg.Done()
//		}()
//
//		wg.Wait()
//	}
//
//	wait(testPlsConfig.Miner.Recommit * 2 / time.Second)
//
//	for _, tx := range txs {
//		r, isPending, err := plsClient.TransactionByHash(context.Background(), tx.Hash())
//		signer := types.NewEIP155Signer(params.MainnetChainConfig.ChainID)
//		msg, _ := r.AsMessage(signer)
//		from := msg.From()
//
//		if isPending {
//			t.Fatalf("Transaction %s is pending (from: %s, nonce: %d)", r.Hash().String(), from.String(), r.Nonce())
//		}
//
//		if err != nil {
//			t.Fatalf("failed to get transaction receipt: %v", err)
//		}
//
//		if r == nil {
//			t.Fatalf("Transaction %s not mined (from: %s, nonce: %d)", r.Hash().String(), from.String(), r.Nonce())
//		}
//	}
//
//	nTxs := 0
//	lastBlockNumber := pls.blockchain.CurrentBlock().NumberU64()
//	for i := 1; i <= int(lastBlockNumber); i++ {
//		block := pls.blockchain.GetBlockByNumber(uint64(i))
//
//		if block.Transactions().Len() == 0 {
//			t.Fatalf("Block#%d has no transaction", i)
//		}
//
//		nTxs += block.Transactions().Len()
//	}
//
//	firstBlock := pls.blockchain.GetBlockByNumber(uint64(1))
//	lastBlock := pls.blockchain.GetBlockByNumber(uint64(lastBlockNumber))
//
//	elapsed := lastBlock.Time() - firstBlock.Time()
//	tps := float64(nTxs) / float64(elapsed)
//	t.Logf("Elapsed time: %s nTxs: %d TPS: %6.3f", elapsed, nTxs, tps)
//
//}

//func TestRestart(t *testing.T) {
//	timer := time.NewTimer(2 * time.Minute)
//	go func() {
//		<-timer.C
//		t.Fatal("Out of time")
//	}()
//
//	pls, rpcServer, dir, err := makePls()
//	defer os.RemoveAll(dir)
//
//	if err != nil {
//		t.Fatalf("Failed to make pls service: %v", err)
//	}
//	defer pls.Stop()
//	defer rpcServer.Stop()
//
//	if err := pls.rootchainManager.Start(); err != nil {
//		t.Fatalf("Failed to start RootChainManager: %v", err)
//	}
//	pls.protocolManager.Start(1)
//
//	rpcClient := rpc.DialInProc(rpcServer)
//
//	// assign to global variable
//	plsClient = plsclient.NewClient(rpcClient)
//
//	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
//	defer plasmaBlockMinedEvents.Unsubscribe()
//
//	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
//	blockSubmitWatchOpts := &bind.WatchOpts{
//		Start:   nil,
//		Context: context.Background(),
//	}
//	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
//	defer blockFilterer.Unsubscribe()
//
//	rcm := pls.rootchainManager
//
//	wait(3)
//
//	log.Info("All backends are set up")
//
//	// ORE#4 make enter request
//	enterAmount := ether(1)
//
//	startETHDeposit(t, rcm, key1, enterAmount)
//	startETHDeposit(t, rcm, key2, enterAmount)
//	startETHDeposit(t, rcm, key3, enterAmount)
//	startETHDeposit(t, rcm, key4, enterAmount)
//
//	// NRE#1 / NRB#1
//	// deploy EtherToken in child chain
//	deployEtherTokenInChildChain(t)
//
//	if err := checkSubmittedBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 1); err != nil {
//		t.Fatal(err)
//	}
//	// NRE#1 / NRB#2
//	//tokenInRootChain, tokenInChildChain, tokenAddrInRootChain, tokenAddrInChildChain := deployTokenContracts(t)
//	deployTokenContracts(t)
//
//	if err := checkSubmittedBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 2); err != nil {
//		t.Fatal(err)
//	}
//
//	// ORE#2 is empty
//
//	// NRE#3 / NRB#3
//	makeSampleTx(pls.rootchainManager)
//	if err := checkSubmittedBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 3); err != nil {
//		t.Fatal(err)
//	}
//	// NRE#3 / NRB#4
//	makeSampleTx(pls.rootchainManager)
//	if err := checkSubmittedBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 4); err != nil {
//		t.Fatal(err)
//	}
//
//	// ORE#4 / ORB#5 : enter request
//	if err := checkSubmittedBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, true, 0, 5); err != nil {
//		t.Fatal(err)
//	}
//
//	// stop pls service
//}

//func TestMinerRestart(t *testing.T) {
//	pls, rpcServer, dir, err := makePls()
//	defer os.RemoveAll(dir)
//
//	if err != nil {
//		t.Fatalf("Failed to make pls service: %v", err)
//	}
//	defer pls.Stop()
//	defer rpcServer.Stop()
//
//	// pls.Start()
//	pls.protocolManager.Start(1)
//
//	if err := pls.rootchainManager.Start(); err != nil {
//		t.Fatalf("Failed to start RootChainManager: %v", err)
//	}
//
//	pls.StartMining(runtime.NumCPU())
//
//	rpcClient := rpc.DialInProc(rpcServer)
//
//	// assign to global variable
//	plsClient = plsclient.NewClient(rpcClient)
//
//	plasmaBlockMinedEvents := pls.rootchainManager.eventMux.Subscribe(core.NewMinedBlockEvent{})
//	defer plasmaBlockMinedEvents.Unsubscribe()
//
//	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
//	blockSubmitWatchOpts := &bind.WatchOpts{
//		Start:   nil,
//		Context: context.Background(),
//	}
//	blockFilterer, _ := pls.rootchainManager.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
//	defer blockFilterer.Unsubscribe()
//
//	log.Info("All backends are set up")
//
//	// NRE#1 / Block#1 (1/2)
//	makeSampleTx(pls.rootchainManager)
//
//	if err := checkSubmittedBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 1); err != nil {
//		t.Fatal(err)
//	}
//
//	pls.StopMining()
//	blockBeforeStop := pls.blockchain.CurrentBlock()
//
//	wait(5)
//	pls.StartMining(runtime.NumCPU())
//
//	// NRE#1 / Block#2 (2/2)
//	makeSampleTx(pls.rootchainManager)
//
//	if err := checkSubmittedBlock(pls, plasmaBlockMinedEvents, blockSubmitEvents, false, 0, 2); err != nil {
//		t.Fatal(err)
//	}
//
//	blockAfterStop := pls.blockchain.CurrentBlock()
//	if diff := blockAfterStop.NumberU64() - blockBeforeStop.NumberU64(); diff != 1 {
//		t.Fatal("failed to resume current epoch", "difference", diff)
//	}
//}

func startETHDeposit(t *testing.T, rcm *RootChainManager, key *ecdsa.PrivateKey, amount *big.Int) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot deposit 0 ETH")
	}

	watchOpt := &bind.WatchOpts{Start: nil, Context: context.Background()}
	event := make(chan *rootchain.RootChainRequestCreated)
	filterer, _ := rcm.rootchainContract.WatchRequestCreated(watchOpt, event)
	defer filterer.Unsubscribe()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	opt := opts[addr]
	setNonce(opt, noncesRootChain[addr])

	trieKey, err := etherToken.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := amount.Bytes()
	trieValue = common.LeftPadBytes(trieValue, 32)
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rcm.rootchainContract.StartEnter(opt, etherTokenAddr, trieKey, trieValue32Bytes[:])

	if err != nil {
		t.Fatalf("Failed to make an ETH deposit request: %v", err)
	}

	if err = waitEthTx(tx.Hash(), ""); err != nil {
		t.Fatalf("Failed to make an ETH deposit request: %v", err)
	}

	timer := time.NewTimer(time.Second * 10)
	select {
	case request := <-event:
		log.Debug("Ether deposit request", "requestId", request.RequestId, "request", request)
	case <-timer.C:
		t.Error("out of time")
	}

	receipt, err := rcm.backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("Failed to send eth deposit tx: %v", err)
	} else if receipt.Status == 0 {
		t.Fatal("ETH deposit tx is reverted")
	}
}

func startTokenDeposit(t *testing.T, rcm *RootChainManager, tokenContract *token.RequestableSimpleToken, tokenAddress common.Address, key *ecdsa.PrivateKey, amount *big.Int) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot deposit 0 Token")
	}

	watchOpt := &bind.WatchOpts{Start: nil, Context: context.Background()}
	event := make(chan *rootchain.RootChainRequestCreated)
	filterer, _ := rcm.rootchainContract.WatchRequestCreated(watchOpt, event)
	defer filterer.Unsubscribe()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	opt := opts[addr]
	setNonce(opt, noncesRootChain[addr])

	trieKey, err := tokenContract.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := amount.Bytes()
	trieValue = common.LeftPadBytes(trieValue, 32)
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rcm.rootchainContract.StartEnter(opt, tokenAddress, trieKey, trieValue32Bytes[:])

	if err != nil {
		log.Error("Failed to make an token deposit request", "err", err, "hash", tx.Hash())
		t.Fatalf("Failed to make an token deposit request: %v", err)
	}

	if err = waitEthTx(tx.Hash(), ""); err != nil {
		t.Fatalf("Failed to make an Token deposit request: %v", err)
	}

	timer := time.NewTimer(time.Second * 10)
	select {
	case request := <-event:
		log.Debug("Token deposit request", "requestId", request.RequestId, "request", request)
	case <-timer.C:
		t.Error("out of time")
	}

	receipt, err := rcm.backend.TransactionReceipt(context.Background(), tx.Hash())

	log.Debug("Token deposit", "hash", tx.Hash().String(), "receipt", receipt)

	if err != nil {
		t.Fatalf("Failed to send token deposit tx: %v", err)
	} else if receipt.Status == 0 {
		t.Fatal("Token deposit tx is reverted")
	}
}

func startETHWithdraw(t *testing.T, rcm *RootChainManager, key *ecdsa.PrivateKey, value, cost *big.Int) {
	if value.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot deposit 0 ETH")
	}

	watchOpt := &bind.WatchOpts{Start: nil, Context: context.Background()}
	event := make(chan *rootchain.RootChainRequestCreated)
	filterer, _ := rcm.rootchainContract.WatchRequestCreated(watchOpt, event)
	defer filterer.Unsubscribe()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	opt := makeTxOpt(key, 3000000, nil, cost)
	setNonce(opt, noncesRootChain[addr])

	trieKey, err := etherToken.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := value.Bytes()
	trieValue = common.LeftPadBytes(trieValue, 32)
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rcm.rootchainContract.StartExit(opt, etherTokenAddr, trieKey, trieValue32Bytes[:])

	if err != nil {
		t.Fatalf("Failed to make an ETH withdraw request: %v", err)
	}

	if err = waitEthTx(tx.Hash(), ""); err != nil {
		t.Fatalf("Failed to make an Token deposit request: %v", err)
	}

	timer := time.NewTimer(time.Second * 10)
	select {
	case request := <-event:
		log.Debug("ETH withdrawal request", "requestId", request.RequestId, "request", request)
	case <-timer.C:
		t.Error("out of time")
	}

	receipt, err := rcm.backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("Failed to send eth deposit tx: %v", err)
	} else if receipt.Status == 0 {
		t.Fatal("ETH withdraw tx is reverted")
	}
}

func startTokenWithdraw(t *testing.T, rcm *RootChainManager, tokenContract *token.RequestableSimpleToken, tokenAddress common.Address, key *ecdsa.PrivateKey, amount, cost *big.Int) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Cannot withdraw 0 Token")
	}

	watchOpt := &bind.WatchOpts{Start: nil, Context: context.Background()}
	event := make(chan *rootchain.RootChainRequestCreated)
	filterer, _ := rcm.rootchainContract.WatchRequestCreated(watchOpt, event)
	defer filterer.Unsubscribe()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	opt := makeTxOpt(key, 3000000, nil, cost)
	setNonce(opt, noncesRootChain[addr])

	trieKey, err := tokenContract.GetBalanceTrieKey(baseCallOpt, addr)
	if err != nil {
		t.Fatalf("Failed to get trie key: %v", err)
	}
	trieValue := amount.Bytes()
	trieValue = common.LeftPadBytes(trieValue, 32)
	trieValue32Bytes := common.BytesToHash(trieValue)

	tx, err := rcm.rootchainContract.StartExit(opt, tokenAddress, trieKey, trieValue32Bytes[:])

	if err != nil {
		t.Fatalf("Failed to make an token withdrawal request: %v", err)
	}

	if err = waitEthTx(tx.Hash(), ""); err != nil {
		t.Fatalf("Failed to make an Token withdrawal request: %v", err)
	}

	timer := time.NewTimer(time.Second * 10)
	select {
	case request := <-event:
		log.Debug("Toke withdrawal request", "requestId", request.RequestId, "request", request)
	case <-timer.C:
		t.Error("out of time")
	}

	receipt, err := rcm.backend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatalf("Failed to send eth deposit tx: %v", err)
	} else if receipt.Status == 0 {
		t.Fatal("Token withdraw tx is reverted")
	}
}

func transferToken(t *testing.T, tokenContract *token.RequestableSimpleToken, key *ecdsa.PrivateKey, to common.Address, amount *big.Int, isRootChain bool) {
	opt := makeTxOpt(key, 0, nil, nil)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	if isRootChain {
		setNonce(opt, noncesRootChain[addr])
	} else {
		setNonce(opt, noncesChildChain[addr])
	}

	_, err := tokenContract.Transfer(opt, to, amount)
	if err != nil {
		t.Fatalf("Failed to transfer toekn: %v", err)
	}
}

func transferETH(key *ecdsa.PrivateKey, to common.Address, amount *big.Int, isRootChain bool) (*types.Transaction, error) {
	opt := makeTxOpt(key, 21000, defaultGasPrice, amount)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	locks[addr].Lock()
	defer locks[addr].Unlock()

	if isRootChain {
		setNonce(opt, noncesRootChain[addr])
	} else {
		setNonce(opt, noncesChildChain[addr])
	}

	tx := types.NewTransaction(opt.Nonce.Uint64(), to, amount, 21000, defaultGasPrice, []byte{})

	var err error

	chainId := params.MainnetChainConfig.ChainID

	if isRootChain {
		chainId = testPlsConfig.TxConfig.ChainId
	}

	signer := types.NewEIP155Signer(chainId)
	signedTx, err := types.SignTx(tx, signer, key)
	if err != nil {
		return nil, err
	}

	if isRootChain {
		err = ethClient.SendTransaction(context.Background(), signedTx)
		return signedTx, err
	}

	err = plsClient.SendTransaction(context.Background(), signedTx)
	return signedTx, err
}

func finalizeBlocks(t *testing.T, rootchainContract *rootchain.RootChain, targetNumber int64) {
	target := big.NewInt(targetNumber)

	last, err := rootchainContract.GetLastFinalizedBlock(baseCallOpt, big.NewInt(0))
	if err != nil {
		t.Fatalf("Failed to GetLastFinalizedBlock: %v", err)
	}

	cp, err := rootchainContract.CPWITHHOLDING(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get CP_WITHHOLDING: %v", err)
	}

	for last.Cmp(target) < 0 {
		log.Info("Try to finalize block", "lastFinalizedBlock", last, "lastBlock", target)

		nTries := 2
		for i := 0; i < nTries; i++ {
			opt := opts[addr1]
			opt.Value = nil
			setNonce(opt, noncesRootChain[addr1])
			opt.GasLimit = 6000000

			tx, err := rootchainContract.FinalizeBlock(opt)
			if err != nil {
				t.Errorf("Failed to fianlize block: %v", err)
			}

			waitEthTx(tx.Hash(), fmt.Sprintf("Finalize Block#%d", last.Uint64()+1))

			receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				t.Errorf("Failed to get transaction receipt: %v", err)
			}

			last, err = rootchainContract.GetLastFinalizedBlock(baseCallOpt, big.NewInt(0))
			if err != nil {
				t.Fatalf("Failed to GetLastFinalizedBlock: %v", err)
			}

			if receipt.Status == 1 {
				break
			}

			if i == nTries-1 {
				t.Fatalf("Failed to finalize block%d up to %d tries", last.Uint64()+1, nTries)
			}
			log.Error("FinalizeBlock transaction is failed")
			wait(time.Duration(cp.Uint64()) + 1)
		}
	}

	log.Info("All blocks are finalized")
}

// apply a single request
func applyRequest(rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey) error {
	requestIdToFinalize, err := rootchainContract.EROIdToFinalize(baseCallOpt)
	if err != nil {
		return err
	}

	numRequests, err := rootchainContract.GetNumEROs(baseCallOpt)
	if err != nil {
		return err
	}

	cp, err := rootchainContract.CPEXIT(baseCallOpt)
	if err != nil {
		return err
	}

	if numRequests.Uint64() == requestIdToFinalize.Uint64() {
		return errors.New("There is no request to finalize")
	}

	log.Info("Finalize request", "targetRequestId", requestIdToFinalize, "numRequests", numRequests)

	n := 2
	var tx *types.Transaction
	for i := 0; i < n; i++ {
		opt := makeTxOpt(key, 5000000, nil, nil)
		addr := crypto.PubkeyToAddress(key.PublicKey)

		setNonce(opt, noncesRootChain[addr])

		tx, err = rootchainContract.FinalizeRequest(opt)

		if err != nil {
			return fmt.Errorf("failed to send transaction to finalize requeest: %v", err)
		}

		err = waitEthTx(tx.Hash(), fmt.Sprintf("finalizeRequest(%d)", requestIdToFinalize.Uint64()))

		if err == nil {
			return nil
		}

		wait(time.Duration(cp.Uint64() + 1))
	}

	if err != nil {
		return fmt.Errorf("Failed to finalize request#%d up to %d tries: %v", requestIdToFinalize.Uint64(), n, err)
	}

	return nil
}

// apply all requests
func applyRequests(t *testing.T, rootchainContract *rootchain.RootChain, key *ecdsa.PrivateKey) {
	opt := makeTxOpt(key, 2000000, nil, nil)
	addr := crypto.PubkeyToAddress(key.PublicKey)

	last, err := rootchainContract.EROIdToFinalize(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get last applied ERO: %v", err)
	}

	target, err := rootchainContract.GetNumEROs(baseCallOpt)
	if err != nil {
		t.Fatalf("Failed to get number of EROs: %v", err)
	}

	target = new(big.Int).Sub(target, big.NewInt(1))

	for last.Cmp(target) < 0 {
		log.Info("Try to apply request", "last", last, "target", target)
		setNonce(opt, noncesRootChain[addr])

		wait(1)

		tx, err := rootchainContract.FinalizeRequest(opt)
		if err != nil {
			t.Fatalf("failed to apply requeest: %v", err)
		}

		if err := waitEthTx(tx.Hash(), ""); err != nil {
			t.Fatalf("failed to apply requeest: %v", err)

		}

		last, _ = rootchainContract.EROIdToFinalize(baseCallOpt)
		target, _ = rootchainContract.GetNumEROs(baseCallOpt)

	}
}

func deployRootChain(genesis *types.Block) (rootchainAddress common.Address, rootchainContract *rootchain.RootChain, err error) {

	dummyDB := rawdb.NewMemoryDatabase()
	defer dummyDB.Close()
	dummyBlock := core.DeveloperGenesisBlock(
		0,
		common.HexToAddress("0xdead"),
		operator,
		params.DefaultStaminaConfig,
	).ToBlock(dummyDB)

	var tx *types.Transaction
	log.Info("Deploying contracts for development mode")

	// 1. deploy MintableToken
	setNonce(operatorOpt, &operatorNonceRootChain)
	mintableTokenAddr, tx, mintableToken, err = mintabletoken.DeployERC20Mintable(operatorOpt, ethClient)

	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy MintableToken contract: %v", err))
	}
	log.Info("Deploy MintableToken contract", "hash", tx.Hash(), "address", mintableTokenAddr)

	log.Info("Wait until deploy transaction is mined")
	waitEthTx(tx.Hash(), "")

	// 2. deploy EtherToken
	setNonce(operatorOpt, &operatorNonceRootChain)
	etherTokenAddr, tx, etherToken, err = ethertoken.DeployEtherToken(operatorOpt, ethClient, development, mintableTokenAddr, swapEnabledInRootChain)

	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy EtherToken contract: %v", err))
	}
	log.Info("Deploy EtherToken contract", "hash", tx.Hash(), "address", etherTokenAddr)

	log.Info("Wait until deploy transaction is mined")
	waitEthTx(tx.Hash(), "")

	// 3. deploy EpochHandler
	setNonce(operatorOpt, &operatorNonceRootChain)
	epochHandlerAddr, tx, _, err := epochhandler.DeployEpochHandler(operatorOpt, ethClient)

	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy EpochHandler contract: %v", err))
	}
	log.Info("Deploy EpochHandler contract", "hash", tx.Hash(), "address", epochHandlerAddr)

	log.Info("Wait until deploy transaction is mined")
	waitEthTx(tx.Hash(), "")

	// 4. deploy SubmitHandler
	setNonce(operatorOpt, &operatorNonceRootChain)
	submitHandlerAddr, tx, _, err := submithandler.DeploySubmitHandler(operatorOpt, ethClient, epochHandlerAddr)
	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy SubmitHandler contract: %v", err))
	}
	log.Info("Deploy SubmitHandler contract", "hash", tx.Hash(), "address", submitHandlerAddr)

	log.Info("Wait until deploy transaction is mined")
	waitEthTx(tx.Hash(), "")

	// 5. deploy RootChain
	setNonce(operatorOpt, &operatorNonceRootChain)
	rootchainAddr, tx, rootchainContract, err := rootchain.DeployRootChain(operatorOpt, ethClient, epochHandlerAddr, submitHandlerAddr, etherTokenAddr, development, NRELength, dummyBlock.Root(), dummyBlock.TxHash(), dummyBlock.ReceiptHash())
	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to deploy RootChain contract: %v", err))
	}
	log.Info("Deploy RootChain contract", "hash", tx.Hash(), "address", rootchainAddr)
	waitEthTx(tx.Hash(), "")

	// 5. initialize EtherToken
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx, err = etherToken.Init(operatorOpt, rootchainAddr)
	if err != nil {
		return common.Address{}, nil, errors.New(fmt.Sprintf("Failed to initialize EtherToken: %v", err))
	}
	log.Info("Initialize EtherToken", "hash", tx.Hash())
	waitEthTx(tx.Hash(), "")

	// 6. mint tokens
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx1, err := mintableToken.Mint(operatorOpt, addr1, ether(100))
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx2, err := mintableToken.Mint(operatorOpt, addr2, ether(100))
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx3, err := mintableToken.Mint(operatorOpt, addr3, ether(100))
	setNonce(operatorOpt, &operatorNonceRootChain)
	tx4, err := mintableToken.Mint(operatorOpt, addr4, ether(100))

	waitEthTx(tx1.Hash(), "")
	waitEthTx(tx2.Hash(), "")
	waitEthTx(tx3.Hash(), "")
	waitEthTx(tx4.Hash(), "")

	log.Info("Mint MintableToken to users")

	// 7. swap MintableToken to EtherToken
	setNonce(opt1, &addr1NonceRootChain)
	setNonce(opt2, &addr2NonceRootChain)
	setNonce(opt3, &addr3NonceRootChain)
	setNonce(opt4, &addr4NonceRootChain)

	if tx1, err = mintableToken.Approve(opt1, etherTokenAddr, ether(100)); err != nil {
		log.Error("Failed to approve MintableToken to EtherToken", "err", err)
	}
	if tx2, err = mintableToken.Approve(opt2, etherTokenAddr, ether(100)); err != nil {
		log.Error("Failed to approve MintableToken to EtherToken", "err", err)
	}
	if tx3, err = mintableToken.Approve(opt3, etherTokenAddr, ether(100)); err != nil {
		log.Error("Failed to approve MintableToken to EtherToken", "err", err)
	}
	if tx4, err = mintableToken.Approve(opt4, etherTokenAddr, ether(100)); err != nil {
		log.Error("Failed to approve MintableToken to EtherToken", "err", err)
	}

	log.Info("MintableToken is approved to EtherToken")

	waitEthTx(tx1.Hash(), "")
	waitEthTx(tx2.Hash(), "")
	waitEthTx(tx3.Hash(), "")
	waitEthTx(tx4.Hash(), "")

	setNonce(opt1, &addr1NonceRootChain)
	setNonce(opt2, &addr2NonceRootChain)
	setNonce(opt3, &addr3NonceRootChain)
	setNonce(opt4, &addr4NonceRootChain)

	tx1, _ = etherToken.Deposit(opt1, ether(100))
	tx2, _ = etherToken.Deposit(opt2, ether(100))
	tx3, _ = etherToken.Deposit(opt3, ether(100))
	tx4, _ = etherToken.Deposit(opt4, ether(100))

	waitEthTx(tx1.Hash(), "")
	waitEthTx(tx2.Hash(), "")
	waitEthTx(tx3.Hash(), "")
	waitEthTx(tx4.Hash(), "")

	log.Info("Swap MintableToken to EtherToken")

	for i, addr := range addrs {
		bal, err := etherToken.BalanceOf(baseCallOpt, addr)
		if err != nil {
			log.Error("Failed to get EtherToken balance", "err", err)
		}

		bal.Div(bal, ether(1))

		log.Info("EtherToken balance", "i", i, "balance", bal)
	}

	testPlsConfig.RootChainContract = rootchainAddr

	return rootchainAddr, rootchainContract, err
}

func newCanonical(n int, full bool) (ethdb.Database, *core.BlockChain, error) {
	gspec := core.DeveloperGenesisBlock(0, common.Address{}, operator, params.DefaultStaminaConfig)
	// Initialize a fresh chain with only a genesis block
	db := rawdb.NewMemoryDatabase()
	genesis := gspec.MustCommit(db)

	blockchain, _ := core.NewBlockChain(db, nil, params.MainnetChainConfig, engine, testVmConfg, nil)
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
	blocks, _ := core.GenerateChain(params.MainnetChainConfig, parent, engine, db, n, func(i int, b *core.BlockGen) {
		b.SetCoinbase(common.Address{0: byte(seed), 19: byte(i)})
	})
	return blocks
}

func newTxPool(blockchain *core.BlockChain) *core.TxPool {
	pool := core.NewTxPool(*testTxPoolConfig, params.MainnetChainConfig, blockchain)

	return pool
}

func tmpKeyStore() (string, *keystore.KeyStore) {
	d, err := ioutil.TempDir("/tmp", "eth-keystore-test")
	if err != nil {
		log.Error("Failed to set temporary keystore directory", "err", err)
	}
	ks := keystore.NewKeyStore(d, 2, 1)

	return d, ks
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

func makePls() (*Plasma, *rpc.Server, string, error) {
	db, blockchain, err := newCanonical(0, true)

	if err != nil {
		log.Error("Failed to creaet blockchain", "err", err)
		return nil, nil, "", err
	}

	config := testPlsConfig
	chainConfig := params.MainnetChainConfig

	rootchainAddress, rootchainContract, err := deployRootChain(blockchain.Genesis())

	if err != nil {
		log.Error("Failed to deploy rootchain contract", "err", err)
		return nil, nil, "", err
	}

	config.RootChainContract = rootchainAddress

	d, ks := tmpKeyStore()

	var oac accounts.Account
	var cac accounts.Account
	if oac, err = ks.ImportECDSA(operatorKey, ""); err != nil {
		log.Error("Failed to import operator account", "err", err)
	}

	if cac, err = ks.ImportECDSA(challengerKey, ""); err != nil {
		log.Error("Failed to import challenger account", "err", err)
	}

	if err = ks.Unlock(oac, ""); err != nil {
		log.Error("Failed to unlock operator account", "err", err)
	}

	if err = ks.Unlock(cac, ""); err != nil {
		log.Error("Failed to unlock challenger account", "err", err)
	}

	for _, key := range keys {
		var acc accounts.Account
		if acc, err = ks.ImportECDSA(key, ""); err != nil {
			log.Error("Failed to import user account", "err", err)
		}

		if err := ks.Unlock(acc, ""); err != nil {
			log.Error("Failed to unlock user  account", "err", err)
		}
	}

	config.Operator = oac
	config.Challenger = cac

	// configure account manager with temporary keystore backend
	backends := []accounts.Backend{
		ks,
	}

	accountConfig := accounts.Config{true}
	accManager := accounts.NewManager(&accountConfig, backends...)

	pls := &Plasma{
		config:         config,
		chainDb:        db,
		eventMux:       new(event.TypeMux),
		accountManager: accManager,
		blockchain:     blockchain,
		engine:         engine,
		shutdownChan:   make(chan bool),
		networkID:      config.NetworkId,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   NewBloomIndexer(db, params.BloomBitsBlocks, params.BloomConfirms),
	}
	pls.bloomIndexer.Start(pls.blockchain)
	pls.txPool = core.NewTxPool(config.TxPool, chainConfig, pls.blockchain)

	cacheConfig := &core.CacheConfig{
		TrieCleanLimit:      config.TrieCleanCache,
		TrieCleanNoPrefetch: config.NoPrefetch,
		TrieDirtyLimit:      config.TrieDirtyCache,
		TrieDirtyDisabled:   config.NoPruning,
		TrieTimeLimit:       config.TrieTimeout,
	}
	cacheLimit := cacheConfig.TrieCleanLimit + cacheConfig.TrieDirtyLimit

	if pls.protocolManager, err = NewProtocolManager(chainConfig, nil, config.SyncMode, config.NetworkId, pls.eventMux, pls.txPool, pls.engine, pls.blockchain, db, cacheLimit, config.Whitelist); err != nil {
		return nil, nil, d, err
	}

	epochEnv := epoch.New()
	pls.miner = miner.New(pls, &config.Miner, chainConfig, pls.EventMux(), pls.engine, epochEnv, db, pls.isLocalBlock)
	pls.miner.SetExtra(makeExtraData(config.Miner.ExtraData))
	pls.APIBackend = &PlsAPIBackend{false, pls, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.Miner.GasPrice
	}
	pls.APIBackend.gpo = gasprice.NewOracle(pls.APIBackend, gpoParams)
	// Dial rootchain provider
	rootchainBackend, err := ethclient.Dial(config.RootChainURL)
	if err != nil {
		return nil, nil, d, err
	}
	log.Info("Rootchain provider connected", "url", config.RootChainURL)

	if err != nil {
		return nil, nil, d, err
	}

	stopFn := func() { pls.Stop() }
	txManager, err := tx.NewTransactionManager(ks, rootchainBackend, db, &config.TxConfig)

	if err != nil {
		return nil, nil, d, err
	}

	if pls.rootchainManager, err = NewRootChainManager(
		config,
		stopFn,
		pls.txPool,
		pls.blockchain,
		rootchainBackend,
		rootchainContract,
		pls.eventMux,
		pls.accountManager,
		txManager,
		pls.miner,
		epochEnv,
	); err != nil {
		return nil, nil, d, err
	}

	handler := rpc.NewServer()
	apis := pls.APIs()

	for _, api := range apis {
		if api.Service == nil || api.Namespace != "eth" {
			log.Debug("InProc skipped to register service", "service", api.Service, "namespace", api.Namespace)
			continue
		}

		if err := handler.RegisterName(api.Namespace, api.Service); err != nil {
			return nil, nil, d, err
		}
		log.Debug("InProc registered", "service", api.Service, "namespace", api.Namespace)
	}

	return pls, handler, d, nil
}

func setNonce(opt *bind.TransactOpts, nonce *uint64) {
	opt.Nonce = big.NewInt(int64(*nonce))
	*nonce++
}

func deployEtherTokenInChildChain(t *testing.T) {
	opt := makeTxOpt(operatorKey, 0, nil, nil)

	setNonce(opt, &operatorNonceChildChain)
	etherTokenAddrInChildChain, _, etherTokenInChildChain, err = ethertoken.DeployEtherToken(
		opt,
		plsClient,
		development,
		common.Address{},
		swapEnabledInChildChain,
	)

	if err != nil {
		t.Fatal("Failed to deploy EtherToken in child chain", "err", err)
	}

	log.Info("EtherToken deployed in child chain", "addr", etherTokenAddrInChildChain)

	return
}

func deployTokenContracts(t *testing.T) (*token.RequestableSimpleToken, *token.RequestableSimpleToken, common.Address, common.Address) {
	setNonce(operatorOpt, &operatorNonceRootChain)
	tokenAddrInRootChain, tx, tokenInRootChain, err := token.DeployRequestableSimpleToken(
		operatorOpt,
		ethClient,
	)
	waitEthTx(tx.Hash(), "")
	if err != nil {
		t.Fatal("Failed to deploy token contract in root chain", "err", err)
	}
	log.Info("Token deployed in root chain", "address", tokenAddrInRootChain)

	setNonce(operatorOpt, &operatorNonceChildChain)
	tokenAddrInChildChain, tx, tokenInChildChain, err := token.DeployRequestableSimpleToken(
		operatorOpt,
		plsClient,
	)
	waitPlsTx(tx.Hash(), "deploy token contract")
	if err != nil {
		t.Fatal("Failed to deploy token contract in child chain", "err", err)
	}
	log.Info("Token deployed in child chain", "address", tokenAddrInChildChain)

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

	_, ks := tmpKeyStore()

	mux := new(event.TypeMux)
	epochEnv := epoch.New()
	miner := miner.New(minerBackend, &testPlsConfig.Miner, params.MainnetChainConfig, mux, engine, epochEnv, db, nil)

	account, err := ks.ImportECDSA(operatorKey, "")
	if err != nil {
		log.Error("Failed to import operator account", "err", err)
	}
	if err = ks.Unlock(account, ""); err != nil {
		log.Error("Failed to unlock operator account", "err", err)
	}
	// configure account manager with temporary keystore backend
	backends := []accounts.Backend{
		ks,
	}

	accConfig := accounts.Config{true}
	accManager := accounts.NewManager(&accConfig, backends...)
	txManager, err := tx.NewTransactionManager(ks, ethClient, db, &testPlsConfig.TxConfig)

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
		accManager,
		txManager,
		miner,
		epochEnv,
	)

	if err != nil {
		return nil, func() {}, err
	}
	// TODO (aiden): there's no need to start miner in here, it starts when rcm connect to root chain contract by reading 1st NRE.
	//go Miner.Start(operator, &miner.NRE)
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

func waitEthTx(hash common.Hash, caption string) error {
	log.Info("Waiting ethereum transaction mined...", "caption", caption, "hash", hash.String())

	var receipt *types.Receipt
	for receipt, _ = ethClient.TransactionReceipt(context.Background(), hash); receipt == nil; {
		<-time.NewTimer(testPlsConfig.TxConfig.Interval).C
		log.Error("Ethereum transaction is pending...", "caption", caption, "hash", hash.String())
		receipt, _ = ethClient.TransactionReceipt(context.Background(), hash)
	}
	log.Info("Ethereum transaction mined", "caption", caption, "blockNumber", receipt.BlockNumber, "hash", hash.String())

	if receipt.Status == 0 {
		log.Error("Ethereum transaction reverted", "caption", caption, "gasUsed", receipt.GasUsed, "hash", hash.String())
		return errors.New("Ethereum transaction reverted")
	}

	return nil
}

func waitPlsTx(hash common.Hash, caption string) error {
	log.Info("Waiting plasma transaction...", "hash", hash, "caption", caption)

	var receipt *types.Receipt
	for receipt, _ = plsClient.TransactionReceipt(context.Background(), hash); receipt == nil; {
		<-time.NewTimer(testPlsConfig.Miner.Recommit).C
		log.Error("Plasma transaction is pending...", "hash", hash, "caption", caption)
		receipt, _ = plsClient.TransactionReceipt(context.Background(), hash)
	}

	log.Info("Plasma transaction mined", "hash", hash, "blockNumber", receipt.BlockNumber, "caption", caption, "receipt", receipt)

	if receipt.Status == 0 {
		log.Error("Plasma transaction reverted", "hash", hash, "caption", caption)
		return errors.New("Plasma transaction reverted")
	}

	return nil
}

// TODO: any user sends tx
func makeSampleTx(rcm *RootChainManager) error {
	pool := rcm.txPool

	// self transfer
	var err error

	tx := types.NewTransaction(operatorNonceChildChain, operator, nil, 21000, nil, []byte{})
	operatorNonceChildChain++

	signer := types.NewEIP155Signer(params.MainnetChainConfig.ChainID)

	tx, err = types.SignTx(tx, signer, operatorKey)
	if err != nil {
		log.Error("Failed to sign sample tx", "err", err)
		return err
	}

	if err = pool.AddLocal(tx); err != nil {
		log.Error("Failed to insert sample tx to tx pool", "err", err)

		return err
	}

	log.Debug("Sample transaction is submitted in child chian")

	return nil
}

func checkMinedBlock(pls *Plasma, pbMinedEvents *event.TypeMuxSubscription, expectedIsRequest bool, expectedForkNumber, expectedBlockNumber int64) error {

	timer := time.NewTimer(time.Second * 15)
	defer timer.Stop()

	log.Warn("Check mined block", "expectedBlockNumber", expectedBlockNumber)

	var block *types.Block

	select {
	case e, ok := <-pbMinedEvents.Chan():
		if !ok {
			return errors.New("cannot read from mined block channel")
		}
		block = e.Data.(core.NewMinedBlockEvent).Block

	case <-timer.C:
		return errors.New("out of time")
	}

	log.Warn("Check Block Number", "expectedBlockNumber", expectedBlockNumber, "minedBlockNumber", block.NumberU64(), "forkNumber", block.Difficulty().Uint64())

	// check block number.
	if expectedBlockNumber != block.Number().Int64() {
		return errors.New(fmt.Sprintf("Expected block number: %d, actual block %d", expectedBlockNumber, block.Number().Int64()))
	}

	// check fork number
	if expectedForkNumber+1 != block.Difficulty().Int64() {
		return errors.New(fmt.Sprintf("Block Expected ForkNumber: %d, Actual ForkNumber %d", expectedForkNumber+1, block.Difficulty().Int64()))
	}

	// check isRequest.
	if block.IsRequest() != expectedIsRequest {
		return errors.New(fmt.Sprintf("Expected isRequest: %t, Actual isRequest %t", expectedIsRequest, block.IsRequest()))
	}

	return nil
}

func checkSubmittedBlock(pls *Plasma, pbSubmitedEvents chan *rootchain.RootChainBlockSubmitted, expectedForkNumber, expectedBlockNumber int64) error {
	setNonce(operatorOpt, &operatorNonceRootChain) // due to block submit

	timer := time.NewTimer(time.Second * 15)
	defer timer.Stop()

	log.Warn("Check mined block", "expectedBlockNumber", expectedBlockNumber)

	select {
	case _, ok := <-pbSubmitedEvents:
		if !ok {
			return errors.New("cannot read BlockSubmitted event")
		}

	case <-timer.C:
		return errors.New("out of time")
	}

	minedBlock := pls.blockchain.GetBlockByNumber(uint64(expectedBlockNumber))

	log.Warn("Check Submitted Block Number", "expectedBlockNumber", expectedBlockNumber, "minedBlockNumber", minedBlock.NumberU64(), "forkNumber", minedBlock.Difficulty().Uint64())

	pb, _ := pls.rootchainManager.getBlock(big.NewInt(int64(pls.rootchainManager.state.currentFork)), minedBlock.Number())

	if pb.Timestamp == 0 {
		log.Debug("Submitted plasma minedBlock", "pb", pb)
		log.Debug("Mined plasma minedBlock", "b", minedBlock)
		return errors.New("Plasma minedBlock is not submitted yet.")
	}

	if pb.IsRequest != minedBlock.IsRequest() {
		return errors.New(fmt.Sprintf("Block isRequest mismatch: (expected: %t, actual: %t)", pb.IsRequest, minedBlock.IsRequest()))
	}

	pbStateRoot := pb.StatesRoot[:]
	bStateRoot := minedBlock.Header().Root.Bytes()
	if bytes.Compare(pbStateRoot, bStateRoot) != 0 {
		return errors.New(fmt.Sprintf("Block state root mismatch: (expected: %s, actual: %s)", common.Bytes2Hex(pbStateRoot), common.Bytes2Hex(bStateRoot)))
	}

	pbTxRoot := pb.TransactionsRoot[:]
	bTxRoot := minedBlock.Header().TxHash.Bytes()
	if bytes.Compare(pbTxRoot, bTxRoot) != 0 {
		return errors.New(fmt.Sprintf("Block transactions root mismatch: (expected: %s, actual: %s)", common.Bytes2Hex(pbTxRoot), common.Bytes2Hex(bTxRoot)))
	}

	pbReceiptsRoot := pb.ReceiptsRoot[:]
	bReceiptsRoot := minedBlock.Header().ReceiptHash.Bytes()
	if bytes.Compare(pbReceiptsRoot, bReceiptsRoot) != 0 {
		return errors.New(fmt.Sprintf("Block receipts root mismatch: (expected: %s, actual: %s)", common.Bytes2Hex(pbReceiptsRoot), common.Bytes2Hex(bReceiptsRoot)))
	}
	return nil
}

func checkEpochAfter(pls *Plasma, epochPreparedEvents chan *rootchain.RootChainEpochPrepared, expectedIsEmpty bool, expectedIsRequest bool, expectedForkNumber int64, expectedEpochNumber int64) error {
	timer := time.NewTimer(time.Second * 15)
	defer timer.Stop()

	log.Warn("Check epoch after submission", "expectedEpochNumber", expectedEpochNumber)

	select {
	case e, ok := <-epochPreparedEvents:
		if !ok {
			return errors.New("cannot read EpochPrepared event")
		}

		log.Info("Next prepared epoch", "e", e)

	case <-timer.C:
		return errors.New("out of time")
	}

	epoch, err := pls.rootchainManager.getEpoch(big.NewInt(expectedForkNumber), big.NewInt(expectedEpochNumber))
	if err != nil {
		return err
	}

	log.Warn("Epoch?", "e", epoch)

	if epoch.Timestamp == 0 {
		return fmt.Errorf("Epoch#%d is not submitted", expectedEpochNumber)
	}

	if !epoch.Initialized {
		return fmt.Errorf("Epoch#%d is not initialized", expectedEpochNumber)
	}

	if epoch.IsRequest != expectedIsRequest {
		return fmt.Errorf("isRequest mismatch: (expected: %t, actual: %t)", expectedIsRequest, epoch.IsRequest)
	}

	if epoch.IsEmpty != expectedIsEmpty {
		return fmt.Errorf("IsEmpty mismatch: (expected: %t, actual: %t)", expectedIsEmpty, epoch.IsEmpty)
	}

	if !epoch.IsRequest {
		setNonce(operatorOpt, &operatorNonceRootChain) // due to block submit
	}

	if epoch.IsRequest {
		return nil
	}

	var blocks types.Blocks

	for i := epoch.StartBlockNumber; i <= epoch.EndBlockNumber; i++ {
		blocks = append(blocks, pls.blockchain.GetBlockByNumber(i))
	}

	expectedStatesRoot := blocks.StatesRoot().Bytes()[:]
	stateRoot := epoch.NRE.EpochStateRoot[:]
	if bytes.Compare(expectedStatesRoot, stateRoot) != 0 {
		return errors.New(fmt.Sprintf("Epoch states root mismatch (expected: %s actual: %s)", common.Bytes2Hex(expectedStatesRoot), common.Bytes2Hex(stateRoot)))
	}

	expectedTxRoot := blocks.TransactionsRoot().Bytes()[:]
	txRoot := epoch.NRE.EpochTransactionsRoot[:]
	if bytes.Compare(expectedTxRoot, txRoot) != 0 {
		return errors.New(fmt.Sprintf("Epoch transactions root mismatch (expected: %s actual: %s)", common.Bytes2Hex(expectedTxRoot), common.Bytes2Hex(txRoot)))
	}

	expectedReceiptsRoot := blocks.ReceiptssRoot().Bytes()[:]
	receiptsRoot := epoch.NRE.EpochReceiptsRoot[:]
	if bytes.Compare(expectedReceiptsRoot, receiptsRoot) != 0 {
		return errors.New(fmt.Sprintf("Epoch receipts root mismatch (expected: %s actual: %s)", common.Bytes2Hex(expectedReceiptsRoot), common.Bytes2Hex(receiptsRoot)))
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
			return errors.New(fmt.Sprintf(caption+"\t : Expected %s (after) == %s (before) + %s (diff), but it isn't", at, bt, dt))
		}
		return nil
	}

	target := new(big.Int).Sub(new(big.Int).Sub(a, b), diff)
	e := new(big.Int).Abs(offset)
	s := new(big.Int).Neg(e)

	// out of range
	if s.Cmp(target) > 0 || target.Cmp(e) > 0 {
		st := toString(truncate(s))
		et := toString(truncate(e))
		targett := toString(truncate(target))

		wait(1)
		return errors.New(fmt.Sprintf(caption+"\t : Expected %s (-offset) < %s (after - before - diff) < %s (+offset), but it isn't", st, targett, et))
	}

	return nil
}

func getETHBalances(addrs []common.Address) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = etherToken.BalanceOf(baseCallOpt, addr)
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

func getEtherTokenBalances(addrs []common.Address) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = etherToken.BalanceOf(baseCallOpt, addr)
	}

	return balances
}

func getPEtherTokenBalances(addrs []common.Address) []*big.Int {
	balances := make([]*big.Int, len(addrs))

	for i, addr := range addrs {
		// balances[i] would be nil if ethClient.BalanceAt fails
		balances[i], _ = etherToken.BalanceOf(baseCallOpt, addr)
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
