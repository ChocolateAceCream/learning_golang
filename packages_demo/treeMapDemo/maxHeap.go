package TreeMapDemo

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}
func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *MaxHeap) Pop() any {
	l := m.Len()
	x := (*m)[l-1]
	(*m) = (*m)[0 : l-1]
	return x
}
func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func InitMaxHeap() {
	mx := &MaxHeap{}
	heap.Init(mx)
	heap.Push(mx, 12)
	heap.Push(mx, 21)
	heap.Push(mx, 444)
	heap.Push(mx, 4)
	heap.Push(mx, 41)
	fmt.Println(mx)
	heap.Pop(mx)
	fmt.Println(mx)
	heap.Remove(mx, 2)
	// heap.Pop(mx)
	fmt.Println(mx)
	// heap.Pop(mx)
	// fmt.Println(mx)
	// heap.Pop(mx)
	fmt.Println(1<<32 + 1)
}
