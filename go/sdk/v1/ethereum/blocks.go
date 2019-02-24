package ethereum

import "github.com/orbs-network/orbs-contract-sdk/go/context"

// returns the current Ethereum block number with some safety margin to avoid forks (about 25 minutes)
func GetBlockNumber() uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEthereumGetBlockNumber(contextId, permissionScope)
}
