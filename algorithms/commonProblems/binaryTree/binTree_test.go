package main

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	tr := newTree()
	if tr.size() != 0 {
		t.Error("Invalid size, got: ", tr.size())
	}
	if ln := len(tr.values()); ln != 0 {
		t.Error("Invalid size, expected 0, got: ", ln)
	}
}

func TestOneElement(t *testing.T) {
	tr := newTree()
	tr.insert(5)
	if tr.size() != 1 {
		t.Error("Invalid size, got: ", tr.size())
	}
	if ln := len(tr.values()); ln != 1 {
		t.Fatal("Invalid size, expected 1, got: ", ln)
	}
	if v := tr.values(); v[0] != 5 {
		t.Error("Invalid element received, got: ", v[0])
	}
}

