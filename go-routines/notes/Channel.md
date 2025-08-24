# Channels

## Introduction
- Channel is a place to communicate Synchronously for Goroutines
- In Channel, there are sender and receiver, in common situations, the sender and receiver are different Goroutines
- When sending data to a Channel, Goroutine will be blocked, until something receives the data
- Channel fits for an alternative mechanism of async await like in JavaScript
- go1 -> channel -> go2
- if go2 have not received the data given by go1 via channel, then go1 will be blocked

## Channel Characteristics
- By default, channel can only store ONE DATA, if you want to add more, then wait for the old data to be taken by a Goroutine
- When creating channel, it needs a data type to store, or use interface{}
- Channel can be used by many go routines, but always just one data, same as the receiver
- Channel must be closed after usage, as it can cause a memory leak

## Creating a Channel
- Channel is represented by the data type chan
- To create a Channel, use make(), like creating a map
- When creating a Channel, we must decide the data type that the channel can store