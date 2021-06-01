package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fileSize := 40000000000.0
	fmt.Printf("%0.2fGB\n", fileSize/GB)
}
