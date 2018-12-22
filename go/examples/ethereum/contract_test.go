package main

import (
	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadValueFromLog(t *testing.T) {

	address := "a"
	abi := "b"
	txid := "c"
	eventName := "d"

	InServiceScope(nil, nil, func(m Mockery) {

		m.MockEthereumLog(address, abi, txid, eventName, func(out interface{}) {
			out.(*event).Value = "foo"
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
			out.(*event).Value = "bar"
		}, arg1, arg2)

		v := callEthereumMethod(address, abi, methodName, arg1, arg2)

		require.Equal(t, "bar", v, "did not get expected Value from method call")

	})
}
