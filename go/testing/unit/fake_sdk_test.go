// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package unit

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type foo struct {
	bar string
}

func TestMockHandler_SdkStateReadBytes_NoValue(t *testing.T) {
	s := aFakeSdk()
	require.Zero(t, s.SdkStateReadBytes(EXAMPLE_CONTEXT_ID, 0, AnAddress()), "read from uninitialized address did not return zero")
}

func TestMockHandler_SdkStateReadBytes_Success(t *testing.T) {
	s := aFakeSdk()
	a := AnAddress()
	v := []byte{42}
	s.SdkStateWriteBytes(EXAMPLE_CONTEXT_ID, 0, a, v)
	require.Equal(t, v, s.SdkStateReadBytes(EXAMPLE_CONTEXT_ID, 0, a), "read from initialized address did not return expected value")
}

func TestMockHandler_SdkAddressGetCallerAddress(t *testing.T) {
	caller := AnAddress()
	s := aFakeSdkFor([]byte{}, caller)
	require.Equal(t, caller, s.SdkAddressGetCallerAddress(EXAMPLE_CONTEXT_ID, 0))
}

func TestMockHandler_SdkAddressGetSignerAddress(t *testing.T) {
	signer := AnAddress()
	s := aFakeSdkFor(signer, []byte{})
	require.Equal(t, signer, s.SdkAddressGetSignerAddress(EXAMPLE_CONTEXT_ID, 0))
}

func TestMockHandler_SdkEthereumCallMethod_NotStubbed(t *testing.T) {
	s := aFakeSdk()
	require.Panics(t, func() {
		s.SdkEthereumCallMethod(EXAMPLE_CONTEXT_ID, 0, "a", "b", 1, "c", "d", nil)
	}, "call to unstubbed method did not panic")
}

func TestMockHandler_SdkEthereumCallMethod_PartialMatch(t *testing.T) {
	s := aFakeSdk()
	address := "a"
	abi := "b"
	methodName := "c"
	s.MockEthereumCallMethod(address, abi, methodName, func(out interface{}) {
		out.(*foo).bar = "baz"
	})
	require.Panics(t, func() {
		var out foo
		s.SdkEthereumCallMethod(EXAMPLE_CONTEXT_ID, 0, address, abi, 1, methodName, &out, 1, 2)
	}, "call to partially stubbed method did not panic")
	require.Panics(t, func() { s.VerifyMocks() }, "missing call to ethereum should have failed verify")
}

func TestMockHandler_SdkEthereumCallMethod_Success(t *testing.T) {
	s := aFakeSdk()
	address := "a"
	abi := "b"
	methodName := "c"
	s.MockEthereumCallMethod(address, abi,  methodName, func(out interface{}) {
		out.(*foo).bar = "baz"
	}, 1, 2)

	var out foo
	s.SdkEthereumCallMethod(EXAMPLE_CONTEXT_ID, 0, address, abi, 1, methodName, &out, 1, 2)

	require.Equal(t, out.bar, "baz", "did not get expected value from stubbed method")
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEthereumGetTransactionLog_NotStubbed(t *testing.T) {
	s := aFakeSdk()
	require.Panics(t, func() {
		s.SdkEthereumGetTransactionLog(EXAMPLE_CONTEXT_ID, 0, "a", "b", "c", "d", nil)
	}, "call to unstubbed method did not panic")
}

func TestMockHandler_SdkEthereumGetTransactionLog_PartialMatch(t *testing.T) {
	s := aFakeSdk()
	address := "a"
	abi := "b"
	txHash := "c"
	s.MockEthereumLog(address, abi, txHash, "e1", 17, 42, func(out interface{}) {
		out.(*foo).bar = "baz"
	})
	require.Panics(t, func() {
		var out foo
		s.SdkEthereumGetTransactionLog(EXAMPLE_CONTEXT_ID, 0, address, abi, txHash, "e2", &out)
	}, "call to partially stubbed method did not panic")
	require.Panics(t, func() { s.VerifyMocks() }, "missing call to ethereum should have failed verify")
}

func TestMockHandler_SdkEthereumGetTransactionLog_Success(t *testing.T) {
	s := aFakeSdk()
	address := "a"
	abi := "b"
	txHash := "c"
	eventName := "e"
	s.MockEthereumLog(address, abi, txHash, eventName, 17, 42, func(out interface{}) {
		out.(*foo).bar = "baz"
	})

	var out foo
	bh, txidx :=  s.SdkEthereumGetTransactionLog(EXAMPLE_CONTEXT_ID, 0, address, abi, txHash, eventName, &out)
	require.Equal(t, out.bar, "baz", "did not get expected value from stubbed method")
	require.EqualValues(t, 17, bh, "block height should be 17")
	require.EqualValues(t, 42, txidx, "transaction index should be 42")

	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkServiceCallMethod_Unstubbed(t *testing.T) {
	s := aFakeSdk()

	require.Panics(t, func() {
		s.SdkServiceCallMethod(EXAMPLE_CONTEXT_ID, 0, "a", "b", "c", 1)
	}, "unstubbed method call did not panic")
}

func TestMockHandler_SdkServiceCallMethod_Partial(t *testing.T) {
	s := aFakeSdk()
	serviceName := "a"
	methodName := "b"

	s.MockServiceCallMethod(serviceName, methodName, nil, "d")

	require.Panics(t, func() {
		s.SdkServiceCallMethod(EXAMPLE_CONTEXT_ID, 0, serviceName, methodName, "c", 1)
	}, "partially stubbed method call did not panic")
	require.Panics(t, func() { s.VerifyMocks() }, "missing method should have failed verify")
}

func TestMockHandler_SdkServiceCallMethod_Success(t *testing.T) {
	s := aFakeSdk()
	serviceName := "a"
	methodName := "b"
	arg1 := "c"
	arg2 := 1
	out := []interface{}{true, "z"}

	s.MockServiceCallMethod(serviceName, methodName, out, arg1, arg2)

	require.Equal(t, out, s.SdkServiceCallMethod(EXAMPLE_CONTEXT_ID, 0, serviceName, methodName, arg1, arg2))
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEventsEmitEvent_Unstubbed(t *testing.T) {
	s := aFakeSdk()

	require.Panics(t, func() {
		s.SdkEventsEmitEvent(EXAMPLE_CONTEXT_ID, 0, func() {}, 1)
	}, "unstubbed event emit did not panic")
}

func TestMockHandler_SdkEventsEmitEvent_Success(t *testing.T) {
	s := aFakeSdk()

	f := func(i int, s string) {}
	arg1 := 1
	arg2 := "c"
	s.MockEmitEvent(f, arg1, arg2)

	require.NotPanics(t, func() { s.SdkEventsEmitEvent(EXAMPLE_CONTEXT_ID, 0, f, arg1, arg2) })
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEventsEmitEvent_Partial(t *testing.T) {
	s := aFakeSdk()

	f := func(i int, s string) {}
	arg1 := 1
	arg2 := "c"
	s.MockEmitEvent(f, arg1, arg2)

	require.Panics(t, func() {
		s.SdkEventsEmitEvent(EXAMPLE_CONTEXT_ID, 0, f, arg1, "d")
	}, "partially stubbed event emit did not panic")
	require.Panics(t, func() { s.VerifyMocks() }, "missing call to emit should have failed verify")
}


func TestMockHandler_SdkEnvGetBlockHeight(t *testing.T) {
	s := aFakeSdk()
	b1 := s.SdkEnvGetBlockHeight(EXAMPLE_CONTEXT_ID, 0)
	b2 := s.SdkEnvGetBlockHeight(EXAMPLE_CONTEXT_ID, 0)

	require.EqualValues(t, 1, b1, "first block should be 1")
	require.EqualValues(t, 2, b2, "second block should be 2")
}

func TestMockHandler_SdkEnvGetBlockTimestamp(t *testing.T) {
	s := aFakeSdk()
	ts := s.SdkEnvGetBlockTimestamp(EXAMPLE_CONTEXT_ID, 0)
	require.InDelta(t, uint64(time.Now().UnixNano()), ts, float64(time.Second), "expected current block time to be around 1 second from current time")
}

