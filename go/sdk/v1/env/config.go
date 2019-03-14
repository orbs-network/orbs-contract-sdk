package env

import (
	"github.com/orbs-network/orbs-contract-sdk/go/context"
)

func GetVirtualChainId() uint32 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEnvGetVirtualChainId(contextId, permissionScope)
}
