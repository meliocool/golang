package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// * The code previously in encode_json_test.go does not comply to the JSON Contract, which is an Object and an Array
// * JSON Object in Go is represented as Struct
// * Where every attribute in JSON Object is an attribute in Struct

type Address struct {
	Street     string
	Country    string
	PostalCode string
}
type Customer struct {
	FirstName, MiddleName, LastName string
	Age                             int
	Married                         bool
	Hobbies                         []string
	Addresses                       []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Melio",
		MiddleName: "Nisha",
		LastName:   "Dhitan",
		Age:        20,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
