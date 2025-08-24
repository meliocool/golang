package go_routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// * Pool is an implementation of design pattern called Object Pool Pattern
// * Design Pattern Pool is used to store data, use the data, but after using the data, we store it back to the Pool
// * The implementation of Pool in Go is RACE CONDITION SAFE
// * For example making Connection to Database

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Default"
		},
	}

	pool.Put("Melio")
	pool.Put("Dhitan")
	pool.Put("Nisha")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("DONE!")
}
