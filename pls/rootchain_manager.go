package pls

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/miner"
	"github.com/Onther-Tech/plasma-evm/miner/epoch"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/tx"
)

const (
	MAX_EPOCH_EVENTS = 0
)

var (
	baseCallOpt               = &bind.CallOpts{Pending: false, Context: context.Background()}
	requestableContractABI, _ = abi.JSON(strings.NewReader(rootchain.RequestableIABI))
	rootchainContractABI, _   = abi.JSON(strings.NewReader(rootchain.RootChainABI))

	ErrKnownTransaction = errors.New("known transaction")
)

type invalidExit struct {
	forkNumber  *big.Int
	blockNumber *big.Int
	receipt     *types.Receipt
	index       int64
	proof       []common.Hash
}

type invalidExits []*invalidExit

type RootChainManager struct {
	config *Config
	stopFn func()

	txPool     *core.TxPool
	blockchain *core.BlockChain

	backend           *ethclient.Client
	rootchainContract *rootchain.RootChain

	eventMux       *event.TypeMux
	accountManager *accounts.Manager
	txManager      *tx.TransactionManager

	miner    *miner.Miner
	minerEnv *epoch.EpochEnvironment
	state    *rootchainState

	// fork => block number => invalidExits
	invalidExits map[uint64]map[uint64]invalidExits

	// channels
	quit             chan struct{}
	epochPreparedCh  chan *rootchain.RootChainEpochPrepared
	blockFinalizedCh chan *rootchain.RootChainBlockFinalized

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and etherbase)
}

func (rcm *RootChainManager) RootchainContract() *rootchain.RootChain { return rcm.rootchainContract }
func (rcm *RootChainManager) NRELength() (*big.Int, error) {
	return rcm.rootchainContract.NRELength(baseCallOpt)
}

func NewRootChainManager(
	config *Config,
	stopFn func(),
	txPool *core.TxPool,
	blockchain *core.BlockChain,
	backend *ethclient.Client,
	rootchainContract *rootchain.RootChain,
	eventMux *event.TypeMux,
	accountManager *accounts.Manager,
	txManager *tx.TransactionManager,
	miner *miner.Miner,
	env *epoch.EpochEnvironment,
) (*RootChainManager, error) {
	code, err := backend.CodeAt(context.Background(), config.RootChainContract, nil)
	if err != nil {
		return nil, err
	}

	n := 10
	for i := 0; i <= n; i++ {
		if len(code) > 0 {
			break
		}

		<-time.NewTimer(time.Second).C
		code, err = backend.CodeAt(context.Background(), config.RootChainContract, nil)
		if err != nil {
			return nil, err
		}

		if len(code) == 0 && i == n {
			return nil, errors.New(fmt.Sprintf("RootChain contract is not deployed at %s", config.RootChainContract.Hex()))
		}
	}

	rcm := &RootChainManager{
		config:            config,
		stopFn:            stopFn,
		txPool:            txPool,
		blockchain:        blockchain,
		backend:           backend,
		rootchainContract: rootchainContract,
		eventMux:          eventMux,
		accountManager:    accountManager,
		txManager:         txManager,
		miner:             miner,
		minerEnv:          env,
		invalidExits:      make(map[uint64]map[uint64]invalidExits),
		quit:              make(chan struct{}),
		epochPreparedCh:   make(chan *rootchain.RootChainEpochPrepared, MAX_EPOCH_EVENTS),
		blockFinalizedCh:  make(chan *rootchain.RootChainBlockFinalized),
	}

	rcm.state = newRootchainState(rcm)

	epochLength, err := rcm.NRELength()
	if err != nil {
		return nil, err
	}

	if config.NodeMode == ModeOperator {
		miner.SetNRBepochLength(epochLength)
	}

	return rcm, nil
}

func (rcm *RootChainManager) Start() error {
	if err := rcm.run(); err != nil {
		return err
	}

	go rcm.pingBackend()
	rcm.txManager.Start()

	if rcm.config.NodeMode == ModeOperator {
		go rcm.miner.Start(rcm.config.Operator.Address, new(rootchain.RootChainEpochPrepared), true)
	}

	return nil
}

func (rcm *RootChainManager) Stop() error {
	rcm.backend.Close()
	rcm.txManager.Stop()
	close(rcm.quit)
	return nil
}

func (rcm *RootChainManager) run() error {
	go rcm.runHandlers()
	go rcm.runSubmitter()
	go rcm.runDetector()

	if err := rcm.watchEvents(); err != nil {
		return err
	}

	return nil
}

// watchEvents watchs RootChain contract events
func (rcm *RootChainManager) watchEvents() error {
	closed := false

	filterer, err := rootchain.NewRootChainFilterer(rcm.config.RootChainContract, rcm.backend)
	if err != nil {
		return err
	}

	startBlockNumber := rcm.blockchain.GetRootchainBlockNumber()
	filterOpts := &bind.FilterOpts{
		Start:   startBlockNumber,
		End:     nil,
		Context: context.Background(),
	}

	// TODO: have to read only NRE1
	// iterate to find previous epoch prepared events
	iteratorForEpochPreparedEvent, err := filterer.FilterEpochPrepared(filterOpts)
	if err != nil {
		return err
	}

	log.Info("Iterating epoch prepared event")
	for iteratorForEpochPreparedEvent.Next() {
		e := iteratorForEpochPreparedEvent.Event
		if e != nil {
			if err := rcm.handleEpochPrepared(e); err != nil {
				log.Error("Failed to handle past EpochPrepared events", "err", err)
			}
		}
	}

	// iterate to find previous block finalized events
	iteratorForBlockFinalizedEvent, err := filterer.FilterBlockFinalized(filterOpts)
	if err != nil {
		return err
	}

	log.Info("Iterating block finalized event")
	for iteratorForBlockFinalizedEvent.Next() {
		e := iteratorForBlockFinalizedEvent.Event
		if e != nil {
			if err := rcm.handleBlockFinalized(e); err != nil {
				log.Error("Failed to handle past BlockFinalized events", "err", err)
			}
		}
	}

	watchOpts := &bind.WatchOpts{
		Context: context.Background(),
		Start:   &startBlockNumber,
	}
	epochPrepareWatchCh := make(chan *rootchain.RootChainEpochPrepared)
	blockFinalizedWatchCh := make(chan *rootchain.RootChainBlockFinalized)

	log.Info("Watching epoch prepared event", "startBlockNumber", startBlockNumber)
	epochPrepareSub, err := filterer.WatchEpochPrepared(watchOpts, epochPrepareWatchCh)
	if err != nil {
		return err
	}

	log.Info("Watching block finalized event", "startBlockNumber", startBlockNumber)
	blockFinalizedSub, err := filterer.WatchBlockFinalized(watchOpts, blockFinalizedWatchCh)
	if err != nil {
		return err
	}

	resubTimer := time.NewTimer(0)
	<-resubTimer.C
	resub := func() {
		if closed {
			return
		}

		startBlockNumber := rcm.blockchain.GetRootchainBlockNumber() + 1
		watchOpts := &bind.WatchOpts{
			Context: context.Background(),
			Start:   &startBlockNumber,
		}

		log.Info("Re-subsribe EpochPrepared event", "startBlockNumber", startBlockNumber)
		epochPrepareSub2, err := filterer.WatchEpochPrepared(watchOpts, epochPrepareWatchCh)
		if err != nil {
			log.Error("Failed to re-subscribe event", "err", err)
			resubTimer.Reset(5 * time.Second)
			return
		}
		epochPrepareSub = epochPrepareSub2

		log.Info("Watching block finalized event", "startBlockNumber", startBlockNumber)
		blockFinalizedSub2, err := filterer.WatchBlockFinalized(watchOpts, blockFinalizedWatchCh)
		if err != nil {
			log.Error("Failed to re-subscribe event", "err", err)
			resubTimer.Reset(5 * time.Second)
			return
		}

		blockFinalizedSub = blockFinalizedSub2
	}

	// TODO: wait untli previous submit transaction is mined.

	go func() {
		for {
			select {
			case e := <-epochPrepareWatchCh:
				if e != nil {
					rcm.epochPreparedCh <- e
				}

			case err := <-epochPrepareSub.Err():
				if err != nil {
					log.Error("Epoch prepared event subscription error", "err", err)
					resub()
				}

			case e := <-blockFinalizedWatchCh:
				if e != nil {
					rcm.blockFinalizedCh <- e
				}

			case err := <-blockFinalizedSub.Err():
				if err != nil {
					log.Error("Block finalized event subscription error", "err", err)
					resub()
				}
				return

			case <-rcm.quit:
				closed = true
				return
			}
		}
	}()

	return nil
}

func makePos(v1 *big.Int, v2 *big.Int) *big.Int {
	a := new(big.Int).Mul(
		v1,
		new(big.Int).Exp(
			big.NewInt(2),
			big.NewInt(128),
			nil,
		),
	)

	return new(big.Int).Add(a, v2)
}

func (rcm *RootChainManager) addEpochSubmitTransaction(blocks types.Blocks) error {
	if rcm.config.NodeMode != ModeOperator {
		return errors.New("only operator node can add submit transaction")
	} else if len(blocks) == 0 {
		return errors.New("empty epoch error.")
	}

	operator := rcm.config.Operator
	funcName := "submitNRE"

	forkNumber := new(big.Int).Set(rcm.minerEnv.CurrentFork)
	epochNumber := new(big.Int).Set(rcm.minerEnv.EpochNumber)

	// pos1 = fork number * 2^128 + epoch number
	pos1 := makePos(forkNumber, epochNumber)

	startBlockNumber := blocks[0].Number()
	endBlockNumber := blocks[len(blocks)-1].Number()

	// pos2 = start block number * 2^128 + end block number
	pos2 := makePos(startBlockNumber, endBlockNumber)

	input, err := rootchainContractABI.Pack(
		funcName,
		pos1,
		pos2,
		blocks.StatesRoot(),
		blocks.TransactionsRoot(),
		blocks.ReceiptsRoot(),
	)

	if err != nil {
		return err
	}

	caption := fmt.Sprintf("%s(%d: [%d-%d])", funcName, epochNumber.Uint64(), startBlockNumber.Uint64(), endBlockNumber.Uint64())
	rawTx := tx.NewRawTransaction(operator.Address, params.SubmitBlockGasLimit, &rcm.config.RootChainContract, big.NewInt(int64(rcm.state.costNRB)), input, false, caption)

	return rcm.txManager.Add(operator, rawTx, false)
}

func (rcm *RootChainManager) addBlockSubmitTransaction(block *types.Block) error {
	if rcm.config.NodeMode != ModeOperator {
		return errors.New("only operator node can add submit transaction")
	}

	operator := rcm.config.Operator
	funcName := "submitORB"

	forkNumber := new(big.Int).Set(rcm.minerEnv.CurrentFork)

	pos := makePos(forkNumber, block.Number())

	input, err := rootchainContractABI.Pack(
		funcName,
		pos,
		block.Header().Root,
		block.Header().TxHash,
		block.Header().ReceiptHash,
	)

	if err != nil {
		return err
	}

	caption := fmt.Sprintf("%s(%d)", funcName, block.NumberU64())
	rawTx := tx.NewRawTransaction(operator.Address, params.SubmitBlockGasLimit, &rcm.config.RootChainContract, big.NewInt(int64(rcm.state.costNRB)), input, false, caption)

	return rcm.txManager.Add(operator, rawTx, false)
}

func (rcm *RootChainManager) runSubmitter() {
	if rcm.config.NodeMode != ModeOperator {
		return
	}

	plasmaBlockMinedEvents := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	submitBlockOrEpoch := func(block *types.Block) error {
		rcm.lock.Lock()
		rcm.minerEnv.Lock()

		defer rcm.lock.Unlock()
		defer rcm.minerEnv.Unlock()

		log.Info("New block is mined", "number", block.Number())

		// if the epoch is completed, stop mining operation and wait next epoch
		if rcm.minerEnv.Completed {
			rcm.miner.Stop()
		}

		bal, err := rcm.backend.BalanceAt(context.Background(), rcm.config.Operator.Address, nil)
		if err != nil {
			log.Error("Failed to get balance of operator account from rootchain", "err", err)
		}

		if bal.Cmp(rcm.config.OperatorMinEther) < 0 {
			log.Warn("Operator account balance on rootchain is too low")
		}

		if rcm.minerEnv.IsRequest {
			err = rcm.addBlockSubmitTransaction(block)
		} else if !rcm.minerEnv.IsRequest && rcm.minerEnv.Completed {
			var blocks types.Blocks

			st := time.Now()
			s := rcm.minerEnv.StartBlockNumber.Uint64()
			e := rcm.minerEnv.EndBlockNumber.Uint64()
			for i := s; i <= e; i++ {
				blocks = append(blocks, rcm.blockchain.GetBlockByNumber(i))
			}
			elapsed := time.Since(st)
			log.Debug("Read blocks for NRE", "epochNumber", rcm.minerEnv.EpochNumber, "numBlocks", e-s+1, "elapsed", elapsed)

			err = rcm.addEpochSubmitTransaction(blocks)
		} else {
			log.Info("Non-request epoch is not completed yet", "epochNumber", rcm.minerEnv.EpochNumber)
			return nil
		}

		if err == tx.ErrDuplicateRaw {
			log.Error("Same block submit transaction was included.")
			return nil
		} else if err != nil {
			return err
		}

		return nil
	}

	for {
		select {
		case ev, ok := <-plasmaBlockMinedEvents.Chan():
			if !ok {
				continue
			}

			if err := submitBlockOrEpoch(ev.Data.(core.NewMinedBlockEvent).Block); err != nil {
				log.Error("Failed to submit block", "err", err)
				rcm.stopFn()
			}

		case <-rcm.quit:
			return
		}
	}
}

func (rcm *RootChainManager) runHandlers() {
	if rcm.config.NodeMode != ModeOperator {
		return
	}

	for {
		select {
		case e := <-rcm.epochPreparedCh:
			if err := rcm.handleEpochPrepared(e); err != nil {
				log.Error("Failed to handle epoch prepared", "err", err)
			} else {
				rcm.blockchain.SetRootchainBlockNumber(e.Raw.BlockNumber)
			}
		case e := <-rcm.blockFinalizedCh:
			if err := rcm.handleBlockFinalized(e); err != nil {
				log.Error("Failed to handle block finazlied", "err", err)
			} else {
				rcm.blockchain.SetRootchainBlockNumber(e.Raw.BlockNumber)
			}
		case <-rcm.quit:
			return
		}
	}
}

// handleEpochPrepared handles EpochPrepared event from RootChain contract after
// plasma chain is SYNCED.
func (rcm *RootChainManager) handleEpochPrepared(ev *rootchain.RootChainEpochPrepared) error {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()

	// Short circuit if event is removed.
	if ev.Raw.Removed {
		return errors.New(fmt.Sprintf("EpochPrepared#%s event is removed. root chain would had been reorganized.", ev.EpochNumber.String()))
	}

	// Short circuit if epoch prepared event is fired due to reorg.
	if rcm.minerEnv.EpochNumber.Cmp(ev.EpochNumber) >= 0 {
		return errors.New(fmt.Sprintf("Epoch#%s is less than current epoch#%s.", ev.EpochNumber.String(), rcm.minerEnv.EpochNumber.String()))
	}

	e := *ev

	length := new(big.Int).Add(new(big.Int).Sub(e.EndBlockNumber, e.StartBlockNumber), big.NewInt(1))

	log.Info("RootChain epoch prepared",
		"epochNumber", e.EpochNumber,
		"startBlockNumber", e.StartBlockNumber,
		"endBlockNumber", e.EndBlockNumber,
		"epochLength", length,
		"isRequest", e.IsRequest,
		"userActivated", e.UserActivated,
		"isEmpty", e.EpochIsEmpty,
		"ForkNumber", e.ForkNumber,
		"isRebase", e.Rebase,
	)

	if e.EpochIsEmpty {
		log.Info("epoch is empty, jump to next epoch")
		return nil
	}

	if rcm.config.NodeMode == ModeOperator {
		go rcm.miner.Start(rcm.config.Operator.Address, &e, false)
	}

	// prepare request tx for ORBs
	if e.IsRequest && !e.EpochIsEmpty {
		events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
		defer events.Unsubscribe()

		numORBs := new(big.Int).Sub(e.EndBlockNumber, e.StartBlockNumber)
		numORBs = new(big.Int).Add(numORBs, big.NewInt(1))

		bodies := make([]types.Transactions, 0, numORBs.Uint64()) // [][]types.Transaction

		currentFork := big.NewInt(int64(rcm.state.currentFork))
		epoch, err := rcm.getEpoch(currentFork, e.EpochNumber)
		if err != nil {
			return err
		}
		log.Debug("rcm.getEpoch", "epoch", epoch)

		// TODO: URE, ORE' should handle requestBlockId in a different way.
		requestBlockId := big.NewInt(int64(epoch.RE.FirstRequestBlockId))

		log.Debug("Num Orbs", "epochNumber", e.EpochNumber, "numORBs", numORBs, "requestBlockId", requestBlockId, "e.EndBlockNumber", e.EndBlockNumber, "e.StartBlockNumber", e.StartBlockNumber)
		for blockNumber := e.StartBlockNumber; blockNumber.Cmp(e.EndBlockNumber) <= 0; {
			begin := time.Now()

			orb, err := rcm.rootchainContract.ORBs(baseCallOpt, requestBlockId)
			if err != nil {
				return err
			}

			numRequests := orb.RequestEnd - orb.RequestStart + 1
			log.Debug("Fetching ORB", "blockNumber", blockNumber, "requestStart", orb.RequestStart, "requestEnd", orb.RequestEnd, "requestBlockId", requestBlockId)

			body := make(types.Transactions, 0, numRequests)

			for requestId := orb.RequestStart; requestId <= orb.RequestEnd; {
				request, err := rcm.rootchainContract.EROs(baseCallOpt, big.NewInt(int64(requestId)))
				if err != nil {
					return err
				}

				log.Debug("Request fetched", "requestId", requestId, "hash", common.Bytes2Hex(request.Hash[:]), "request", request)

				var (
					to    common.Address
					value *big.Int
					input []byte
				)

				if request.IsTransfer && !request.IsExit {
					to = request.Requestor
					value = new(big.Int).SetBytes(request.TrieValue[:])
				} else {
					to, _ = rcm.rootchainContract.RequestableContracts(baseCallOpt, request.To)
					value = request.Value
					input, err = requestableContractABI.Pack("applyRequestInChildChain",
						request.IsExit,
						big.NewInt(int64(requestId)),
						request.Requestor,
						request.TrieKey,
						request.TrieValue,
					)
					if err != nil {
						log.Error("Failed to pack applyRequestInChildChain", "err", err)
					}

					log.Debug("Request tx.data", "payload", common.Bytes2Hex(input))
				}

				requestTx := types.NewTransaction(0, to, value, params.RequestTxGasLimit, params.RequestTxGasPrice, input)

				log.Debug("Request Transaction", "tx", requestTx)

				body = append(body, requestTx)
				requestId += 1
			}

			log.Info("Request txs fetched", "blockNumber", blockNumber, "requestBlockId", requestBlockId, "numRequests", len(body), "elapsed", time.Since(begin))

			bodies = append(bodies, body)

			blockNumber = new(big.Int).Add(blockNumber, big.NewInt(1))
			requestBlockId = new(big.Int).Add(requestBlockId, big.NewInt(1))
		}

		var numMinedORBs uint64 = 0

		// TODO: handle ModeUser
		if rcm.config.NodeMode != ModeOperator {
			return nil
		}

		// Unlock mutex and make submit loop to process
		rcm.lock.Unlock()
		for numMinedORBs < numORBs.Uint64() {
			if err := rcm.txPool.EnqueueReqeustTxs(bodies[numMinedORBs]); err != nil {
				return err
			}

			log.Info("Waiting new request block mined event...")

			e := <-events.Chan()
			block := e.Data.(core.NewMinedBlockEvent).Block

			log.Info("New request block is mined", "blockNumber", block.Number(), "txs", block.Transactions().Len())

			if !block.IsRequest() {
				return errors.New("Invalid request block type.")
			}

			receipts := rcm.blockchain.GetReceiptsByHash(block.Hash())

			for _, receipt := range receipts {
				if receipt.Status == 0 {
					log.Error("Request transaction is reverted", "blockNumber", block.Number(), "hash", receipt.TxHash)
				}
			}

			numMinedORBs += 1
		}

		// Re-lock
		rcm.lock.Lock()
	}

	return nil
}

// Challenge on invalid exits
func (rcm *RootChainManager) handleBlockFinalized(ev *rootchain.RootChainBlockFinalized) error {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()

	e := *ev

	log.Info("RootChain block finalized", "forkNumber", e.ForkNumber, "blockNubmer", e.BlockNumber)

	callerOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}

	w, err := rcm.accountManager.Find(rcm.config.Challenger)
	if err != nil {
		log.Error("Failed to get challenger account", "err", err)
		return err
	}

	block, err := rcm.rootchainContract.GetBlock(callerOpts, e.ForkNumber, e.BlockNumber)
	if err != nil {
		return err
	}

	if block.IsRequest {
		invalidExits := rcm.invalidExits[e.ForkNumber.Uint64()][e.BlockNumber.Uint64()]
		for i := 0; i < len(invalidExits); i++ {

			var proofs []byte
			for j := 0; j < len(invalidExits[i].proof); j++ {
				proof := invalidExits[i].proof[j].Bytes()
				proofs = append(proofs, proof...)
			}
			// TODO: ChallengeExit receipt check
			input, err := rootchainContractABI.Pack("challengeExit", e.ForkNumber, e.BlockNumber, big.NewInt(invalidExits[i].index), invalidExits[i].receipt.GetRlp(), proofs)
			if err != nil {
				log.Error("Failed to pack challengeExit", "err", err)
			}

			nonce, err := rcm.backend.NonceAt(context.Background(), rcm.config.Challenger.Address, nil)
			if err != nil {
				log.Error("Failed to get challenger nonce", "err", err)
			}

			// TODO: use tx.TransactionManager
			challengeTx := types.NewTransaction(uint64(nonce), rcm.config.RootChainContract, big.NewInt(0), params.SubmitBlockGasLimit, params.SubmitBlockGasPrice, input)

			signedTx, err := w.SignTx(rcm.config.Challenger, challengeTx, big.NewInt(int64(rcm.config.RootChainNetworkID)))
			if err != nil {
				log.Error("Failed to sign challengeTx", "err", err)
			}

			err = rcm.backend.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Error("Failed to send challengeTx", "err", err)
			} else {
				log.Info("challengeExit is submitted", "exit request number", invalidExits[i].index, "hash", signedTx.Hash().Hex())
			}
		}
	}

	return nil
}

func (rcm *RootChainManager) runDetector() {
	if rcm.config.NodeMode == ModeUser {
		return
	}

	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	callerOpts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}

	for {
		select {
		case ev, ok := <-events.Chan():
			if !ok {
				continue
			}
			rcm.lock.Lock()

			block := ev.Data.(core.NewMinedBlockEvent).Block

			if block.IsRequest() {
				var invalidExitsList invalidExits

				forkNumber, err := rcm.rootchainContract.CurrentFork(callerOpts)
				if err != nil {
					log.Warn("failed to get current fork number", "err", err)
					rcm.lock.Unlock()
					continue
				}

				// TODO: read and write to DB
				if rcm.invalidExits[forkNumber.Uint64()] == nil {
					rcm.invalidExits[forkNumber.Uint64()] = make(map[uint64]invalidExits)
				}

				receipts := rcm.blockchain.GetReceiptsByHash(block.Hash())

				// TODO: should check if the request[i] is enter or exit request. Undo request will make posterior enter request.
				for i := 0; i < len(receipts); i++ {
					if receipts[i].Status == types.ReceiptStatusFailed {
						invalidExit := &invalidExit{
							forkNumber:  forkNumber,
							blockNumber: block.Number(),
							receipt:     receipts[i],
							index:       int64(i),
							proof:       types.GetMerkleProof(receipts, i),
						}
						invalidExitsList = append(invalidExitsList, invalidExit)

						log.Info("Invalid Exit Detected", "invalidExit", invalidExit, "forkNumber", forkNumber, "blockNumber", block.Number())
					}
				}
				// TODO: read and write to DB
				rcm.invalidExits[forkNumber.Uint64()][block.NumberU64()] = invalidExitsList
			}
			rcm.lock.Unlock()

		case <-rcm.quit:
			return
		}
	}
}

func (rcm *RootChainManager) getEpoch(forkNumber, epochNumber *big.Int) (rootchain.DataEpoch, error) {
	return rcm.rootchainContract.GetEpoch(baseCallOpt, forkNumber, epochNumber)
}
func (rcm *RootChainManager) getBlock(forkNumber, blockNumber *big.Int) (rootchain.DataPlasmaBlock, error) {
	return rcm.rootchainContract.GetBlock(baseCallOpt, forkNumber, blockNumber)
}
func (rcm *RootChainManager) lastBlock(forkNumber *big.Int, pending bool) (*big.Int, error) {
	baseCallOpt := &bind.CallOpts{Pending: pending, Context: context.Background()}

	num, err := rcm.rootchainContract.LastBlock(baseCallOpt, forkNumber)
	if err != nil {
		return nil, err
	}
	return num, nil
}

// pingBackend checks rootchain backend is alive.
func (rcm *RootChainManager) pingBackend() {
	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-ticker.C:
			if _, err := rcm.backend.SyncProgress(context.Background()); err != nil {
				log.Error("Rootchain provider doesn't respond", "err", err)
			}
		case <-rcm.quit:
			ticker.Stop()
			return
		}
	}
}
