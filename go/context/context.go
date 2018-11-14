package context

import (
	"sync"
	"github.com/orbs-network/orbs-contract-sdk/go/context/g"
	"unsafe"
)

type ContextId uint32

type context struct {
	contextId ContextId
	handler SdkHandler
}

var mutex = &sync.RWMutex{}
var activeContexts = make(map[unsafe.Pointer]*context)

func CreateContext(contextId ContextId, handler SdkHandler) {
	gid := g.G()
	mutex.Lock()
	defer mutex.Unlock()
	activeContexts[gid] = &context{contextId, handler}
}

func DestroyContext() {
	gid := g.G()
	mutex.Lock()
	defer mutex.Unlock()
	delete(activeContexts, gid)
}

func GetContext() (contextId ContextId, handler SdkHandler) {
	gid := g.G()
	mutex.RLock()
	defer mutex.RUnlock()
	c := activeContexts[gid]
	if c != nil {
		return c.contextId, c.handler
	} else {
		return 0, nil
	}
}