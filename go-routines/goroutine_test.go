package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World!")
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld() // * RunHelloWorld() will run in Goroutine
	fmt.Println("This is the next line of code")

	time.Sleep(3 * time.Second)
}

// * Go is VERY LIGHT
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(2 * time.Second)
}
