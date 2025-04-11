package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	h := fnv.New32()
	fmt.Println(h)
	h.Write([]byte("asd"))
	fmt.Println(h.Sum32())
	h.Write([]byte("asd"))
	fmt.Println(h.Sum())

}
