package searching

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// node keys can be anything - usually some kind of comparable value
// then we can search other things than just integers in a tree and use it as an index in db
func TestBinarySearchTree(t *testing.T) {
	newBst := func() *bst[intWrapper]{
		return &bst[intWrapper]{}
	}

	t.Run("empty", func(t *testing.T) {
		b := newBst()
		assert.Equal(t, []intWrapper{}, b.traverse())
		
		_, ok := b.max()
		assert.False(t, ok)
		
		_, ok = b.min()
		assert.False(t, ok)
	})

	t.Run("multiple", func(t *testing.T) {
		b := newBst()
		for _, v := range []intWrapper{7,5,3,2,5,7,1,2} {
			b.add(v)
		}
		assert.Equal(t, []intWrapper{1,2,3,5,7}, b.traverse())
		
		max, ok := b.max()
		assert.True(t, ok)
		assert.Equal(t, 7, int(max))

		min, ok := b.min()
		assert.True(t, ok)
		assert.Equal(t, 1, int(min))
	})
		
	t.Run("min max interleaved", func(t *testing.T) {
		b := newBst()
		for _, v := range []intWrapper{7,5,3,2,5,7,1,2} {
			b.add(v)
		}
		max, ok := b.max()
		assert.True(t, ok)
		assert.Equal(t, 7, int(max))

		min, ok := b.min()
		assert.True(t, ok)
		assert.Equal(t, 1, int(min))

		for _, v := range []intWrapper{88,-123,32,12,-5} {
			b.add(v)
		}

		max, ok = b.max()
		assert.True(t, ok)
		assert.Equal(t, 88, int(max))

		min, ok = b.min()
		assert.True(t, ok)
		assert.Equal(t, -123, int(min))
	})

	t.Run("get", func(t *testing.T) {
		b := newBst()
		for _, v := range []intWrapper{7,5,3,2,5,7,1,2} {
			b.add(v)
		}
		got, ok := b.get(5)
		assert.True(t, ok)
		assert.Equal(t, got, intWrapper(5))

		_, ok = b.get(125)
		assert.False(t, ok)
	})
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		a int
		b int
		exp int		
	}{
		{2,5,-1},
		{5,2,1},
		{5,5,0},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v %v", tC.a, tC.b), func(t *testing.T) {
			assert.Equal(t, tC.exp, intWrapper(tC.a).cmp(intWrapper(tC.b)))
		})
	}
}

func TestBstDelete(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.delete(123)
		b.delete(1)
		b.delete(3)

		assert.Equal(t, []intWrapper{}, b.traverse())
	})

	t.Run("single el", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(123)
		b.delete(123)

		assert.Equal(t, []intWrapper{}, b.traverse())
	})

	t.Run("root when 1 left child", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(5)
		b.add(1)
		b.delete(5)

		assert.Equal(t, []intWrapper{1}, b.traverse())
	})

	t.Run("root when 1 right child", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(5)
		b.add(10)
		b.delete(5)

		assert.Equal(t, []intWrapper{10}, b.traverse())
	})

	t.Run("root when 2 children", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(5)
		b.add(10)
		b.add(1)
		b.delete(5)

		assert.Equal(t, []intWrapper{1,10}, b.traverse())
	})

	t.Run("not empty", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []int{5,3,7,4,1,8} {
			b.add(intWrapper(v))
		}
		b.delete(123)
		b.delete(1)
		b.delete(7)
		b.delete(5)

		assert.Equal(t, []intWrapper{3,4,8}, b.traverse())
	})

	t.Run("root when many", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []int{5,3,7,4,1,8} {
			b.add(intWrapper(v))
		}
		b.delete(5)

		assert.Equal(t, []intWrapper{1,3,4,7,8}, b.traverse())
	})

	t.Run("tall tree", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []int{1,2,3,10,11,12, 5,6,7} {
			b.add(intWrapper(v))
		}
		b.delete(10)
		b.delete(11)

		assert.Equal(t, []intWrapper{1,2,3,5,6,7,12}, b.traverse())
	})

	t.Run("interleaved", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []int{5,3,7,4,1,8} {
			b.add(intWrapper(v))
		}

		b.delete(1)
		b.delete(7)
		b.delete(5)

		for _, v := range []int{2,10,-14,1} {
			b.add(intWrapper(v))
		}

		assert.Equal(t, []intWrapper{-14,1,2,3,4,8,10}, b.traverse())
	})
}

func TestBstRange(t *testing.T) {
	testCases := []struct {
		min int
		max int
		exp []intWrapper
	}{
		{ 1,10,[]intWrapper{1,2,3,4,5,6,7,8,9,10} },
		{ 1,5,[]intWrapper{1,2,3,4,5} },
		{ 1,100,[]intWrapper{1,2,3,4,5,6,7,8,9,10} },
		{ -100,100,[]intWrapper{1,2,3,4,5,6,7,8,9,10} },
		{ -100,3,[]intWrapper{1,2,3} },
		{ 5,10,[]intWrapper{5,6,7,8,9,10} },
		{ 10,8, []intWrapper{}},
		{ 1,1,[]intWrapper{1}},
		{ 1,2, []intWrapper{1,2}},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tC.min, tC.max), func(t *testing.T) {
			b := &bst[intWrapper]{}
			// 1,2,3,4,5,6,7,8,9,10
			for _, v := range []intWrapper{4,10,9,2,5,6,3,8,1,7} {
				b.add(v)
			}
			assert.Equal(t, tC.exp, b.inRange(intWrapper(tC.min),intWrapper(tC.max)))
		})
	}
}

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

func (b *bst[T]) traverse() []T {
	out := []T{}
	var fn func(*node[T])
	fn = func(n *node[T]) {
		if n == nil {
			return
		}
		fn(n.left)
		out = append(out, n.val)
		fn(n.right)
	}
	fn(b.root)
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

// 3 cases:
// delete a leaf
// node with 1 child
// node with 2 children
func (b *bst[T]) delete(v T) {
	var parent *node[T] = nil
	ptr := b.root

	for ptr != nil && v.cmp(ptr.val) != 0 {
		if v.cmp(ptr.val) < 0 {
			parent = ptr
			ptr = ptr.left
		} else if v.cmp(ptr.val) > 0 {
			parent = ptr
			ptr = ptr.right
		}
	}

	if ptr == nil {
		return
	} else if parent == nil { // 1 element, root itself matched
		b.root = nil
		return
	}

	if ptr.left != nil && ptr.right != nil { // 2 children

	} else if ptr.left == nil && ptr.right == nil { // 0 child
		if v.cmp(parent.val) < 0 {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else { // 1 child
		var nextNode *node[T]
		if ptr.left != nil {
			nextNode = ptr.left
		} else {
			nextNode = ptr.right
		}

		if v.cmp(parent.val) < 0 {
			parent.left = nextNode
		} else {
			parent.right = nextNode
		}
		ptr = nil
	}
}