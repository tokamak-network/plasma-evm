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

var (
	isRequest      = false
	numNRBmined    = big.NewInt(0)
	numORBmined    = big.NewInt(0)
	NRBepochLength = big.NewInt(2)
	ORBepochLength = big.NewInt(2)
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
}

func New(pls Backend, config *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine, recommit time.Duration, gasFloor, gasCeil uint64) *Miner {
	miner := &Miner{
		pls:      pls,
		mux:      mux,
		engine:   engine,
		exitCh:   make(chan struct{}),
		worker:   newWorker(config, engine, pls, mux, recommit, gasFloor, gasCeil),
		canStart: 2,
	}
	go miner.update()
	go miner.operate()

	return miner
}

func (self *Miner) operate() {
	events := self.mux.Subscribe(NRBEpochCompleted{}, ORBEpochCompleted{}, EpochPrepared{})
	defer events.Unsubscribe()

	for {
		select {
		case ev := <-events.Chan():
			if ev == nil {
				return
			}
			switch ev.Data.(type) {
			case NRBEpochCompleted:
				self.Stop()
				numNRBmined.Set(big.NewInt(0))
				atomic.StoreInt32(&self.canStart, 2)
				log.Info("NRB epoch is completed, Waiting for preparing next epoch")
			case ORBEpochCompleted:
				self.Stop()
				numORBmined.Set(big.NewInt(0))
				atomic.StoreInt32(&self.canStart, 2)
				log.Info("ORB epoch is completed, Waiting for preparing next epoch")
			case EpochPrepared:
				// start mining only when the epoch is prepared
				atomic.StoreInt32(&self.canStart, 1)
				payload := ev.Data.(EpochPrepared).Payload
				switch payload.IsRequest {
				case true:
					ORBepochLength.Set(big.NewInt(0))
					if payload.EpochIsEmpty == true {
						isRequest = false
						self.Start(params.Operator)
						log.Info("ORB epoch is empty, NRB epoch is started")
					} else {
						isRequest = true
						ORBepochLength.Add(payload.EndBlockNumber.Sub(payload.EndBlockNumber, payload.StartBlockNumber), big.NewInt(1))
						self.Start(params.Operator)
						log.Info("ORB epoch is prepared, ORB epoch is started")
					}
				case false:
					isRequest = false
					self.Start(params.Operator)
					log.Info("NRB epoch is prepared, NRB epoch is started")
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
