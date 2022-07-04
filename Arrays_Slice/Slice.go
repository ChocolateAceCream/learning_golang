package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
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
	RuneDemo()
	SliceToArrDemo()
	RotateNToFront()
	ReverseSliceDemo()
}

func SliceToArrDemo() {
	fmt.Println("---------SliceToArrDemo---------")

	// need go version >= 1.7
	//convert slice to array
	s := make([]byte, 2, 4)
	fmt.Println("type: ", reflect.TypeOf(s))
	fmt.Println(s)
	s0 := (*[2]byte)(s) // s0 != nil
	fmt.Println(s0)
	fmt.Println("type: ", reflect.TypeOf(s0))
	s3 := [3]uint8{1, 2, 3}
	fmt.Println("type: ", reflect.TypeOf(&s3))
	// s1 := (*[1]byte)(s[1:]) // &s1[0] == &s[1]
	// s2 := (*[2]byte)(s)     // &s2[0] == &s[0]

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

	fmt.Println("-----------iteration demo-----------")
	// keep caution when append slice to a slice, because slice store pointers, not value, so once the appended slice is changed, it will also affect the original slice
	ss := []int{1, 2, 3, 4}
	fmt.Println("ss: ", ss[:len(ss)-1])
	for i := 0; i < len(ss); i++ {
		c := make([]int, len(ss))
		copy(c, ss)
		c = append(c[:i], c[i+1:]...)
		fmt.Println(c)
	}

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

func RuneDemo() {
	fmt.Println("-----------RuneDemo-----------")

	x := strings.Repeat("0", 10)
	fmt.Println("x: ", x)
	r := []rune{}

	fmt.Println("r: ", r)

	for i := 1; i < 10; i++ {
		r = append(r, rune(i+'0'))
	}
	fmt.Println("r: ", string(r))
}

func RotateNToFront() {
	fmt.Println("-----------RotateNToFront-----------")
	//e.g [1,2,3,4,5,6], rotate 4 to front => [4,1,2,3,5,6]
	s := []int{1, 2, 3, 4, 5, 6}
	start := 2
	end := 4
	tmp := s[end]
	for i := end; i > start; i-- {
		s[i] = s[i-1]
	}
	s[start] = tmp
	fmt.Println(s)

}

func ReverseSliceDemo() {
	fmt.Println("-------------reverse slice demo-------------")
	s := []int{1,5,51,5,5125}
	fmt.Println("s: ", s)
	ReverseSlice(s)
	fmt.Println("s after reverse: ", s)
}
func ReverseSlice[T comparable](s []T) {
    sort.SliceStable(s, func(i, j int) bool {
        return i > j
    })
}