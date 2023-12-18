package Alg

/*
╔═══════════════════╦════════════════════════╗
║ vertex            ║ index                  ║
╠═══════════════════╬════════════════════════╣
║ root              ║ 0                      ║
║ current           ║ i                      ║
║ parent            ║ (i - 1) / 2            ║
║ left child        ║ 2*i + 1                ║
║ right child       ║ 2*i + 2                ║
║ the last non-leaf ║ (array length - 2) / 2 ║
╚═══════════════════╩════════════════════════╝
*/

type MinHeap struct {
	Heap []int
}

func InitMinHeap() *MinHeap {
	return &MinHeap{
		Heap: []int{},
	}
}

func (m *MinHeap) Push(n int) {
	m.Heap = append(m.Heap, n)
	curIdx := len(m.Heap) - 1
	parentIdx := (curIdx - 1) / 2
	for parentIdx != curIdx {
		if m.Heap[parentIdx] > m.Heap[curIdx] {
			m.Heap[parentIdx], m.Heap[curIdx] = m.Heap[curIdx], m.Heap[parentIdx]
			curIdx, parentIdx = parentIdx, (parentIdx-1)/2
		} else {
			break
		}
	}
}

func (m *MinHeap) Pop() int {
	l := len(m.Heap)
	m.Heap[l-1], m.Heap[0] = m.Heap[0], m.Heap[l-1]
	r := m.Heap[l-1]
	m.Heap = m.Heap[:l-1]
	m.SiftDown()
	return r
}

func (m *MinHeap) SiftDown() {
	l := len(m.Heap)
	if l <= 1 {
		return
	}
	endIdx := l - 1
	curIdx := 0
	for 2*curIdx+1 <= endIdx {
		if 2*curIdx+2 > endIdx {
			//no right child
			if m.Heap[2*curIdx+1] < m.Heap[curIdx] {
				m.Heap[2*curIdx+1], m.Heap[curIdx] = m.Heap[curIdx], m.Heap[2*curIdx+1]
			}
			return
		}
		smallerChildIdx := 2*curIdx + 2
		if m.Heap[2*curIdx+1] < m.Heap[2*curIdx+2] {
			smallerChildIdx = (2*curIdx + 1)
		}
		if m.Heap[smallerChildIdx] < m.Heap[curIdx] {
			m.Heap[smallerChildIdx], m.Heap[curIdx] = m.Heap[curIdx], m.Heap[smallerChildIdx]
			curIdx = smallerChildIdx
		} else {
			return
		}
	}

}
