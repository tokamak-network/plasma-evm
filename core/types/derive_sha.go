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

package types

import (
	"bytes"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/rlp"
	"github.com/Onther-Tech/plasma-evm/trie"
	"github.com/ethereum/go-ethereum/crypto"
)

type DerivableList interface {
	Len() int
	GetRlp(i int) []byte
}

func DeriveSha(list DerivableList) common.Hash {
	keybuf := new(bytes.Buffer)
	trie := new(trie.Trie)
	for i := 0; i < list.Len(); i++ {
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i))
		trie.Update(keybuf.Bytes(), list.GetRlp(i))
	}
	return trie.Hash()
}

func DeriveShaFromBMT(list DerivableList) common.Hash {
	var level [][]byte
	for i := 0; i < list.Len(); i++ {
		level = append(level, list.GetRlp(i))
	}
	return getBinaryMerkleRoot(level)
}

// TODO : needs to be refactored because of performance issue.
func getBinaryMerkleRoot(level [][]byte) common.Hash {
	if len(level) == 1 {
		root := common.BytesToHash(level[0])
		return root
	}
	var nextlevel = make([][]byte, (len(level)+1)/2)
	var i int

	for i = 0; i+1 < len(level); i += 2 {
		hash := crypto.Keccak256(level[i], level[i+1])
		nextlevel[i/2] = hash
	}

	if len(level)%2 == 1 {
		nextlevel[i/2] = crypto.Keccak256(level[len(level)-1], level[len(level)-1])
	}

	return getBinaryMerkleRoot(nextlevel)
}

func GetMerkleProof(list DerivableList, index int) []common.Hash {
	var proof []common.Hash
	var leafLevel []common.Hash
	var tree [][]common.Hash
	var depth int

	// convert value of leaf nodes to hash.
	for i := 0; i < list.Len(); i++ {
		leafLevel = append(leafLevel, common.BytesToHash(crypto.Keccak256(list.GetRlp(i))))
	}

	tree = append(tree, leafLevel)
	createTree := func(level []common.Hash) {
		// length == 1 체크 부분은 빼도 될듯
		if len(level) == 1 {
			return
		}
		var nextLevel = make([]common.Hash, (len(level)+1)/2)
		var i int

		for i = 0; i+1 < len(level); i += 2 {
			hash := common.BytesToHash(crypto.Keccak256(level[i].Bytes(), level[i+1].Bytes()))
			nextLevel[i/2] = hash
		}

		if len(level)%2 == 1 {
			nextLevel[i/2] = common.BytesToHash(crypto.Keccak256(level[len(level)-1].Bytes(), level[len(level)-1].Bytes()))
		}

		tree = append(tree, nextLevel)
	}

	// create merkle tree first
	for depth = 0; len(tree[depth]) > 1 ; depth++ {
		createTree(tree[depth])
	}

	nodeIndex := index

	// create merkle proof
	for i := 0; i < depth; i++ {
		if nodeIndex % 2 == 0 {
			siblingIndex := nodeIndex + 1
			proof = append(proof, tree[i][siblingIndex])
		} else {
			siblingIndex := nodeIndex - 1
			proof = append(proof, tree[i][siblingIndex])
		}
		nodeIndex = nodeIndex / 2
	}

	return proof
}

