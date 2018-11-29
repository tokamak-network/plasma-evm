package types

import (
	"testing"
	"math/big"

	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/Onther-Tech/plasma-evm/common"
)

// tx hash = [77 165 128 253 46 76 4 243 40 217 249 71 236 243 86 65 30 184 228 163 165 199 69 243 131 179 204 215 156 54 168 212]
var tx, _ = NewTransaction(
		3,
		common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
		big.NewInt(10),
		2000,
		big.NewInt(1),
		common.FromHex("5544"),
	).WithSignature(
		HomesteadSigner{},
		common.Hex2Bytes("98ff921201554726367d2be8c804a7ff89ccf285ebc57dff8ae4c44b9c19ac4a8887321be575c8095f789dd4c743dfe42c1820f9231f98a962b210e3ac2452a301"),
	)

func setListAndTarget(size int, target int) (DerivableList, int) {
	var list Transactions
	for i := 0; i < size; i++ {
		list = append(list, tx)
	}
	return list, target
}


func TestCheckMembership(t *testing.T) {
	list, index := setListAndTarget(8, 0)
	root := DeriveShaFromBMT(list)
	proof := GetMerkleProof(list, index)
	computedHash := crypto.Keccak256(list.GetRlp(index))
	nodeIndex := index

	for i := 0; i < len(proof); i++ {
		if nodeIndex % 2 == 0 {
			computedHash = crypto.Keccak256(computedHash, proof[i].Bytes())
		} else {
			computedHash = crypto.Keccak256(proof[i].Bytes(), computedHash)
		}
		nodeIndex = nodeIndex / 2
	}

	assert.Equal(t, computedHash, root)
}