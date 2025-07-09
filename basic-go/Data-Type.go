package main

import "fmt"

func main() {
	// Integer -> int8, int16, int32, int64 -> minimum value can be negative
	// Unsigned Integer -> uint8, uint16, uint32, uint64 -> minimum value is 0
	// Floating Point -> float32, float64, complex64, complex128 (complex can include real and imaginary)
	// Alias -> byte: uint8, rune: int32, int: Minimal int32, uint: Minimal uint32
	fmt.Println("-- Numbers:")
	fmt.Println("One = ", 1)
	fmt.Println("Two = ", 2)
	fmt.Println("Three Point Five = ", 3.5)
	fmt.Println()

	// Boolean in Golang is Bool
	// true - false
	fmt.Println("-- Boolean:")
	fmt.Println("True = ", true)
	fmt.Println("False = ", false)
	fmt.Println()

	// -- STRING -- //
	// String in Golang always uses double quotes -> ""
	// Function for String:
	// len("string") -> Counts the number of characters in the String
	// "string"[number] -> Takes a character from a position based on the number (starts from 0) -> in Byte Form
	fmt.Println("-- String:")
	fmt.Println("Hi")
	fmt.Println("Her Name Is")
	fmt.Println("Lee Jae-In")
	fmt.Println(len("Park Gyu-Young"))
	fmt.Println("Cha Woo-Min"[0]) // Outputs the byte form which is 67
}
