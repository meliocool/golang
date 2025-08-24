package go_routines

import (
	"fmt"
	"testing"
	"time"
)

// * The most common problem encountered in Concurrency or Parallel programming is RACE CONDITION
// * When using Goroutine, it doesn't only run concurrent, it can also run in parallel
// * This can be dangerous if we manipulate the same variable by different goroutine at the same time

func TestRaceCondition(t *testing.T) {
	x := 0 // * x is shared with multiple go routines
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Count:", x) // * It will NOT be 100000
}

// * To solve this, use sync.Mutex or Mutual Exclusion (See mutex_test)
