package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

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
	fmt.Println("---------change slice demo---------")
	s4 := []int{10, 20, 33}
	ChangeSlice(s4)
	fmt.Println(s4)

	fmt.Println("---------sort slice demo---------")
	sortSlice()
	pointerDemo()
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

// Caution: change slice in function will also change the original slice
func ChangeSlice(slice []int) {
	slice[1] = 1000
}

func sortSlice() {
	slice := []int{5, 1, 87, 42, 3}
	sort.Ints(slice)
	fmt.Println(slice)
}

func pointerDemo() {
	fmt.Println("---------pointerDemo---------")
	tmp2 := make([]int, 10)
	fmt.Println("tmp2: ", tmp2)
	fmt.Println("tmp2 length: ", len(tmp2))
	tmp := make([]int, 0)
	fmt.Println("tmp: ", tmp)
	fmt.Println("tmp length: ", len(tmp))
	fmt.Println("type: ", reflect.TypeOf(tmp))
	ptr := &tmp
	fmt.Println("type: ", reflect.TypeOf(ptr))
	*ptr = append(*ptr, 1, 2, 3, 4)
	fmt.Println("*ptr: ", *ptr)
	item := make([]int, len(*ptr))
	fmt.Println("item: ", item)
	fmt.Println("type: ", reflect.TypeOf(item))
	copy(item, *ptr)
	fmt.Println("after copy item: ", item)
	fmt.Println("type: ", reflect.TypeOf(item))

	// remove last two elements
	*ptr = (*ptr)[:len(*ptr)-2]
	fmt.Println("*ptr: ", *ptr)

	fmt.Println("-----------SortAndRemoveDup-----------")
	s := []int{61}
	s = SortAndRemoveDup(s)
	fmt.Println("s: ", s)

	s4 := []int{1, 2, 3, 4}
	s4 = s4[:len(s4)-1]
	fmt.Println("s4: ", s4)

	fmt.Println("-----------FilterDemo-----------")
	s5 := []int{3, 2, 1, 4, 4, -1, 1}
	fmt.Println("s5 after filter: ", FilterDemo(s5))
	s5 = []int{2, 1, 4, 2, 4, 4, -1, 6}
	fmt.Println(firstMissingPositive(s5))

}

func SortAndRemoveDup(s []int) []int {
	sort.Ints(s)
	prev := 1
	for curr := 1; curr < len(s); curr++ {
		if s[curr-1] != s[curr] {
			s[prev] = s[curr]
			prev++
		}
	}
	return s[:prev]
}

func FilterDemo(s []int) []int {
	tmp := make([]int, 0)
	for _, v := range s {
		if v > 0 {
			fmt.Println("v: ", v)
			if len(tmp) > 0 {
				for i := 0; i < len(tmp); i++ {
					if v < tmp[i] {
						tmp = append(tmp[:i+1], tmp[i:]...)
						tmp[i] = v
						break
					} else if v == tmp[i] {
						break
					} else {
						if i == len(tmp)-1 {
							tmp = append(tmp, v)
							break
						}
					}
				}
			} else {
				tmp = append(tmp, v)
			}
		}
	}
	return tmp
}

func firstMissingPositive(nums []int) int {
	sz := len(nums)

	for i := 0; i < sz; i++ {
		// ensure no negative
		if nums[i] <= 0 {
			nums[i] = sz + 1
		}
	}
	fmt.Println("nums: ", nums)

	for i := 0; i < sz; i++ {
		n := int(math.Abs(float64(nums[i])))

		if n > sz {
			continue
		}

		if nums[n-1] > 0 {
			nums[n-1] *= -1
		}
	}

	fmt.Println("nums: ", nums)

	for i := 1; i <= sz; i++ {
		if nums[i-1] > 0 {
			return i
		}
	}

	return sz + 1
}
