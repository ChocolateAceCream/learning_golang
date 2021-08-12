/*
字符串替换问题
问题描述
请编写⼀个⽅法，将字符串中的空格全部替换为“%20”。 假定该字符串有⾜够的空间存
放新增的字符，并且知道字符串的真实⻓度(⼩于等于1000)，同时保证字符串由【⼤⼩
写的英⽂字⺟组成】。 给定⼀个string为原始的串，返回替换后的string。
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "aaspkop SA AS AFA A AFA sd fa sf asdfasd fa sf asdfasd fa sf asdfasd fa sf asdfasd fa sf asdfasd fa sf asdfasd fa sf asdfasd fa sf asdfa"

	for i := 0; i < 10; i++ {
		r, str := replaceString(s)
		if r {
			fmt.Println(str)
		} else {
			fmt.Println("illegal string")
		}
	}
}

// func replaceString(str string) (bool, string) {
// 	s := strings.Split(str, "")
// 	if len(s) > 1000 {
// 		return false, str
// 	}
// 	for i, v := range s {
// 		if v == " " {
// 			s[i] = "%20"
// 		} else if !unicode.IsLetter([]rune(v)[0]) {
// 			return false, str
// 		}
// 	}
// 	return true, strings.Join(s, "")
// }

alternatively, we can use strings.Replace(str, " ", "%20", -1) method, but this approach is slower
func replaceString(str string) (bool, string) {
	s := []rune(str)
	if len(s) > 1000 {
		return false, str
	}
	for _, v := range str {
		if string(v) != " " && !unicode.IsLetter(v) {
			return false, str
		}
	}
	return true, strings.Replace(str, " ", "%20", -1)
}
