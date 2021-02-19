package binarySearchTree

import "testing"

func assertSize(t *testing.T, tree *tree, expected []int) {
	if size := tree.size(); size != len(expected) {
		t.Errorf("Size should be %v, got %v", len(expected), size)
	}
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
	assertSize(t, tree, []int{})
}

func TestAdd(t *testing.T) {
	tree := createTree()
	tree.insert(15)
	assertSize(t, tree, []int{15})
}