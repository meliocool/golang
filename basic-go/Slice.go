package main

// * Slice -> is taking a slice from an Array
// * Slice is NOT FIXED SIZE
// * Slice is a data that access a part of an Array or as a whole

// * Slice has 3 Data: Pointer, Length, and Capacity
// * Pointer -> Points to the first data in the slice
// * Length -> Length of the Slice
// * Capacity -> Capacity of the slice, where length cannot exceed the Capacity
// * array[low:high] -> index low to before high
// * array[low:] -> index 0 to the end
// * array[:high] -> index 0 to index before high
// * array[:] -> Take all

// * Example:
// * data array of months from 0 to 11 (January to December)
// * array[4:7] -> pointer = 4, length = 3, capacity = 8 (11 - 4 [4 is included])

import "fmt"

func main() {
	month := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	slice1 := month[4:7]
	fmt.Println("First Slice from [4:7] -> ", slice1)
	slice2 := month[:7]
	fmt.Println("Second Slice from [:7] -> ", slice2)
	slice3 := month[7:]
	fmt.Println("Third Slice from [7:] -> ", slice3)
	var slice4 []string = month[:] // Other form to declare a slice
	fmt.Println("Fourth Slice from [:] -> ", slice4)

	// * Functions in Slice
	// * len() -> The Length of a Slice
	slice1Length := len(slice1)
	fmt.Println("Length of the First Slice: ", slice1Length)

	// * cap() -> The Capacity of a Slice
	slice2cap := cap(slice2)
	fmt.Println("Cap of the second Slice: ", slice2cap)

	// * Append -> Creates a new Slice and adds data to the end of the Slice
	// * If the capacity is full, then create a new array
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice1 := days[5:]
	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "New Day") // * Creates a new array which is daySlice2, independent of days
	daySlice2[0] = "First Day"
	fmt.Println(daySlice2) // * [First Day New Sunday New Day]

	// * Make Slice
	newSlice := make([]string, 2, 5) // * make([]Type, Length, Capacity)
	newSlice[0] = "First Data"
	newSlice[1] = "Second Data"
	fmt.Println(newSlice)
	fmt.Println("Length of newSlice: ", len(newSlice))
	fmt.Println("Capacity of newSlice: ", cap(newSlice))

	newSliceAppend := append(newSlice, "Third Data")
	fmt.Println(newSliceAppend)
	fmt.Println("Length of newSliceAppend: ", len(newSliceAppend))
	fmt.Println("Capacity of newSliceAppend: ", cap(newSliceAppend))

	// * Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice) // copy(destination, source)
	fmt.Println("From Slice: ", fromSlice)
	fmt.Println("To Slice: ", toSlice)

	// * Careful when creating an Array
	thisisArray := [...]int{1, 2, 3}
	thisisSlice := []int{1, 2, 3}
	fmt.Println("This is an Array", thisisArray)
	fmt.Println("This is a Slice", thisisSlice)
}
