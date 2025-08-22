package main

// * RegEXP is a package to search based on Regular Expression
// * uses a C library made by Google called RE2

import (
	"fmt"
	"regexp"
)

func main() {
	// * regexp.MustCompile(string) -> Create Regexp
	var regex *regexp.Regexp = regexp.MustCompile("m([a-z])+o")
	str := "melio"
	str2 := "mendo"
	str3 := "mEnDo"

	// * regex.MatchString(string) bool -> Check if the RegExp matches with the string
	fmt.Println(regex.MatchString(str))
	fmt.Println(regex.MatchString(str2))
	fmt.Println(regex.MatchString(str3))

	// * regex.FindAllString(string) bool -> Check if the regexp matches all given string
	fmt.Println(regex.FindAllString("melio mendo mEnDo", 10))
}
