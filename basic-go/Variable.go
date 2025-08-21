package main

// * Variables -> A Place to Store Data
// * Constants -> Variables that cannot be changed
// * Type Declarations -> Create a new data type from an existing data type (alias)

import "fmt"

func main() {
	// * var variableName variableDataType
	// * var name string
	var name string

	name = "Lee Jae-In"
	fmt.Println()
	fmt.Println("Regular Variable Declaration: ")
	fmt.Println(name)

	// * var age = 19 -> Go Automatically makes the variable Number or Int because of the data inside
	var age = 19
	fmt.Println()
	fmt.Println("Automatic Number: ")
	fmt.Println(age)

	// * var is optional if it is immediately declared
	university := "Bina Nusantara University"
	fmt.Println()
	fmt.Println("Optional var: ")
	fmt.Println(university)

	// * Declaring Multiple Variables
	var (
		firstName = "Lee"
		lastName  = "Jae-In"
	)
	fmt.Println()
	fmt.Println("Declaring Multiple Variables: ")
	fmt.Println(firstName, lastName)

	// * const constantName constantDataType = value
	// * const constantName = value
	const NIM = "111"
	const noKTP = 111
	fmt.Println()
	fmt.Println("Constants: ")
	fmt.Println(NIM)
	fmt.Println(noKTP)

	// * Declaring Multiple Constants
	const (
		test1 = "Test1"
		test2 = "Test2"
	)
	fmt.Println()
	fmt.Println("Declaring Multiple Constants: ")
	fmt.Println(test1, test2)

	// * Data Type Conversion: Number
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	fmt.Println()
	fmt.Println("Data Type Conversion: Number")
	fmt.Println(value64)

	// * Data Type Conversion: String
	var secondName = "Dhitan"
	var firstLetterByte uint8 = uint8(secondName[0])
	var firstLetterString = string(firstLetterByte)
	fmt.Println()
	fmt.Println("Data Type Conversion: String")
	fmt.Println(firstLetterString)

	// * Type Declaration Example
	type mynumber int // Alternative name to number

	var herNumber mynumber = 111
	fmt.Println()
	fmt.Println("Type Declaration:")
	fmt.Println(herNumber)
	fmt.Println(mynumber(555))
}
