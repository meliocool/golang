package go_routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// * Timer is a representation of a moment
// * When the timer expires, then the event will be sent to the channel
// * To create a Timer, use time.NewTimer(duration)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

// * To execute a function after a brief delay, use time.AfterFunc()
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())
	group.Wait()
}
