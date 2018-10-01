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

var txHash = common.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000dead")

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

func TestGetMerkleRoot(t *testing.T) {
	var txHashes [][]byte
	for i := 0; i < 100; i++ {
		txHashes = append(txHashes, txHash)

		if i == 0 {
			fmt.Println(getMerkleRoot(txHashes))
			// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 222 173]
		} else if i == 1 {
			fmt.Println(getMerkleRoot(txHashes))
			// [10 243 254 172 103 165 159 138 108 131 158 94 125 133 231 170 22 216 86 154 11 190 216 90 226 32 79 164 101 48 13 222]
		} else if i == 9 {
			fmt.Println(getMerkleRoot(txHashes))
			// [64 240 161 254 60 96 35 250 193 54 62 138 185 243 3 66 42 134 241 125 241 215 197 26 138 69 164 111 167 107 54 117]
		} else if i == 99 {
			fmt.Println(getMerkleRoot(txHashes))
			// [9 128 149 2 140 90 91 209 3 173 57 132 170 252 80 206 44 4 237 207 101 181 253 189 195 89 252 157 13 74 6 24]
		}
	}
}



