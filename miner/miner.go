// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package miner implements Plasma block creation and mining.
package miner

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/consensus"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/state"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/pls/downloader"
	"math/big"
)

// Backend wraps all methods required for mining.
type Backend interface {
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
}

// Miner creates blocks and searches for proof-of-work values.
type Miner struct {
	mux      *event.TypeMux
	worker   *worker
	coinbase common.Address
	pls      Backend
	engine   consensus.Engine
	exitCh   chan struct{}

	canStart    int32 // can start indicates whether we can start the mining operation
	shouldStart int32 // should start indicates whether we should start after sync

	env *EpochEnvironment

	lock sync.Mutex
}

type EpochEnvironment struct {
	IsRequest          bool
	UserActivated      bool
	Rebase             bool
	Completed          bool
	NumBlockMined      *big.Int
	EpochLength        *big.Int
	CurrentFork        *big.Int
	LastFinalizedBlock *big.Int

	lock sync.Mutex
}

// TODO (aiden): it's only for test
var NRE = rootchain.RootChainEpochPrepared{
	ForkNumber:       big.NewInt(0),
	EpochNumber:      big.NewInt(1),
	StartBlockNumber: big.NewInt(1),
	EndBlockNumber:   big.NewInt(2),
	RequestStart:     new(big.Int),
	RequestEnd:       new(big.Int),
	EpochIsEmpty:     false,
	IsRequest:        false,
	UserActivated:    false,
	Rebase:           false,
}

func New(pls Backend, config *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine, env *EpochEnvironment, recommit time.Duration, gasFloor, gasCeil uint64, isLocalBlock func(block *types.Block) bool) *Miner {
	miner := &Miner{
		pls:      pls,
		mux:      mux,
		engine:   engine,
		env:      env,
		exitCh:   make(chan struct{}),
		worker:   newWorker(config, engine, pls, env, mux, recommit, gasFloor, gasCeil, isLocalBlock),
		canStart: 2,
	}
	go miner.update()

	return miner
}

func NewEpochEnvironment() *EpochEnvironment {
	return &EpochEnvironment{
		IsRequest:          false,
		UserActivated:      false,
		Rebase:             false,
		Completed:          false,
		NumBlockMined:      big.NewInt(0),
		EpochLength:        big.NewInt(0),
		CurrentFork:        big.NewInt(0),
		LastFinalizedBlock: big.NewInt(0),
	}
}

// update keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your mining operation for as long as the DOS continues.
func (self *Miner) update() {
	events := self.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
	defer events.Unsubscribe()

	for {
		select {
		case ev := <-events.Chan():
			if ev == nil {
				return
			}
			switch ev.Data.(type) {
			case downloader.StartEvent:
				atomic.StoreInt32(&self.canStart, 0)
				if self.Mining() {
					self.Stop()
					atomic.StoreInt32(&self.shouldStart, 1)
					log.Info("Mining aborted due to sync")
				}
			case downloader.DoneEvent, downloader.FailedEvent:
				shouldStart := atomic.LoadInt32(&self.shouldStart) == 1

				atomic.StoreInt32(&self.canStart, 1)
				atomic.StoreInt32(&self.shouldStart, 0)
				if shouldStart {
					// TODO (aiden): there's no need to start miner in here, it starts when rcm connect to root chain contract by reading 1st NRE.
					//self.Start(self.coinbase, &NRE)
				}
				// stop immediately and ignore all further pending events
				return
			}
		case <-self.exitCh:
			return
		}
	}
}

func (self *Miner) Start(coinbase common.Address, epoch *rootchain.RootChainEpochPrepared) {

	if epoch.EpochIsEmpty {
		log.Info("ORB epoch is empty, NRB epoch will be started")
		return
	}

	atomic.StoreInt32(&self.shouldStart, 1)
	self.SetEtherbase(coinbase)

	if atomic.LoadInt32(&self.canStart) == 0 {
		log.Info("Network syncing, will start miner afterwards")
		return
	}

	self.env.setIsRequest(epoch.IsRequest)
	self.env.setUserActivated(epoch.UserActivated)
	self.env.setRebase(epoch.Rebase)
	self.env.setCompleted(false)
	self.env.setNumBlockMined(big.NewInt(0))
	self.env.setEpochLength(new(big.Int).Add(new(big.Int).Sub(epoch.EndBlockNumber, epoch.StartBlockNumber), big.NewInt(1)))
	self.env.setCurrentFork(epoch.ForkNumber)

	if epoch.IsRequest && !epoch.UserActivated {
		log.Info("ORB epoch is prepared, ORB epoch is started", "ORBepochLength", self.env.EpochLength)
	} else if epoch.IsRequest && epoch.UserActivated {
		log.Info("URB epoch is prepared, URB epoch is started", "URBepochLength", self.env.EpochLength)
	} else if !epoch.IsRequest {
		log.Info("NRB epoch is prepared, NRB epoch is started", "NRBepochLength", self.env.EpochLength)
	}

	self.worker.start()
}

func (self *Miner) Stop() {
	self.worker.stop()
	atomic.StoreInt32(&self.shouldStart, 0)
}

func (self *Miner) Close() {
	self.worker.close()
	close(self.exitCh)
}

func (self *Miner) Mining() bool {
	return self.worker.isRunning()
}

func (self *Miner) HashRate() uint64 {
	if pow, ok := self.engine.(consensus.PoW); ok {
		return uint64(pow.Hashrate())
	}
	return 0
}

func (self *Miner) SetExtra(extra []byte) error {
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("Extra exceeds max length. %d > %v", len(extra), params.MaximumExtraDataSize)
	}
	self.worker.setExtra(extra)
	return nil
}

// SetRecommitInterval sets the interval for sealing work resubmitting.
func (self *Miner) SetRecommitInterval(interval time.Duration) {
	self.worker.setRecommitInterval(interval)
}

// Pending returns the currently pending block and associated state.
func (self *Miner) Pending() (*types.Block, *state.StateDB) {
	return self.worker.pending()
}

// PendingBlock returns the currently pending block.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (self *Miner) PendingBlock() *types.Block {
	return self.worker.pendingBlock()
}

func (self *Miner) SetEtherbase(addr common.Address) {
	self.coinbase = addr
	self.worker.setEtherbase(addr)
}

func (self *Miner) SetNRBepochLength(length *big.Int) {
	self.env.lock.Lock()
	defer self.env.lock.Unlock()

	self.env.EpochLength = length
}

func (env *EpochEnvironment) setIsRequest(b bool) {
	env.lock.Lock()
	defer env.lock.Unlock()

	env.IsRequest = b
}

func (env *EpochEnvironment) setUserActivated(b bool) {
	env.lock.Lock()
	defer env.lock.Unlock()

	env.UserActivated = b
}

func (env *EpochEnvironment) setRebase(b bool) {
	env.lock.Lock()
	defer env.lock.Unlock()

	env.Rebase = b
}

func (env *EpochEnvironment) setCompleted(b bool) {
	env.lock.Lock()
	defer env.lock.Unlock()

	env.Completed = b
}

func (env *EpochEnvironment) setNumBlockMined(n *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()

	env.NumBlockMined = n
}

func (env *EpochEnvironment) setEpochLength(l *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()

	env.EpochLength = l
}

func (env *EpochEnvironment) setCurrentFork(f *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()

	env.CurrentFork = f
}

func (env *EpochEnvironment) SetLastFinalizedBlock(n *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()

	env.LastFinalizedBlock = n
}
