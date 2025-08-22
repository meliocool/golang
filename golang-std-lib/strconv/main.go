package main

// * Package strconv allows complex data type conversion

import (
	"fmt"
	"strconv"
)

func main() {
	str := "12345"
	strBool := "true"
	strFloat := "42.69"

	boolz := true
	intz := 420
	floatz := 6.9

	// * strconv.ParseBool(string) -> String to Boolean
	resultBool, err := strconv.ParseBool(strBool)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	// * strconv.ParseFloat(string, bitSize) -> String to float64 (bitSize = 32 or 64)
	resultFloat, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(resultFloat)
	}

	// * strconv.ParseInt(string, base, bitSize) -> String to int64
	// * base = 10 means decimal, bitSize = 0/8/16/32/64
	resultInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	// * strconv.Atoi(string) -> String to int (shortcut for base 10, 64 bit)
	resultAtoi, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(resultAtoi)
	}

	// * strconv.FormatBool(bool) -> Boolean to String
	fmt.Println(strconv.FormatBool(boolz))

	// * strconv.FormatFloat(float64, fmt, prec, bitSize) -> float64 to String
	// * fmt: 'f' = decimal, 'e' = scientific, 'g' = compact
	fmt.Println(strconv.FormatFloat(floatz, 'f', 2, 64)) // "6.90"

	// * strconv.FormatInt(int64, base) -> int64 to String
	fmt.Println(strconv.FormatInt(int64(intz), 10)) // decimal
	fmt.Println(strconv.FormatInt(int64(intz), 16)) // hex

	// * strconv.Itoa(int) -> int to String (shortcut for base 10)
	fmt.Println(strconv.Itoa(intz))

	// * strconv.Quote(string) -> Add quotes around string + escape special chars
	fmt.Println(strconv.Quote("Hello\nWorld"))

	// * strconv.Unquote(string) -> Remove quotes / unescape string
	quoted := strconv.Quote("Lee \"Jae-In\"")
	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(unquoted)
	}
}
