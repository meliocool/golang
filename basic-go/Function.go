package main

import "fmt"

// Function -> A Block of Code to be used or re-used multiple times

func sayHello() {
	fmt.Println("Hello")
}

// Data can be inserted into a function, called Parameter
func sayHelloTo(name string, age int) {
	fmt.Println("Hello", name, "You are", age, "years old")
}

// A Function can also return a value
func plus(a int, b int) int {
	return a + b
}

// A Function could return more than 1 value
func getUser(name string, age int, height float32) (string, int, float32) {
	return name, age, height
}

// Named Return Values
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "FirstName"
	middleName = "MiddleName"
	lastName = "LastName"
	return firstName, middleName, lastName
}

// Variadic Functions (VarArgs)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function is also Data Type, and can be stored inside a Variable
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// Function as Parameter
type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func sayHelloWithFilter2(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Nate" {
		return "..."
	} else {
		return name
	}
}

// Anonymous Function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User is blacklisted")
	} else {
		fmt.Println("Welcome!")
	}
}

// Recursive Function
// Factorial without Recursion
func factorialLoop(value int) int {
	result := 1
	for i := 1; i <= value; i++ {
		result *= i
	}
	return result
}

// Factorial with Recursion
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	sayHello()
	sayHelloTo("Lee Jae-In", 22)
	fmt.Println(plus(1, 3))
	userName, userAge, userHeight := getUser("Dhitan", 19, 168.5)
	userNameOnly, _, _ := getUser("Dhitan", 19, 168.5)
	fmt.Println(userName, userAge, userHeight)
	fmt.Println(userNameOnly)
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)

	// Variadic Function
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))

	// Function as a Value
	goodbye := getGoodBye
	fmt.Println(goodbye("Lee Jae-In"))

	// Function as an argument
	sayHelloWithFilter("Nate", spamFilter)

	// Anonymous Function
	blacklist := func(name string) bool {
		return name == "Nate"
	}

	registerUser("Lee Jae-in", blacklist)
	registerUser("Nate", func(name string) bool {
		return name == "Nate"
	})

	// Recursive Function
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))

	// Closures -> Function interacting with data in the same scope
	// Don't use too much

	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	// If codebase is complex, it gets very confusing, because:
	// increment() actually changes the value of counter
	// Without the programmer passing in a value in the parameter
	increment()
	increment()
	fmt.Println(counter)
}
