package main

import (
	"fmt"
	"math"
)

func main() {
	n1 := 0.1
	n2 := 0.123
	fmt.Println("---------compare func 1---------------")
	compare1(n1) //true
	compare1(n2) // false since sqrt result will round up, which is not accurate
	fmt.Println("---------optimized func 2---------------")
	compare2(n1)
	compare2(n2)
}

func compare1(number float64) {
	if math.Pow(math.Sqrt(number), 2) == number {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}
}

func compare2(number float64) {
	if math.Abs(math.Pow(math.Sqrt(number), 2)/number-1) < 0.0001 {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}
}
