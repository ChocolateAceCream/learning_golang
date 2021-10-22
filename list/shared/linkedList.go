package shared // All the files inside lib will have the same package name

import "fmt"

// This func must be Exported, Capitalized, and comment added.
func Demo() {
	fmt.Println("HI")
}

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func InitList() *List {
	//create an empty List struct and return its address
	return &List{}
}

func (l *List) AddFront(value interface{}) {
	node := &Node{
		Value: value,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	return
}

func (l *List) AddBack(value interface{}) {
	node := &Node{
		Value: value,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = l.Tail.Next
	}
	return
}

func (l *List) PrintList() error {
	fmt.Println("----start printing list---")
	if l.Head == nil {
		return fmt.Errorf("----list is empty---")
	}
	current := l.Head
	for current.Next != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Printf("%v\n", current.Value)
	return nil
}

func (l *List) Size() int {
	step := 0
	c := l.Head
	for c != nil {
		step++
		c = c.Next
	}
	return step

}

func (l *List) RemoveFront() {
	if l.Size() < 1 {
		fmt.Println("list is empty")
	} else {
		l.Head = l.Head.Next
	}
}

func (l *List) RemoveBack() {
	if l.Size() < 1 {
		fmt.Println("list is empty")
	} else {
		current := l.Head
		for current.Next != l.Tail {
			current = current.Next
		}
		l.Tail = current
	}
}

func MergeList(l1 *List, l2 *List) *List {
	result := InitList()
	c1 := l1.Head
	c2 := l2.Head
	for c1 != nil && c2 != nil {
		if c1.Value.(int) > c2.Value.(int) {
			result.AddBack(c2.Value)
			c2 = c2.Next
		} else {
			result.AddBack(c1.Value)
			c1 = c1.Next
		}
	}
	if c1 == nil {
		fmt.Println("c2", c2.Value)
		result.Tail.Next = c2
	} else {
		fmt.Println("c1", c1.Value)
		result.Tail.Next = c1
	}
	result.Tail = result.Tail.Next
	fmt.Println("---size---", result.Size())
	return result
}
