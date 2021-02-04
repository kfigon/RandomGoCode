package main

type tree struct {
	top *node
}

func (t *tree) size() int {
	return len(t.values())
}

func (t *tree) values() []int {
	out := make([]int,0)
	if t.top != nil {
		out = append(out, t.top.val)
	}
	return out
}

func (t *tree) insert(val int) {
	if t.top == nil {
		t.top = newNode(val)
		return
	}
}

func newTree() *tree {
	return &tree{}
}

type node struct {
	val int
	left *node
	right *node
}

func newNode(val int) *node {
	return &node{val:val}
}