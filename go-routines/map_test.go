package go_routines

import (
	"fmt"
	"sync"
	"testing"
)

// * sync.Map is different from regular map
// * This map is concurrent safe using goroutine
// * A few function that can be used in Map:
// * Store(key, value) -> To store data into the Map
// * Load(key) -> To get data from the Map using Key
// * Delete(key) -> To delete data from the Map using key
// * Range(function(key, value)) -> To iterate through the whole Map

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
