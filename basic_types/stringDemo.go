/*
* @fileName stringDemo.go
* @author Di Sheng
* @date 2022/06/18 10:56:42
* @description
 */

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	SpecialCaseDemo()
	stringIterationDemo()
	StringSplitDemo()
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
