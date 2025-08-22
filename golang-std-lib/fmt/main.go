package main

// * fmt -> Package to format things, including input and output

import (
	"fmt"
	"strings"
)

func main() {
	// * Println -> simple print with space + newline
	fmt.Println("Hello this is the fmt package!")

	firstName := "Muhammad"
	lastName := "Dhitan"

	fmt.Println("Hi", firstName, lastName)

	// * Printf -> formatted print (C-style)
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)

	// * Sprintf -> store formatted string instead of printing
	fullName := fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println("Stored in variable:", fullName)

	// * Errorf -> create an error with formatted text
	err := fmt.Errorf("something went wrong with user: %s", fullName)
	fmt.Println("Error created:", err)

	// * Print numbers with formatting
	pi := 3.1415926535
	fmt.Printf("Pi default: %v\n", pi) // %v → default format
	fmt.Printf("Pi with 2 decimals: %.2f\n", pi)
	fmt.Printf("Pi scientific: %e\n", pi)

	// * Padding and alignment
	num := 42
	fmt.Printf("Padded number: %05d\n", num)   // 00042
	fmt.Printf("Right aligned: |%6d|\n", num)  // '    42'
	fmt.Printf("Left aligned:  |%-6d|\n", num) // '42    '

	// * Booleans
	truth := true
	fmt.Printf("Boolean: %t\n", truth)

	// * Different base representations
	dec := 255
	fmt.Printf("Decimal: %d\n", dec)
	fmt.Printf("Binary: %b\n", dec)
	fmt.Printf("Hexadecimal: %x\n", dec)

	// * Type info
	fmt.Printf("Type of pi: %T\n", pi)
	fmt.Printf("Type of firstName: %T\n", firstName)

	// * Scanln -> read user input (simple demo)
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Printf("You entered: %d\n", age)

	// * Fprint -> print to somewhere else (not console), like a file or buffer
	// * Example with a string builder
	var sb = &strings.Builder{}
	fmt.Fprint(sb, "Writing into a string builder instead of stdout!")
	fmt.Println("\nString builder contains:", sb.String())
}
