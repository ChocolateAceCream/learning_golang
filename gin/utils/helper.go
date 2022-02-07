package utils

import "fmt"

// check if a slice of words includes certain word
func Contains(sliceOfWords []string, singleWord string) bool {
	for _, v := range sliceOfWords {
		fmt.Println("value -> ", v)
		if v == singleWord {
			return true
		}
	}
	return false
}
