package main

import "fmt"

// * Defer -> Scheduling a function to be executed after another function is done executing
// * Used to always be executed even though there is an error
func logging() {
	fmt.Println("Function called successfully!")
}

func runApplication() {
	defer logging()
	fmt.Println("Application Started!")
}

// * Panic -> Function that can be used to TERMINATE the program
// * Defer function will still be executed

func endApp() {
	fmt.Println("Application Ended!")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error Occurred!")
	}
}

// * Recover -> Function that catches the Panic Data
// * With Recover, panic will stop and program will continue to run

// ! WRONG WAY TO USE RECOVER
func runApp2(error bool) {
	defer endApp()
	if error {
		panic("ERROR!!!")
	}
	message := recover()
	fmt.Println("Panic Occured", message)
}

// * CORRECT WAY TO USE RECOVER !
func endAppRecover() {
	fmt.Println("Application Ended!")
	message := recover()
	fmt.Println("Panic Occured", message)
}

func runApp3(error bool) {
	defer endAppRecover()
	if error {
		panic("ERROR!!!")
	}
}

func main() {
	//runApplication()
	//runApp(false)
	//runApp2(true)
	runApp3(true)
	fmt.Println("Hi!") // * This will run because of recover()
}
