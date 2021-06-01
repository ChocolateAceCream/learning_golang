package main

import (
	"fmt"
	"strconv"
)

func f() {
	var i int = 32
	var j float32 = float32(i)
	fmt.Printf("%v, %f, %T", j, j, j)

	k, err := strconv.Atoi("12")
	if err == nil {
		fmt.Printf("%v,%T", k, k)
	}

	//boolean type
	fmt.Println("\n------boolean-----------")
	var boolean bool = true
	fmt.Printf("%v, %T", boolean, boolean)
	m := 1 == 2
	fmt.Printf("%v, %T", m, m)

	//int
	fmt.Println("\n------integer-----------")
	var a int16 = 14
	var b int8 = 5
	fmt.Printf("%v, %T", int8(a)/b, int8(a)/b)

	fmt.Println("\n------bit shift-----------")
	c := 8
	fmt.Printf("%v\n", c<<3) // 2^3 * 2^3
	fmt.Printf("%v\n", c>>3) // 2^3 / 2^3

	//float
	fmt.Println("\n------float-----------")
	d := 3.14
	d = 12e300
	var dd float32 = 1.1
	fmt.Printf("%v, %T\n", d, d)

	//complex numbers
	fmt.Println("\n------complex numbers-----------")
	e := 1 + 2i
	var f complex128 = 2 + 5.1i
	fmt.Printf("%v\n", e+f)
	fmt.Printf("%v\n", e-f)
	fmt.Printf("%v\n", e*f)
	fmt.Printf("%v\n", e/f)
	fmt.Printf("%v %T\n", real(e), real(e)) //get real part, return float64 if complex128, return float32 if complex64
	fmt.Printf("%v %T\n", imag(e), imag(e)) //get imaginary part, return float64 if complex128, return float32 if complex64

	g := complex(dd, dd) //complex() can either take two float64 and return a complex128, or two float32 and return a complex64, but cannot take a float32 and a float64
	fmt.Printf("%v %T\n", g, g)

	//string
	fmt.Println("\n------String-----------")
	s := " 1 this is a string"
	fmt.Printf("%v, %T\n", s[1], s[1]) // will print 32, unit8 since string is alias for bits
	s2 := "       string number 2"
	fmt.Printf("%v\n", s+s2) // string contamination
	bytes := []byte(s)       // convert string into collection of bytes(ea)
	fmt.Printf("%v, %T\n", bytes, bytes)

	//rune
	fmt.Println("\n------rune-----------")
	s3 := '⌘'
	fmt.Printf("%v, %T\n", s3, s3)
}
