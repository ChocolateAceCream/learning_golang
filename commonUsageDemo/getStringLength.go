package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "ABCDEFAGHIJKLMNOPQRST"
	fmt.Println(str[1:2])
	str2 := "BBBBB"
	fmt.Println(strings.Count(str2, "BB"))
	//Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
	fmt.Println(strings.Count("世界⌘", ""))
	fmt.Println(len("⌘"))
	fmt.Println("Hello, 世界", len("世界"), utf8.RuneCountInString("世界"))

	b := "这是个测试⌘"
	len1 := len([]rune(b))
	len2 := bytes.Count([]byte(b), nil) - 1
	len3 := strings.Count(b, "") - 1
	len4 := utf8.RuneCountInString(b)
	fmt.Println(len1)
	fmt.Println(len2)
	fmt.Println(len3)
	fmt.Println(len4)
}
