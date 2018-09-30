package types

import (
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/crypto"
)

func getMerkleRoot(level [][]byte) common.Hash {
	if (len(level) == 1) {
		root := common.BytesToHash(level[0])
		return root
	}
	var nextlevel = make([][]byte, (len(level) + 1) / 2)
	var i int

	for i = 0; i + 1 < len(level); i += 2 {
		hash := crypto.Keccak256(level[i], level[i+1])
		nextlevel[i/2] = hash
	}

	if (len(level) % 2 == 1) {
		nextlevel[i/2] = crypto.Keccak256(level[len(level) - 1], level[len(level) - 1])
	}

	return getMerkleRoot(nextlevel)
}

//GetTransactionRoot is getter for Merkle Root of transaction hashes.
func GetTransactionRoot(txs []*Transaction) common.Hash {
	var level [][]byte
	for _, tx := range txs {
		level = append(level, tx.Hash().Bytes())
	}
	return getMerkleRoot(level)
}

//GetIntermediateStateRoot is getter for Merkle Root of intermediateState hashes.
func GetIntermediateStateRoot(receipts []*Receipt) common.Hash {
	var level [][]byte
	for _, receipt := range receipts {
		level = append(level, receipt.PostState)
	}
	return getMerkleRoot(level)
}
