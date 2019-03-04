package unit

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"github.com/pkg/errors"
	"math/rand"
	"reflect"
)

type stateMap map[string][]byte

var EXAMPLE_CONTEXT_ID = []byte{0x43}

type Mockery interface {
	MockEthereumGetBlockNumber(block int)                                                                                                           // TODO get return values
	MockEthereumLog(ethTxHash string, abiJson string, eventName string, outEthBlockNumber int, outEthTxIndex int, outMutator func(out interface{})) // TODO get return values
	MockEthereumCallMethod(address string, abiJson string, methodName string, outMutator func(out interface{}), args ...interface{})
	MockEthereumCallMethodAtBlock(blockNumber uint64, address string, abiJson string, methodName string, outMutator func(out interface{}), args ...interface{})
	MockServiceCallMethod(serviceName string, methodName string, out []interface{}, args ...interface{})
	MockEmitEvent(eventFunctionSignature interface{}, args ...interface{})
	VerifyMocks()
}

type ethereumStub struct {
	key         []interface{}
	outMutator  func(interface{})
	blockHeight uint64
	txIndex     uint32
	satisfied   bool
}

type eventStub struct {
	key       []interface{}
	satisfied bool
}

type serviceStub struct {
	key       []interface{}
	out       []interface{}
	satisfied bool
}

type mockHandler struct {
	signerAddress []byte
	callerAddress []byte

	state         stateMap
	ethereumStubs []*ethereumStub
	serviceStubs  []*serviceStub
	eventStubs    []*eventStub
}

func (m *mockHandler) SdkStateReadBytes(ctx context.ContextId, permissionScope context.PermissionScope, key []byte) []byte {
	return m.state[hex.EncodeToString(key)]
}

func (m *mockHandler) SdkStateWriteBytes(ctx context.ContextId, permissionScope context.PermissionScope, key []byte, value []byte) {
	m.state[hex.EncodeToString(key)] = value
}

func (m *mockHandler) SdkServiceCallMethod(ctx context.ContextId, permissionScope context.PermissionScope, serviceName string, methodName string, args ...interface{}) []interface{} {
	var key []interface{}
	key = append(key, serviceName, methodName)
	key = append(key, args...)

	for _, stub := range m.serviceStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			return stub.out
		}
	}

	panic(errors.Errorf("No service call stubbed for service %s, method name %s, args %+v", serviceName, methodName, args))
}

func (m *mockHandler) SdkEthereumCallMethod(ctx context.ContextId, permissionScope context.PermissionScope, ethContractAddress string, jsonAbi string, ethBlockNumber uint64, methodName string, out interface{}, args ...interface{}) {
	var key []interface{}
	key = append(key, ethContractAddress, jsonAbi, methodName, ethBlockNumber)
	key = append(key, args...)

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			stub.outMutator(out)
			return
		}
	}

	panic(errors.Errorf("No Ethereum call stubbed for address %s, jsonAbi %s, method name %s, block number %d, args %+v", ethContractAddress, jsonAbi, methodName, ethBlockNumber, args))
}

func (m *mockHandler) SdkEthereumGetBlockNumber(ctx context.ContextId, permissionScope context.PermissionScope) (ethBlockNumber uint64) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockNumber")

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			return stub.blockHeight
		}
	}

	panic(errors.Errorf("No Ethereum call stubbed for GetBlockNumber"))
}

func (m *mockHandler) SdkEthereumGetTransactionLog(ctx context.ContextId, permissionScope context.PermissionScope, ethTxHash string, jsonAbi string, eventName string, out interface{}) (ethBlockNumber uint64, ethTxIndex uint32) {
	var key []interface{}
	key = append(key, ethTxHash, jsonAbi, eventName)

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			stub.outMutator(out)
			return stub.blockHeight, stub.txIndex
		}
	}
	panic(errors.Errorf("No Ethereum logs stubbed for jsonAbi %s, txHash %s, event name %s", jsonAbi, ethTxHash, eventName))
}

func (m *mockHandler) SdkAddressGetSignerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	return m.signerAddress
}

func (m *mockHandler) SdkAddressGetCallerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	return m.callerAddress
}

func (m *mockHandler) SdkAddressGetOwnAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	panic("Not implemented")
	return []byte{}
}

func (m *mockHandler) SdkAddressGetContractAddress(ctx context.ContextId, permissionScope context.PermissionScope, contractName string) []byte {
	panic("Not implemented")
	return []byte{}
}

func (m *mockHandler) SdkEventsEmitEvent(ctx context.ContextId, permissionScope context.PermissionScope, eventFunctionSignature interface{}, args ...interface{}) {
	var key []interface{}
	key = append(key, args...)

	for _, stub := range m.eventStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			return
		}
	}

	panic(errors.Errorf("No Emit Event stubbed for func %s, arguments %v", eventFunctionSignature, args))
}

func (m *mockHandler) SdkEnvGetBlockHeight(ctx context.ContextId, permissionScope context.PermissionScope) uint64 {
	panic("Not implemented")
	return 0
}

func (m *mockHandler) SdkEnvGetBlockTimestamp(ctx context.ContextId, permissionScope context.PermissionScope) uint64 {
	panic("Not implemented")
	return 0
}

func (m *mockHandler) MockEthereumLog(ethTxHash string, jsonAbi string, eventName string, outEthBlockNumber int, outEthTxIndex int, outMutator func(out interface{})) {
	var key []interface{}
	key = append(key, ethTxHash, jsonAbi, eventName)
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: outMutator, blockHeight: uint64(outEthBlockNumber), txIndex: uint32(outEthTxIndex)})
}

func (m *mockHandler) MockEthereumCallMethod(address string, jsonAbi string, methodName string, outMutator func(out interface{}), args ...interface{}) {
	var key []interface{}
	key = append(key, address, jsonAbi, methodName, uint64(0))
	key = append(key, args...)
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: outMutator})
}

func (m *mockHandler) MockEthereumCallMethodAtBlock(blockNumber uint64, address string, jsonAbi string, methodName string, outMutator func(out interface{}), args ...interface{}) {
	var key []interface{}
	key = append(key, address, jsonAbi, methodName, blockNumber)
	key = append(key, args...)
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: outMutator})
}

func (m *mockHandler) MockEthereumGetBlockNumber(block int) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockNumber")
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: nil, blockHeight: uint64(block)})
}

func (m *mockHandler) MockServiceCallMethod(serviceName string, methodName string, out []interface{}, args ...interface{}) {
	var key []interface{}
	key = append(key, serviceName, methodName)
	key = append(key, args...)
	m.serviceStubs = append(m.serviceStubs, &serviceStub{key: key, out: out})
}

func (m *mockHandler) MockEmitEvent(eventFunctionSignature interface{}, args ...interface{}) {
	var key []interface{}
	key = append(key, args...)
	m.eventStubs = append(m.eventStubs, &eventStub{key: key})
}

func (m *mockHandler) VerifyMocks() {
	for _, stub := range m.eventStubs {
		if !stub.satisfied {
			panic(errors.Errorf("emit event mock set but not called"))
		}
	}
	for _, stub := range m.serviceStubs {
		if !stub.satisfied {
			panic(errors.Errorf("service call mock set but not called"))
		}
	}
	for _, stub := range m.ethereumStubs {
		if !stub.satisfied {
			panic(errors.Errorf("ethereum mock set but not called"))
		}
	}
}

func InSystemScope(signerAddress []byte, callerAddress []byte, f func(mockery Mockery)) {
	inScope(signerAddress, callerAddress, context.PERMISSION_SCOPE_SYSTEM, f)
}

func InServiceScope(signerAddress []byte, callerAddress []byte, f func(mockery Mockery)) {
	inScope(signerAddress, callerAddress, context.PERMISSION_SCOPE_SERVICE, f)
}

func inScope(signerAddress []byte, callerAddress []byte, scope context.PermissionScope, f func(mockery Mockery)) {
	if signerAddress == nil {
		signerAddress = AnAddress()
	}
	if callerAddress == nil {
		callerAddress = AnAddress()
	}
	handler := aFakeSdkFor(signerAddress, callerAddress)
	cid := context.ContextId([]byte{byte(rand.Int())})
	context.PushContext(cid, handler, scope)
	f(handler)
	handler.VerifyMocks()
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
