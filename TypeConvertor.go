package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func demo() {
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

	//string formatting
	fmt.Println("stirng formatting wiht fmt.Sprintf")
	const name, age = "kim", 22
	f1 := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(f1)

	//rune
	fmt.Println("\n------rune-----------")
	s3 := '⌘'
	fmt.Printf("%v, %T\n", s3, s3)

	b1 := []byte("1,2,3,4,5")
	for i := 0; i < len(b1); i++ {
		if b1[i] != ',' {
			fmt.Println(string(b1[i]))
		}
	}

	s5 := "21"
	fmt.Println("convertor: ", convertor(s5))

	i = 123
	i *= -1
	fmt.Println("type: ", reflect.TypeOf(i))
	i = int(math.Abs(float64(i)))
	fmt.Println("i: ", i)
	fmt.Println("type: ", reflect.TypeOf(i))

	ConvertIntSliceToString()
	ConvertStringBinaryToBase2Int64()
}

func convertor(s string) string {
	var r []uint8
	var count = '0'
	curr := s[0]
	for i := 0; i < len(s); i++ {
		if curr == s[i] {
			count++
		} else {
			r = append(r, uint8(count), curr)
			count = '1'
			curr = s[i]
		}
	}
	if count > 0 {
		r = append(r, uint8(count), curr)
	}
	return string(r)
}

func ConvertIntSliceToString() {
	fmt.Println("------ConvertIntSliceToString demo-----------")
	i := []int{1, 2, 3, 4, 5}
	s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(i)), ""), "[]")
	fmt.Println(s)
}

// limitation: be careful when value is over 64 byte
func ConvertStringBinaryToBase2Int64() {
	fmt.Println("------ConvertStringBinaryToBase2Int64 demo-----------")
	s := "101010101"
	fmt.Println("s: ", s)
	r, _ := strconv.ParseInt(s, 2, 64)
	fmt.Println("after convert to base 2 Int64: ", r)
	fmt.Println("print binary form: ", strconv.FormatInt(r, 2))
}
