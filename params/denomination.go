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

import "math/big"

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
	return new(big.Int).Mul(big.NewInt(int64(v*GWei)), big.NewInt(Ether))
}

func ToEtherBigInt(v float64) *big.Int {
	return new(big.Int).Mul(big.NewInt(int64(v*GWei)), big.NewInt(GWei))
}

func ToGWeiBigInt(v float64) *big.Int {
	return big.NewInt(int64(v * GWei))
}

func ToRayFloat64(v *big.Int) float64 {
	b := new(big.Int).Div(v, big.NewInt(GWei))
	b = new(big.Int).Div(b, big.NewInt(GWei))
	return float64(b.Uint64()) / GWei
}

func ToEtherFloat64(v *big.Int) float64 {
	b := new(big.Int).Div(v, big.NewInt(GWei))
	return float64(b.Uint64()) / GWei
}

func ToGWeiFloat64(v *big.Int) float64 {
	return float64(v.Uint64()) / GWei
}
