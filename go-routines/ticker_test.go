package go_routines

import (
	"fmt"
	"testing"
	"time"
)

// * Ticker is a representation of a looping event
// * When ticker is expire, event will be sent to the channel
// * To make a ticker use time.NewTicker(duration)
// * To stop a ticker use Ticker.Stop()

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Println("ticker tick", i+1)
	}
}

func TestTick(t *testing.T) {
	ch := time.Tick(500 * time.Millisecond)

	for i := 0; i < 5; i++ {
		<-ch
		fmt.Println("tick tick", i+1)
	}
}
