package miner

import (
	"math/big"

	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
)

type invalidExit struct {
	forkNumber 	*big.Int
	blockNumber *big.Int
	receipt 	*types.Receipt
	index 		int
	proof 		[]common.Hash
}

type invalidExits []*invalidExit

func (s invalidExits) Len() int { return len(s) }

func (s invalidExits) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(s[i])
	return enc
}

var InvalidExits map[*big.Int]invalidExits

func collectInvalidExits(receipts types.Receipts, num *big.Int) {
	var exits invalidExits
	for i := 0; i < len(receipts); i++ {
		if receipts[i].Status == types.ReceiptStatusFailed {
			exit := &invalidExit{
				// TODO: get forkNumber from root chain.
				forkNumber: big.NewInt(0),
				blockNumber: num,
				receipt: receipts[i],
				index: i,
				proof: types.GetMerkleProof(receipts, i),
			}
			exits = append(exits, exit)
		}
	}
	InvalidExits[num] = exits
}