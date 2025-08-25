package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// * Parent and Child Context
// * We can create child context from an existing context
// * Parent context can have many child, but a child can only have 1 parent context
// * Similar to Inheritance
// * Parent and child is always connected, if the parent is cancelled, then the child and sub-child will be cancelled
// * If A has a child B and B has a child C, then if B is cancelled, A will NOT get cancelled, only B and C is cancelled
// * If the parent has a data, then the child also has the same data
// * If A has a child B and B has a child C, then if B has data D, A will NOT get the data, only B and C has data D

// * Context is an object that is IMMUTABLE, that means when Context is made, it cannot change anymore
// * If we insert a value to a context or "change" the data, it automatically creates a child context
// * It does not CHANGE the original Context

// * Initially when creating a context it does not have a value
// * We can add a pair of data (key - value) inside the context
// * When adding value to a context, it automatically creates a child context and the OG does not change at all
// * To add value to a context use the function: context.WithValue(parent, key, value)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// * Getting a value
	fmt.Println(contextF.Value("f")) // * Get
	fmt.Println(contextF.Value("c")) // * Get the parent's value
	fmt.Println(contextF.Value("b")) // * Cannot Get Different Parent <nil>
	fmt.Println(contextA.Value("b")) // * Cannot get the child value <nil>

}

// * Context with Cancel
// * We can also send cancel signal to a process (mostly goroutines)
// * Goroutine that uses context, must also do a checking to the context
// * To send a context with cancelSignal, use context.WithCancel(parent)
// * This solves the goroutine leak problem

func CreateCounterLEAK() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestContextWithoutCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	destination := CreateCounterLEAK()
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Final Goroutine:", runtime.NumGoroutine()) // * There will be 3 goroutine instead of 2
}

func CreateCounterContext(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			case destination <- counter:
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterContext(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}
	cancel() // * Sends a cancel signal to context
	time.Sleep(5 * time.Second)

	fmt.Println("Final Goroutine:", runtime.NumGoroutine()) // * There will be 2 goroutines
}

// * With timeout, we can cancel automatically in a certain duration
// * For example querying to a database and has a fixed duration, if it exceeds then timeout to cancel it
// * use the function context.WithTimeout(parent, duration)

func CreateCounterTimeout(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(2 * time.Second) // * Slow simulation
			}
		}
	}()
	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterTimeout(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter:", n)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final Goroutine:", runtime.NumGoroutine()) // * There will be 2 goroutines
}

// * Deadline
// * Timeout gives time based on the current time + the timeframe
// * For deadline, it is determined when the deadline is, so for example 12 midnight
// * use context.WithDeadline(parent, time)

func CreateCounterDeadline(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(2 * time.Second) // * Slow simulation
			}
		}
	}()
	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounterDeadline(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter:", n)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final Goroutine:", runtime.NumGoroutine()) // * There will be 2 goroutines
}
