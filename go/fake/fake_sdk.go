package fake

import (
	"encoding/hex"
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"github.com/pkg/errors"
	"math/rand"
	"strings"
)

type stateMap map[string][]byte
type ethereumLogStubs map[string]func(interface{})

type Mockery interface {
	MockEthereumLog(address string, abiJson string, ethTxHash string, eventName string, f func(out interface{}))
}

type mockHandler struct {
	mock.Mock

	signerAddress []byte
	callerAddress []byte

	state            stateMap
	ethereumLogStubs ethereumLogStubs
}

func (m *mockHandler) SdkStateReadBytesByAddress(ctx context.ContextId, permissionScope context.PermissionScope, address []byte) []byte {
	return m.state[hex.EncodeToString(address)]
}

func (m *mockHandler) SdkStateWriteBytesByAddress(ctx context.ContextId, permissionScope context.PermissionScope, address []byte, value []byte) {
	m.state[hex.EncodeToString(address)] = value
}

func (m *mockHandler) SdkServiceCallMethod(ctx context.ContextId, permissionScope context.PermissionScope, serviceName string, methodName string, args ...interface{}) []interface{} {
	panic("implement me")
}

func (m *mockHandler) SdkEthereumCallMethod(ctx context.ContextId, permissionScope context.PermissionScope, contractAddress string, jsonAbi string, methodName string, out interface{}, args ...interface{}) {
	panic("implement me")
}

func (m *mockHandler) SdkEthereumGetTransactionLog(ctx context.ContextId, permissionScope context.PermissionScope, contractAddress string, jsonAbi string, ethTransactionId string, eventName string, out interface{}) {
	if f, ok := m.ethereumLogStubs[mapKeyOf(contractAddress, jsonAbi, ethTransactionId, eventName)]; ok {
		f(out)
	} else {
		panic(errors.Errorf("No Ethereum logs stubbed for address %s, jsonAbi %s, txHash %s, event name %s", contractAddress, jsonAbi, ethTransactionId, eventName))
	}
}

func (m *mockHandler) SdkAddressGetSignerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	return m.signerAddress
}

func (m *mockHandler) SdkAddressGetCallerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	return m.callerAddress
}

func (m *mockHandler) MockEthereumLog(address string, abiJson string, ethTxHash string, eventName string, f func(out interface{})) {
	m.ethereumLogStubs[mapKeyOf(address, abiJson, ethTxHash, eventName)] = f
}

func InSystemScope(signerAddress []byte, f func(mockery Mockery)) {
	inScope(signerAddress, []byte{}, context.PERMISSION_SCOPE_SYSTEM, f)

}

func InServiceScope(callerAddress []byte, f func(mockery Mockery)) {
	inScope([]byte{}, callerAddress, context.PERMISSION_SCOPE_SERVICE, f)
}

func inScope(signerAddress []byte, callerAddress []byte, scope context.PermissionScope, f func(mockery Mockery)) {
	handler := &mockHandler{
		signerAddress:    signerAddress,
		callerAddress:    callerAddress,
		state:            make(stateMap),
		ethereumLogStubs: make(ethereumLogStubs),
	}
	cid := context.ContextId(43)
	context.PushContext(cid, handler, scope)
	f(handler)
	context.PopContext(cid)
}

func AnAddress() (address []byte) {
	address = make([]byte, 20)
	rand.Read(address)
	return
}

func mapKeyOf(ss ...string) string {
	return strings.Join(ss, "|")
}

