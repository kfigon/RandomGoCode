package searching

type nodeKind int
const (
	twoNode nodeKind = iota
	threeNode
)

type btreeNode[T comparable[T]] struct{
	kind nodeKind
	twoNode *twoBTreeNode[T]
	threeNode *threeBTreeNode[T]
}

type twoBTreeNode[T comparable[T]] struct{
	v T
	
	left *btreeNode[T]
	right *btreeNode[T]
}

type threeBTreeNode[T comparable[T]] struct{
	vLow T
	vHigh T

	left *btreeNode[T]
	middle *btreeNode[T]
	right *btreeNode[T]
}

func new2Node[T comparable[T]](v T) *btreeNode[T] {
	return &btreeNode[T]{
		kind: twoNode,
		twoNode: &twoBTreeNode[T]{
			v: v,
		},
	}
}

func new3Node[T comparable[T]](vLow, vHigh T) *btreeNode[T] {
	return &btreeNode[T]{
		kind: threeNode,
		threeNode: &threeBTreeNode[T]{
			vLow: vLow,
			vHigh: vHigh,
		},
	}
}

type btree struct{}