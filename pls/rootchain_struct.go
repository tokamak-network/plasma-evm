package pls

import (
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
)

// TODO: update solidity@0.6.0 and remove below structs
// NOTE: abigen sometimes mixes struct nubmering. it should be manually checked.
// Structs for PlasmaEpoch and PlasmaBlock
type PlasmaEpoch rootchain.Struct3

func newPlasmaEpoch(e rootchain.Struct3) *PlasmaEpoch {
	e2 := PlasmaEpoch(e)
	return &e2
}

type PlasmaBlock rootchain.Struct0

func newPlasmaBlock(b rootchain.Struct0) *PlasmaBlock {
	b2 := PlasmaBlock(b)
	return &b2
}
