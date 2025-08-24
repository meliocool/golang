package go_routines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// * Creating a channel
	// * channel := make(chan string)
	channel := make(chan string)

	// * Sending data to channel
	// * channel <- data
	//channel <- "Melio"

	// * Receiving data from channel
	// * data <- channel

	//data := <-channel
	//fmt.Println(<-channel)

	// * Use close() to close the channel
	defer close(channel)

	// * Example
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Dhitan Imam Sakti"
		fmt.Println("Done sending data to Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// * Channel is sent to a function via parameter
// * By default function parameter are passed by value, therefore we use pointer to pass by reference
// * However, for channel its automatically by reference, so no need for pointer

func GiveMeAResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Dhitan Imam Sakti"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeAResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

// * Channel In and Out
// * When sending channel as a parameter, the function can send or receive data from the channel
// * We can tell the function whether the channel is only used to send or receive data
// * This can be done in parameter by marking the channel if its for in or out

// * This function is only used to send data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Melio"
}

// * This function is only used to receive data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// * Buffered Channel
// * Channel can only receive 1 data, so if we add a second data, then it must wait for the first one
// * But sometimes, the sender is faster than the receiver, so that will make the sender slower too
// * We use Buffered Channel to create a queue for channel

// * Buffer Capacity -> We can store as many capacity, but be wise about it
// * For example 5
// * This fits perfectly when the goroutine that receives data is slower than the goroutine sending the data

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Melio"
		channel <- "Meliodash"
		channel <- "Yeah"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(cap(channel)) // * Capacity of the Buffered Channel
	fmt.Println(len(channel)) // * To see the data inside the Buffered Channel

	fmt.Println("DONE")
}

// * Range Channel
// * Sometimes there is a case where a channel is sent data continuously
// * We can use range when receiving data from a channel
// * When the channel is closed, then the loop stops
// * This is more simple compared to manually checking channel

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Loop -" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receiving Data", data)
	}

	fmt.Println("DONE")
}

// * Select Channel
// * We can create a few channel, and run a few goroutine
// * If we want every data from every channel, we can use select channel
// * With select channel, we can choose the fastest data from a few channel
// * If the data came at the same time, it will be chosen randomly

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeAResponse(channel1)
	go GiveMeAResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// * Default Select
// * If selecting to a channel with non-existent data, then a deadlock error will appear
// * But, using Default Select we can do something if there is no data

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeAResponse(channel1)
	go GiveMeAResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		default:
			fmt.Println("Waiting For Data")
		}
		if counter == 2 {
			break
		}
	}
}
