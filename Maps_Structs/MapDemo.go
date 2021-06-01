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
		"va": 3,
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

	//check if a key exist in map
	va, ok := m2["va"]
	fmt.Println(va, ok)
	fmt.Println("length of m2 is:", len(m2))
	delete(m2, "va")
	_, ok = m2["va"]
	fmt.Println(ok)
	//key is deleted, return 0, and false since this 0 is not the key value, but the default value.
	//use _ if we only need to confirm if key is present in map

	//use len() function to get the length of map
	fmt.Println("length of m2 after delete is: ", len(m2))

	// change a map will also other maps that copyed from orignal map, since it's passing by reference, not data
	m3 := m2
	fmt.Println("m2 is:", m2)
	fmt.Println("m3 is:", m3)
	m3["add"] = 321
	fmt.Println("m2 is:", m2)
	fmt.Println("m3 is:", m3)
}
