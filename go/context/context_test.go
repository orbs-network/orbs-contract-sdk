package context

import (
	"sync"
	"testing"
	"time"
)

func TestPushPop(t *testing.T) {
	PushContext(17, nil, PERMISSION_SCOPE_SERVICE)

	cid, h, perm := GetContext()
	if cid != 17 || h != nil || perm != PERMISSION_SCOPE_SERVICE {
		t.Fatalf("Read context (1) is incorrect")
	}

	PushContext(17, nil, PERMISSION_SCOPE_SYSTEM)

	cid, h, perm = GetContext()
	if cid != 17 || h != nil || perm != PERMISSION_SCOPE_SYSTEM {
		t.Fatalf("Read context (2) is incorrect")
	}

	PopContext(17)

	cid, h, perm = GetContext()
	if cid != 17 || h != nil || perm != PERMISSION_SCOPE_SERVICE {
		t.Fatalf("Read context (1) is incorrect")
	}

	PopContext(17)
}

func TestPushDifferentContextIdsOnSameGoroutinePanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("No panic although pushed different context ids")
		}
	}()

	PushContext(17, nil, PERMISSION_SCOPE_SERVICE)
	PushContext(29, nil, PERMISSION_SCOPE_SYSTEM)
}

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {

		contextId := ContextId(i + 17)
		wg.Add(1)

		go func() {

			t.Logf("Starting goroutine with context id %d", contextId)

			PushContext(contextId, nil, PERMISSION_SCOPE_SERVICE)
			defer PopContext(contextId)

			time.Sleep(5 * time.Millisecond)

			PushContext(contextId, nil, PERMISSION_SCOPE_SERVICE)
			defer PopContext(contextId)

			time.Sleep(5 * time.Millisecond)

			cid, _, _ := GetContext()
			if cid != contextId {
				t.Fatalf("GetContext returned wrong context id")
			}

			t.Logf("Read value %d on goroutine with context id %d", cid, contextId)

			wg.Done()

		}()
	}

	wg.Wait()
}
