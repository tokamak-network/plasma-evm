package tx

import (
	"math/big"
	"time"

	"github.com/Onther-Tech/plasma-evm/params"
)

var DefaultConfig = &Config{
	MinGasPrice: new(big.Int).SetInt64(1 * params.GWei),
	MaxGasPrice: new(big.Int).SetInt64(100 * params.GWei),
	Resubmit:    10 * time.Second,
	ChainId:     new(big.Int).SetInt64(1),
}

type Config struct {
	MinGasPrice *big.Int
	MaxGasPrice *big.Int
	ChainId     *big.Int
	Resubmit    time.Duration
}
