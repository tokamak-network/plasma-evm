package miner

import (
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/contract"
)

type NRBEpochCompleted struct{}

type ORBEpochCompleted struct{}

type EpochPrepared struct {
	Payload *contract.RootChainEpochPrepared
}
