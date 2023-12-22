package TreeMapDemo

import (
	"container/heap"
	"fmt"
)

// proporty queue item: []int{a,b,c},a is value, b is index and c is priority, order by c and descending
type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][2] > pq[j][2]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i][1] = i
	pq[j][1] = j
}

func (pq *PriorityQueue) Pop() any {
	l := pq.Len()
	x := (*pq)[l-1]
	x[1] = -1
	*pq = (*pq)[0 : l-1]
	return x
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([]int))
}

func InitPQ() {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, []int{1, 2, 3})
	heap.Push(pq, []int{1, 22, 32})
	heap.Push(pq, []int{11, 2, 13})
	heap.Push(pq, []int{11, 2, 23})
	fmt.Println(pq)
	heap.Remove(pq, 2)
	fmt.Println(pq)
	heap.Remove(pq, 2)
	fmt.Println(pq)
	heap.Remove(pq, 2)
	fmt.Println(pq)
}
