package ethereum

import "github.com/orbs-network/orbs-contract-sdk/go/context"

// returns a single log (single emit), in future we will add a version that returns multiple
// ethContractAddress: string like "0x66c8bC6e162e45Da2Fc3337cF2164CA5E43CA4c5"
// ethTxHash: string like "0x734c2bb544c90d7b178cdfa18e60d8d002c4e9158e716000f67dbaed72d1a093"
func GetTransactionLog(ethContractAddress string, jsonAbi string, ethTxHash string, eventName string, out interface{}) (ethBlockNumber uint64, ethTxIndex uint32) {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEthereumGetTransactionLog(contextId, permissionScope, ethContractAddress, jsonAbi, ethTxHash, eventName, out)
}
