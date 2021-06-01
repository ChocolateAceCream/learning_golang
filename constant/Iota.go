package main

import (
	"fmt"
)

const (
	_ = iota // _ is blank identifier, which avoids the need to create a dummy variable and makes it clear that the value is to be discarded.
	cat
	dog
	monkey
)

//check if var is assign to constant yet
func main() {
	const animal1 int = cat
	var animal2 int // default value will be 0
	fmt.Printf("%v\n", cat)
	fmt.Printf("%v\n", dog)
	fmt.Printf("%v\n", animal1)
	fmt.Printf("%v\n", animal2)
	fmt.Printf("%v\n", 1<<10)
}
