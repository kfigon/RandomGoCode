package binarySearchTree

type node struct {
	val int
	left *node
	right *node
}
func createNode(val int) *node {
	return &node{val:val}
}

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
	if t.root == nil {
		t.root = createNode(value)
	}
}