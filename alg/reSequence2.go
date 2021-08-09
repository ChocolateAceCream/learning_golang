package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "asdfgasdfgasdfgasdfgasdfgasdfgasdfgasdfgasdfg"
	str2 := "gfdsagfdsagfdsagfdsagfdsagfdsagfdsagfdsagfdsa"
	fmt.Println(reSequence(str1, str2))
}

func reSequence(s1, s2 string) bool {
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
