// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package context

type SdkHandler interface {

	// state
	SdkStateReadBytes(ctx ContextId, permissionScope PermissionScope, key []byte) []byte
	SdkStateWriteBytes(ctx ContextId, permissionScope PermissionScope, key []byte, value []byte)

	// service
	SdkServiceCallMethod(ctx ContextId, permissionScope PermissionScope, serviceName string, methodName string, args ...interface{}) []interface{}

	// events
	SdkEventsEmitEvent(ctx ContextId, permissionScope PermissionScope, eventFunctionSignature interface{}, args ...interface{})

	// ethereum
	SdkEthereumCallMethod(ctx ContextId, permissionScope PermissionScope, ethContractAddress string, jsonAbi string, ethBlockNumber uint64, methodName string, out interface{}, args ...interface{})
	SdkEthereumGetTransactionLog(ctx ContextId, permissionScope PermissionScope, ethContractAddress string, jsonAbi string, ethTxHash string, eventName string, out interface{}) (ethBlockNumber uint64, ethTxIndex uint32)
	SdkEthereumGetBlockNumber(ctx ContextId, permissionScope PermissionScope) (ethBlockNumber uint64)
	SdkEthereumGetBlockNumberByTime(ctx ContextId, permissionScope PermissionScope, ethBlockTimestamp uint64) (ethBlockNumber uint64)
	SdkEthereumGetBlockTime(ctx ContextId, permissionScope PermissionScope) (ethBlockTimestamp uint64)
	SdkEthereumGetBlockTimeByNumber(ctx ContextId, permissionScope PermissionScope, ethBlockNumber uint64) (ethBlockTimestamp uint64)

	// address
	SdkAddressGetSignerAddress(ctx ContextId, permissionScope PermissionScope) []byte
	SdkAddressGetCallerAddress(ctx ContextId, permissionScope PermissionScope) []byte
	SdkAddressGetOwnAddress(ctx ContextId, permissionScope PermissionScope) []byte
	SdkAddressGetContractAddress(ctx ContextId, permissionScope PermissionScope, contractName string) []byte

	// env
	SdkEnvGetBlockHeight(ctx ContextId, permissionScope PermissionScope) uint64
	SdkEnvGetBlockTimestamp(ctx ContextId, permissionScope PermissionScope) uint64
	SdkEnvGetBlockProposerAddress(ctx ContextId, permissionScope PermissionScope) []byte
	SdkEnvGetBlockCommittee(ctx ContextId, permissionScope PermissionScope) [][]byte
	SdkEnvGetVirtualChainId(ctx ContextId, permissionScope PermissionScope) uint32
}
