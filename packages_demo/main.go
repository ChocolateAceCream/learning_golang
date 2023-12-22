package main

import (
	"fmt"
	TreeMapDemo "package_demo/treeMapDemo"
)

func main() {
	fmt.Println("-----min heap----")
	TreeMapDemo.InitMinHeap()
	fmt.Println("-----max heap----")
	TreeMapDemo.InitMaxHeap()
	fmt.Println("-----priority queue----")
	TreeMapDemo.InitPQ()
}
