package types

import (
	"math/big"
	"testing"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/crypto"
	"github.com/stretchr/testify/assert"
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

var txHash = common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000dead")

func setListAndTarget(size int, target int) (DerivableList, int) {
	var list Transactions
	for i := 0; i < size; i++ {
		list = append(list, tx)
	}
	return list, target
}

func TestGetBinaryMerkleRoot(t *testing.T) {
	var txHashes []common.Hash
	for i := 0; i < 100; i++ {
		txHashes = append(txHashes, common.BytesToHash(txHash))

		// 1 txHash
		if len(txHashes) == 1 {
			actual := getBinaryMerkleRoot(txHashes).Hex()
			expected := "0x000000000000000000000000000000000000000000000000000000000000dead"
			if !assert.Equal(t, actual, expected) {
				t.Fatal("both hash should be equal, but they aren't", actual, expected)
			}

			// 2 txHashes
		} else if len(txHashes) == 2 {
			actual := getBinaryMerkleRoot(txHashes).Hex()
			expected := "0x0af3feac67a59f8a6c839e5e7d85e7aa16d8569a0bbed85ae2204fa465300dde"
			if !assert.Equal(t, actual, expected) {
				t.Fatal("both hash should be equal, but they aren't", actual, expected)
			}

			// 10 txHashes
		} else if len(txHashes) == 10 {
			actual := getBinaryMerkleRoot(txHashes).Hex()
			expected := "0x40f0a1fe3c6023fac1363e8ab9f303422a86f17df1d7c51a8a45a46fa76b3675"
			if !assert.Equal(t, actual, expected) {
				t.Fatal("both hash should be equal, but they aren't", actual, expected)
			}

			// 100 txHashes
		} else if len(txHashes) == 100 {
			actual := getBinaryMerkleRoot(txHashes).Hex()
			expected := "0x098095028c5a5bd103ad3984aafc50ce2c04edcf65b5fdbdc359fc9d0d4a0618"
			if !assert.Equal(t, actual, expected) {
				t.Fatal("both hash should be equal, but they aren't", actual, expected)
			}
		}
	}
}

func TestCheckMembership(t *testing.T) {
	list, index := setListAndTarget(8, 0)
	root := DeriveShaFromBMT(list)
	proof := GetMerkleProof(list, index)
	computedHash := crypto.Keccak256(list.GetRlp(index))
	nodeIndex := index

	for i := 0; i < len(proof); i++ {
		if nodeIndex%2 == 0 {
			computedHash = crypto.Keccak256(computedHash, proof[i].Bytes())
		} else {
			computedHash = crypto.Keccak256(proof[i].Bytes(), computedHash)
		}
		nodeIndex = nodeIndex / 2
	}
	cH := common.BytesToHash(computedHash).Hex()
	r := root.Hex()

	if !assert.Equal(t, cH, r) {
		t.Fatal("both hash should be equal, but they aren't", cH, r)
	}
}
