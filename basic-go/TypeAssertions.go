package main

import "fmt"

// * Type Assertions is the abillity to change data type into desired type
// * Often used on empty interface

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// ! Panic
	// ! resultInt := result.(int)
	// ! fmt.Println(resultInt)

	// * Best Usage:
	value := random()
	switch val := value.(type) {
	case string:
		fmt.Println("String", val)
	case int:
		fmt.Println("Int", val)
	default:
		fmt.Println("Unknown Type")
	}
}
