package main

// * Array is fixed size when declared
// * Index starts from 0

import "fmt"

func main() {
	var names [3]string
	names[0] = "Lee Jae-In"
	names[1] = "Go Youn-Jung"
	names[2] = "Kotone"

	fmt.Println("All Names: ", names)
	fmt.Println("Name at Index 0: ", names[0])
	fmt.Println("Name at Index 1: ", names[1])
	fmt.Println("Name at Index 2: ", names[2])

	var values = [3]int{
		90,
		10,
		// * This is 0
	}
	fmt.Println("Values: ", values)

	// Functions in Array
	var newValue = [...]int{
		80,
		70,
		60,
	}
	fmt.Println("New Value: ", newValue)
	fmt.Println("Length of New Value Array: ", len(newValue))
}
