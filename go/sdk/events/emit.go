package events

import (
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"reflect"
	"runtime"
	"strings"
)

func EmitEvent(function interface{}, args ...interface{}) {
	contextId, handler, permissionScope := context.GetContext()
	functionName := functionNameFromFunction(function)
	handler.SdkEventsEmitEvent(contextId, permissionScope, functionName, args...)
}

func functionNameFromFunction(function interface{}) string {
	v := reflect.ValueOf(function)
	if v.Kind() != reflect.Func {
		panic("first argument of EmitEvent must be a function")
	}
	fullPackageName := runtime.FuncForPC(v.Pointer()).Name()
	parts := strings.Split(fullPackageName, ".")
	if len(parts) == 0 {
		panic("first argument of EmitEvent returned an invalid function name")
	} else {
		return parts[len(parts)-1]
	}
}
