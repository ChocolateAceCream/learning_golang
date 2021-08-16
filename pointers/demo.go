package main

import "fmt"

type Location struct {
	x int
	y int
}

func main() {
	l := Location{x: 0, y: 0}
	fmt.Println("-----modifyPointer------")
	modifyPointer(&l)
	fmt.Println(l)
	fmt.Println("-----modifyReference------")
	l = Location{x: 0, y: 0}
	modifyReference(&l)

	fmt.Println(l)
}

func modifyPointer(ptr *Location) {
	ptr.x = 1
	ptr.y = 2
}

//the pointer you modified here is a copy of the pointer passed in arguments
//so modify the value of ptr will not affect original value in main()
func modifyReference(ptr *Location) {
	ptr = &Location{x: 4, y: 4}
}
