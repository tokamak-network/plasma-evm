// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/Onther-Tech/plasma-evm/log"
)

// These are the multipliers for ether denominations.
// Example: To get the wei value of an amount in 'gwei', use
//
//    new(big.Int).Mul(value, big.NewInt(params.GWei))
//
const (
	Wei   = 1
	GWei  = 1e9
	Ether = 1e18
)

func ToRayBigInt(v float64) *big.Int {
	str := strconv.FormatFloat(v, 'f', -1, 64)
	return parseFloatString(str, 27)
}

func ToEtherBigInt(v float64) *big.Int {
	str := strconv.FormatFloat(v, 'f', -1, 64)
	return parseFloatString(str, 18)
}

func ToGWeiBigInt(v float64) *big.Int {
	str := strconv.FormatFloat(v, 'f', -1, 64)
	return parseFloatString(str, 9)
}

func ToRayFloat64(v *big.Int) float64 {
	isNeg := v.Cmp(big.NewInt(0)) < 0

	abs := new(big.Int).Abs(v)

	b := new(big.Int).Div(abs, big.NewInt(GWei))
	b = new(big.Int).Div(b, big.NewInt(GWei))

	if isNeg {
		return -float64(b.Uint64()) / GWei
	}

	return float64(b.Uint64()) / GWei
}

func ToEtherFloat64(v *big.Int) float64 {
	isNeg := v.Cmp(big.NewInt(0)) < 0

	abs := new(big.Int).Abs(v)

	b := new(big.Int).Div(abs, big.NewInt(GWei))

	if isNeg {
		return -float64(b.Uint64()) / GWei
	}

	return float64(b.Uint64()) / GWei
}

func ToGWeiFloat64(v *big.Int) float64 {
	isNeg := v.Cmp(big.NewInt(0)) < 0

	abs := new(big.Int).Abs(v)

	if isNeg {
		return -float64(abs.Uint64()) / GWei
	}

	return float64(abs.Uint64()) / GWei
}

func parseFloatString(str string, decimals int) *big.Int {
	i := strings.Index(str, ".")

	v := str
	n := decimals

	// split string with "."
	if i >= 0 {
		a := str[:i]
		b := str[i+1:]
		n = n - len(b)

		if n < 0 {
			log.Crit("out of decimals precision", "decimals", decimals, "length", len(b))
			return nil
		}

		v = a + b
	}

	v = v + strings.Repeat("0", n)

	bi, ok := big.NewInt(0).SetString(v, 10)

	if !ok {
		log.Crit("failed to parse string", "string", str)
		return nil
	}

	return bi
}

func ParseFloatString(str string, decimals int) *big.Int {
	return parseFloatString(str, decimals)
}
