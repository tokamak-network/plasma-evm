// Copyright 2018 The go-ethereum Authors
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

package rawdb

import (
	"math/big"
	"testing"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethdb"
)

// Tests that positional lookup metadata can be stored and retrieved.
func TestLookupStorage(t *testing.T) {
	db := ethdb.NewMemDatabase()

	tx1 := types.NewTransaction(1, common.BytesToAddress([]byte{0x11}), big.NewInt(111), 1111, big.NewInt(11111), []byte{0x11, 0x11, 0x11})
	tx2 := types.NewTransaction(2, common.BytesToAddress([]byte{0x22}), big.NewInt(222), 2222, big.NewInt(22222), []byte{0x22, 0x22, 0x22})
	tx3 := types.NewTransaction(3, common.BytesToAddress([]byte{0x33}), big.NewInt(333), 3333, big.NewInt(33333), []byte{0x33, 0x33, 0x33})
	txs := []*types.Transaction{tx1, tx2, tx3}

	block := types.NewBlock(&types.Header{Number: big.NewInt(314)}, txs, nil, nil)

	// Check that no transactions entries are in a pristine database
	for i, tx := range txs {
		if txn, _, _, _ := ReadTransaction(db, tx.Hash()); txn != nil {
			t.Fatalf("tx #%d [%x]: non existent transaction returned: %v", i, tx.Hash(), txn)
		}
	}
	// Insert all the transactions into the database, and verify contents
	WriteBlock(db, block)
	WriteTxLookupEntries(db, block)

	for i, tx := range txs {
		if txn, hash, number, index := ReadTransaction(db, tx.Hash()); txn == nil {
			t.Fatalf("tx #%d [%x]: transaction not found", i, tx.Hash())
		} else {
			if hash != block.Hash() || number != block.NumberU64() || index != uint64(i) {
				t.Fatalf("tx #%d [%x]: positional metadata mismatch: have %x/%d/%d, want %x/%v/%v", i, tx.Hash(), hash, number, index, block.Hash(), block.NumberU64(), i)
			}
			if tx.Hash() != txn.Hash() {
				t.Fatalf("tx #%d [%x]: transaction mismatch: have %v, want %v", i, tx.Hash(), txn, tx)
			}
		}
	}
	// Delete the transactions and check purge
	for i, tx := range txs {
		DeleteTxLookupEntry(db, tx.Hash())
		if txn, _, _, _ := ReadTransaction(db, tx.Hash()); txn != nil {
			t.Fatalf("tx #%d [%x]: deleted transaction returned: %v", i, tx.Hash(), txn)
		}
	}
}

func TestInvalidExitReceiptsLookupStorage(t *testing.T) {
	db := ethdb.NewMemDatabase()

	tx1 := types.NewTransaction(1, common.BytesToAddress([]byte{0x11}), big.NewInt(111), 1111, big.NewInt(11111), []byte{0x11, 0x11, 0x11})
	tx2 := types.NewTransaction(2, common.BytesToAddress([]byte{0x22}), big.NewInt(222), 2222, big.NewInt(22222), []byte{0x22, 0x22, 0x22})
	tx3 := types.NewTransaction(3, common.BytesToAddress([]byte{0x33}), big.NewInt(333), 3333, big.NewInt(33333), []byte{0x33, 0x33, 0x33})
	txs := []*types.Transaction{tx1, tx2, tx3}

	block := types.NewBlock(&types.Header{Number: big.NewInt(314)}, txs, nil, nil)
	txIndices := []uint64{0, 2}

	WriteInvalidExitReceiptsLookupEntry(db, block.CurrentFork(), block.Hash(), block.NumberU64(), txIndices)

	hash, num, indices := ReadInvalidExitReceiptsLookupEntry(db, block.CurrentFork(), block.NumberU64())
	if hash != block.Hash() || num != block.NumberU64() || len(indices) == 0 {
		t.Fatalf("invalid exit receipt lookup entries mismatch")
		for i, v := range indices {
			if txIndices[i] != v {
				t.Fatal("invalid exit receipt index mismatch")
			}
		}
	}

	DeleteInvalidExitReceiptsLookupEntry(db, block.CurrentFork(), block.NumberU64())
	_, _, indices = ReadInvalidExitReceiptsLookupEntry(db, block.CurrentFork(), block.NumberU64())
	if len(indices) != 0 {
		t.Fatalf("invalid exit receipt indices exist")
	}
}

func TestBlockInvalidExitReceiptsStorage(t *testing.T) {
	db := ethdb.NewMemDatabase()

	fork := uint64(1)
	num := uint64(2)
	hash := common.BytesToHash([]byte{0x01, 0x15})
	receipt1 := &types.Receipt{
		Status:            types.ReceiptStatusFailed,
		CumulativeGasUsed: 1,
		Logs: []*types.Log{
			{Address: common.BytesToAddress([]byte{0x11})},
			{Address: common.BytesToAddress([]byte{0x01, 0x11})},
		},
		TxHash:          common.BytesToHash([]byte{0x11, 0x11}),
		ContractAddress: common.BytesToAddress([]byte{0x01, 0x11, 0x11}),
		GasUsed:         111111,
	}
	receipt2 := &types.Receipt{
		CumulativeGasUsed: 2,
		Logs: []*types.Log{
			{Address: common.BytesToAddress([]byte{0x22})},
			{Address: common.BytesToAddress([]byte{0x02, 0x22})},
		},
		TxHash:          common.BytesToHash([]byte{0x22, 0x22}),
		ContractAddress: common.BytesToAddress([]byte{0x02, 0x22, 0x22}),
		GasUsed:         222222,
	}
	receipt3 := &types.Receipt{
		Status:            types.ReceiptStatusFailed,
		CumulativeGasUsed: 3,
		Logs: []*types.Log{
			{Address: common.BytesToAddress([]byte{0x33})},
			{Address: common.BytesToAddress([]byte{0x03, 0x33})},
		},
		TxHash:          common.BytesToHash([]byte{0x33, 0x33}),
		ContractAddress: common.BytesToAddress([]byte{0x03, 0x33, 0x33}),
		GasUsed:         333333,
	}
	receipts := []*types.Receipt{receipt1, receipt2, receipt3}

	WriteReceipts(db, hash, num, receipts)
	WriteInvalidExitReceiptsLookupEntry(db, fork, hash, num, []uint64{0, 2})

	iers := ReadInvalidExitReceipts(db, fork, num)
	if iers == nil || len(iers) != 2 {
		t.Fatalf("invalid invalid exit receipts returned")
	} else {
		if (*iers[0]).TxHash != (*receipts[0]).TxHash || (*iers[2]).TxHash != (*receipts[2]).TxHash {
			t.Fatalf("different receipt: read receipt is %v, saved receipt is %v", *iers[0], *receipts[0])
		}
	}

	// Delete the receipts
	DeleteReceipts(db, hash, 0)
}
