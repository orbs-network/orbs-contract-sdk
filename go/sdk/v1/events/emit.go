package events

import (
	"github.com/orbs-network/orbs-contract-sdk/go/context"
)

func EmitEvent(eventFunctionSignature interface{}, args ...interface{}) {
	contextId, handler, permissionScope := context.GetContext()
	handler.SdkEventsEmitEvent(contextId, permissionScope, eventFunctionSignature, args...)
}
