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

func New(pls Backend, config *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine, env *EpochEnvironment, recommit time.Duration, gasFloor, gasCeil uint64) *Miner {
	miner := &Miner{
		pls:      pls,
		mux:      mux,
		engine:   engine,
		env:      env,
		exitCh:   make(chan struct{}),
		worker:   newWorker(config, engine, pls, env, mux, recommit, gasFloor, gasCeil),
		canStart: 2,
	}
	go miner.update()
	go miner.operate()

	return miner
}

// operate manages mining operation according to the events sent by rootchain manager.
func (self *Miner) operate() {
	events := self.mux.Subscribe(core.NewMinedBlockEvent{}, EpochPrepared{})
	defer events.Unsubscribe()

	for {
		select {
		case ev := <-events.Chan():
			if ev == nil {
				return
			}
			switch ev.Data.(type) {
			case core.NewMinedBlockEvent:
				// stop mining when the epoch is completed
				if self.env.Completed == true {
					self.Stop()
					atomic.StoreInt32(&self.canStart, 2)
					switch self.env.IsRequest {
					case true:
						if !self.env.IsUserActivated {
							self.env.setNumORBmined(big.NewInt(0))
							log.Info("ORB epoch is completed, Waiting for preparing next epoch")
						} else {
							self.env.setNumURBmined(big.NewInt(0))
							log.Info("URB epoch is completed, Waiting for preparing next epoch")
						}
					case false:
						self.env.setNumNRBmined(big.NewInt(0))
						log.Info("NRB epoch is completed, Waiting for preparing next epoch")
					}
				}
			case EpochPrepared:
				// start mining only when the epoch is prepared
				atomic.StoreInt32(&self.canStart, 1)
				self.env.setCompletedFalse()
				payload := ev.Data.(EpochPrepared).Payload
				switch payload.IsRequest {
				case true:
					// start mining ORBs
					if !payload.UserActivated {
						self.env.setORBepochLength(big.NewInt(0))
						if payload.EpochIsEmpty == true {
							self.env.setIsRequest(false)
							self.Start(params.Operator)
							log.Info("ORB epoch is empty, NRB epoch is started", "NRBepochLength", self.env.NRBepochLength)
						} else {
							self.env.setIsRequest(true)
							ORBepochLength := new(big.Int).Add(new(big.Int).Sub(payload.EndBlockNumber, payload.StartBlockNumber), big.NewInt(1))
							self.env.setORBepochLength(ORBepochLength)
							self.Start(params.Operator)
							log.Info("ORB epoch is prepared, ORB epoch is started", "ORBepochLength", self.env.ORBepochLength)
						}
						// stop current mining operation and restart to mine URBs
					} else {
						// stop and reset current mining operation
						self.Stop()
						atomic.StoreInt32(&self.canStart, 2)
						switch self.env.IsRequest {
						case true:
							self.env.setNumORBmined(big.NewInt(0))
							log.Info("current ORB epoch is canceled, URB epoch is prepared")
						case false:
							self.env.setNumNRBmined(big.NewInt(0))
							log.Info("current NRB epoch is canceled, URB epoch is prepared")
						}
						// start mining URBs
						self.env.SetEmergency(false) // emergency is now relieved
						self.env.setURBepochLength(big.NewInt(0))
						// TODO: should we check the URB epoch is empty?
						if payload.EpochIsEmpty == true {
							self.env.setIsRequest(false)
							self.Start(params.Operator)
							log.Info("URB epoch is empty, NRB epoch is started")
						} else {
							self.env.setIsRequest(true)
							URBepochLength := new(big.Int).Add(new(big.Int).Sub(payload.EndBlockNumber, payload.StartBlockNumber), big.NewInt(1))
							self.env.setURBepochLength(URBepochLength)
							self.Start(params.Operator)
							log.Info("URB epoch is started", "URBepochLength", URBepochLength)
						}
					}
				case false:
					self.env.setIsRequest(false)
					self.Start(params.Operator)
					log.Info("NRB epoch is prepared, NRB epoch is started", "NRBepochLength", self.env.NRBepochLength)
				}
			}
		case <-self.exitCh:
			return
		}
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
					self.Start(self.coinbase)
				}
				// stop immediately and ignore all further pending events
				return
			}
		case <-self.exitCh:
			return
		}
	}
}

func (self *Miner) Start(coinbase common.Address) {
	atomic.StoreInt32(&self.shouldStart, 1)
	self.SetEtherbase(coinbase)

	if atomic.LoadInt32(&self.canStart) == 0 {
		log.Info("Network syncing, will start miner afterwards")
		return
	}
	if atomic.LoadInt32(&self.canStart) == 2 {
		log.Info("Preparing epoch, will start miner afterwards")
		return
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

func (self *Miner) SetNRBepochLength(NRBepochLength *big.Int) {
	self.env.setNRBepochLength(NRBepochLength)
}

type EpochEnvironment struct {
	IsRequest       bool
	IsUserActivated bool
	NumNRBmined     *big.Int
	NumORBmined     *big.Int
	NumURBmined     *big.Int
	NRBepochLength  *big.Int
	ORBepochLength  *big.Int
	URBepochLength  *big.Int
	LastFinalized   *big.Int
	Completed       bool
	Emergency       bool

	lock sync.Mutex
}

func NewEpochEnvironment() *EpochEnvironment {
	return &EpochEnvironment{
		IsRequest:       false,
		IsUserActivated: false,
		NumNRBmined:     big.NewInt(0),
		NumORBmined:     big.NewInt(0),
		NumURBmined:     big.NewInt(0),
		NRBepochLength:  big.NewInt(0),
		ORBepochLength:  big.NewInt(0),
		URBepochLength:  big.NewInt(0),
		LastFinalized:   big.NewInt(0),
		Completed:       false,
		Emergency:       false,
	}
}

func (env *EpochEnvironment) setCompletedTrue() {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.Completed = true
}

func (env *EpochEnvironment) setCompletedFalse() {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.Completed = false
}

func (env *EpochEnvironment) setIsRequest(IsRequest bool) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.IsRequest = IsRequest
}

func (env *EpochEnvironment) setIsUserActivated(IsUserActivated bool) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.IsUserActivated = IsUserActivated
}

func (env *EpochEnvironment) setNumNRBmined(NumNRBmined *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.NumNRBmined = NumNRBmined
}

func (env *EpochEnvironment) setNumORBmined(NumORBmined *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.NumORBmined = NumORBmined
}

func (env *EpochEnvironment) setNumURBmined(NumURBmined *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.NumURBmined = NumURBmined
}

func (env *EpochEnvironment) setNRBepochLength(NRBepochLength *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.NRBepochLength = NRBepochLength
}

func (env *EpochEnvironment) setORBepochLength(ORBepochLength *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.ORBepochLength = ORBepochLength
}

func (env *EpochEnvironment) setURBepochLength(URBepochLength *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.URBepochLength = URBepochLength
}

func (env *EpochEnvironment) SetLastFinalized(n *big.Int) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.LastFinalized = n
}

func (env *EpochEnvironment) SetEmergency(e bool) {
	env.lock.Lock()
	defer env.lock.Unlock()
	env.Emergency = e
}
