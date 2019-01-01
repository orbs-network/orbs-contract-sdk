package env

import "github.com/orbs-network/orbs-contract-sdk/go/context"

func GetBlockHeight() uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEnvGetBlockHeight(contextId, permissionScope)
}

func GetBlockTimestamp() uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEnvGetBlockTimestamp(contextId, permissionScope)
}
