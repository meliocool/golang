package main

import "fmt"

// * Interface is an Abstract Data Type, and doesnt have real implementation
// * Interface is filled with method definition
// * Interface is mostly used as Contract

type HasName interface {
	GetName() string // * if something wants to implement HasName, then it must have GetName()
}

func SayHi(value HasName) {
	fmt.Println("Hi, ", value.GetName())
}

// * Implementing Interface is automatic as long as it fits the contract
type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

// * Empty Interface -> No Method declaration in it
// * Makes every data type implement it
// * Type alias of any
// * Example -> fmt.Println(a ...interface{}) panic(v interface{})

func Oops() interface{} {
	return 1
}

func Oops2() any {
	return 1
}

func main() {
	person := Person{Name: "Lee Jae-In"}
	animal := Animal{Name: "Doggy"}
	SayHi(person)
	SayHi(animal)

	var empty any = Oops()
	var empty2 any = Oops2()
	fmt.Println(empty)
	fmt.Println(empty2)
}
