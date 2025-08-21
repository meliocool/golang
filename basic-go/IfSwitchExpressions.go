package main

import "fmt"

// * If -> Decision, if it is true then do something
// * else if -> Alternative Decision
// * else -> If all If statements are false, then do this

func main() {
	name := "Lee Jae-In"
	// * If Expression
	if name == "Lee Jae-In" {
		fmt.Println("Hello Lee Jae-In")
	} else if name == "Go Youn-Jung" {
		fmt.Println("Hello Go Youn-Jung")
	} else {
		fmt.Println("Hello Random Person!")
	}

	// * Short Statement if
	if length := len(name); length >= 10 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Name is good")
	}

	// * Switch Expression
	switch name {
	case "Lee Jae-In":
		fmt.Println("Hello Lee Jae-In")
	case "Go Youn-Jung":
		fmt.Println("Hello Go Youn-Jung")
	default:
		fmt.Println("Hello Random Person!")
	}

	// * Short Statement If
	switch length := len(name); length > 10 {
	case true:
		fmt.Println("Name is too long")
	case false:
		fmt.Println("Name is good")
	default:
		fmt.Println("This is the default part")
	}

	// * Adding Condition in each case
	length := len(name)
	switch {
	case length > 15:
		fmt.Println("Name is too long")
	case length > 8:
		fmt.Println("Name kind of long")
	default:
		fmt.Println("Name is good")
	}
}
