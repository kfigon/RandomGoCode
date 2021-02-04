package main

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	tree := newTree()
	if tree.size() != 0 {
		t.Error("Invalid size, got: ", tree.size())
	}
	if ln := len(tree.values()); ln != 0 {
		t.Error("Invalid size, expected 0, got: ", ln)
	}
}

