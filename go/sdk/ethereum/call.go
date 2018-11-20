package ethereum

import "github.com/orbs-network/orbs-contract-sdk/go/context"

func CallMethod(contractAddress string, jsonAbi string, methodName string, out interface{}, args ...interface{}) {
	contextId, handler, permissionScope := context.GetContext()
	handler.SdkEthereumCallMethod(contextId, permissionScope, contractAddress, jsonAbi, methodName, out, args...)
}
