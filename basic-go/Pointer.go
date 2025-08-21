package main

import "fmt"

// * By Default, All Variables in GoLang are passed by VALUE, not by REFERENCE
// * If we pass in a variable to a function or another variable,
// * We are sending a Duplicate of the value If there are changes, the OG variable will not change

type Address struct {
	City, Province, Country string
}

// * Pointer allows the programmer to create a reference to the data location IN MEMORY, without duplicating
// * Pass By Reference

// ! Without Pointer
func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

// * With Pointer
func ChangeAddressToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia"
}

// * Pointer in Method
// * Even though methods are attached to struct, accessing data in the struct are passed by value
// * Recommended to use Pointer in Method

type Man struct {
	Name string
}

// ! Without Pointer
func (man Man) Married() {
	man.Name = "Mr." + man.Name
}

// * With Pointer
func (man *Man) PointerMarried() {
	man.Name = "Mr." + man.Name
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := address1 // * Value Address 1 is COPIED to address2

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	// * Operator &
	addressPointer := &address1 // * References to address1, so they share the same data without duplication
	addressPointer.City = "Seoul"
	fmt.Println(address1)
	fmt.Println(addressPointer)

	// * In full form, will look like this:
	var addressOG Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var addressPointerOG *Address = &addressOG
	addressPointerOG.Country = "South Korea"
	fmt.Println(addressOG)
	fmt.Println(addressPointerOG)

	// * Operator * (Asterisk) allows to change all pointer variable that points to a data
	// * Without *
	var addressPointer1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var addressPointer2 *Address = &addressPointer1

	addressPointer2.Country = "South Korea"
	fmt.Println(addressPointer1) // * Changes
	fmt.Println(addressPointer2) // * South Korea

	addressPointer2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(addressPointer1)
	fmt.Println(addressPointer2)

	// * With *
	addressPointer3 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	addressPointer4 := &addressPointer3

	addressPointer4.City = "Bandung"
	*addressPointer4 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(addressPointer3)
	fmt.Println(addressPointer4)

	// * Operator new to create pointers
	newAddress := new(Address)
	newAddress2 := newAddress

	newAddress.City = "Bandung"
	fmt.Println(newAddress)
	fmt.Println(newAddress2)

	// * Pointer in Functions
	addressPassToFunc := Address{}
	ChangeAddressToIndonesia(addressPassToFunc)
	fmt.Println(addressPassToFunc)

	addressPassToFuncPointer := &Address{}
	ChangeAddressToIndonesiaPointer(addressPassToFuncPointer)
	fmt.Println(addressPassToFuncPointer)

	// * Pointer in Method
	Jack := Man{"Jack"}
	Jack.Married()

	fmt.Println(Jack.Name)

	Jack.PointerMarried()
	fmt.Println(Jack.Name)
}
