package TreeMapDemo

import (
	"container/heap"
	"fmt"
)

type IntMinHeap []int

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	l := h.Len()
	x := (*h)[l-1]
	*h = (*h)[0 : l-1]
	return x
}

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func InitMinHeap() {
	h := &IntMinHeap{}
	heap.Init(h)
	heap.Push(h, 22)
	heap.Push(h, 44)
	heap.Push(h, 2)
	heap.Push(h, 6)
	heap.Push(h, 12)
	fmt.Println(h)
	heap.Pop(h)
	fmt.Println(h)
	heap.Pop(h)
	fmt.Println(h)
	heap.Pop(h)
	fmt.Println(h)
}
