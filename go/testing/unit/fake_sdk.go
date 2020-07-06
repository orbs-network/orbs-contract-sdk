// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package unit

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"github.com/pkg/errors"
	"math/rand"
	"reflect"
	"time"
)

type stateMap map[string][]byte

var EXAMPLE_CONTEXT_ID = []byte{0x43}

type Mockery interface {
	MockEthereumGetBlockNumber(block int)
	MockEthereumGetBlockNumberByTime(block int, timestamp int)
	MockEthereumGetBlockTime(timestamp int)
	MockEthereumGetBlockTimeByNumber(block int, timestamp int)
	MockEthereumLog(address string, abiJson string, ethTxHash string, eventName string, outEthBlockNumber int, outEthTxIndex int, outMutator func(out interface{})) // TODO get return valus
	MockEthereumCallMethod(address string, abiJson string, methodName string, outMutator func(out interface{}), args ...interface{})
	MockEthereumCallMethodAtBlock(blockNumber uint64, address string, abiJson string, methodName string, outMutator func(out interface{}), args ...interface{})
	MockServiceCallMethod(serviceName string, methodName string, out []interface{}, args ...interface{})
	MockEnvBlockHeight(height int)
	MockEnvBlockTimestamp(timestamp int)
	MockEnvBlockProposerAddress(addr []byte)
	MockEnvGetBlockCommittee(committee [][]byte)
	MockEnvGetNextBlockCommittee(committee [][]byte)
	MockEmitEvent(eventFunctionSignature interface{}, args ...interface{})
	MockCallContractAddress(name string, value []byte)
	VerifyMocks()
}

type ethereumStub struct {
	key            []interface{}
	outMutator     func(interface{})
	blockHeight    uint64
	blockTimestamp uint64
	txIndex        uint32
	satisfied      bool
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

type addressStub struct {
	key       []interface{}
	out       []byte
	satisfied bool
}

type mockHandler struct {
	signerAddress []byte
	callerAddress []byte

	blockProposerAddress []byte
	blockHeight          uint64
	blockTimestamp       uint64

	state         stateMap
	stateKeyOrder []string
	stateReads    uint64
	stateWrites   uint64

	ethereumStubs []*ethereumStub
	serviceStubs  []*serviceStub
	eventStubs    []*eventStub
	addressStubs  []*addressStub
	committee     [][]byte
	nextCommittee [][]byte
}

type StateDiff struct {
	Key   []byte
	Value []byte
}

func (m *mockHandler) SdkStateReadBytes(ctx context.ContextId, permissionScope context.PermissionScope, key []byte) []byte {
	m.stateReads += 1
	return m.state[hex.EncodeToString(key)]
}

func (m *mockHandler) SdkStateWriteBytes(ctx context.ContextId, permissionScope context.PermissionScope, key []byte, value []byte) {
	m.stateWrites += 1
	hexKey := hex.EncodeToString(key)

	shouldUpdate := true
	for _, key := range m.stateKeyOrder {
		if key == hexKey {
			shouldUpdate = false
		}
	}

	if shouldUpdate {
		m.stateKeyOrder = append(m.stateKeyOrder, hexKey)
	}
	m.state[hexKey] = value
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
	var keyWithoutBlockNumber []interface{}
	key = append(key, ethContractAddress, jsonAbi, methodName, ethBlockNumber)
	key = append(key, args...)
	keyWithoutBlockNumber = append(keyWithoutBlockNumber, ethContractAddress, jsonAbi, methodName)
	keyWithoutBlockNumber = append(keyWithoutBlockNumber, args...)

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) || keyEquals(stub.key, keyWithoutBlockNumber) {
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

func (m *mockHandler) SdkEthereumGetBlockNumberByTime(ctx context.ContextId, permissionScope context.PermissionScope, ethTimestamp uint64) (ethBlockNumber uint64) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockNumberByTime", ethTimestamp)

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			return stub.blockHeight
		}
	}

	panic(errors.Errorf("No Ethereum call stubbed for GetBlockNumberByTime"))
}

func (m *mockHandler) SdkEthereumGetBlockTime(ctx context.ContextId, permissionScope context.PermissionScope) (ethTimestamp uint64) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockTime")

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			return stub.blockTimestamp
		}
	}

	panic(errors.Errorf("No Ethereum call stubbed for GetBlockTime"))
}

func (m *mockHandler) SdkEthereumGetBlockTimeByNumber(ctx context.ContextId, permissionScope context.PermissionScope, ethBlockNumber uint64) (ethTimestamp uint64) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockTimeByNumber", ethBlockNumber)

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			return stub.blockTimestamp
		}
	}

	panic(errors.Errorf("No Ethereum call stubbed for GetBlockNumber"))
}

func (m *mockHandler) SdkEthereumGetTransactionLog(ctx context.ContextId, permissionScope context.PermissionScope, contractAddress string, jsonAbi string, ethTransactionId string, eventName string, out interface{}) (ethBlockNumber uint64, ethTxIndex uint32) {
	var key []interface{}
	key = append(key, contractAddress, jsonAbi, ethTransactionId, eventName)

	for _, stub := range m.ethereumStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			stub.outMutator(out)
			return stub.blockHeight, stub.txIndex
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

func (m *mockHandler) SdkAddressGetOwnAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	panic("Not implemented")
	return []byte{}
}

func (m *mockHandler) SdkAddressGetContractAddress(ctx context.ContextId, permissionScope context.PermissionScope, contractName string) []byte {
	var key []interface{}
	key = append(key, "contract name", contractName)

	for _, stub := range m.addressStubs {
		if keyEquals(stub.key, key) {
			stub.satisfied = true
			return stub.out
		}
	}
	panic(errors.Errorf("No contract address stubed for name %s", contractName))
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

	panic(errors.Errorf("No Emit Event stubbed for func %s, arguments %v", reflect.TypeOf(eventFunctionSignature).String(), args))
}

func (m *mockHandler) SdkEnvGetBlockHeight(ctx context.ContextId, permissionScope context.PermissionScope) uint64 {
	if m.blockHeight == 0 {
		panic(errors.Errorf("No block height set"))
	}
	return m.blockHeight
}

func (m *mockHandler) SdkEnvGetBlockTimestamp(ctx context.ContextId, permissionScope context.PermissionScope) uint64 {
	if m.blockTimestamp != 0 {
		return m.blockTimestamp
	}
	return uint64(time.Now().UnixNano())
}

func (m *mockHandler) SdkEnvGetBlockProposerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	if m.blockProposerAddress == nil || len(m.blockProposerAddress) == 0 {
		panic(errors.Errorf("No block proposer address set"))
	}
	return m.blockProposerAddress
}

func (m *mockHandler) SdkEnvGetBlockCommittee(ctx context.ContextId, permissionScope context.PermissionScope) [][]byte {
	return m.committee
}

func (m *mockHandler) SdkEnvGetNextBlockCommittee(ctx context.ContextId, permissionScope context.PermissionScope) [][]byte {
	return m.nextCommittee
}

func (m *mockHandler) SdkEnvGetVirtualChainId(ctx context.ContextId, permissionScope context.PermissionScope) uint32 {
	return 42
}

func (m *mockHandler) MockEthereumLog(address string, abiJson string, ethTxHash string, eventName string, outEthBlockNumber int, outEthTxIndex int, outMutator func(out interface{})) {
	var key []interface{}
	key = append(key, address, abiJson, ethTxHash, eventName)
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: outMutator, blockHeight: uint64(outEthBlockNumber), txIndex: uint32(outEthTxIndex)})
}

func (m *mockHandler) MockEthereumCallMethod(address string, abiJson string, methodName string, outMutator func(out interface{}), args ...interface{}) {
	var key []interface{}
	key = append(key, address, abiJson, methodName)
	key = append(key, args...)
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: outMutator})
}

func (m *mockHandler) MockEthereumCallMethodAtBlock(blockNumber uint64, address string, abiJson string, methodName string, outMutator func(out interface{}), args ...interface{}) {
	var key []interface{}
	key = append(key, address, abiJson, methodName, blockNumber)
	key = append(key, args...)
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: outMutator})
}

func (m *mockHandler) MockEthereumGetBlockNumber(block int) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockNumber")
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: nil, blockHeight: uint64(block)})
}

func (m *mockHandler) MockEthereumGetBlockNumberByTime(block int, timestamp int) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockNumberByTime", uint64(timestamp))
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: nil, blockHeight: uint64(block), blockTimestamp: uint64(timestamp)})
}

func (m *mockHandler) MockEthereumGetBlockTime(timestamp int) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockTime")
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: nil, blockTimestamp: uint64(timestamp)})
}

func (m *mockHandler) MockEthereumGetBlockTimeByNumber(block int, timestamp int) {
	var key []interface{}
	key = append(key, "SdkEthereumGetBlockTimeByNumber", uint64(block))
	m.ethereumStubs = append(m.ethereumStubs, &ethereumStub{key: key, outMutator: nil, blockHeight: uint64(block), blockTimestamp: uint64(timestamp)})
}

func (m *mockHandler) MockServiceCallMethod(serviceName string, methodName string, out []interface{}, args ...interface{}) {
	var key []interface{}
	key = append(key, serviceName, methodName)
	key = append(key, args...)
	m.serviceStubs = append(m.serviceStubs, &serviceStub{key: key, out: out})
}

func (m *mockHandler) MockEnvBlockHeight(block int) {
	m.blockHeight = uint64(block)
}

func (m *mockHandler) MockEnvBlockTimestamp(time int) {
	m.blockTimestamp = uint64(time)
}

func (m *mockHandler) MockEnvGetBlockCommittee(committee [][]byte) {
	m.committee = committee
}

func (m *mockHandler) MockEnvGetNextBlockCommittee(committee [][]byte) {
	m.nextCommittee = committee
}

func (m *mockHandler) MockEnvBlockProposerAddress(addr []byte) {
	m.blockProposerAddress = addr
}

func (m *mockHandler) MockEmitEvent(eventFunctionSignature interface{}, args ...interface{}) {
	var key []interface{}
	key = append(key, args...)
	m.eventStubs = append(m.eventStubs, &eventStub{key: key})
}

func (m *mockHandler) MockCallContractAddress(name string, value []byte) {
	var key []interface{}
	key = append(key, "contract name", name)
	m.addressStubs = append(m.addressStubs, &addressStub{key: key, out: value})
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
	for _, stub := range m.addressStubs {
		if !stub.satisfied {
			panic(errors.Errorf("addresses mock set but not called"))
		}
	}
}

func (m *mockHandler) getStateDiffs() []*StateDiff {
	var diffs []*StateDiff

	for _, k := range m.stateKeyOrder {
		byteKey, _ := hex.DecodeString(k)
		diffs = append(diffs, &StateDiff{
			Key:   byteKey,
			Value: m.state[k],
		})
	}
	return diffs
}

func InSystemScope(signerAddress []byte, callerAddress []byte, f func(mockery Mockery)) (diffs []*StateDiff, reads uint64, writes uint64) {
	return inScope(signerAddress, callerAddress, context.PERMISSION_SCOPE_SYSTEM, f)
}

func InServiceScope(signerAddress []byte, callerAddress []byte, f func(mockery Mockery)) (diffs []*StateDiff, reads uint64, writes uint64) {
	return inScope(signerAddress, callerAddress, context.PERMISSION_SCOPE_SERVICE, f)
}

func inScope(signerAddress []byte, callerAddress []byte, scope context.PermissionScope, f func(mockery Mockery)) (diffs []*StateDiff, reads uint64, writes uint64) {
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

	return handler.getStateDiffs(), handler.stateReads, handler.stateWrites
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
