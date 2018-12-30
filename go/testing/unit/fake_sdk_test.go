package unit

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type foo struct {
	bar string
}

func TestMockHandler_SdkStateReadBytesByAddress_NoValue(t *testing.T) {
	s := aFakeSdk()
	require.Zero(t, s.SdkStateReadBytesByAddress(0, 0, AnAddress()), "read from uninitialized address did not return zero")
}

func TestMockHandler_SdkStateReadBytesByAddress_Success(t *testing.T) {
	s := aFakeSdk()
	a := AnAddress()
	v := []byte{42}
	s.SdkStateWriteBytesByAddress(0, 0, a, v)
	require.Equal(t, v, s.SdkStateReadBytesByAddress(0, 0, a), "read from initialized address did not return expected value")
}

func TestMockHandler_SdkAddressGetCallerAddress(t *testing.T) {
	caller := AnAddress()
	s := aFakeSdkFor([]byte{}, caller)
	require.Equal(t, caller, s.SdkAddressGetCallerAddress(0, 0))
}

func TestMockHandler_SdkAddressGetSignerAddress(t *testing.T) {
	signer := AnAddress()
	s := aFakeSdkFor(signer, []byte{})
	require.Equal(t, signer, s.SdkAddressGetSignerAddress(0, 0))
}

func TestMockHandler_SdkEthereumCallMethod_NotStubbed(t *testing.T) {
	s := aFakeSdk()
	require.Panics(t, func() {
		s.SdkEthereumCallMethod(0, 0, "a", "b", "c", "d", nil)
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
		s.SdkEthereumCallMethod(0, 0, address, abi, methodName, &out, 1, 2)
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
	s.SdkEthereumCallMethod(0, 0, address, abi, methodName, &out, 1, 2)

	require.Equal(t, out.bar, "baz", "did not get expected value from stubbed method")
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEthereumGetTransactionLog_NotStubbed(t *testing.T) {
	s := aFakeSdk()
	require.Panics(t, func() {
		s.SdkEthereumGetTransactionLog(0, 0, "a", "b", "c", "d", nil)
	}, "call to unstubbed method did not panic")
}

func TestMockHandler_SdkEthereumGetTransactionLog_PartialMatch(t *testing.T) {
	s := aFakeSdk()
	address := "a"
	abi := "b"
	txHash := "c"
	s.MockEthereumLog(address, abi, txHash, "e1", func(out interface{}) {
		out.(*foo).bar = "baz"
	})
	require.Panics(t, func() {
		var out foo
		s.SdkEthereumGetTransactionLog(0, 0, address, abi, txHash, "e2", &out)
	}, "call to partially stubbed method did not panic")
	require.Panics(t, func() { s.VerifyMocks() }, "missing call to ethereum should have failed verify")
}

func TestMockHandler_SdkEthereumGetTransactionLog_Success(t *testing.T) {
	s := aFakeSdk()
	address := "a"
	abi := "b"
	txHash := "c"
	eventName := "e"
	s.MockEthereumLog(address, abi, txHash, eventName, func(out interface{}) {
		out.(*foo).bar = "baz"
	})

	var out foo
	s.SdkEthereumGetTransactionLog(0, 0, address, abi, txHash, eventName, &out)
	require.Equal(t, out.bar, "baz", "did not get expected value from stubbed method")
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkServiceCallMethod_Unstubbed(t *testing.T) {
	s := aFakeSdk()

	require.Panics(t, func() {
		s.SdkServiceCallMethod(0, 0, "a", "b", "c", 1)
	}, "unstubbed method call did not panic")
}

func TestMockHandler_SdkServiceCallMethod_Partial(t *testing.T) {
	s := aFakeSdk()
	serviceName := "a"
	methodName := "b"

	s.MockServiceCallMethod(serviceName, methodName, nil, "d")

	require.Panics(t, func() {
		s.SdkServiceCallMethod(0, 0, serviceName, methodName, "c", 1)
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

	require.Equal(t, out, s.SdkServiceCallMethod(0, 0, serviceName, methodName, arg1, arg2))
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEventsEmitEvent_Unstubbed(t *testing.T) {
	s := aFakeSdk()

	require.Panics(t, func() {
		s.SdkEventsEmitEvent(0, 0, func() {}, 1)
	}, "unstubbed event emit did not panic")
}

func TestMockHandler_SdkEventsEmitEvent_Success(t *testing.T) {
	s := aFakeSdk()

	f := func(i int, s string) {}
	arg1 := 1
	arg2 := "c"
	s.MockEmitEvent(f, arg1, arg2)

	require.NotPanics(t, func() { s.SdkEventsEmitEvent(0, 0, f, arg1, arg2) })
	require.NotPanics(t, func() { s.VerifyMocks() })
}

func TestMockHandler_SdkEventsEmitEvent_Partial(t *testing.T) {
	s := aFakeSdk()

	f := func(i int, s string) {}
	arg1 := 1
	arg2 := "c"
	s.MockEmitEvent(f, arg1, arg2)

	require.Panics(t, func() {
		s.SdkEventsEmitEvent(0, 0, f, arg1, "d")
	}, "partially stubbed event emit did not panic")
	require.Panics(t, func() { s.VerifyMocks() }, "missing call to emit should have failed verify")
}
