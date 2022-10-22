package priorityqueue

import "testing"

// heap (binary heap) - structure where each key is guaranteed to be larger
// to the 2 children. Array or binary tree(+ link to parrent)
// # BinaryHeap is kind of binarySearchTree (2 nodes)
// # maxBinaryHeap - parent is always larger than children nodes
// # minBinaryHeap - parent is always smaller than childen nodes
// heap is always balanced (binary tree is not, can be tall)

func TestHeap(t *testing.T) {
	t.Fatal("todo")	
}

type heap[T any] interface {
	insert(pair[int, T])
	max() (pair[int, T], bool)
	delMax() (pair[int, T], bool)
}

type heapEl[T any] pair[int, T]
type arrayHeap[T any] struct {
	tab []*heapEl[T]
}

func newArrayHeap[T any]() *arrayHeap[T] {
	return &arrayHeap[T]{tab: []*heapEl[T]{}}
}

func (a *arrayHeap[T]) insert(v heapEl[T]) {

}

func (a *arrayHeap[T]) max() (*heapEl[T], bool) {
	return nil, false
}

func (a *arrayHeap[T]) delMax() (*heapEl[T], bool) {
	return nil, false
}

func (a *arrayHeap[T]) children(idx int) (int,int) {
	x := idx*2
	return x, x+1
}

func (a *arrayHeap[T]) parent(idx int) int {
	return idx/2
}