package main

import "fmt"

//the whole func(int) int is the closure function that returned by adder()
// equalized to func adder() (func() int) {
func adder() func(*int) int {
	sum := 0
	var a *int = &sum
	//same here, the whole func(x int) int {} is the return value
	return func(x *int) int {
		*a++
		return *a
	}
}

func main() {
	pos, neg := adder(), adder()
	b := 2
	var d *int = &b
	fmt.Println(pos(d), neg(d))
}
