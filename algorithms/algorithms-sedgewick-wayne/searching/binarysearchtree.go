package searching

// node keys can be anything - usually some kind of comparable value
// then we can search other things than just integers in a tree and use it as an index in db

type intWrapper int
func (i intWrapper) cmp(other intWrapper) int {
	// def(a,b): 
	//   return a-b
	
	// a.cmp(b) < 0		a < b
	// a.cmp(b) > 0 	a > b
	if int(i) < int(other) {
		return -1
	} else if int(i) > int(other) {
		return 1
	}
	return 0
}

type comparable[T any] interface {
	cmp(T) int
}

type node[T comparable[T]] struct {
	val T
	left *node[T]
	right *node[T]
}

func newNode[T comparable[T]](val T) *node[T] {
	return &node[T]{val: val}
}

type bst[T comparable[T]] struct {
	root *node[T]
}

func (b *bst[T]) add(v T) {
	if b.root == nil {
		b.root = newNode(v)
		return
	}

	ptr := b.root 
	for ptr != nil {
		if v.cmp(ptr.val) < 0 {
			if ptr.left != nil {
				ptr = ptr.left
			} else {
				ptr.left = newNode(v)
				return
			}
		} else if v.cmp(ptr.val) > 0 {
			if ptr.right != nil {
				ptr = ptr.right
			} else {
				ptr.right = newNode(v)
				return
			}
		} else {
			ptr.val = v
			return
		}
	}
}

func (b *bst[T]) addRec(v T) {
	var fn func(*node[T]) *node[T]
	fn = func(n *node[T]) *node[T] {
		if n == nil {
			return newNode(v)
		}

		if v.cmp(n.val) < 0 {
			n.left = fn(n.left)
		} else if v.cmp(n.val) > 0 {
			n.right = fn(n.right)
		} else {
			n.val = v
		}
		return n
	}

	b.root = fn(b.root)
}

func (b *bst[T]) traverseDfs(fn func(T)) {
	var travFn func(*node[T])
	travFn = func(n *node[T]) {
		if n == nil {
			return
		}
		travFn(n.left)
		fn(n.val)
		travFn(n.right)
	}
	travFn(b.root)
}

func (b *bst[T]) traverseBfs(fn func(T)) {
	
}

func (b *bst[T]) collectBfs() []T {
	out := []T{}
	b.traverseBfs(func(v T) {
		out = append(out, v)
	})
	return out
}

func (b *bst[T]) collect() []T {
	out := []T{}
	b.traverseDfs(func(v T) {
		out = append(out, v)
	})
	return out
}

func (b *bst[T]) get(val T) (T, bool) {
	ptr := b.root
	for ptr != nil {
		if val.cmp(ptr.val) < 0 {
			ptr = ptr.left
		} else if val.cmp(ptr.val) > 0 {
			ptr = ptr.right
		} else {
			return ptr.val, true
		}
	}
	var out T
	return out, false
}

func (b *bst[T]) getRec(val T) (T, bool) {
	var fn func(*node[T]) (T, bool)
	fn = func(n *node[T]) (T, bool) {
		if n == nil {
			var out T
			return out, false
		}
		
		if val.cmp(n.val) < 0 {
			return fn(n.left)
		} else if val.cmp(n.val) > 0 {
			return fn(n.right)
		}
		return n.val, true
	}

	return fn(b.root)
}

func (b *bst[T]) max() (T, bool) {
	if b.root == nil {
		var out T
		return out, false
	}
	ptr := b.root
	for ptr.right != nil {
		ptr = ptr.right
	}
	return ptr.val, true
}

func (b *bst[T]) min() (T, bool) {
	// or recursive:
	// def min(n):
	// 		if n.left == null: return n.val
	// 		return min(n.left)
	if b.root == nil {
		var out T
		return out, false
	}
	ptr := b.root
	for ptr.left != nil {
		ptr = ptr.left
	}
	return ptr.val, true
}

func (b *bst[T]) inRange(min, max T) []T {
	out := []T{}
	if min.cmp(max) > 0 {
		return out
	}

	var fn func(*node[T])
	fn = func(n *node[T]) {
		if n == nil {
			return
		}
		cmpMin := n.val.cmp(min)
		cmpMax := n.val.cmp(max)
		if cmpMin > 0 { // v > min
			fn(n.left)
		}
		if cmpMin >= 0 && cmpMax <= 0 { // min <= v && v <= max
			out = append(out, n.val)	
		}
		if cmpMax < 0 { // v < max
			fn(n.right)
		}
	}
	fn(b.root)
	return out
}

func (b *bst[T]) delMin() {
	var fn func(*node[T]) *node[T]
	fn = func(n *node[T]) *node[T] {
		if n.left == nil {
			return n.right
		}
		// go left until theres no more
		// then set parent.left to right of the deleted node
		n.left = fn(n.left) 
		return n
	}
	b.root = fn(b.root)
}

// 3 cases:
// delete a leaf
// node with 1 child
// node with 2 children
func (b *bst[T]) delete(v T) {
	toBeDeleted := b.root
    var parent *node[T]
 
	// recursive solution is tricky, this one also :O

    for toBeDeleted != nil && toBeDeleted.val.cmp(v) != 0 {
        parent = toBeDeleted
        if toBeDeleted.val.cmp(v) < 0 {
            toBeDeleted = toBeDeleted.right
		} else {
			toBeDeleted = toBeDeleted.left
		}
	}
 
    if toBeDeleted == nil {
        return
	}
 
	// at most 1 child
	if toBeDeleted.left == nil || toBeDeleted.right == nil { 
        var newCurr *node[T]
        if toBeDeleted.left == nil {
            newCurr = toBeDeleted.right
		} else {
            newCurr = toBeDeleted.left
		}

        if parent == nil {
			b.root = newCurr
            return
		}

        if toBeDeleted == parent.left {
            parent.left = newCurr
		} else {
            parent.right = newCurr
		}
        toBeDeleted = nil
		return
	}

	// 2 children
	var ptr *node[T]
	succesor := toBeDeleted.right

	for succesor.left != nil {
		ptr = succesor
		succesor = succesor.left
	}

	if ptr != nil {
		ptr.left = succesor.right
	} else {
		toBeDeleted.right = succesor.right
	}
	toBeDeleted.val = succesor.val
	succesor = nil
}

// O(n)
func (b *bst[T]) height() int {
	max := func(a,b int) int {
		if a > b {
			return a
		}
		return b
	}

	var fn func(*node[T]) int
	fn = func(n *node[T]) int {
		if n == nil {
			return 0
		}
		return 1 + max(fn(n.left), fn(n.right))
	}
	return fn(b.root)
}