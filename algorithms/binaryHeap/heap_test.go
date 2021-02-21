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

func TestBubbleUp_simpleCase(t *testing.T) {
	heap := newHeap()
	heap.insert(5)
	heap.insert(3)
	assertNode(t, heap.root, valueNode(5))
	assertNode(t, heap.root.left, valueNode(3))

	assertNode(t, heap.root.right, emptyNode())
}

func TestBubbleUp_complicated(t *testing.T) {
	heap := newHeap()
	heap.insert(5)
	heap.insert(3)
	heap.insert(8)
	assertNode(t, heap.root, valueNode(8))
	assertNode(t, heap.root.left, valueNode(3))
	assertNode(t, heap.root.right, valueNode(5))

	assertNode(t, heap.root.left.left, emptyNode())
	assertNode(t, heap.root.left.right, emptyNode())
	assertNode(t, heap.root.right.left, emptyNode())
	assertNode(t, heap.root.right.right, emptyNode())
}

type opt struct {
	val   int
	empty bool
}

func emptyNode() opt        { return opt{empty: true} }
func valueNode(val int) opt { return opt{empty: false, val: val} }

func assertNode(t *testing.T, n *node, expectedNode opt) {
	if expectedNode.empty && n != nil {
		t.Errorf("Expected node to be empty, has %v", n.val)
	} else if !expectedNode.empty {
		if n == nil {
			t.Error("Node empty, exp", expectedNode.val)
		} else if n.val != expectedNode.val {
			t.Errorf("Invalid node val got %v, exp %v", n.val, expectedNode.val)
		}
	}
}
