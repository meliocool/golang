package go_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// * Streaming Decoder
// * In reality JSON mostly came from an input like io.Reader or a Request Body
// * Package json has a feature to support this, which is json.Decoder
// * use json.NewDecoder(reader)
// * To read the input reader, convert to GoLang data using the function Decode(interface{})

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{
		FirstName:  "Melio",
		MiddleName: "Nisha",
		LastName:   "Dhitan",
	}

	decoder.Decode(customer)

	fmt.Println(customer)
}
