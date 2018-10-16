// Copyright 2015 The go-ethereum Authors
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

package pls

import (
	"context"
	"math/big"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/common/math"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/bloombits"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/core/state"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/core/vm"
	"github.com/Onther-Tech/plasma-evm/pls/downloader"
	"github.com/Onther-Tech/plasma-evm/pls/gasprice"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/event"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/rpc"
)

// PlsAPIBackend implements ethapi.Backend for full nodes
type PlsAPIBackend struct {
	pls *Plasma
	gpo *gasprice.Oracle
}

// ChainConfig returns the active chain configuration.
func (b *PlsAPIBackend) ChainConfig() *params.ChainConfig {
	return b.pls.chainConfig
}

func (b *PlsAPIBackend) CurrentBlock() *types.Block {
	return b.pls.blockchain.CurrentBlock()
}

func (b *PlsAPIBackend) SetHead(number uint64) {
	b.pls.protocolManager.downloader.Cancel()
	b.pls.blockchain.SetHead(number)
}

func (b *PlsAPIBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.pls.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.pls.blockchain.CurrentBlock().Header(), nil
	}
	return b.pls.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *PlsAPIBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.pls.blockchain.GetHeaderByHash(hash), nil
}

func (b *PlsAPIBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.pls.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.pls.blockchain.CurrentBlock(), nil
	}
	return b.pls.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *PlsAPIBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.pls.miner.Pending()
		return state, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.pls.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *PlsAPIBackend) GetBlock(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.pls.blockchain.GetBlockByHash(hash), nil
}

func (b *PlsAPIBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	if number := rawdb.ReadHeaderNumber(b.pls.chainDb, hash); number != nil {
		return rawdb.ReadReceipts(b.pls.chainDb, hash, *number), nil
	}
	return nil, nil
}

func (b *PlsAPIBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	number := rawdb.ReadHeaderNumber(b.pls.chainDb, hash)
	if number == nil {
		return nil, nil
	}
	receipts := rawdb.ReadReceipts(b.pls.chainDb, hash, *number)
	if receipts == nil {
		return nil, nil
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

func (b *PlsAPIBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.pls.blockchain.GetTdByHash(blockHash)
}

func (b *PlsAPIBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.pls.BlockChain(), nil)
	return vm.NewEVM(context, state, b.pls.chainConfig, vmCfg), vmError, nil
}

func (b *PlsAPIBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.pls.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *PlsAPIBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.pls.BlockChain().SubscribeChainEvent(ch)
}

func (b *PlsAPIBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.pls.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *PlsAPIBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.pls.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *PlsAPIBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.pls.BlockChain().SubscribeLogsEvent(ch)
}

func (b *PlsAPIBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.pls.txPool.AddLocal(signedTx)
}

func (b *PlsAPIBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.pls.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *PlsAPIBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.pls.txPool.Get(hash)
}

func (b *PlsAPIBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.pls.txPool.State().GetNonce(addr), nil
}

func (b *PlsAPIBackend) Stats() (pending int, queued int) {
	return b.pls.txPool.Stats()
}

func (b *PlsAPIBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.pls.TxPool().Content()
}

func (b *PlsAPIBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.pls.TxPool().SubscribeNewTxsEvent(ch)
}

func (b *PlsAPIBackend) Downloader() *downloader.Downloader {
	return b.pls.Downloader()
}

func (b *PlsAPIBackend) ProtocolVersion() int {
	return b.pls.EthVersion()
}

func (b *PlsAPIBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *PlsAPIBackend) ChainDb() ethdb.Database {
	return b.pls.ChainDb()
}

func (b *PlsAPIBackend) EventMux() *event.TypeMux {
	return b.pls.EventMux()
}

func (b *PlsAPIBackend) AccountManager() *accounts.Manager {
	return b.pls.AccountManager()
}

func (b *PlsAPIBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.pls.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *PlsAPIBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.pls.bloomRequests)
	}
}
