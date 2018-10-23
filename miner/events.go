package miner

import (
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
)

type NRBEpochCompleted struct{}

type ORBEpochCompleted struct{}

type EpochPrepared struct {
	Payload *rootchain.RootChainEpochPrepared
}
