package main

import "fmt"

// * Nil can only be used on a few Data Types:
// * interface, function, map, slice, pointer, channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data is Nil")
	} else {
		fmt.Println(data["name"])
	}
}
