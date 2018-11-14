package ethereum

import "github.com/orbs-network/orbs-contract-sdk/go/context"

func CallMethod(contractAddress string, jsonAbi string, methodName string, out interface{}, args ...interface{}) {
	contextId, handler := context.GetContext()
	handler.SdkEthereumCallMethod(contextId, contractAddress, jsonAbi, methodName, out, args...)
}