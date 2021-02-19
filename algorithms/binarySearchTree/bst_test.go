package binarySearchTree

import (
	"testing"
	"fmt"
)

func assertValues(t *testing.T, tree *tree, expected []int) {
	vals := tree.values()
	if len(vals) != len(expected) {
		t.Errorf("Value len should be %v, got %v", len(expected), len(vals))
	}

	for i := 0; i < len(vals); i++ {
		exp := expected[i]
		got := vals[i]
		if exp != got {
			t.Errorf("Invalid element in idx %v, got %v, exp %v", i, got, exp)
		}
	}
}
func TestCreateTree(t *testing.T) {
	tree := createTree()
	assertValues(t, tree, []int{})
}

func TestInsertAndTraverse(t *testing.T) {
	tdt := [] struct {
		elements []int
		expected []int
	} {
		{ []int{15}, []int{15}},
	}
	for _, tc := range tdt {
		t.Run(fmt.Sprint(tc.elements), func(t *testing.T) {
			tree := createTree()
			for _, v := range tc.elements {
				tree.insert(v)	
			}
			assertValues(t, tree, []int{15})
		})
	}
	
}