// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package g

import (
	"testing"
	"unsafe"
)

func TestPrintG(t *testing.T) {
	ch := make(chan unsafe.Pointer)
	g0 := G()
	t.Logf("G0 value is %p", g0)

	for i := 1; i < 5; i++ {
		go func() {
			ch <- G()
		}()
	}

	for i := 1; i < 5; i++ {
		gp := <-ch
		t.Logf("G%d value is %p", i, gp)
		if gp == g0 {
			t.Fatalf("Goroutine G is identical to original one.")
		}
	}

	gp := G()
	t.Logf("G0 value is %p", gp)
	if gp != g0 {
		t.Fatalf("Original goroutine G is different than the original one.")
	}
}

func TestG(t *testing.T) {
	gp1 := G()

	if gp1 == nil {
		t.Fatalf("fail to get G.")
	}

	t.Run("G in another goroutine", func(t *testing.T) {
		gp2 := G()

		if gp2 == nil {
			t.Fatalf("fail to get G.")
		}

		if gp2 == gp1 {
			t.Fatalf("every living G must be different. [gp1:%p] [gp2:%p]", gp1, gp2)
		}
	})
}
