package context

type SdkHandler interface {

	// state
	SdkStateReadBytesByAddress(ctx ContextId, address []byte) []byte
	SdkStateWriteBytesByAddress(ctx ContextId, address []byte, value []byte)

	// service
	SdkServiceCallMethod(ctx ContextId, serviceName string, methodName string, args ...interface{}) []interface{}

	// ethereum
	SdkEthereumCallMethod(ctx ContextId, contractAddress string, jsonAbi string, methodName string, out interface{}, args ...interface{})

	// address
	SdkAddressGetSignerAddress(ctx ContextId) []byte
	SdkAddressGetCallerAddress(ctx ContextId) []byte

}