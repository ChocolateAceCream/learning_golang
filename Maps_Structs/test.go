package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	length int
	head   *Node
	tail   *Node
}

func initList() *List {
	fmt.Println("----list---", &List{})
	return &List{}
}

func main() {
	l := initList()
	node := &Node{
		value: 1,
	}
	l.head = node
	l.tail = node
	if l.head == l.tail {
		fmt.Printf("111")
	} else {
		fmt.Printf("222")
	}

}
