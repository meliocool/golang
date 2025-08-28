package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

// * Streaming Encoder
// * Encoder can be creating and write JSON straight to io.Writer
// * With that, there is no need to store it in a variable string or []byte
// * use json.Encoder and use json.NewEncoder(writer) with json.Encode(interface{})

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("Customer_Output.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName:  "Melio",
		MiddleName: "Nisha",
		LastName:   "Dhitan",
	}
	encoder.Encode(customer)
}
