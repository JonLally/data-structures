package main

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) Sort() {
	for i := (len(h) / 2); i >= 0; i-- {
		h.maxHeapify(i, len(h))
	}
	for i := h.Len() - 1; i > 0; i-- {
		// Move current root to end
		h.Swap(0, i)
		h.maxHeapify(0, i)
	}
}

func (h IntHeap) leaf(index int, size int) bool {
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

func (h IntHeap) maxHeapify(position int, size int) {
	if h.leaf(position, size) {
		return
	}
	maximum := position
	leftChild := 2*position + 1
	rightChild := leftChild + 1
	if leftChild < size && h[leftChild] > h[position] {
		maximum = leftChild
	}
	if rightChild < size && h[rightChild] > h[maximum] {
		maximum = rightChild
	}

	if position != maximum {
		h.Swap(position, maximum)
		h.maxHeapify(maximum, size)
	}
}
