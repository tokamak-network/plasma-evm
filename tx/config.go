package tx

import (
	"math/big"
	"time"

	"github.com/Onther-Tech/plasma-evm/params"
)

var DefaultConfig = &Config{
	GasPrice:    new(big.Int).SetInt64(10 * params.GWei),
	MinGasPrice: new(big.Int).SetInt64(1 * params.GWei),
	MaxGasPrice: new(big.Int).SetInt64(100 * params.GWei),
	Interval:    10 * time.Second,
	ChainId:     new(big.Int).SetInt64(1),
}

type Config struct {
	GasPrice    *big.Int
	MinGasPrice *big.Int
	MaxGasPrice *big.Int
	ChainId     *big.Int
	Interval    time.Duration
}
