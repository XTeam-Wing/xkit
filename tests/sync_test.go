package tests

import (
	"github.com/XTeam-Wing/xkit/kits/sync"
	"sync/atomic"
	"testing"
)

func TestWg(t *testing.T) {

	swg, err := sync.New(sync.WithSize(10))
	if err != nil {
		t.Fatal(err)
	}

	var c uint32

	for i := 0; i < 10000; i++ {
		swg.Add()
		go func(c *uint32) {
			defer swg.Done()
			atomic.AddUint32(c, 1)
		}(&c)
	}

	swg.Wait()

	if c != 10000 {
		t.Fatalf("%d, not all routines have been executed.", c)
	}
}
