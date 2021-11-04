package main

import "fmt"

type Id map[string]interface{}
type Show struct {
	Id
}

func main() {
	a := new(Show)
	a.Id = map[string]interface{}{
		"asdf": 123,
	}
	fmt.Println(*a)

	//you can not use following statement directly without initiating a.Id as a map
	a.Id["ddd"] = 321
	fmt.Println(*a)
}
