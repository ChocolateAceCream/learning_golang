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
