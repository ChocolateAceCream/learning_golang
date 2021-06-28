package main

import (
	"fmt"
)

func main() {
	a := 42
	b := a
	fmt.Println(a, b)
	a = 28
	fmt.Println("after changed a, a,b =>", a, b)
	fmt.Println("---pointers example----")
	c := 1
	e := 22
	var d *int = &c
	fmt.Printf("address c: %v, d: %v, e: %v\n", &c, d, &e) // same address for c and d
	fmt.Printf("value c: %v, d: %v\n", c, *d)              //print out the value stored in that same address
	c++
	fmt.Printf("after c++, now c: %v, d: %v\n", c, *d)
	*d++
	fmt.Printf("after d++, now c: %v, d: %v\n", c, *d)

	fmt.Println("---arr example----")
	arr := [3]int{1, 2, 3}
	p1 := &arr[0]
	p2 := &arr[1]
	fmt.Printf("address arr: %v, p1: %v, p2: %v\n", &arr, p1, p2) // same address for c and d
	fmt.Printf("value arr: %v, p1: %v, p2: %v\n", arr, *p1, *p2)  //print out the value stored in that same address

	fmt.Println("---create pointers to objects example----")

	//1. using new keywords: cannot initialize the field at the same time
	var ms *myStruct
	ms = new(myStruct)
	//syntax sugar, original syntax is to first dereference the point then assign/obtain value from the struct it points to
	//(*ms).foo
	//fmt.Println((*ms).foo)
	ms.foo = 42
	fmt.Println(ms.foo)

	// 2. can use & if value type already exist
	ms2 := myStruct{foo: 41}
	fmt.Println(ms2.foo)

	//3. use & before initializer
	ms3 := &myStruct{foo: 40}
	fmt.Println(ms3.foo)

}

type myStruct struct {
	foo int
}
