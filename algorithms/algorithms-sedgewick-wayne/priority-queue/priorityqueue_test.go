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

func TestPriorityQueue(t *testing.T) {
	t.Fatal("todo")
}


type pair[T any, V any] struct {
	a T
	b V
}

type priorityQueue[T any] interface {
	insert(pair[int, T])
	max() pair[int, T]
	delMax() pair[int, T]
	size() int
	empty() bool
}

