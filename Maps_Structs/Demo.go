package main

import (
	"fmt"
)

func main() {

	//initialize map
	mapper := map[string]int{
		"CA": 12,
		"BS": 32,
	}
	fmt.Println(mapper)

	//initialize map with make()
	m2 := make(map[string]int)
	m2 = map[string]int{
		"bb": 1,
		"c":  2,
	}
	fmt.Println(m2)

	//retrive element value from map
	fmt.Println(m2["c"])

	//assign value to map element. however, the return order of map is not guaranteed, so change of element value may also change the order of map
	m2["c"] = 123
	fmt.Println(m2)

	//delete an element from map
	delete(m2, "c")
	fmt.Println(m2)
}
