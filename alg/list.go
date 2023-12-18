package Alg

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func initList() *List {
	//create an empty List struct and return its address
	return &List{}
}

func (l *List) AddFront(value interface{}) {
	node := &Node{
		value: value,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
	return
}

func (l *List) AddBack(value interface{}) {
	node := &Node{
		value: value,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = l.tail.next
	}
	return
}

func (l *List) PrintList() error {
	fmt.Println("----start printing list---")
	if l.head == nil {
		return fmt.Errorf("----list is empty---")
	}
	current := l.head
	for current.next != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Printf("%v\n", current.value)
	return nil
}

func (l *List) Size() int {
	step := 0
	c := l.head
	for c != nil {
		step++
		c = c.next
	}
	return step

}

func (l *List) RemoveFront() {
	if l.Size() < 1 {
		fmt.Println("list is empty")
	} else {
		l.head = l.head.next
	}
}

func (l *List) RemoveBack() {
	if l.Size() < 1 {
		fmt.Println("list is empty")
	} else {
		current := l.head
		for current.next != l.tail {
			current = current.next
		}
		l.tail = current
	}
}

func Runner() {
	l := initList()
	l.AddBack(1)
	l.AddBack(3)
	l.AddBack(5)
	l.AddBack(7)
	l.AddBack(10)
	l.AddBack(15)
	l2 := initList()
	l2.AddBack(2)
	l2.AddBack(4)
	l2.AddBack(6)
	l2.AddBack(8)
	l.PrintList()
	l2.PrintList()
	fmt.Println("---result----")
	r := mergeList(l, l2)
	r.PrintList()
}

func mergeList(l1 *List, l2 *List) *List {
	result := initList()
	c1 := l1.head
	c2 := l2.head
	for c1 != nil && c2 != nil {
		if c1.value.(int) > c2.value.(int) {
			result.AddBack(c2.value)
			c2 = c2.next
		} else {
			result.AddBack(c1.value)
			c1 = c1.next
		}
	}
	if c1 == nil {
		fmt.Println("c2", c2.value)
		result.tail.next = c2
	} else {
		fmt.Println("c1", c1.value)
		result.tail.next = c1
	}
	result.tail = result.tail.next
	fmt.Println("---size---", result.Size())
	return result
}
