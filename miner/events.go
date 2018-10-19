package miner

import (
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/contract"
	"math/big"
	"github.com/Onther-Tech/plasma-evm/core/types"
)

type NRBEpochCompleted struct{}

type ORBEpochCompleted struct{}

type EpochPrepared struct {
	Payload *contract.RootChainEpochPrepared
}

type BlockMined struct {
	Payload *blockMined
}

type blockMined struct {
	IsRequest bool
	BlockNumber *big.Int
	Remaining *big.Int
	Header *types.Header
}