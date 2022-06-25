/*
* @fileName bitOperationDemo.go
* @author Di Sheng
* @date 2022/06/22 22:39:27
* @description
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a uint8 = 1
	a = a << 1 // shift a to left 1 time
	b := a << 2
	fmt.Println(a)
	fmt.Println(b)
	c := 'c'
	fmt.Println(reflect.TypeOf(c))
	m := map[int32]int{}
	for i := 'a'; i <= 'z'; i++ {
		m[i] = 1 << (i - 'a')
	}
	fmt.Println(m)
}
