package main

import (
	"fmt"
)

var (
	a int     = 123
	b string  = "123123"
	c float32 = 1.2312312
)

var shadowExample int = 1

func main() {
	fmt.Println(shadowExample)
	shadowExample := 2
	fmt.Println(shadowExample)

	fmt.Println("Hello world")

	var i int
	i = 1
	fmt.Printf("%v, %T", i, i)

	var j int = 2
	fmt.Printf("%v, %T", j, j)

	k := 2.123123123
	fmt.Printf("%v, %T", k, k)
	fmt.Printf("%v, %T", a, a)
	fmt.Printf("%v, %T", b, b)
	fmt.Printf("%v, %T", c, c)

	fmt.Println("\n---------test-------- ")
	str := "aaaaabbcccd"
	m2 := make(map[string]int)
	for _, v := range str {
		m2[string(v)] += 1
	}
	fmt.Println(m2)
}
