package searching


type btreeNode[T comparable[T]] interface{
	dummy()
}

type twoBTreeNode[T comparable[T]] struct{
	v T
	
	left *btreeNode[T]
	right *btreeNode[T]
}

func (_ *twoBTreeNode[T]) dummy(){}

type threeBTreeNode[T comparable[T]] struct{
	vLow T
	vHigh T

	left *btreeNode[T]
	middle *btreeNode[T]
	right *btreeNode[T]
}
func (_ *threeBTreeNode[T]) dummy(){}

func new2Node[T comparable[T]](v T) btreeNode[T] {
	return &twoBTreeNode[T]{
		v: v,
	}
}

func new3Node[T comparable[T]](vLow, vHigh T) btreeNode[T] {
	return &threeBTreeNode[T]{
		vLow: vLow,
		vHigh: vHigh,
	}
}

type btree struct{}