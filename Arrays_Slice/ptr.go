package main

import (
	"fmt"
	"os"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func main() {
	f, err := os.Open("Arrays_Slice/Demo.go")
	if err != nil {
		fmt.Println("err open the file:", err)
		return
	}
	defer f.Close()

}
