package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Int() % 10000
	fmt.Println(a)
}
