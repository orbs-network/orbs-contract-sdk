package context

import (
	"github.com/orbs-network/orbs-contract-sdk/go/context/g"
	"sync"
	"unsafe"
)

type context struct {
	contextId       ContextId
	handler         SdkHandler
	permissionStack []PermissionScope
}

var mutex = &sync.RWMutex{}
var activeContexts = make(map[unsafe.Pointer]*context)

func PushContext(contextId ContextId, handler SdkHandler, permissionScope PermissionScope) {
	gid := g.G()
	mutex.Lock()
	defer mutex.Unlock()

	activeContext := activeContexts[gid]
	if activeContext != nil {
		if activeContext.contextId != contextId {
			panic("PushContext: multiple contexts found")
		}
		activeContext.permissionStack = append(activeContext.permissionStack, permissionScope)
	} else {
		activeContexts[gid] = &context{
			contextId:       contextId,
			handler:         handler,
			permissionStack: []PermissionScope{permissionScope},
		}
	}
}

func PopContext(contextId ContextId) {
	gid := g.G()
	mutex.Lock()
	defer mutex.Unlock()

	activeContext := activeContexts[gid]
	if activeContext != nil {
		if activeContext.contextId != contextId {
			panic("PopContext: multiple contexts found")
		}
		if len(activeContext.permissionStack) <= 1 {
			delete(activeContexts, gid)
		} else {
			activeContext.permissionStack = activeContext.permissionStack[:len(activeContext.permissionStack)-1]
		}
	} else {
		panic("PopContext: context not found")
	}
}

func GetContext() (contextId ContextId, handler SdkHandler, permissionScope PermissionScope) {
	gid := g.G()
	mutex.RLock()
	defer mutex.RUnlock()

	activeContext := activeContexts[gid]
	if activeContext != nil && len(activeContext.permissionStack) >= 1 {
		return activeContext.contextId, activeContext.handler, activeContext.permissionStack[len(activeContext.permissionStack)-1]
	} else {
		panic("GetContext: context not found")
	}
}
