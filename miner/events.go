package miner

import (
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"math/big"
)

type NRBEpochCompleted struct{}

type ORBEpochCompleted struct{}

type EpochPrepared struct {
	Payload *rootchain.RootChainEpochPrepared
}

type LastFinalizedBlock struct {
	Number *big.Int
}

type CurrentFork struct {
	Number *big.Int
}
