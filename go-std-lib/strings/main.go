package main

// * Package strings allows to access functions that manipulates strings

import (
	"fmt"
	"strings"
)

func main() {
	str := "  Hi My Name is Lee Jae-In  "
	str2 := "lee jae-in, go youn-jung, han so-hee"

	// * Usage Examples:

	// * strings.Trim(string, cutset) -> Cut in beginning or end of string
	fmt.Println(strings.Trim(str, " "))

	// * strings.ToLower(string) -> Turns all the characters in a string to lowercase
	fmt.Println(strings.ToLower(str))

	// * strings.ToUpper(string) -> Turns all the characters in a string to uppercase
	fmt.Println(strings.ToUpper(str))

	// * strings.Split(string, separator) -> Cut string based on separator
	fmt.Println(strings.Split(str, "is"))

	// * strings.Contains(string, search) -> Check if a string contains other string
	fmt.Println(strings.Contains(str, "Lee"))

	// * strings.ReplaceAll(string, from, to) -> Change all string from "from" to "to"
	fmt.Println(strings.ReplaceAll(str, "Lee Jae-In", "Go Youn-Jung"))

	// * strings.HasPrefix(string, prefix) -> Check if string starts with prefix
	fmt.Println(strings.HasPrefix(str, "  Hi"))

	// * strings.HasSuffix(string, suffix) -> Check if string ends with suffix
	fmt.Println(strings.HasSuffix(str, "Jae-In  "))

	// * strings.Repeat(string, count) -> Repeat string n times
	fmt.Println(strings.Repeat("GoLang ", 3))

	// * strings.Count(string, substr) -> Count occurrences of substr in string
	fmt.Println(strings.Count(str, "a"))

	// * strings.Index(string, substr) -> Find first index of substr, -1 if not found
	fmt.Println(strings.Index(str, "Name"))

	// * strings.LastIndex(string, substr) -> Find last index of substr
	fmt.Println(strings.LastIndex(str, "i"))

	// * strings.Fields(string) -> Split on whitespace (like words)
	fmt.Println(strings.Fields(str))

	// * strings.Join([]string, separator) -> Join string slices with separator
	parts := strings.Split(str2, ", ")
	fmt.Println(strings.Join(parts, " | "))

	// * strings.Compare(a, b) -> Lexicographically compare two strings (-1, 0, 1)
	fmt.Println(strings.Compare("abc", "abc")) // 0
	fmt.Println(strings.Compare("abc", "abd")) // -1
	fmt.Println(strings.Compare("abd", "abc")) // 1

	// * strings.EqualFold(a, b) -> Case-insensitive string comparison
	fmt.Println(strings.EqualFold("GoLang", "golang"))
}
