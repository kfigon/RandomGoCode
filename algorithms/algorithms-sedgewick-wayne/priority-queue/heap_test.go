package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// heap (binary heap) - structure where each key is guaranteed to be larger
// to the 2 children. Array or binary tree(+ link to parrent)

// BinaryHeap is kind of binarySearchTree (2 nodes)
// maxBinaryHeap - parent is always larger than children nodes
// minBinaryHeap - parent is always smaller than childen nodes

// heap is always balanced (binary tree is not, can be tall). Here we fill it sequentially

func TestHeap(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		hip := newArrayHeap[string]()
		_, ok := hip.max()
		assert.False(t, ok)
	})

	t.Run("single", func(t *testing.T) {
		hip := newArrayHeap[string]()

		hip.insert(heapEl[string]{5, "foo"})

		v, ok := hip.max()
		assert.True(t, ok)
		assert.Equal(t, heapEl[string]{5, "foo"}, *v)

		_, ok = hip.max()
		assert.False(t, ok)
	})

	t.Fatal("todo more")
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
	return &arrayHeap[T]{
		tab: []*heapEl[T]{nil},  // first element nil to help with index math
	}
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

func (a *arrayHeap[T]) hasParrent(idx int) bool {
	return idx > 1
}

func (a *arrayHeap[T]) hasLeftChild(idx int) bool {
	return (idx*2) < len(a.tab)
}

func (a *arrayHeap[T]) hasRightChild(idx int) bool {
	return (idx*2 +1) < len(a.tab)
}