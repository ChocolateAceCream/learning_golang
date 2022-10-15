package main

import "fmt"

func main() {
	AnonymousStructureDemo()
	AnonymousFieldDemo()
	InterfaceAsAnonymousFieldDemo()
	StructAsAnonymousFieldDemo()
}

// Anonymous Structure is used to create a one-time use struct, cheaper and safer than using map[string]interface{}.
func AnonymousStructureDemo() {
	fmt.Println("---------AnonymousStructureDemo----------------")
	p := struct {
		name string
		age  int
	}{
		name: "asdf",
		age:  11,
	}

	fmt.Println(p)
}

type CP struct {
	string
	int
}

// Anonymous Field: use type as the name of the field.
// Caution: cannot use same type of anonymous fields in one struct!
func AnonymousFieldDemo() {
	fmt.Println("---------Anonymous Field Demo----------------")
	c := CP{"cobal", 12312}
	fmt.Println(c)
	fmt.Println("c.string: ", c.string)
	fmt.Println("c.int: ", c.int)
}

/*
	use interface as struct anonymous field.
	Benefits:
	1. flexible, you can rewrite the implementation on anything that implemented the interface
	2. restrained, when construct a struct instance, that struct must implement the anonymous interface. However, that struct can also contain other fields
	e.g.
	type Sortable struct {
		sort.Interface
		// other field
		Type string
	}


*/

type ArrInterface interface {
	Len() int
	Less(int, int) bool
}

type MyArr []int

func (myArr MyArr) Len() int {
	return len(myArr)
}

func (myArr MyArr) Less(i, j int) bool {
	return myArr[i] < myArr[j]
}

// now MyArr implement ArrInterface
type reverse struct {
	ArrInterface
}

func (r reverse) Less(i, j int) bool {
	return r.ArrInterface.Less(j, i)
}

func NewReverse(myArr MyArr) ArrInterface {
	return &reverse{myArr}
}

func InterfaceAsAnonymousFieldDemo() {
	fmt.Println("---------Use Interface as Anonymous Field Demo----------------")
	arr1 := MyArr{1, 2, 3, 4}
	arr2 := NewReverse([]int{1, 2, 3})
	fmt.Println("arr1: ", arr1.Less(0, 1))
	fmt.Println("arr2: ", arr2.Less(0, 1))
}

// ----------------use anonymous struct as field demo--------------------

type reverse2 struct {
	MyArr
}

func NewReverse2(arr MyArr) ArrInterface {
	return &reverse2{arr}
}

func (r reverse2) Less(i, j int) bool {
	return r.MyArr.Less(j, i)
}
func StructAsAnonymousFieldDemo() {
	fmt.Println("---------Use Struct as Anonymous Field Demo----------------")
	arr3 := MyArr{1, 2, 3, 4}
	arr4 := NewReverse2(MyArr{1, 2, 3})
	fmt.Println("arr3: ", arr3.Less(0, 1))
	fmt.Println("arr4: ", arr4.Less(0, 1))
}
