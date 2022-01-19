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

	s1 := []int{1, 2, 3}
	s2 := []int{2, 3, 4}
	// cannot use slice directly as second param in append()
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
