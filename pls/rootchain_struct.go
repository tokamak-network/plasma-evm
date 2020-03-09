package pls

import (
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
)

// TODO: update solidity@0.6.0 and remove below structs
// NOTE: abigen sometimes mixes struct nubmering. it should be manually checked.
// Structs for PlasmaEpoch and PlasmaBlock
type PlasmaEpoch rootchain.Struct2

func newPlasmaEpoch(e rootchain.Struct2) *PlasmaEpoch {
	e2 := PlasmaEpoch(e)
	return &e2
}

type PlasmaBlock rootchain.Struct3

func newPlasmaBlock(b rootchain.Struct3) *PlasmaBlock {
	b2 := PlasmaBlock(b)
	return &b2
}
