// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package safeuint64

import "math"

func Add(x uint64, y uint64) uint64 {
	if y > math.MaxUint64-x {
		panic("integer overflow on add")
	}
	return x + y
}

func Sub(x uint64, y uint64) uint64 {
	if x < y {
		panic("integer overflow on sub")
	}
	return x - y
}

func Mul(x uint64, y uint64) uint64 {
	if x == 0 || y == 0 {
		return 0
	}
	if y > math.MaxUint64/x {
		panic("integer overflow on mul")
	}
	return x * y
}

func Div(x uint64, y uint64) uint64 {
	if y == 0 {
		panic("division by zero")
	}
	return x / y
}

func Mod(x uint64, y uint64) uint64 {
	if y == 0 {
		panic("division by zero")
	}
	return x % y
}
