package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "Validation Error, ID Not Found!"}
	}
	if id != "Lee Jae-In" {
		return &notFoundError{
			Message: "Data Not Found",
		}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		// * Using If Else
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation Error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not Found Error:", notFoundErr.Error())
		} else {
			fmt.Print("Unknown Error!", err.Error())
		}

		// * Using Switch Case
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", finalError.Error())
		default:
			fmt.Print("Unknown Error!", finalError.Error())
		}
	} else {
		fmt.Println("Success!")
	}
}
