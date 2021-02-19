package binarySearchTree

import "testing"

func TestCreateTree(t *testing.T) {
	tree := createTree()
	if tree.size() != 0 {
		t.Error("Size should be 0, got", tree.size())
	}
}