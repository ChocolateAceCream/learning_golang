package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("------e.g.1----------------")
	i := 22
	switch i {
	case 1, 2, 3, 4, 5:
		fmt.Println("1, 2, 3, 4, 5")
	case 22, 33, 44, 55: //test case must be uniq, unless you use comparion operators like example below
		fmt.Println("22, 33, 44, 55")
	default:
		fmt.Println("default")
	}

	fmt.Println("------e.g.2----------------")
	n := rand.Intn(100)
	n = 88
	switch {
	case n > 80 && n < 90:
		fmt.Println(" 80 < n 90", n)
		fallthrough //use fallthrough to jump to next case, no matter that case is failed or not
	case n < 20:
		fmt.Println("fallthrough case: n<20", n)
	default:
		fmt.Println("default", n)
	}

	fmt.Println("------e.g.2--type switch--------------")
	var i2 interface{} = [3]int{}
	switch i2.(type) {
	case int:
		fmt.Println("int")
		break
		fmt.Println("int") //this statement wont execute
	case float64, float32:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	case [2]int:
		fmt.Println("[2]int{}")
	case [3]int: //array must have same type and same length in order to be equal
		fmt.Println("[3]int{}")
	default:
		fmt.Println("sth else")
	}
}
