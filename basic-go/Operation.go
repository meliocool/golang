package main

import "fmt"

func main() {
	// +, -, *, /, %
	var a = 10
	var b = 10
	var aplusb int = a + b
	fmt.Println(aplusb)
	var aminb int = a - b
	fmt.Println(aminb)

	// Augmented Assignment
	a += 10
	fmt.Println(a)

	// Unary Operator
	b++
	fmt.Println(b)

	var boolA = true
	var boolB = !boolA
	var boolC = !boolB
	fmt.Println(boolB)

	// Comparison Operator
	// > < >= <= == !=
	var compare bool = boolA == boolC
	fmt.Println(compare)

	// Boolean Operator
	// &&
	// true true -> true
	// true false -> false
	// false true -> false
	// false false -> false
	// ||
	// true true -> true
	// true false -> true
	// false true -> true
	// false false -> false
	// !
	// true -> false
	// false -> true

	var finalGrade = 90
	var absent = 3
	var passFinalGrade bool = finalGrade > 80
	var passAbsent bool = absent < 5
	var pass bool = passFinalGrade && passAbsent
	fmt.Println(pass)
}
