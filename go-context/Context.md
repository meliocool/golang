# Go Context

## Introduction
- Context is a data that brings value, cancel signal, timeout signal, and deadline signal
- Context is commonly made per-request
- For example when there is an incoming request to a web server via HTTP Request
- Context is used to ease forwarding value and signal from process to process

## Why is it necessary?
- Context in Golang is used to send request data or signal to other process
- With context, when we want to cancel ALL process, we just need to send a signal to the context, and automatically will cancel all process
- Almost all aspect of Golang uses context, for example database, http server, http client, etc
- Even in google when they use golang, context is a must use and always sent to functions

## How it works
- Process A has a context in it and sends said context to another process B
- So Process B has context from process A
- If process A decides to cancel, it will send a cancel signal to the context in process B
- Thus, both process will be cancelled

## Context
- Context is represented as an Interface
- Interface Context is inside the package context
- Because it is an Interface, we need to make a struct that fits the interface contract
- But this is done automatically by the package

## Function to Create a Context
- context.Background(): Create an empty context, will never be cancelled, never timeout, and never have any value, usually used in main function or in a test or at the start of a request process
- context.TODO(): Create an empty context, but is used when its not clear what context is gonna be used