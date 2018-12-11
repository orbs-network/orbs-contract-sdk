package context

type SdkHandler interface {

	// state
	SdkStateReadBytesByAddress(ctx ContextId, permissionScope PermissionScope, address []byte) []byte
	SdkStateWriteBytesByAddress(ctx ContextId, permissionScope PermissionScope, address []byte, value []byte)

	// service
	SdkServiceCallMethod(ctx ContextId, permissionScope PermissionScope, serviceName string, methodName string, args ...interface{}) []interface{}

	// events
	SdkEventsEmitEvent(ctx ContextId, permissionScope PermissionScope, eventFunctionSignature interface{}, args ...interface{})

	// ethereum
	SdkEthereumCallMethod(ctx ContextId, permissionScope PermissionScope, contractAddress string, jsonAbi string, methodName string, out interface{}, args ...interface{})
	SdkEthereumGetTransactionLog(ctx ContextId, permissionScope PermissionScope, contractAddress string, jsonAbi string, ethTransactionId string, eventName string, out interface{})

	// address
	SdkAddressGetSignerAddress(ctx ContextId, permissionScope PermissionScope) []byte
	SdkAddressGetCallerAddress(ctx ContextId, permissionScope PermissionScope) []byte
}
