package main

import (
	"fmt"
)

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

func (t *tree) isPresent(v int) bool {
	return false
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


func main()  {
	t := newTree()
	els := []int{1,6,3,1,45,6,3,1,3,4,6,2,7,8}
	for _,v := range els{
		t.insert(v)
	}
	out := t.values()
	for _,v := range out {
		fmt.Println(v)
	}
}