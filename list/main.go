package main

import (
	"fmt"
	L "list/shared"
)

func main() {
	L.Demo()
	// mergeDemo()
	findCommonNode()
}

/*
Q: 输入两个链表，找出它们的第一个公共结点
*/
func findCommonNode() {
	l := L.InitList()
	l.AddBack(1)
	l.AddBack(3)
	l.AddBack(5)
	l.AddBack(7)

	l2 := L.InitList()
	l2.AddBack(2)
	l2.AddBack(4)

	l3 := L.InitList()
	l3.AddBack(12)
	l3.AddBack(14)
	l3.AddBack(16)

	l.Tail.Next = l3.Head
	l.Tail = l3.Tail

	l2.Tail.Next = l3.Head
	l2.Tail = l3.Tail

	fmt.Println("---list 1----")
	l.PrintList()
	fmt.Println("---list 2----")
	l2.PrintList()

	pt1 := l.Head
	pt2 := l2.Head
	for pt1 != pt2 {
		pt1 = pt1.Next
		pt2 = pt2.Next
		if pt1 == nil {
			pt1 = l2.Head
		}
		if pt2 == nil {
			pt2 = l.Head
		}
	}
	fmt.Println("common node: ", pt1.Value)
}

/*
Q:
输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按照递增排序的，例如：
链表1：1 -> 3 -> 5 -> 7
链表2: 2 -> 4 -> 6 -> 8
合并后的链表3:
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8
*/
func mergeDemo() {
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
