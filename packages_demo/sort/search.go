package sort_demo

import (
	"fmt"
	"sort"
)

//demo for sort pkg usage

//1. sort.Search: func Search(n int, f func(int) bool) int
// return the first idx i in [0,n) at which f(i) is true

func SortSearchDemo() {
	arr := []int{41, 63, 12, 42, 77, 11}
	sort.Ints(arr)
	fmt.Println("sorted arr: ", arr)

	result := sort.Search(len(arr), func(a int) bool {
		return arr[a] > 44
	})
	fmt.Println("first element that greater than 44 in sorted arr is: ", arr[result])
}

func SortByDemo() {
	arr := []int{41, 63, 12, 42, 77, 11}
	sort.Ints(arr)
	fmt.Println(arr)
}

type Person struct {
	Age  int
	Name string
}

type SortablePerson []Person

func (s SortablePerson) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}
func (s SortablePerson) Swap(i, j int) {
	if s[i].Age < s[j].Age {
		s[i], s[j] = s[j], s[i]
	}
}
func (s SortablePerson) Len() int {
	return len(s)
}

func SortByKeyDemo(p []Person) {
	sort.Sort(SortablePerson(p))
	fmt.Println(p)
}
