package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// * To convert data to JSON (Encode) use json.Marshal(interface{})
// * This takes in any type of data into the Marshal function

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJson("Melio")
	logJson(1)
	logJson(true)
	logJson([]string{"Melio", "Nisha", "Dhitan"})
}
