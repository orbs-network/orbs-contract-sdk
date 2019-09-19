// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package main

import (
	. "github.com/orbs-network/orbs-contract-sdk/go/sdk/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadValueFromLog(t *testing.T) {

	address := "a"
	abi := "b"
	txid := "c"
	eventName := "d"
	outEthBlockNumber := 1
	outhEthTxIndex := 0

	InServiceScope(nil, nil, func(m Mockery) {
		m.MockEthereumLog(address, abi, txid, eventName, outEthBlockNumber, outhEthTxIndex, func(out interface{}) {
			Value = "foo"
		})

		v := readValueFromLog(address, abi, txid, eventName)

		require.Equal(t, "foo", v, "did not get expected Value from log")
	})
}

func TestCallEthereumMethod(t *testing.T) {
	address := "a"
	abi := "b"
	methodName := "c"
	arg1 := 1
	arg2 := true

	InServiceScope(nil, nil, func(m Mockery) {

		m.MockEthereumCallMethod(address, abi, methodName, func(out interface{}) {
			Value = "bar"
		}, arg1, arg2)

		v := callEthereumMethod(address, abi, methodName, arg1, arg2)

		require.Equal(t, "bar", v, "did not get expected Value from method call")

	})
}
