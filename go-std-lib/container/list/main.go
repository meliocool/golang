package main

import (
	"container/list"
	"fmt"
)

// * Package container/list is an implementation of the Double Linked List Data Structure

func main() {
	data := list.New()
	data.PushBack("Muhammad")
	data.PushBack("Dhitan")
	data.PushBack("Imam")
	data.PushBack("Sakti")
	data.PushFront("Is")
	data.PushFront("Name")
	data.PushFront("My")
	data.PushFront("Hi")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	var head *list.Element = data.Front()
	var tail *list.Element = data.Back()

	fmt.Println("Head:", head.Value)
	fmt.Println("After Head:", head.Next().Value)
	fmt.Println("Tail:", tail.Value)
}
