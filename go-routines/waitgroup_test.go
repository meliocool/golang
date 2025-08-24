package go_routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// * WaitGroup is a feature to wait a process until it is finished
// * Used when we want to wait for all goroutines to finish before moving on
// * Use Add(int) and then when the goroutine is finished, use Done()
// * To wait all process use Wait()

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hi")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Done!")
}
