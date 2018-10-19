package pls

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/contract"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/miner"
)

const MAX_EPOCH_EVENTS = 0

var baseCallOpt *bind.CallOpts = &bind.CallOpts{Pending: false, Context: context.Background()}

type RootChainManager struct {
	config *Config
	stopFn func()

	txPool     *core.TxPool
	blockchain *core.BlockChain

	backend           *ethclient.Client
	rootchainContract *contract.RootChain

	eventMux       *event.TypeMux
	accountManager *accounts.Manager

	miner *miner.Miner

	// channels
	quit            chan struct{}
	epochPreparedCh chan *contract.RootChainEpochPrepared

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and etherbase)
}

func (rcm *RootChainManager) RootchainContract() *contract.RootChain { return rcm.rootchainContract }
func (rcm *RootChainManager) NRBEpochLength() (*big.Int, error) {
	return rcm.rootchainContract.NRBEpochLength(baseCallOpt)
}

func NewRootChainManager(
	config *Config,
	stopFn func(),
	txPool *core.TxPool,
	blockchain *core.BlockChain,
	backend *ethclient.Client,
	rootchainContract *contract.RootChain,
	eventMux *event.TypeMux,
	accountManager *accounts.Manager,
	miner *miner.Miner,
) *RootChainManager {
	rcm := &RootChainManager{
		config:            config,
		stopFn:            stopFn,
		txPool:            txPool,
		blockchain:        blockchain,
		backend:           backend,
		rootchainContract: rootchainContract,
		eventMux:          eventMux,
		accountManager:    accountManager,
		miner:             miner,
		quit:              make(chan struct{}),
		epochPreparedCh:   make(chan *contract.RootChainEpochPrepared, MAX_EPOCH_EVENTS),
	}

	return rcm
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

	if err := rcm.watchEvents(); err != nil {
		return err
	}

	return nil
}

// watch RootChain contract events
func (rcm *RootChainManager) watchEvents() error {
	filterer, err := contract.NewRootChainFilterer(rcm.config.RootChainContract, rcm.backend)
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

	// watch events from now
	watchOpts := &bind.WatchOpts{
		Context: context.Background(),
		Start:   &startBlockNumber, // read events from rootchain block#1
	}

	epochPrepareWatchCh := make(chan *contract.RootChainEpochPrepared)
	epochPrepareSub, err := filterer.WatchEpochPrepared(watchOpts, epochPrepareWatchCh)
	if err != nil {
		return err
	}

	log.Info("Watching EpochPrepared event", "startBlockNumber", startBlockNumber)

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

			case <-rcm.quit:
				return
			}
		}
	}()

	return nil
}

func (rcm *RootChainManager) runSubmitter() {
	events := rcm.eventMux.Subscribe(miner.BlockMined{})
	defer events.Unsubscribe()

	privKey, err := crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	if err != nil {
		log.Error("Failed to get operator private key")
		return
	}
	transactOpts := bind.NewKeyedTransactor(privKey)

	for {
		select {
		case ev := <-events.Chan():
			blockInfo := ev.Data.(miner.BlockMined).Payload

			// send block to root chain contract
			if blockInfo.IsRequest == false {
				_, err := rcm.rootchainContract.SubmitNRB(transactOpts, blockInfo.Header.Root, blockInfo.Header.TxHash, blockInfo.Header.IntermediateStateHash)
				if err != nil {
					log.Warn("Failed to submit non request block", "error", err)
				}
			} else {
				_, err := rcm.rootchainContract.SubmitORB(transactOpts, blockInfo.Header.Root, blockInfo.Header.TxHash, blockInfo.Header.IntermediateStateHash)
				if err != nil {
					log.Warn("Failed to submit request block", "error", err)
				}
			}
		case <-rcm.quit:
			return
		}
	}
}

func (rcm *RootChainManager) runHandlers() {
	events := rcm.eventMux.Subscribe(miner.BlockMined{})
	defer events.Unsubscribe()

	// var epoch []miner.BlockMined
	for {
		select {
		case e := <-rcm.epochPreparedCh:
			if err := rcm.handleEpochPrepared(e); err != nil {
				log.Error("Failed to handle epoch prepared", "err", err)
			}
		case <-rcm.quit:
			return
		}
	}
}

func (rcm *RootChainManager) handleEpochPrepared(ev *contract.RootChainEpochPrepared) error {
	rcm.lock.Lock()
	defer rcm.lock.Unlock()

	e := *ev

	log.Info("RootChain epoch prepared", "epochNunber", e.EpochNumber, "isRequest", e.IsRequest, "userActivated", e.UserActivated)

	// prepare request tx for ORBs
	if e.IsRequest && !e.EpochIsEmpty {
		numORBs := new(big.Int).Sub(e.EndBlockNumber, e.StartBlockNumber)
		numORBs = new(big.Int).Add(numORBs, big.NewInt(1))

		bodies := make([]types.Transactions, 0, numORBs.Uint64()) // [][]types.Transaction

		log.Error("Num Orbs", "epochNumber", e.EpochNumber, "numORBs", numORBs, "e.EndBlockNumber", e.EndBlockNumber, "e.StartBlockNumber", e.StartBlockNumber)

		for blockNumber := e.StartBlockNumber; blockNumber.Cmp(e.EndBlockNumber) <= 0; {
			callOpts := &bind.CallOpts{
				Pending: false,
				Context: context.Background(),
			}

			currentFork, err := rcm.rootchainContract.CurrentFork(callOpts)
			if err != nil {
				return err
			}

			pb, err := rcm.rootchainContract.Blocks(callOpts, currentFork, blockNumber)
			if err != nil {
				return err
			}

			orb, err := rcm.rootchainContract.ORBs(callOpts, big.NewInt(int64(pb.RequestBlockId)))
			if err != nil {
				return err
			}

			numRequests := orb.RequestEnd - orb.RequestStart + 1
			log.Error("Fetching ORB", "requestBlockId", pb.RequestBlockId, "numRequests", numRequests)

			body := make(types.Transactions, 0, numRequests)

			for requestId := orb.RequestStart; requestId <= orb.RequestEnd; {
				request, err := rcm.rootchainContract.EROs(callOpts, big.NewInt(int64(requestId)))
				if err != nil {
					return err
				}

				requestTx := rcm.toTransaction(request)
				body = append(body, requestTx)

				requestId += 1
			}

			log.Info("Request txs fetched", "blockNumber", blockNumber, "requestBlockId", pb.RequestBlockId, "body", body)

			bodies = append(bodies, body)

			blockNumber = new(big.Int).Add(blockNumber, big.NewInt(1))
		}

		events := rcm.eventMux.Subscribe(miner.BlockMined{})
		defer events.Unsubscribe()

		var numMinedORBs uint64 = 0

		for numMinedORBs < numORBs.Uint64() {
			go rcm.txPool.EnqueueReqeustTxs(bodies[numMinedORBs])

			<-events.Chan()

			numMinedORBs += 1
		}
	}

	go rcm.eventMux.Post(miner.EpochPrepared{Payload: &e})

	return nil
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

type Request struct {
	Timestamp  uint64
	IsExit     bool
	Applied    bool
	Finalized  bool
	Challenged bool
	Value      *big.Int
	Requestor  common.Address
	To         common.Address
	TrieKey    [32]byte
	TrieValue  [32]byte
}

func (rcm *RootChainManager) toTransaction(data interface{}) *types.Transaction {
	request, ok := data.(Request)

	if !ok {
		return &types.Transaction{}
	}

	isTransfer := request.Requestor.Hex() == request.To.Hex()
	var to common.Address

	if isTransfer {
		to = request.Requestor
	} else {
		callOpts := &bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		}

		to, _ = rcm.rootchainContract.RequestableContracts(callOpts, request.To)
	}

	// TODO: check data field
	return types.NewTransaction(0, to, request.Value, uint64(100000), big.NewInt(1000000000), nil)
}

type rootchainParameters struct {
	costERO        *big.Int
	costERU        *big.Int
	costURBPrepare *big.Int
	costURB        *big.Int
	costORB        *big.Int
	costNRB        *big.Int
	maxRequests    *big.Int
	requestGas     *big.Int
	currentEpoch   *big.Int
	currentFork    *big.Int
}

func (rp *rootchainParameters) setCostERO(rootchainContract *contract.RootChain) *big.Int {
	rp.costERO, _ = rootchainContract.COSTERO(baseCallOpt)
	return rp.costERO
}
func (rp *rootchainParameters) setCostERU(rootchainContract *contract.RootChain) *big.Int {
	rp.costERU, _ = rootchainContract.COSTERU(baseCallOpt)
	return rp.costERU
}
func (rp *rootchainParameters) setCostURBPrepare(rootchainContract *contract.RootChain) *big.Int {
	rp.costURBPrepare, _ = rootchainContract.COSTURBPREPARE(baseCallOpt)
	return rp.costURBPrepare
}
func (rp *rootchainParameters) setCostURB(rootchainContract *contract.RootChain) *big.Int {
	rp.costURB, _ = rootchainContract.COSTURB(baseCallOpt)
	return rp.costURB
}
func (rp *rootchainParameters) setCostORB(rootchainContract *contract.RootChain) *big.Int {
	rp.costORB, _ = rootchainContract.COSTORB(baseCallOpt)
	return rp.costORB
}
func (rp *rootchainParameters) setCostNRB(rootchainContract *contract.RootChain) *big.Int {
	rp.costNRB, _ = rootchainContract.COSTNRB(baseCallOpt)
	return rp.costNRB
}
func (rp *rootchainParameters) setMaxRequests(rootchainContract *contract.RootChain) *big.Int {
	rp.maxRequests, _ = rootchainContract.MAXREQUESTS(baseCallOpt)
	return rp.maxRequests
}
func (rp *rootchainParameters) setRequestGas(rootchainContract *contract.RootChain) *big.Int {
	rp.requestGas, _ = rootchainContract.REQUESTGAS(baseCallOpt)
	return rp.requestGas
}
func (rp *rootchainParameters) setCurrentEpoch(rootchainContract *contract.RootChain) *big.Int {
	rp.currentEpoch, _ = rootchainContract.CurrentEpoch(baseCallOpt)
	return rp.currentEpoch
}
func (rp *rootchainParameters) setCurrentFork(rootchainContract *contract.RootChain) *big.Int {
	rp.currentFork, _ = rootchainContract.CurrentFork(baseCallOpt)
	return rp.currentFork
}
