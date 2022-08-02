/*
* @fileName main.go
* @author Di Sheng
* @date 2022/07/29 09:41:29
* @description generics demo
 */

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	reverseArr(&arr)
	fmt.Println(arr)
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := MapDemo(m)
	fmt.Println(keys)
	StructDemo()

	UnionOperatorDemo()

	TypeApproximationDemo()
}

// any is equivalent to {}interface
func reverseArr[T any](arr *[]T) {
	fmt.Println("----------arr demo--------")
	l := len(*arr)
	for idx, val := range *arr {
		// cannot index array, so first dereference pointer.
		(*arr)[l-idx-1] = val
	}
}

// the key of map must be of type comparable, value could be any
func MapDemo[K comparable, V any](m map[K]V) []K {
	fmt.Println("----------MapDemo--------")
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// struct demo
type MyStruct[T any] struct {
	inner T
	outer string
}

func (m *MyStruct[T]) GetInner() T {
	return m.inner
}
func (m *MyStruct[T]) GetOuter() string {
	return m.outer
}

func (m *MyStruct[T]) SetInner(inner T) {
	m.inner = inner
}
func (m *MyStruct[T]) SetOuter(outer string) {
	m.outer = outer
}

func StructDemo() {
	fmt.Println("----------StructDemo--------")
	s := MyStruct[int]{
		inner: 123,
		outer: "abc",
	}
	fmt.Println(s.GetInner())
}

type Num interface {
	int | int8 | int16 | float32 | float64
}

// return min
func UnionOperatorDemo() {
	fmt.Println("----------UnionOperatorDemo--------")
	x := 1
	y := 2
	z := Min(x, y)
	fmt.Printf("x: %v, y: %v, min: %v\n", x, y, z)

}

func Min[T Num](x, y T) T {
	if x > y {
		return y
	} else {
		return x
	}
}

// Type approximation: ~ operators allow us to specify that interface also supports types with the same underlying types.

// Any Type with given underlying type will be supported by this interface
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Point int

func TypeApproximationDemo() {
	fmt.Println("----------TypeApproximationDemo--------")
	var x Point = 1
	var y Point = 2
	z := Max(x, y)
	fmt.Printf("x: %v, y: %v, max: %v", x, y, z)
}

func Max[T Number](x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}

// Nesting constrains
type Integer interface {
	~int | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

// Number is build from Integer and Float
type Num2 interface {
	Integer | Float
}
