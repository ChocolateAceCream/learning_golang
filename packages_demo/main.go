package main

import (
	"fmt"
	Alg "learning_go/alg"
)

func main() {
	m := Alg.InitMinHeap()
	fmt.Println(m.Heap)
	m.Push(15)
	m.Push(99)
	m.Push(22)
	m.Push(55)
	m.Push(42)
	m.Push(59)
	m.Push(5)
	fmt.Println(m.Heap)
	m.Pop()
	fmt.Println(m.Heap)
	m.Pop()
	fmt.Println(m.Heap)
}
