package types

import (
	"github.com/Onther-Tech/plasma-evm/common"
)

// GetTransactionRoot is getter for Merkle Root of transaction hashes.
func GetTransactionRoot(txs []*Transaction) common.Hash {
	var level [][]byte
	for _, tx := range txs {
		level = append(level, tx.Hash().Bytes())
	}
	return getBinaryMerkleRoot(level)
}

// GetIntermediateStateRoot is getter for Merkle Root of intermediateState hashes.
func GetIntermediateStateRoot(receipts []*Receipt) common.Hash {
	var level [][]byte
	for _, receipt := range receipts {
		level = append(level, receipt.PostState)
	}
	return getBinaryMerkleRoot(level)
}