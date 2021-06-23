package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("defer msg")
	panic("something bad happened")
	fmt.Println("end")
}
