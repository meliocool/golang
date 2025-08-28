package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// * JSON Array
// * JSON also uses the Array data type
// * Array in JSON is similar to arrays in JavaScript
// * Array can contain primitive data or complex data type (Object or Array)
// * In Go, JSON Array is represented as Slice

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Melio",
		MiddleName: "Nisha",
		LastName:   "Dhitan",
		Age:        20,
		Married:    false,
		Hobbies:    []string{"Gaming", "Coding", "Reading"},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Melio","MiddleName":"Nisha","LastName":"Dhitan","Age":20,"Married":false,"Hobbies":["Gaming","Coding","Reading"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
}

// * Complex JSON Array
func TestComplexJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Melio",
		MiddleName: "Nisha",
		LastName:   "Dhitan",
		Age:        20,
		Married:    false,
		Hobbies:    []string{"Gaming", "Coding", "Reading"},
		Addresses: []Address{
			{
				Street:     "Pancoran",
				Country:    "Indonesia",
				PostalCode: "696969",
			},
			{
				Street:     "Mampang",
				Country:    "Japan",
				PostalCode: "123123",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestComplexJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Melio","MiddleName":"Nisha","LastName":"Dhitan","Age":20,"Married":false,"Hobbies":["Gaming","Coding","Reading"],"Addresses":[{"Street":"Pancoran","Country":"Indonesia","PostalCode":"696969"},{"Street":"Mampang","Country":"Japan","PostalCode":"123123"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

// * Encode Plain JSON Array
func TestJSONPlainArrayEncode(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Pancoran",
			Country:    "Indonesia",
			PostalCode: "696969",
		}, {
			Street:     "Mampang",
			Country:    "Japan",
			PostalCode: "123123",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}

// * Decode Plain JSON Array
func TestJSONPlainArrayDecode(t *testing.T) {
	jsonString := `[{"Street":"Pancoran","Country":"Indonesia","PostalCode":"696969"},{"Street":"Mampang","Country":"Japan","PostalCode":"123123"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
