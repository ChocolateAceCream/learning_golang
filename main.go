package main

import (
	"fmt"
)

type Demo interface {
	Logger()
	With(d, string)
}

type d struct {
	t string
}

func (d1 d) Logger() {
	fmt.Println(d1.t)
}
func (d1 d) With(d2 d, s string) d {
	d1.t = s
	return d1
}

var logger d

func main() {
	T1()
	T2()
}

func T1() {
	fmt.Println("T1")

	// Block 1 without {}
	logger.Logger()
	logger = logger.With(logger, "Hello")
	fmt.Println("After Block 1:")
	logger.Logger()

	// Block 2 without {}
	logger.Logger()
	logger = logger.With(logger, "World")
	fmt.Println("After Block 2:")
	logger.Logger()

	fmt.Println("T1 end")
}

func T2() {
	fmt.Println("T2")

	// Block 1 with {}
	{
		logger.Logger()
		logger = logger.With(logger, "123")
	}
	fmt.Println("After Block 1:")
	logger.Logger()

	// Block 2 with {}
	{
		logger.Logger()
		logger = logger.With(logger, "456")
	}
	fmt.Println("After Block 2:")
	logger.Logger()

	fmt.Println("T2 end")
}
