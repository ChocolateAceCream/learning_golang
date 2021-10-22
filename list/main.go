package main

import (
	"fmt"
	L "list/shared"
)

func main() {
	L.Demo()
	l := L.InitList()
	l.AddBack(1)
	l.AddBack(3)
	l.AddBack(5)
	l.AddBack(7)
	l.AddBack(10)
	l.AddBack(15)
	l2 := L.InitList()
	l2.AddBack(2)
	l2.AddBack(4)
	l2.AddBack(6)
	l2.AddBack(8)
	l.PrintList()
	l2.PrintList()
	fmt.Println("---result----")
	r := L.MergeList(l, l2)
	r.PrintList()
}
