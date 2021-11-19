package main

import (
	"fmt"
	"strconv"
)

func main() {
	//convert int to string
	println("a" + strconv.Itoa(123)) // a32

	//convert string to int
	i, err := strconv.Atoi("123")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println(1 + i)
	}

	//Parse: parse string to target type
	//e.g. ParseBool()、ParseFloat()、ParseInt()、ParseUint()

	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("%v, %T\n", b, b)
	}

	//64 refers to the bitsize
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("%v, %T\n", f, f)
	}

	g, err := strconv.ParseFloat("infinity", 32)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("%v, %T\n", g, g)
	}

	// second argument refer to the base, if base = 0, then parse the string based on its prefix
	//e.g. 0x means base 16, 0means base 8, otherwise use base 10
	i2, _ := strconv.ParseInt("10000", 2, 64)
	fmt.Printf("%v, %T\n", i2, i2)
	i3, _ := strconv.ParseInt("0x12", 0, 64)
	fmt.Printf("%v, %T\n", i3, i3)
	i3, _ = strconv.ParseInt("012", 0, 64)
	fmt.Printf("%v, %T\n", i3, i3)

	//Append: convert type to string then append result to a slice
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
	fmt.Println(strconv.IntSize)

}
