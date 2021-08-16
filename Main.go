package main

import (
	"fmt"
)

var (
	a int     = 123
	b string  = "123123"
	c float32 = 1.2312312
)

type Student struct {
	id   int
	name string
}

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
	v1 := new(Student)
	v1.name = "v1"
	fmt.Println(v1)

	v2 := Student{name: "v2"}
	fmt.Println(v2)

	str := "12d"
	for _, s := range str {
		r := s - '0'
		fmt.Printf("%v\n", r)
		fmt.Printf("%T\n", r)
		// r = r*10 + int(s) - '0'
		// fmt.Println(r)
		// fmt.Println(r > 10)
	}
}
