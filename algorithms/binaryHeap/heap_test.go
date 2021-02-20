package binaryHeap

import (
	"fmt"
	"testing"
)

func assertElements(t *testing.T, h *heap, expected []int) {
	vals := h.values()
	if len(vals) != len(expected) {
		t.Fatalf("Invalid len, got: %v, exp %v", len(vals), len(expected))
	}
	for i := range expected {
		got := vals[i]
		exp := expected[i]
		if got != exp {
			t.Errorf("Error in idx %v, got %v, exp %v", i, got, exp)
			t.Fatalf("%v != %v", vals, expected)
		}
	}
}

func TestInsert(t *testing.T) {
	tdt := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 5}, []int{1, 2, 5}},
		{[]int{1, 4, 2, 5}, []int{1, 2, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	}
	for _, tc := range tdt {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			heap := newHeap()
			for _, v := range tc.input {
				heap.insert(v)
			}
			assertElements(t, heap, tc.expected)
		})
	}
}

func TestBubbleUp(t *testing.T) {
	t.Fatal("todo - bubble up when adding")
}
