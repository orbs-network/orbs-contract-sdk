package ethereum

import "github.com/orbs-network/orbs-contract-sdk/go/context"

// returns a single log (single emit), in future we will add a version that returns multiple
func GetTransactionLog(contractAddress string, jsonAbi string, ethTransactionId string, eventName string, out interface{}) {
	contextId, handler, permissionScope := context.GetContext()
	handler.SdkEthereumGetTransactionLog(contextId, permissionScope, contractAddress, jsonAbi, ethTransactionId, eventName, out)
}
