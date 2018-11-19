package service

import "github.com/orbs-network/orbs-contract-sdk/go/context"

func CallMethod(serviceName string, methodName string, args ...interface{}) []interface{} {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkServiceCallMethod(contextId, permissionScope, serviceName, methodName, args...)
}