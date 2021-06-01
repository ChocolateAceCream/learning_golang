package main

import (
	"fmt"
)

//iota: only scoped in current constant block
const (
	aa = iota
	bb
	cc
)

const (
	a2 = iota
)

func main() {
	fmt.Println("----constant------")
	const myConst int = 123
	fmt.Printf("%v, %T\n", myConst, myConst)

	const a = 42 //compiler can infer the type of constant
	var b int16 = 7
	fmt.Printf("%v, %T\n", a+b, a+b) //compiler will look up where constant a was used, then replace symbol a with its value 42, then infer its type

	// iota
	fmt.Println("----iota------")
	fmt.Printf("%v, %T\n", aa, aa)
	fmt.Printf("%v, %T\n", bb, bb)
	fmt.Printf("%v, %T\n", cc, cc)
	fmt.Printf("%v, %T\n", a2, a2)
}
