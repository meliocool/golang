package main

import "fmt"

// * Struct -> Template Data used to combine 0 or more data type into 1 unity
// * Struct is usually a representation of the data in program
// * Data in struct are stored in Field
// * Struct -> a group of Fields
// * Struct fields are typed with Pascal Case or Capitalize the first letter

type Customer struct {
	Name, Address string
	Age           int
}

// * Struct can also be used as a parameter for a function
// * We can add Method in struct, so its a function attached to a struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello,", name, "My Name is", customer.Name)
}

func main() {
	var customer Customer
	customer.Name = "Lee Jae-in"
	customer.Address = "South Korea"
	customer.Age = 20

	// Struct Literals
	customer2 := Customer{
		Name:    "Go Youn-Jung",
		Address: "South Korea",
		Age:     30,
	}

	customer3 := Customer{"Jung Somin", "South Korea", 36}

	fmt.Println(customer)
	fmt.Println(customer2)
	fmt.Println(customer3)

	customer4 := Customer{Name: "Kotone"}
	customer4.sayHello("Dhitan")
}
