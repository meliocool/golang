package main

// * Package Reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(data any) (result bool) {
	result = true
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			datas := reflect.ValueOf(data).Field(i).Interface()
			result = datas != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Dhitan"})
	readField(Person{"Melio", "Pancoran", "deiton@gmail"})

	person := Person{
		Name:    "Dhitan",
		Address: "JKT",
		Email:   "deiton@gmail",
	}

	fmt.Println(IsValid(person))
}
