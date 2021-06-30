package main

import "fmt"

func main() {
	r, err := divideFunc(1.4, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}

func divideFunc(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}
