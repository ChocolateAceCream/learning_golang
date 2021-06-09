package main

import (
	"fmt"
)

//when struct member key is in lower case, member is package level, you can see struct(since it's capitalized) but cannot access to its member
type Car struct {
	number int
	make   string
	plate  []int
	car    Car2
}

// when capitalize the member key, the struct members will turn to global as well
type Car2 struct {
	Number int
	Make   string
	Plate  []int
}

func main() {
	car := Car2{
		Number: 4,
		Make:   "bbbb",
		Plate:  []int{4, 5, 6},
	}
	s := Car{
		number: 3,
		make:   "chevy",
		plate:  []int{1, 2, 3},
		car:    car,
	}
	fmt.Println("struc is: ", s)

	//retrive key of struct
	fmt.Println("make is ", s.make)
	fmt.Println("last plate is ", s.plate[len(s.plate)-1])

	// 	positional syntax: no need to specify the key, but you have to assign value by orders
	//	pros: less code
	s2 := Car2{
		23,
		"1chevy",
		[]int{11, 22, 33},
	}
	fmt.Println("s2 is ", s2)

	//struct is passing by data, when passing struct around, you are passsing a copy of data
	//p.s. when member in type of array etc.. still passing by reference
	s3 := s2
	s3.Plate[0] = 1111
	fmt.Println("-------------alter array member-----------")
	fmt.Println("s3 is ", s3)
	fmt.Println("s2 is ", s2)

	s3.Make = "bbb"
	fmt.Println("-------------alter normal member-----------")
	fmt.Println("s3 is ", s3)
	fmt.Println("s2 is ", s2)

	//if you want to pass struct by reference, use & like arrays
	s4 := &s3
	s4.Make = "s4"
	fmt.Println("s4 is ", s4)
	fmt.Println("s3 is ", s3)

	//struct in struct
	fmt.Println("-------------struct in struct-----------")
	s5 := s
	s6 := Car2{
		Number: 123,
		Make:   "s6",
	}
	fmt.Println("s is ", s)
	fmt.Println("s5 is ", s5)
	s5.car = s6
	fmt.Println("after change, s is ", s)
	fmt.Println("after change, s5 is ", s5)
	//anonymous struct
	//use case: generate a JSON res for a service call
	fmt.Println("-------------anonymous struct-----------")
	singleton := struct{ name string }{name: "john Wick"}
	fmt.Println("anonymous struct is ", singleton)
	fmt.Println("anonymous struct name is ", singleton.name)

}
