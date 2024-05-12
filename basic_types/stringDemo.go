/*
* @fileName stringDemo.go
* @author Di Sheng
* @date 2022/06/18 10:56:42
* @description
 */

package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	SpecialCaseDemo()
	stringIterationDemo()
	StringSplitDemo()
	StringCountRepeatCharDemo()
	StringRepeatDemo()
	StringSliceSortDemo()
}

func SpecialCaseDemo() {
	fmt.Println("-------StringSortDemo------")
	StringSortDemo()
	fmt.Println("-------SpecialCaseDemo------")
	s1 := ""
	s2 := ""
	fmt.Println("s1 == s2 ? ", s1 == s2)
	fmt.Println("len(s1)", len(s1))
	s3 := "0123"
	s3 = s3[len(s3):] // return nothing, s3 is holding nil value
	fmt.Println("s3: ", s3)
	fmt.Println("s3 == s1 ? ", s3 == s1)
	regex := "***"
	x := strings.Repeat("*", len(regex))
	fmt.Printf("regex: %v, x: %v, regex == x ? %v \n", regex, x, regex == x)

	a := "012"
	fmt.Printf("a: %v, a': %v\n", a, a[3:])
	a = "*ac"
	fmt.Println("a[0] == 'a'", a[0] == 'a')

	for idx, val := range a {
		if val != '*' {
			fmt.Println(a[idx:])
			break
		}
	}

	// c := "b"
	// d := c[3:] // will cause error
	// fmt.Println("d: ", d)

}

func stringIterationDemo() {
	fmt.Println("-------stringIterationDemo------")
	s := ""
	for idx, val := range s {
		fmt.Printf("idx: %v, value: %v, s[idx]: %v\n", idx, val, s[idx])
	}
}

func StringSortDemo() {
	str := "msaduih"
	fmt.Println("str before sort:", str)
	tmp := []byte(str)
	sort.Slice(tmp, func(a, b int) bool { return tmp[a] < tmp[b] })
	str = string(tmp)
	fmt.Println("tmp after sort:", str)
}

func StringSplitDemo() {
	fmt.Println("-------StringSplitDemo------")
	s := "Welcome, to the, online portal, of GeeksforGeeks"
	res1 := strings.Split(s, " ")
	fmt.Println(res1)
}

func StringCountRepeatCharDemo() {
	fmt.Println("-------StringCountRepeatCharDemo------")
	s := "-"
	res1 := strings.Count(s[1:], "1")
	fmt.Println("s: ", s[1:])
	fmt.Println("how many 1 in s: ", res1)
}

func StringRepeatDemo() {
	fmt.Println("-------StringRepeatDemo------")
	s := "a"
	s2 := strings.Repeat(s, 3)
	fmt.Println(s2)
}

func StringSliceSortDemo() {
	fmt.Println("-------StringSliceSortDemo------")
	s := []string{"a", "c", "b"}
	sort.Strings(s)
	fmt.Println(s)
}

func StringBuilder() string {
	builder := strings.Builder{}
	for i := 0; i < 100000; i++ {
		builder.WriteString("data ")
		builder.Write([]byte{'1', '2'})
		builder.WriteByte('3')

	}
	return builder.String()
}
func StringConcatenation() string {
	s := ""
	for i := 0; i < 100000; i++ {
		s += "data "
	}
	return s
}

func BitOperation() {
	var a uint8 = 1
	a = a << 1 // shift a to left 1 time
	b := a << 2
	fmt.Println(a)
	fmt.Println(b)
	c := 'c'
	fmt.Println(reflect.TypeOf(c))
	m := map[int32]int{}
	for i := 'a'; i <= 'z'; i++ {
		m[i] = 1 << (i - 'a')
	}
	fmt.Println(m)
}
