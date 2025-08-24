# GOROUTINES

## Introduction
- Goroutine is a LIGHT THREAD that is managed by Go Runtime
- Goroutine is very small, around 2kb, way smaller than thread (1 mb or 1000 kb)
- Unlike thread that runs in Parallel, Goroutine runs in Concurrent

## How Goroutines Work
- Goroutine is run by Go Scheduler inside a thread, where the number of threads is GOMAXPROCS (Number of CPU Cores)
- Goroutine cannot be called a "replacement" of thread, since it runs on top of thread
- We don't have to manage thread manually, because it is handled by Go Scheduler
- in Go Scheduler there are a few Terminology:
  - G: GoRoutine
  - M: Thread (Machine)
  - P: Processor
- When creating a Goroutine, it is inserted into the Local Queue or Global queue, and is handled by Go Scheduler
- Thread will take the Goroutines from Local Queue, or Global Queue, if it is empty, it "steals" from the other threads to help
- The Concurrent happens when the thread takes a Goroutine, if it is long it will be "suspended" or "paused" and give it back to the queue, and the thread will take another Goroutine that is faster

## How to Create Goroutine
- use go before calling a function in Goroutine -> go RunHelloWorld()
- When a function runs in Goroutine, it will run ASYNCHRONOUSLY, so it will not wait until the function ends
- The App will continue to run the next code without waiting for the Goroutine to finish