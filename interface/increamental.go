package main

import "fmt"

func main() {
	var i IntCounter = IntCounter(0)
	var f FloatCounter = FloatCounter(1.0)
	for n := 0; n < 10; n++ {
		i.Increase()
		f.Increase()
		fmt.Println("---i-----f--", i, f)
	}
}

//interface describe behaviors, store method definitions
type Increaser interface {
	Increase(int) (int, error) //a method take []bytes as input and return int and error
}

type IntCounter int
type FloatCounter float64

func (i *IntCounter) Increase() {
	*i++
}

func (f *FloatCounter) Increase() {
	*f = *f * 2
}
