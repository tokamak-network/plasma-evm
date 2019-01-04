package pls

import (
	"bytes"
	"context"
	"errors"
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
	"github.com/Onther-Tech/plasma-evm/params"
)

const MAX_EPOCH_EVENTS = 0

var (
	baseCallOpt               = &bind.CallOpts{Pending: false, Context: context.Background()}
	requestableContractABI, _ = abi.JSON(strings.NewReader(rootchain.RequestableContractIABI))
	rootchainContractABI, _   = abi.JSON(strings.NewReader(rootchain.RootChainABI))

	//TODO (aiden): sholud delete this after fixing rcm.backend.NetworkId
	rootchainNetworkId = big.NewInt(1337)
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

	miner    *miner.Miner
	minerEnv *miner.EpochEnvironment
	state    *rootchainState

	lastFinalizedBlock *big.Int
	currentFork        *big.Int

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
	miner *miner.Miner,
	env *miner.EpochEnvironment,
) (*RootChainManager, error) {
	rcm := &RootChainManager{
		config:             config,
		stopFn:             stopFn,
		txPool:             txPool,
		blockchain:         blockchain,
		backend:            backend,
		rootchainContract:  rootchainContract,
		eventMux:           eventMux,
		accountManager:     accountManager,
		miner:              miner,
		minerEnv:           env,
		lastFinalizedBlock: big.NewInt(0),
		currentFork:        big.NewInt(0),
		invalidExits:       make(map[uint64]map[uint64]invalidExits),
		quit:               make(chan struct{}),
		epochPreparedCh:    make(chan *rootchain.RootChainEpochPrepared, MAX_EPOCH_EVENTS),
		blockFinalizedCh:   make(chan *rootchain.RootChainBlockFinalized),
	}

	rcm.state = newRootchainState(rcm)

	epochLength, err := rcm.NRELength()
	if err != nil {
		return nil, err
	}

	miner.SetNRBepochLength(epochLength)

	return rcm, nil
}

func (rcm *RootChainManager) Start() error {
	if err := rcm.run(); err != nil {
		return err
	}

	go rcm.pingBackend()

	return nil
}

func (rcm *RootChainManager) Stop() error {
	rcm.backend.Close()
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
	filterer, err := rootchain.NewRootChainFilterer(rcm.config.RootChainContract, rcm.backend)
	if err != nil {
		return err
	}

	// rootchain block#1
	startBlockNumber := uint64(1)

	filterOpts := &bind.FilterOpts{
		Start:   startBlockNumber,
		End:     nil,
		Context: context.Background(),
	}

	// iterate previous events
	// TODO: the events fired while syncing should be dealt with in different way.
	iterator, err := filterer.FilterEpochPrepared(filterOpts)
	if err != nil {
		return err
	}

	log.Info("Iterating EpochPrepared event")

	for iterator.Next() {
		e := iterator.Event
		if e != nil {
			rcm.handleEpochPrepared(e)
		}
	}

	iterator2, err := filterer.FilterBlockFinalized(filterOpts)
	if err != nil {
		return err
	}

	log.Info("Iterating BlockFinalized event")

	for iterator2.Next() {
		e := iterator2.Event
		if e != nil {
			rcm.handleBlockFinalzied(e)
		}
	}

	// watch events from now
	watchOpts := &bind.WatchOpts{
		Context: context.Background(),
		Start:   &startBlockNumber, // read events from rootchain block#1
	}

	epochPrepareWatchCh := make(chan *rootchain.RootChainEpochPrepared)
	epochPrepareSub, err := filterer.WatchEpochPrepared(watchOpts, epochPrepareWatchCh)
	if err != nil {
		return err
	}

	log.Info("Watching EpochPrepared event", "startBlockNumber", startBlockNumber)

	blockFinalizedWatchCh := make(chan *rootchain.RootChainBlockFinalized)
	blockFinalizedSub, err := filterer.WatchBlockFinalized(watchOpts, blockFinalizedWatchCh)
	if err != nil {
		return err
	}

	log.Info("watching BlockFinalized event", "startBlockNumber", startBlockNumber)

	go func() {
		for {
			select {
			case e := <-epochPrepareWatchCh:
				if e != nil {
					rcm.epochPreparedCh <- e
				}

			case err := <-epochPrepareSub.Err():
				log.Error("EpochPrepared event subscription error", "err", err)
				rcm.stopFn()
				return

			case e := <-blockFinalizedWatchCh:
				if e != nil {
					rcm.blockFinalizedCh <- e
				}

			case err := <-blockFinalizedSub.Err():
				log.Error("BlockFinalized event subscription error", "err", err)
				rcm.stopFn()
				return

			case <-rcm.quit:
				return
			}
		}
	}()

	return nil
}

// TODO (aiden): should not submit block mined just now when the URB epoch event is received. It's considered to be canceled.
func (rcm *RootChainManager) runSubmitter() {
	plasmaBlockMinedEvents := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer plasmaBlockMinedEvents.Unsubscribe()

	blockSubmitEvents := make(chan *rootchain.RootChainBlockSubmitted)
	blockSubmitWatchOpts := &bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	}
	blockFilterer, _ := rcm.rootchainContract.WatchBlockSubmitted(blockSubmitWatchOpts, blockSubmitEvents)
	defer blockFilterer.Unsubscribe()

	w, err := rcm.accountManager.Find(rcm.config.Operator)
	if err != nil {
		log.Error("Failed to get operator wallet", "err", err)
		return
	}

	for {
		select {
		case ev := <-plasmaBlockMinedEvents.Chan():
			if ev == nil {
				return
			}
			// if the epoch is completed, stop mining operation and wait next epoch
			if rcm.minerEnv.Completed {
				rcm.miner.Stop()
			}

			rcm.lock.Lock()

			blockInfo := ev.Data.(core.NewMinedBlockEvent)

			if blockInfo.IsURB != rcm.minerEnv.IsUserActivated {
				return
			}

			Nonce := rcm.state.getNonce()

			//TODO (aiden): rcm.backend.NetworkID does not work as intended. It should return 1337, not 1. And it should moved to rcm.config.RootchainNetworkId.
			networkID, err := rcm.backend.NetworkID(context.Background())
			log.Info("network Id", "id", networkID)
			if err != nil {
				log.Error("NetworkId error", "err", err)
			}

			funcName := "submitNRB"
			if rcm.minerEnv.IsRequest {
				funcName = "submitORB"
			}

			input, err := rootchainContractABI.Pack(
				funcName,
				big.NewInt(int64(rcm.state.currentFork)),
				blockInfo.Block.Header().Root,
				blockInfo.Block.Header().TxHash,
				blockInfo.Block.Header().ReceiptHash,
			)

			if err != nil {
				log.Error("Failed to pack "+funcName, "err", err)
			}
			submitTx := types.NewTransaction(Nonce, rcm.config.RootChainContract, big.NewInt(int64(rcm.state.costNRB)), params.SubmitBlockGasLimit, params.SubmitBlockGasPrice, input)

			signedTx, err := w.SignTx(rcm.config.Operator, submitTx, rootchainNetworkId)
			if err != nil {
				log.Error("Failed to sign "+funcName, "err", err)
			}

			err = rcm.backend.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Error("Failed to send "+funcName, "err", err)
			}

			// wait root chain block is mined
			<-blockSubmitEvents

			receipt, err := rcm.backend.TransactionReceipt(context.Background(), signedTx.Hash())
			log.Debug("signed tx receipt", "receipt", receipt, "hash", signedTx.Hash().String())

			if err != nil {
				log.Error("Failed to send "+funcName, "err", err)
			} else if receipt.Status == 0 {
				log.Error(funcName+" is reverted", "hash", signedTx.Hash().Hex())
			} else {
				log.Info("Block is submitted", "funcName", funcName, "blockNumber", blockInfo.Block.NumberU64(), "hash", signedTx.Hash().String())
			}

			rcm.state.incNonce()

			rcm.lock.Unlock()
		case <-rcm.quit:
			return
		}
	}
}

func (rcm *RootChainManager) runHandlers() {
	for {
		select {
		case e := <-rcm.epochPreparedCh:
			if err := rcm.handleEpochPrepared(e); err != nil {
				log.Error("Failed to handle epoch prepared", "err", err)
			}
		case e := <-rcm.blockFinalizedCh:
			if err := rcm.handleBlockFinalzied(e); err != nil {
				log.Error("Failed to handle block finazlied", "err", err)
			}
		case <-rcm.quit:
			return
		}
	}
}

// handleEpochPrepared handles EpochPrepared event from RootChain contract after
// plasma chain is *SYNCED*.
// TODO (aiden): change call method(rcm.rootchaincontract.[someMethod(baseCallOpt, ...]) according to the change in the rootchain contract.
// TODO (aiden): merge ORB, rebased ORB and URB logic if possible
func (rcm *RootChainManager) handleEpochPrepared(ev *rootchain.RootChainEpochPrepared) error {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()

	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	e := *ev

	log.Info("RootChain epoch prepared", "epochNumber", e.EpochNumber, "isRequest", e.IsRequest, "userActivated", e.UserActivated, "isEmpty", e.EpochIsEmpty)

	// stop miner immediately in case of URB epoch
	if e.UserActivated {
		rcm.miner.Stop()
	}

	// start miner
	go rcm.miner.Start(params.Operator, &e)

	// prepare txs for the epoch
	switch e.UserActivated {
	case false:
		// ORB, rebased ORB
		if e.IsRequest && !e.EpochIsEmpty {
			numORBs := new(big.Int).Sub(e.EndBlockNumber, e.StartBlockNumber)
			numORBs = new(big.Int).Add(numORBs, big.NewInt(1))

			bodies := make([]types.Transactions, 0, numORBs.Uint64()) // [][]types.Transaction

			log.Debug("Num Orbs", "epochNumber", e.EpochNumber, "numORBs", numORBs, "e.EndBlockNumber", e.EndBlockNumber, "e.StartBlockNumber", e.StartBlockNumber)

			currentFork := big.NewInt(int64(rcm.state.currentFork))
			epoch, err := rcm.getEpoch(currentFork, e.EpochNumber)
			if err != nil {
				return err
			}

			// TODO: URE, ORE' should handle requestBlockId in a different way.
			requestBlockId := big.NewInt(int64(epoch.FirstRequestBlockId))
			for blockNumber := e.StartBlockNumber; blockNumber.Cmp(e.EndBlockNumber) <= 0; {

				orb, err := rcm.rootchainContract.ORBs(baseCallOpt, requestBlockId)
				if err != nil {
					return err
				}

				numRequests := orb.RequestEnd - orb.RequestStart + 1
				log.Debug("Fetching ORB", "requestBlockId", requestBlockId, "numRequests", numRequests)

				body := make(types.Transactions, 0, numRequests)

				for requestId := orb.RequestStart; requestId <= orb.RequestEnd; {
					request, err := rcm.rootchainContract.EROs(baseCallOpt, big.NewInt(int64(requestId)))
					if err != nil {
						return err
					}

					log.Debug("Request fetched", "requestId", requestId, "hash", common.Bytes2Hex(request.Hash[:]), "request", request)

					var to common.Address
					var input []byte

					if request.IsTransfer {
						to = request.Requestor
					} else {
						to, _ = rcm.rootchainContract.RequestableContracts(baseCallOpt, request.To)
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

						log.Debug("Request tx.data", "payload", input)
					}

					requestTx := types.NewTransaction(0, to, request.Value, params.RequestTxGasLimit, params.RequestTxGasPrice, input)

					log.Debug("Request Transaction", "tx", requestTx)

					eroBytes, err := rcm.rootchainContract.GetEROBytes(baseCallOpt, big.NewInt(int64(requestId)))
					if err != nil {
						log.Error("Failed to get ERO bytes", "err", err)
					}

					// TODO: check only in test
					if !bytes.Equal(eroBytes, requestTx.GetRlp()) {
						log.Error("ERO TX and request tx are different", "requestId", requestId, "eroBytes", common.Bytes2Hex(eroBytes), "requestTx.GetRlp()", common.Bytes2Hex(requestTx.GetRlp()))
					}

					body = append(body, requestTx)

					requestId += 1
				}

				log.Info("Request txs fetched", "blockNumber", blockNumber, "requestBlockId", requestBlockId, "body", body)

				bodies = append(bodies, body)

				blockNumber = new(big.Int).Add(blockNumber, big.NewInt(1))
				requestBlockId = new(big.Int).Add(requestBlockId, big.NewInt(1))
			}

			var numMinedORBs uint64 = 0

			for numMinedORBs < numORBs.Uint64() {
				rcm.txPool.EnqueueRequestTxs(bodies[numMinedORBs])

				log.Info("Waiting new request block mined event...")

				e := <-events.Chan()
				block := e.Data.(core.NewMinedBlockEvent).Block

				log.Info("New request block is mined", "block", block)

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
		} else {
			// NRB
			//TODO (Aiden): change e.IsRequest to e.IsRebase
			switch e.IsRequest {
			case false:
				// ordinary NRB
				return nil
			case true:
				// rebased NRB
				numNRBs := new(big.Int).Sub(e.EndBlockNumber, e.StartBlockNumber)
				numNRBs = new(big.Int).Add(numNRBs, big.NewInt(1))

				bodies := make([]types.Transactions, 0, numNRBs.Uint64()) // [][]types.Transaction

				log.Debug("Num Nrbs", "epochNumber", e.EpochNumber, "numNRBs", numNRBs, "e.EndBlockNumber", e.EndBlockNumber, "e.StartBlockNumber", e.StartBlockNumber)

				for blockNumber := e.StartBlockNumber; blockNumber.Cmp(e.EndBlockNumber) <= 0; {
					//TODO (aiden): should change to new method 'GetBlockByForkAndNumber'
					block := rcm.blockchain.GetBlockByNumber(blockNumber.Uint64())

					txs := block.Transactions()
					bodies = append(bodies, txs)
					blockNumber = new(big.Int).Add(blockNumber, big.NewInt(1))
				}

				var numMinedNRBs uint64 = 0

				for numMinedNRBs < numNRBs.Uint64() {
					rcm.txPool.EnqueueRebasedTxs(bodies[numMinedNRBs])

					log.Info("Waiting new block mined event...")

					<-events.Chan()

					numMinedNRBs += 1
				}
			}
		}

	case true:
		// prepare request tx for URBs
		if !e.IsRequest {
			log.Error("URB epoch is non-request epoch")
		}

		if e.EpochIsEmpty {
			log.Error("URB epoch is empty")
		}

		numURBs := new(big.Int).Sub(e.EndBlockNumber, e.StartBlockNumber)
		numURBs = new(big.Int).Add(numURBs, big.NewInt(1))

		bodies := make([]types.Transactions, 0, numURBs.Uint64()) // [][]types.Transaction

		log.Debug("Num URBs", "epochNumber", e.EpochNumber, "numURBs", numURBs, "e.EndBlockNumber", e.EndBlockNumber, "e.StartBlockNumber", e.StartBlockNumber)

		currentFork := big.NewInt(int64(rcm.state.currentFork))
		epoch, err := rcm.getEpoch(currentFork, e.EpochNumber)
		if err != nil {
			return err
		}

		// TODO: URE, ORE' should handle requestBlockId in a different way.
		requestBlockId := big.NewInt(int64(epoch.FirstRequestBlockId))
		for blockNumber := e.StartBlockNumber; blockNumber.Cmp(e.EndBlockNumber) <= 0; {

			urb, err := rcm.rootchainContract.URBs(baseCallOpt, requestBlockId)

			if err != nil {
				return err
			}

			numRequests := urb.RequestEnd - urb.RequestStart + 1
			log.Debug("Fetching URB", "requestBlockId", requestBlockId, "numRequests", numRequests)

			body := make(types.Transactions, 0, numRequests)

			for requestId := urb.RequestStart; requestId <= urb.RequestEnd; {
				request, err := rcm.rootchainContract.ERUs(baseCallOpt, big.NewInt(int64(requestId)))
				if err != nil {
					return err
				}

				log.Debug("Request fetched", "requestId", requestId, "hash", common.Bytes2Hex(request.Hash[:]), "request", request)

				var to common.Address
				var input []byte

				if request.IsTransfer {
					to = request.Requestor
				} else {
					to, _ = rcm.rootchainContract.RequestableContracts(baseCallOpt, request.To)
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

					log.Debug("Request tx.data", "payload", input)
				}

				requestTx := types.NewTransaction(0, to, request.Value, params.RequestTxGasLimit, params.RequestTxGasPrice, input)
				log.Debug("Request Transaction", "tx", requestTx)

				// TODO (aiden): add GetERUBytes method in root chain contract
				eruBytes, err := rcm.rootchainContract.GetEROBytes(baseCallOpt, big.NewInt(int64(requestId)))

				if err != nil {
					log.Error("Failed to get ERO bytes", "err", err)
				}

				// TODO: check only in test
				if !bytes.Equal(eruBytes, requestTx.GetRlp()) {
					log.Error("ERU TX and request tx are different", "requestId", requestId, "eruBytes", common.Bytes2Hex(eruBytes), "requestTx.GetRlp()", common.Bytes2Hex(requestTx.GetRlp()))

					body = append(body, requestTx)

					requestId += 1
				}

				log.Info("Request txs fetched", "blockNumber", blockNumber, "requestBlockId", requestBlockId, "body", body)

				bodies = append(bodies, body)

				blockNumber = new(big.Int).Add(blockNumber, big.NewInt(1))
				requestBlockId = new(big.Int).Add(requestBlockId, big.NewInt(1))
			}

			var numMinedURBs uint64 = 0

			for numMinedURBs < numURBs.Uint64() {
				rcm.txPool.EnqueueRequestTxs(bodies[numMinedURBs])

				log.Info("Waiting new request block mined event...")

				e := <-events.Chan()
				block := e.Data.(core.NewMinedBlockEvent).Block

				log.Info("New request block is mined", "block", block)

				if !block.IsRequest() {
					return errors.New("Invalid request block type.")
				}

				receipts := rcm.blockchain.GetReceiptsByHash(block.Hash())

				for _, receipt := range receipts {
					if receipt.Status == 0 {
						log.Error("Request transaction is reverted", "blockNumber", block.Number(), "hash", receipt.TxHash)
					}
				}
				numMinedURBs += 1
			}
		}
	}
	return nil
}

func (rcm *RootChainManager) handleBlockFinalzied(ev *rootchain.RootChainBlockFinalized) error {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()

	e := *ev

	log.Info("RootChain block finalized", "forkNumber", e.ForkNumber, "blockNubmer", e.BlockNumber)

	// send LastFinalizedBlock event to miner
	go rcm.eventMux.Post(miner.LastFinalizedBlock{Number: e.BlockNumber})

	callerOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}

	w, err := rcm.accountManager.Find(rcm.config.Operator)
	if err != nil {
		log.Error("Failed to get operator wallet", "err", err)
	}

	block, err := rcm.rootchainContract.GetBlock(callerOpts, e.ForkNumber, e.BlockNumber)
	if err != nil {
		return err
	}

	// exit challenge if any
	if block.IsRequest {
		invalidExits := rcm.invalidExits[e.ForkNumber.Uint64()][e.BlockNumber.Uint64()]
		for i := 0; i < len(invalidExits); i++ {

			var proofs []byte
			for j := 0; j < len(invalidExits[i].proof); j++ {
				proof := invalidExits[i].proof[j].Bytes()
				proofs = append(proofs, proof...)
			}
			// TODO (aiden): ChallengeExit receipt check
			input, err := rootchainContractABI.Pack("challengeExit", e.ForkNumber, e.BlockNumber, big.NewInt(invalidExits[i].index), invalidExits[i].receipt.GetRlp(), proofs)
			if err != nil {
				log.Error("Failed to pack challengeExit", "error", err)
			}

			Nonce := rcm.state.getNonce()
			challengeTx := types.NewTransaction(Nonce, rcm.config.RootChainContract, big.NewInt(0), params.SubmitBlockGasLimit, params.SubmitBlockGasPrice, input)

			signedTx, err := w.SignTx(rcm.config.Operator, challengeTx, rootchainNetworkId)
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
	events := rcm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	defer events.Unsubscribe()

	caller, err := rootchain.NewRootChainCaller(rcm.config.RootChainContract, rcm.backend)
	if err != nil {
		log.Warn("failed to make new root chain caller", "error", err)
		return
	}

	callerOpts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}

	for {
		select {
		case ev := <-events.Chan():
			rcm.lock.Lock()
			if rcm.minerEnv.IsRequest {
				var invalidExitsList invalidExits

				forkNumber, err := caller.CurrentFork(callerOpts)
				if err != nil {
					log.Warn("failed to get current fork number", "error", err)
				}

				if rcm.invalidExits[forkNumber.Uint64()] == nil {
					rcm.invalidExits[forkNumber.Uint64()] = make(map[uint64]invalidExits)
				}

				blockInfo := ev.Data.(core.NewMinedBlockEvent)
				blockNumber := blockInfo.Block.Number()
				receipts := rcm.blockchain.GetReceiptsByHash(blockInfo.Block.Hash())

				// TODO (aiden): should check if the request[i] is enter or exit request. Undo request will make posterior enter request.
				for i := 0; i < len(receipts); i++ {
					if receipts[i].Status == types.ReceiptStatusFailed {
						invalidExit := &invalidExit{
							forkNumber:  forkNumber,
							blockNumber: blockNumber,
							receipt:     receipts[i],
							index:       int64(i),
							proof:       types.GetMerkleProof(receipts, i),
						}
						invalidExitsList = append(invalidExitsList, invalidExit)

						log.Info("Invalid Exit Detected", "invalidExit", invalidExit, "forkNumber", forkNumber, "blockNumber", blockNumber)
					}
				}
				rcm.invalidExits[forkNumber.Uint64()][blockNumber.Uint64()] = invalidExitsList
			}
			rcm.lock.Unlock()

		case <-rcm.quit:
			return
		}
	}
}

func (rcm *RootChainManager) getEpoch(forkNumber, epochNumber *big.Int) (*PlasmaEpoch, error) {
	b, err := rcm.rootchainContract.GetEpoch(baseCallOpt, forkNumber, epochNumber)

	if err != nil {
		return nil, err
	}

	return newPlasmaEpoch(b), nil
}
func (rcm *RootChainManager) getBlock(forkNumber, blockNumber *big.Int) (*PlasmaBlock, error) {
	b, err := rcm.rootchainContract.GetBlock(baseCallOpt, forkNumber, blockNumber)

	if err != nil {
		return nil, err
	}

	return newPlasmaBlock(b), nil
}

// pingBackend checks rootchain backend is alive.
func (rcm *RootChainManager) pingBackend() {
	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-ticker.C:
			if _, err := rcm.backend.SyncProgress(context.Background()); err != nil {
				log.Error("Rootchain provider doesn't respond", "err", err)
				ticker.Stop()
				rcm.stopFn()
				return
			}
		case <-rcm.quit:
			ticker.Stop()
			return
		}
	}
}
