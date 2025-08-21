package main

import (
	_ "basic-go/internal" // * Only for the init function, uses Blank Identifier
	"basic-go/packageexample"
	"basic-go/packageinit"
	"fmt"
)

func main() {
	result := packageexample.SayHelloFromOtherPackage("Lee Jae-In")
	fmt.Println(result)

	database := packageinit.GetDatabase() // * init is executed here
	fmt.Println(database)
}
