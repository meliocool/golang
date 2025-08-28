package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// * JSON Tag
// * By default attribute in Struct and JSON will be mapped based on the same attribute (case sensitive)
// * There are different styles, for example in JSON uses snake_case but struct uses PascalCase
// * Package JSON supports Tag Reflection

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTagEncode(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Macbook Pro",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Macbook Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	// * When reading, its actually case insensitive, but for snake case like image_url it does matter so the _ must be there
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
