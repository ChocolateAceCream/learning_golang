package main

import "fmt"

func main() {

	fmt.Println("---------start---------")
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
	Iterator()

}

// when slice contains large potion of data and value is of type string, try to use slice[key] to retrieve the value in loop, because every loop iteration will re-assign to the value and key
//e.g.
func RangeSlice(slice []int) {
	for index := range slice {
		fmt.Println("value: ", slice[index])
	}
}

// i and val are all settle down when iteration starts, no matter what changed to the original slice
func Iterator() {
	v := []int{1, 2, 3}
	for i, val := range v {
		v = append(v, i)
		v[2] = 20
		fmt.Println("----val---", val)
	}
	fmt.Println("----v---", v)
}

func CheckSliceKeyExist() {
	m1 := map[int]int{}
	m1[1] = 2
	fmt.Println(m1[23])
	if v, ok := m1[22]; ok {
		fmt.Println("---------key exist---------", v)

	} else {
		fmt.Println("---------key not exist---------")
	}
}
