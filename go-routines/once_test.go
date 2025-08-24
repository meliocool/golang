package go_routines

import (
	"fmt"
	"sync"
	"testing"
)

// * sync.Once is a feature in Go that can be used to determine that a function is executed ONLY ONCE
// * So no matter how many goroutine that is accessing it, only the FIRST goroutine can execute the function

var counter = 0

func YOLO() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(YOLO) // * MUST NOT BE A FUNCTION WITH A PARAMETER
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter) // * Prints out 1
}
