package binarySearchTree

import "testing"

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

func TestAdd(t *testing.T) {
	tree := createTree()
	tree.insert(15)
	assertValues(t, tree, []int{15})
}