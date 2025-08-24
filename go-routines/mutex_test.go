package go_routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// * sync.Mutex (Mutual Exclusion)
// * To Lock or Unlock a variable or a process, so if it is lock then only 1 goroutine that can use the variable until it unlocks it

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Count:", x) // * It will be 100000
}
