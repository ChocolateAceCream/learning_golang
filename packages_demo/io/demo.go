package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("cnt/123", 0744)

	if err != nil {
		fmt.Println(err)
	}

	// create the file and its path folder if not exist
	os.MkdirAll("ccc/aaa/bbb", 0744)

	// remove dir and all its children
	err = os.RemoveAll("cnt")
	if err != nil {
		fmt.Println(err)
	}
}
