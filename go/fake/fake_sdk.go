package fake

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"github.com/pkg/errors"
	"math/rand"
	"reflect"
)

type stateMap map[string][]byte

type Mockery interface {
	MockEthereumLog(address string, abiJson string, ethTxHash string, eventName string, outMutator func(out interface{}))
	MockEthereumCallMethod(address string, abiJson string, methodName string, outMutator func(out interface{}), args ...interface{})
	MockServiceCallMethod(serviceName string, methodName string, out []interface{}, args ...interface{})
}

type ethereumStub struct {
	key        []interface{}
	outMutator func(interface{})
}

type serviceStub struct {
	key []interface{}
	out []interface{}
}

type mockHandler struct {
	signerAddress []byte
	callerAddress []byte

	state         stateMap
	ethereumStubs []ethereumStub
	serviceStubs  []serviceStub
}

func (m *mockHandler) SdkStateReadBytesByAddress(ctx context.ContextId, permissionScope context.PermissionScope, address []byte) []byte {
	return m.state[hex.EncodeToString(address)]
}

func (m *mockHandler) SdkStateWriteBytesByAddress(ctx context.ContextId, permissionScope context.PermissionScope, address []byte, value []byte) {
	m.state[hex.EncodeToString(address)] = value
}

func (m *mockHandler) SdkServiceCallMethod(ctx context.ContextId, permissionScope context.PermissionScope, serviceName string, methodName string, args ...interface{}) []interface{} {
	var key []interface{}
	key = append(key, serviceName, methodName)
	key = append(key, args...)

	for _, stub := range m.serviceStubs {
		if keyEquals(stub.key, key) {
			return stub.out
		}
	}

	panic(errors.Errorf("No service call stubbed for service %s, method name %s, args %+v", serviceName, methodName, args))
}

func (m *mockHandler) SdkEthereumCallMethod(ctx context.ContextId, permissionScope context.PermissionScope, contractAddress string, jsonAbi string, methodName string, out interface{}, args ...interface{}) {
	var key []interface{}
	key = append(key, contractAddress, jsonAbi, methodName)
	key = append(key, args...)

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.outMutator(out)
			return
		}
	}

	panic(errors.Errorf("No Ethereum call stubbed for address %s, jsonAbi %s, method name %s, args %+v", contractAddress, jsonAbi, methodName, args))
}

func (m *mockHandler) SdkEthereumGetTransactionLog(ctx context.ContextId, permissionScope context.PermissionScope, contractAddress string, jsonAbi string, ethTransactionId string, eventName string, out interface{}) {
	var key []interface{}
	key = append(key, contractAddress, jsonAbi, ethTransactionId, eventName)

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.outMutator(out)
			return
		}
	}
	panic(errors.Errorf("No Ethereum logs stubbed for address %s, jsonAbi %s, txHash %s, event name %s", contractAddress, jsonAbi, ethTransactionId, eventName))

}

func (m *mockHandler) SdkAddressGetSignerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	return m.signerAddress
}

func (m *mockHandler) SdkAddressGetCallerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	return m.callerAddress
}

func (m *mockHandler) MockEthereumLog(address string, abiJson string, ethTxHash string, eventName string, outMutator func(out interface{})) {
	var key []interface{}
	key = append(key, address, abiJson, ethTxHash, eventName)
	m.ethereumStubs = append(m.ethereumStubs, ethereumStub{key: key, outMutator: outMutator})
}

func (m *mockHandler) MockEthereumCallMethod(address string, abiJson string, methodName string, outMutator func(out interface{}), args ...interface{}) {
	var key []interface{}
	key = append(key, address, abiJson, methodName)
	key = append(key, args...)
	m.ethereumStubs = append(m.ethereumStubs, ethereumStub{key: key, outMutator: outMutator})
}

func (m *mockHandler) MockServiceCallMethod(serviceName string, methodName string, out []interface{}, args ...interface{}) {
	var key []interface{}
	key = append(key, serviceName, methodName)
	key = append(key, args...)
	m.serviceStubs = append(m.serviceStubs, serviceStub{key: key, out: out})
}

func InSystemScope(signerAddress []byte, f func(mockery Mockery)) {
	inScope(signerAddress, []byte{}, context.PERMISSION_SCOPE_SYSTEM, f)
}

func InServiceScope(callerAddress []byte, f func(mockery Mockery)) {
	inScope([]byte{}, callerAddress, context.PERMISSION_SCOPE_SERVICE, f)
}

func inScope(signerAddress []byte, callerAddress []byte, scope context.PermissionScope, f func(mockery Mockery)) {
	handler := aFakeSdkFor(signerAddress, callerAddress)
	cid := context.ContextId(43)
	context.PushContext(cid, handler, scope)
	f(handler)
	context.PopContext(cid)
}

func aFakeSdkFor(signerAddress []byte, callerAddress []byte) *mockHandler {
	handler := &mockHandler{
		signerAddress: signerAddress,
		callerAddress: callerAddress,
		state:         make(stateMap),
	}
	return handler
}

func aFakeSdk() *mockHandler {
	return aFakeSdkFor([]byte{}, []byte{})
}

func AnAddress() (address []byte) {
	address = make([]byte, 20)
	rand.Read(address)
	return
}

func keyEquals(k1 []interface{}, k2 []interface{}) bool {
	return reflect.DeepEqual(k1, k2)
}
