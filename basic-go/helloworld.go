package main // Package main, the main function needs to be in the main package
import "fmt" // Module fmt or Format, built in Go-Lang

func main() { // To declare a function, type func
	fmt.Println("Hello World!") // Outputs Hello World!
}

// Use go build to create a .exe file
// Or use go run helloworld.go to simply run the go file
// There can only be ONE main function in one package
