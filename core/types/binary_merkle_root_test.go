package types

import (
	"testing"
	"math/big"
	"github.com/Onther-Tech/plasma-evm/common"
	"fmt"
)


// tx1 hash = [77 165 128 253 46 76 4 243 40 217 249 71 236 243 86 65 30 184 228 163 165 199 69 243 131 179 204 215 156 54 168 212]
var (
	tx1, _ = NewTransaction(
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

	tx2, _ = NewTransaction(
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

	tx3, _ = NewTransaction(
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

	receipt1 = NewReceipt(common.FromHex("1"), false, 0)

	receipt2 = NewReceipt(common.FromHex("2"), false, 0)
)


func TestGetTxHash(t *testing.T) {
	fmt.Println(tx1.Hash())
	fmt.Println(tx2.Hash())
	fmt.Println(tx3.Hash())
}

func TestGetPostState(t *testing.T) {
	fmt.Println(receipt1.PostState)
	fmt.Println(receipt2.PostState)
}

func TestGetTransactionRootOdd(t *testing.T) {
	var oddTxs []*Transaction
	oddTxs = append(oddTxs, tx1, tx2, tx3)
	fmt.Println(GetTransactionRoot(oddTxs))
}

func TestGetTransactionRootEven(t *testing.T) {
	var evenTxs []*Transaction
	evenTxs = append(evenTxs, tx1, tx2)
	fmt.Println(GetTransactionRoot(evenTxs))
}

func TestGetIntermediateStateRootOdd(t *testing.T) {

}

func TestGetIntermediateStateRootEven(t *testing.T) {

}



var txHash = common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000dead")

func TestGetBinaryMerkleRoot(t *testing.T) {
	var txHashes [][]byte
	for i := 0; i < 100; i++ {
		txHashes = append(txHashes, txHash)

		// 1 txHash
		if len(txHashes) == 1 {
			fmt.Println(getBinaryMerkleRoot(txHashes).Hex())
			// 0x000000000000000000000000000000000000000000000000000000000000dead

		// 2 txHashes
		} else if len(txHashes) == 2 {
			fmt.Println(getBinaryMerkleRoot(txHashes).Hex())
			// 0x0af3feac67a59f8a6c839e5e7d85e7aa16d8569a0bbed85ae2204fa465300dde

		// 10 txHashes
		} else if len(txHashes) == 10 {
			fmt.Println(getBinaryMerkleRoot(txHashes).Hex())
			// 0x40f0a1fe3c6023fac1363e8ab9f303422a86f17df1d7c51a8a45a46fa76b3675

		// 100 txHashes
		} else if len(txHashes) == 100 {
			fmt.Println(getBinaryMerkleRoot(txHashes).Hex())
			// 0x098095028c5a5bd103ad3984aafc50ce2c04edcf65b5fdbdc359fc9d0d4a0618
		}
	}
}








