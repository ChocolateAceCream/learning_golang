/*
* @fileName 2DSliceDemo.go
* @author Di Sheng
* @date 2022/06/10 11:05:33
* @description
 */

package main

import "fmt"

func main() {
	x := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("len(x): ", len(x)) // [[1,2,3],[4,5,6]]
	fmt.Println(x[0:2])             // [[1,2,3],[4,5,6]]
	fmt.Println(x[0:2][0])          // [1,2,3]
	column := []int{}
	for _, row := range x[0:2] {
		fmt.Println("row: ", row)
		column = append(column, row[0:1]...)
	}
	fmt.Println(column)

	var i int
	for i = 0; i < 10; i++ {
	}
	fmt.Println(i)

	arr := []int{1, 2, 3, 5, 6}
	changeArr(arr, 1)
	fmt.Println("arr: ", arr)
	arr4 := append(arr, 5)
	fmt.Println("arr4: ", arr4)
	fmt.Println("arr4: ", arr4)

	arr3 := [][]int{{1, 2, 3}}
	fmt.Println("arr3: ", arr3)          // arr3:  []
	fmt.Println("len arr3: ", len(arr3)) // arr3:  []

}

func changeArr(arr []int, i int) {
	arr[i] = 100
	arr2 := make([]int, len(arr))
	copy(arr2, arr)
	if i != 2 {
		changeArr(arr2, 2)
	}
}
