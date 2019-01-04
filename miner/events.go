package miner

import (
	"math/big"
)

type LastFinalizedBlock struct {
	Number *big.Int
}

type CurrentFork struct {
	Number *big.Int
}
