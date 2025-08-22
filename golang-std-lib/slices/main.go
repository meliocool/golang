package main

import (
	"fmt"
	"slices"
)

// * In Recent GoLang versions, Generic Programming was Introduced
// * Allows programmer to create parameters of different types without using interface{} or any
// * One of the packages that uses such concept is Slices

func main() {
	names := []string{"Melio", "Dhitan", "Nisha"}
	values := []int{100, 95, 80, 90}
	sortedVals := []int{80, 90, 95, 100}

	// * slices.Min() -> Finds smallest value
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))

	// * slices.Max() -> Finds largest value
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))

	// * slices.Contains() -> Finds if a value is inside a slice (returns a bool)
	fmt.Println(slices.Contains(names, "Nate"))

	// * slices.Index() -> Finds the index of a value inside a slice
	fmt.Println(slices.Index(names, "Melio"))

	// * slices.BinarySearch -> Performs a binary search to look for a value inside a SORTED Slice
	fmt.Println(slices.BinarySearch(sortedVals, 95))
}
