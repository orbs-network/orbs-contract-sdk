package service

import "github.com/orbs-network/orbs-contract-sdk/go/context"

func CallMethod(serviceName string, methodName string, args ...interface{}) []interface{} {
	contextId, handler := context.GetContext()
	return handler.SdkServiceCallMethod(contextId, serviceName, methodName, args...)
}