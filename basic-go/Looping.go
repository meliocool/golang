package main

import "fmt"

// For Loop -> Repeating the code inside the loop

func main() {
	counter := 1
	for counter <= 10 { // For as long as this is true
		fmt.Println("Looping:", counter) // This is printed
		counter++                        // Add counter by 1
	}
	fmt.Println("Done")

	// Init statement, statement before for loop is executed
	// Post statement, statement after for loop is executed

	for i := 0; i < 10; i++ {
		fmt.Println("Looping:", i)
	}

	// For Range -> Iterate through a Data Collection (Array, Slice, Map)
	names := []string{"Lee Jae-In", "Dhitan", "Go Youn-Jung"}

	// Manual way with regular for
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Using for range
	// index, variable to store each data := range dataCollection
	// If dont need Index or idx, use _
	for idx, name := range names {
		fmt.Println("Index:", idx, "=", name)
	}

	// Break -> Completely stops the entire loop
	// Continue -> Stops the current iteration, and move on to the next iteration
	for i := 0; i < 10; i++ {
		if i == 5 { // if i is 5 then exit the loop
			break
		}
		fmt.Println("Looping:", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 1 { // Skips Odd Numbers
			continue
		}
		fmt.Println("Only Even:", i)
	}
}
