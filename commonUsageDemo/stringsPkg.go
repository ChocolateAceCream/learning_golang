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

	//strings.Index(str, substr) int
	//Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	fmt.Println("------strings.Index------")
	str = "asdfzx"
	fmt.Println(strings.Index(str, "d"))
	fmt.Println(strings.Index(str, "asd"))
	fmt.Println(strings.Index(str, "aasd")) //-1

	//func LastIndex(s, substr string) int
	//LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
	fmt.Println("------strings.LastIndex------")
	str = "asdfdsa"
	fmt.Println(strings.LastIndex(str, "d"))
	fmt.Println(strings.LastIndex(str, "aasd")) //-1

	// func Replace(s, old, new string, n int) string
	//Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
	// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
	// If n < 0, there is no limit on the number of replacements.
	fmt.Println("------strings.Replace-------------")
	str = "aa bb cc dd ee aa bb cc dd aa"
	fmt.Println(strings.Replace(str, "aa", "AA", 1))
	fmt.Println(strings.Replace(str, "aa", "AA", 2))
	fmt.Println(strings.Replace(str, "aa", "AA", -1))
	fmt.Println(strings.Replace(str, "", "AA", -1))
}