package binarySearchTree

type node struct {
	val int
	left *node
	right *node
}
func createNode(val int) *node {
	return &node{val:val}
}

// binary - 2 children
// search tree - sorted in specific way. Data can be compared
// left child is smaller, right is greater
type tree struct{
	root *node
}

func createTree() *tree {
	return &tree{}
}

func (t *tree) values() []int {
	out := make([]int,0)
	if t.root != nil {
		out = append(out, t.root.val)
	}
	return out
}

func (t *tree) insert(value int)  {
	newNode := createNode(value)
	if t.root == nil {
		t.root = newNode
		return
	}
	

}