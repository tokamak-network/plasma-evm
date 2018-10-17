package pls

import (
	"context"
	"sync"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/contract"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/miner"
	"github.com/Onther-Tech/plasma-evm/params"
)

type RootChainManager struct {
	config      *Config
	chainConfig *params.ChainConfig

	txPool     *core.TxPool
	blockchain *core.BlockChain

	backend           *ethclient.Client
	rootchainContract *contract.RootChain

	eventMux       *event.TypeMux
	accountManager *accounts.Manager

	miner *miner.Miner

	// channels
	quit chan struct{}

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and etherbase)
}

func NewRootChainManager(
	config *Config,
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
		txPool:            txPool,
		blockchain:        blockchain,
		backend:           backend,
		rootchainContract: rootchainContract,
		eventMux:          eventMux,
		accountManager:    accountManager,
		miner:             miner,
		quit:              make(chan struct{}),
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
	close(rcm.quit)
	return nil
}

func (rcm *RootChainManager) run() error {
	if err := rcm.watchEvents(); err != nil {
		return err
	}

	return nil
}

// watch RootChain contract events
func (rcm *RootChainManager) watchEvents() error {
	// rootchain block#1
	var startBlockNumber uint64 = 1

	filterer, err := contract.NewRootChainFilterer(rcm.config.RootChainContract, rcm.backend)
	if err != nil {
		return err
	}

	watchOpts := &bind.WatchOpts{
		Context: context.Background(),
		Start:   &startBlockNumber, // read events from rootchain block#1
	}

	epochPrepareEventCh := make(chan *contract.RootChainEpochPrepared)
	epochPrepareSub, err := filterer.WatchEpochPrepared(watchOpts, epochPrepareEventCh)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case e := <-epochPrepareEventCh:
				if e != nil {
					log.Info("RootChain epoch prepared", "forkNumber", e.ForkNumber, "epochNunber", e.EpochNumber, "isRequest", e.IsRequest, "userActivated", e.UserActivated)
				}

			case err := <-epochPrepareSub.Err():
				log.Error("EpochPrepared event subscription error", err)
				return

			case <-rcm.quit:
				return
			}
		}
	}()

	return nil
}

// pingBackend checks rootchain backend is alive.
func (rcm *RootChainManager) pingBackend() {
	ticker := time.NewTicker(3 * time.Second)

	for {
		<-ticker.C
		if _, err := rcm.backend.SyncProgress(context.Background()); err != nil {
			log.Error("Rootchain provider doesn't respond", err)
			ticker.Stop()
			rcm.Stop()
			return
		}
	}
}
