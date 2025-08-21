package main

// * GoLang has an interface that can be used as Contract to create an Error

import (
	"errors"
	"fmt"
)

func Division(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Cannot Divide by Zero")
	} else {
		return value / divider, nil // * error is an Interface, therefore can be nil
	}
}

func main() {
	result, err := Division(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
