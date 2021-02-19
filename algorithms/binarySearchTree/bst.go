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
	var traverse func(n *node)

	traverse = func(n *node) {
		if n == nil {
			return
		}
		traverse(n.left)
		out = append(out, n.val)
		traverse(n.right)
	}

	traverse(t.root)
	return out
}

func (t *tree) insert(value int)  {
	newNode := createNode(value)
	if t.root == nil {
		t.root = newNode
		return
	}
	
	ptr := t.root
	for ptr != nil {
		if value < ptr.val  {
			if ptr.left == nil {
				ptr = newNode
				break
			}
			ptr = ptr.left
		} else {
			if ptr.right == nil {
				ptr = newNode
				break
			}
			ptr = ptr.right
		}
	}
}