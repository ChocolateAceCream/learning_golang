package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max: "10"`
	Origin string
}

func main() {
	cat := Animal{
		Origin: "room",
	}
	fmt.Println("cat: ", cat)

	t := reflect.TypeOf((Animal{}))
	field, _ := t.FieldByName("Name") //retrive the tag value of that field "name", then leave it for the validation function to determine if that field is validated.
	fmt.Println(field.Tag)
}
