package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// * Decode JSON
// * To decode, use json.Unmarshal(byte[], interface{})
// * byte[] is the JSON data, and interface{} is a place to store the converted value, could be a pointer

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Melio","MiddleName":"Nisha","LastName":"Dhitan","Age":20,"Married":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
