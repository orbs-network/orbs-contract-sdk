package fake

import (
	"encoding/hex"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"math/rand"
)

type stateMap map[string][]byte
type mockHandler struct {
	signerAddress []byte
	callerAddress []byte

	state         stateMap
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
	panic("implement me")
}

func (m *mockHandler) SdkAddressGetSignerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	return m.signerAddress
}

func (m *mockHandler) SdkAddressGetCallerAddress(ctx context.ContextId, permissionScope context.PermissionScope) []byte {
	return m.callerAddress
}

func InSystemScope(signerAddress []byte, f func()) {
	handler := &mockHandler{
		signerAddress: signerAddress,
		callerAddress: []byte{},
		state: make(stateMap),
	}
	cid := context.ContextId(42)
	perm := context.PERMISSION_SCOPE_SYSTEM
	context.PushContext(cid, handler, perm)
	f()
	context.PopContext(cid)
}

func InServiceScope(callerAddress []byte, f func()) {
	handler := &mockHandler{
		signerAddress: []byte{},
		callerAddress: callerAddress,
		state: make(stateMap),
	}
	cid := context.ContextId(43)
	perm := context.PERMISSION_SCOPE_SERVICE
	context.PushContext(cid, handler, perm)
	f()
	context.PopContext(cid)
}

func AnAddress() (address []byte) {
	address = make([]byte, 20)
	rand.Read(address)
	return
}

