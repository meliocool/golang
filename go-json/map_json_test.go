package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// * Map
// * Fit for dynamic JSON data, which means the attribute is not determined
// * Map Data type does not support JSON Tag

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Apple Macbook Pro",
		"price": 20000000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Macbook Pro","price": 20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}
