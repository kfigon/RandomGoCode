package main

type tree struct {
	top *node
}

func (t *tree) size() int {
	return len(t.values())
}

func (t *tree) values() []int {
	out := make([]int,0)
	var dfs func(n *node)
	dfs = func(n *node) {
		if n == nil {
			return
		}
		dfs(n.left)
		out = append(out, n.val)
		dfs(n.right)
	}
	
	dfs(t.top)
	return out
}

func (t *tree) insert(val int) {
	n := newNode(val)
	if t.top == nil {
		t.top = n
		return
	}
	ptr := t.top
	for {
		if n.val < ptr.val {
			if ptr.left == nil {
				ptr.left = n
				break
			}
			ptr = ptr.left
		} else {
			if ptr.right == nil {
				ptr.right = n
				break
			}
			ptr = ptr.right
		}
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