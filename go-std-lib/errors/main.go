package main

// * fmt -> Package to define and check errors

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound   = errors.New("not found error")
)

func getById(id string) error {
	if id == "" {
		return fmt.Errorf("getByID: %w", ErrValidation)
	}

	if id != "Lee Jae-in" {
		return fmt.Errorf("getByID: %w", ErrNotFound)
	}

	return nil
}

func main() {
	if err := getById(""); err != nil {
		if errors.Is(err, ErrValidation) {
			fmt.Println("Validation Error!: ", err)
		} else if errors.Is(err, ErrNotFound) {
			fmt.Println("Not Found Error!: ", err)
		} else {
			fmt.Println("Unknown Error: ", err)
		}
	} else {
		fmt.Println("Success!")
	}
}
