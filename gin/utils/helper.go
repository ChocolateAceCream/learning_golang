package utils

import (
	"fmt"
	"math/rand"
)

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

func RandToken(length int) string {
	result := make([]rune, length)
	var letters = []rune("abcdefghijklmnopqrstu1234567890vwxyz-_:!$ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
