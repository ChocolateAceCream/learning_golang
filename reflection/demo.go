package main

import (
	"fmt"
	"reflect"
)

type order struct {
	id   int
	name string
}

//prints the concrete type and the value of the interface.
//Type  main.order
//Value  {456 56}
func createOrder(query interface{}) {
	fmt.Println("---------------start--------")
	t := reflect.TypeOf(query)
	k := t.Kind()
	fmt.Println("type, ", t)
	fmt.Println("kind, ", k)
	// Type represents the actual type of the interface{}, in this case main.Order
	// Kind represents the specific kind of the type. In this case, it's a struct.
	v := reflect.ValueOf(query)
	fmt.Println("value, ", v)

	// NumField() method returns the number of fields in a struct
	// Field(i int) method returns the reflect.Value of the ith field.
	if k == reflect.Struct {
		n := v.NumField()
		fmt.Println("NumField, ", n)
		for i := 0; i < n; i++ {
			fmt.Printf("Field: %d, Type: %T, Value: %v\n", i, v.Field(i), v.Field(i))
		}
	} else if k == reflect.Int {
		// Int() extract the the reflect.Value as an int64
		i := reflect.ValueOf(query).Int()
		fmt.Printf("Type: %T, Value: %v\n", i, i)
	} else if k == reflect.String {
		i := reflect.ValueOf(query).String()
		fmt.Printf("Type: %T, Value: %v\n", i, i)
	}
	fmt.Println("---------------done--------")

	// The methods Int and String help extract the reflect.Value as an int64 and string

}

func main() {
	o := order{
		id:   12,
		name: "aa",
	}
	createOrder(o)
	b := 123
	createOrder(b)
	c := "hello world"
	createOrder(c)
}

// Reflection is a very powerful and advanced concept in Go and it should be used with care. It is very difficult to write clear and maintainable code using reflection. It should be avoided wherever possible and should be used only when absolutely necessary
