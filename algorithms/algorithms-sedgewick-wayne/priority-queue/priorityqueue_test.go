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

// it's often useful to store a key and a value in the pq
// key is used to sort
func TestPriorityQueue(t *testing.T) {
	// just use heap to interleave inserts and deletions
}