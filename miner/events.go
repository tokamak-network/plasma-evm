package miner

import (
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/contract"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"math/big"
)

type NRBEpochCompleted struct{}

type ORBEpochCompleted struct{}

type EpochPrepared struct {
	Payload *contract.RootChainEpochPrepared
}

type BlockMined struct {
	IsRequest   bool
	BlockNumber *big.Int
	Remaining   *big.Int
	Header      *types.Header
}
