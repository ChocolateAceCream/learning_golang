package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	m := map[int]int{}
	for k := range slice {
		m[k] = slice[k]
	}
	for key, value := range m {
		fmt.Println(key, "->", value)
	}
}
