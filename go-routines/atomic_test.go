package go_routines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// * GoLang has a package called sync/atomic
// * Atomic is a package that allows for manipulating primitive data types safely with concurrency
// * An alternative for Mutex or RWMutex, so no need to lock the variable

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := &sync.WaitGroup{}
	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Count:", x) // * It will NOT be 100000
}
