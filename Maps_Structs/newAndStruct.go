package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	v1 := new(Student)
	v1.name = "v1"
	v1.id = 1

	v2 := Student{name: "v2", id: 2}
	fmt.Println(v1, v2)
}
