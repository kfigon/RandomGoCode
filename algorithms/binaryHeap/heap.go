package binaryHeap

// another category of trees
// max binary heap - parent nodes are always larger than children
// no order like in BST (left smaller, right bigger) - NOT true in bin heap

// heap is more evenly distributed (more compact)
// heaps are used for priority queue
// used for graph traversal

type heap struct {
}

func newHeap() *heap {
	return &heap{}
}

func (h *heap) values() []int {
	return []int{}
}

func (h *heap) insert(v int) {

}
