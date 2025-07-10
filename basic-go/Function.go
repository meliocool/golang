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
}
