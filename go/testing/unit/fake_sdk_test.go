// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package unit

import (
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/state"
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
	s.MockEthereumCallMethod(address, abi, methodName, func(out interface{}) {
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
	bh, txidx := s.SdkEthereumGetTransactionLog(EXAMPLE_CONTEXT_ID, 0, address, abi, txHash, eventName, &out)
	require.Equal(t, out.bar, "baz", "did not get expected value from stubbed method")
	require.EqualValues(t, 17, bh, "block height should be 17")
	require.EqualValues(t, 42, txidx, "transaction index should be 42")

	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEthereumGetBlock_NotStubbed(t *testing.T) {
	s := aFakeSdk()
	require.Panics(t, func() {
		s.SdkEthereumGetBlockNumber(EXAMPLE_CONTEXT_ID, 0)
	}, "call to unstubbed method did not panic")
}

func TestMockHandler_SdkEthereumGetBlock_Success(t *testing.T) {
	s := aFakeSdk()
	blockNumber := 7
	s.MockEthereumGetBlockNumber(blockNumber)

	out := s.SdkEthereumGetBlockNumber(EXAMPLE_CONTEXT_ID, 0)
	require.EqualValues(t, blockNumber, out, "block number should be 7")
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEthereumGetBlockByTime_NotStubbed(t *testing.T) {
	s := aFakeSdk()
	require.Panics(t, func() {
		s.SdkEthereumGetBlockNumberByTime(EXAMPLE_CONTEXT_ID, 0, 7)
	}, "call to unstubbed method did not panic")
}

func TestMockHandler_SdkEthereumGetBlockByTime_PartialStubbed(t *testing.T) {
	s := aFakeSdk()
	blockNumber := 7
	blockTime := 9
	s.MockEthereumGetBlockNumberByTime(blockNumber, blockTime)

	require.Panics(t, func() {
		s.SdkEthereumGetBlockNumberByTime(EXAMPLE_CONTEXT_ID, 0, uint64(12))
	}, "call to partially stubbed method did not panic")
	require.Panics(t, func() { s.VerifyMocks() }, "missing call to ethereum should have failed verify")
}

func TestMockHandler_SdkEthereumGetBlockByTime_Success(t *testing.T) {
	s := aFakeSdk()
	blockNumber := 7
	blockTime := 9
	s.MockEthereumGetBlockNumberByTime(blockNumber, blockTime)

	out := s.SdkEthereumGetBlockNumberByTime(EXAMPLE_CONTEXT_ID, 0, uint64(blockTime))
	require.EqualValues(t, blockNumber, out, "block number should be 7")
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEthereumGetTime_NotStubbed(t *testing.T) {
	s := aFakeSdk()
	require.Panics(t, func() {
		s.SdkEthereumGetBlockTime(EXAMPLE_CONTEXT_ID, 0)
	}, "call to unstubbed method did not panic")
}

func TestMockHandler_SdkEthereumGetTime_Success(t *testing.T) {
	s := aFakeSdk()
	blockTime := 7
	s.MockEthereumGetBlockTime(blockTime)

	out := s.SdkEthereumGetBlockTime(EXAMPLE_CONTEXT_ID, 0)
	require.EqualValues(t, blockTime, out, "block time should be 7")
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEthereumGetTimeByBlock_NotStubbed(t *testing.T) {
	s := aFakeSdk()
	require.Panics(t, func() {
		s.SdkEthereumGetBlockTimeByNumber(EXAMPLE_CONTEXT_ID, 0, 7)
	}, "call to unstubbed method did not panic")
}

func TestMockHandler_SdkEthereumGetTimeByBlock_PartialStubbed(t *testing.T) {
	s := aFakeSdk()
	blockNumber := 7
	blockTime := 9
	s.MockEthereumGetBlockTimeByNumber(blockNumber, blockTime)

	require.Panics(t, func() {
		s.SdkEthereumGetBlockTimeByNumber(EXAMPLE_CONTEXT_ID, 0, uint64(12))
	}, "call to partially stubbed method did not panic")
	require.Panics(t, func() { s.VerifyMocks() }, "missing call to ethereum should have failed verify")
}

func TestMockHandler_SdkEthereumGetTimeByBlock_Success(t *testing.T) {
	s := aFakeSdk()
	blockNumber := 7
	blockTime := 9
	s.MockEthereumGetBlockTimeByNumber(blockNumber, blockTime)

	out := s.SdkEthereumGetBlockTimeByNumber(EXAMPLE_CONTEXT_ID, 0, uint64(blockNumber))
	require.EqualValues(t, blockTime, out, "block time should be 9")
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

func TestMockHandler_SdkEnvGetBlockHeight_Unset(t *testing.T) {
	s := aFakeSdk()

	require.Panics(t, func() {
		s.SdkEnvGetBlockHeight(EXAMPLE_CONTEXT_ID, 0)
	}, "when not set block height fails")
}

func TestMockHandler_SdkEnvGetBlockHeight_Success(t *testing.T) {
	s := aFakeSdk()

	s.MockEnvBlockHeight(5)

	b := s.SdkEnvGetBlockHeight(EXAMPLE_CONTEXT_ID, 0)
	require.EqualValues(t, 5, b)
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEnvGetBlockTimestamp_Success(t *testing.T) {
	s := aFakeSdk()

	timestamp := time.Unix(0, int64(s.SdkEnvGetBlockTimestamp(EXAMPLE_CONTEXT_ID, 0)))
	require.WithinDuration(t, time.Now(), timestamp, 5*time.Millisecond)
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEnvGetBlockProposer_Unset(t *testing.T) {
	s := aFakeSdk()

	require.Panics(t, func() {
		s.SdkEnvGetBlockProposerAddress(EXAMPLE_CONTEXT_ID, 0)
	}, "when not set block proposer fails")
}

func TestMockHandler_SdkEnvGetBlockProposer_Success(t *testing.T) {
	s := aFakeSdk()

	addr := []byte{0x01}
	s.MockEnvBlockProposerAddress(addr)

	res := s.SdkEnvGetBlockProposerAddress(EXAMPLE_CONTEXT_ID, 0)
	require.EqualValues(t, addr, res)
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

func Test_InScope(t *testing.T) {
	caller := AnAddress()

	emptyStateDiffs, zeroReads, zeroWrites := inScope(nil, caller, context.PERMISSION_SCOPE_SERVICE, func(mockery Mockery) {
	})
	require.Empty(t, emptyStateDiffs)
	require.Zero(t, zeroReads, zeroWrites)

	stateDiffs, singleRead, singleWrite := inScope(nil, caller, context.PERMISSION_SCOPE_SERVICE, func(mockery Mockery) {
		state.ReadString([]byte("some key"))
		state.WriteString([]byte("hello"), "world")
	})

	require.EqualValues(t, 1, singleRead)
	require.EqualValues(t, 1, singleWrite)
	require.Len(t, stateDiffs, 1)
	require.EqualValues(t, &StateDiff{
		Key:   []byte("hello"),
		Value: []byte("world"),
	}, stateDiffs[0])

	orderedStateDiffs, multipleReads, multipleWrites := inScope(nil, caller, context.PERMISSION_SCOPE_SERVICE, func(mockery Mockery) {
		state.WriteString([]byte("1974"), "Diamond Dogs")
		state.WriteString([]byte("1976"), "Station to Station")
		state.WriteString([]byte("1969"), "Hunky Dory")
		state.WriteString([]byte("1969"), "Space Oddity")

		state.ReadString([]byte("1976"))
		state.ReadString([]byte("1969"))
	})

	require.EqualValues(t, 2, multipleReads)
	require.EqualValues(t, 4, multipleWrites)
	require.Len(t, orderedStateDiffs, 3)
	require.EqualValues(t, []byte("1974"), orderedStateDiffs[0].Key)
	require.EqualValues(t, []byte("1976"), orderedStateDiffs[1].Key)
	require.EqualValues(t, []byte("1969"), orderedStateDiffs[2].Key)
}
