// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package safeuint256

import (
	"math/big"
)

var zero = big.NewInt(0)
var maxUint256 = big.NewInt(0).Sub(big.NewInt(0).Lsh(big.NewInt(1), 256), big.NewInt(1))

func Validate(n *big.Int) {
	if n.Cmp(zero) < 0 {
		panic("negative result")
	}
	if n.Cmp(maxUint256) > 0 {
		panic("result overflow")
	}
}

func Add(x *big.Int, y *big.Int) *big.Int {
	res := big.NewInt(0).Add(x, y)
	Validate(res)
	return res
}

func Sub(x *big.Int, y *big.Int) *big.Int {
	res := big.NewInt(0).Sub(x, y)
	Validate(res)
	return res
}

func Mul(x *big.Int, y *big.Int) *big.Int {
	res := big.NewInt(0).Mul(x, y)
	Validate(res)
	return res
}

func Div(x *big.Int, y *big.Int) *big.Int {
	if y.Cmp(zero) == 0 {
		panic("division by zero")
	}
	res := big.NewInt(0).Div(x, y)
	Validate(res)
	return res
}

func Mod(x *big.Int, y *big.Int) *big.Int {
	if y.Cmp(zero) == 0 {
		panic("division by zero")
	}
	res := big.NewInt(0).Mod(x, y)
	Validate(res)
	return res
}
