package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "aAbBcC"
	str2 := "a@@@AbBcC"
	fmt.Println(isStringLetterOnly(str))
	fmt.Println(isStringLetterOnly(str2))
	s := "21"
	fmt.Println(unicode.IsNumber(s))
}

func isStringLetterOnly(str string) bool {
	for _, v := range str {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}
