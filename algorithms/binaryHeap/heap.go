package binaryHeap

// another category of trees
// max binary heap - parent nodes are always larger than children
// no order like in BST (left smaller, right bigger) - NOT true in bin heap

// heap is more evenly distributed (more compact)
// heaps are used for priority queue
// used for graph traversal

type node struct {
	val    int
	left   *node
	right  *node
	parent *node
}

func newNode(v int, parent *node) *node {
	return &node{val: v, parent: parent}
}

type heap struct {
	root *node
}

func newHeap() *heap {
	return &heap{}
}

func (h *heap) values() []int {
	out := make([]int, 0)
	var traverse func(n *node)
	traverse = func(n *node) {
		if n == nil {
			return
		}
		traverse(n.left)
		out = append(out, n.val)
		traverse(n.right)
	}

	traverse(h.root)
	return out
}

func (h *heap) insert(v int) {
	node := newNode(v, nil)
	if h.root == nil {
		h.root = node
		return
	}

	ptr := h.root
	// just put it to the first available place (
	for ptr != nil {
		if ptr.left == nil {
			ptr.left = node
			node.parent = ptr
			break
		} else if ptr.right == nil {
			ptr.right = node
			node.parent = ptr
			break
		}
		ptr = ptr.right
	}

	// bubble up
	ptr = node
	parent := ptr.parent
	for ptr != nil && parent != nil && ptr.val > parent.val {
		tmp := ptr.val
		ptr.val = parent.val
		parent.val = tmp

		ptr = parent
		parent = ptr.parent
	}
}
