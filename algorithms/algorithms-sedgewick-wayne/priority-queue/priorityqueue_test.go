package priorityqueue

import "testing"

// a list that can process element with the highest key
// and supports adding new elements while maintaining the order
// or X max elements from infinite stream

// 'list that sorts itself'

// implementations:
// DS				insert	removeMax
// ordered arr		  N			1
// unordered arr	  1			N
// heap			  	 logN      logN

// heap (binary heap) - structure where each key is guaranteed to be larger
// to the 2 children. Array or binary tree(+ link to parrent)
// # BinaryHeap is kind of binarySearchTree (2 nodes)
// # maxBinaryHeap - parent is always larger than children nodes
// # minBinaryHeap - parent is always smaller than childen nodes
// heap is always balanced (binary tree is not, can be tall)
func TestPriorityQueue(t *testing.T) {
	t.Fatal("todo")
}


type comparable[T any] interface {
	less(a, b T) bool
}

type pair[T any, V any] struct {
	a T
	b V
}

type priorityQueue[T comparable[T]] interface {
	insert(comparable[T])
	max() comparable[T]
	delMax() comparable[T]
	size() int
	empty() bool
}

