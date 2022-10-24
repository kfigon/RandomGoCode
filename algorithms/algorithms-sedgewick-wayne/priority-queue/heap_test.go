package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// heap (binary heap) - structure where each key is guaranteed to be larger
// to the 2 children. Array or binary tree(+ link to parrent)

// BinaryHeap is kind of binarySearchTree (2 nodes)
// maxBinaryHeap - parent is always larger than children nodes
// minBinaryHeap - parent is always smaller than childen nodes

// heap is always balanced (binary tree is not, can be tall). Here we fill it sequentially

func TestHeap(t *testing.T) {

	popAssert := func(t *testing.T, hip *arrayHeap, expected int) {
		v, ok := hip.delMax()
		assert.True(t, ok)
		assert.Equal(t, expected, v)
	}

	t.Run("empty", func(t *testing.T) {
		hip := newArrayHeap()
		_, ok := hip.delMax()
		assert.False(t, ok)
		assert.Equal(t, 0, hip.size)

		_, ok = hip.max()
		assert.False(t, ok)
	})

	t.Run("single", func(t *testing.T) {
		hip := newArrayHeap()

		hip.insert(5)
		assert.Equal(t, 1, hip.size)

		popAssert(t, hip, 5)
		assert.Equal(t, 0, hip.size)

		_, ok := hip.delMax()
		assert.False(t, ok)
		assert.Equal(t, 0, hip.size)
	})

	t.Run("multiple", func(t *testing.T) {
		hip := newArrayHeap()

		for i, v := range []int{4,3,6,8,32,2,1} {
			hip.insert(v)
			assert.Equal(t, i+1, hip.size)
		}
		assert.Equal(t, 7, hip.size)
		// assert.Equal(t, []int{-900, 32,8,4,3,6,2,1,-1,-1,-1,-1,-1,-1,-1,-1,-1}, hip.tab) // not so great assert, but

		
		expectedSorted := []int{32,8,6,4,3,2,1}
		got := []int{}
		for {
			v, ok := hip.delMax()
			if !ok {
				break
			}
			got = append(got, v)
			assert.Equal(t, len(expectedSorted) - len(got), hip.size)
		}
		assert.Equal(t, expectedSorted, got)
		assert.Equal(t, 0, hip.size)

		_, ok := hip.max()
		assert.False(t, ok)
	})

	t.Run("inserts between pops", func(t *testing.T) {
		hip := newArrayHeap()
		hip.insert(5)
		hip.insert(8)
		hip.insert(1)
		hip.insert(13)

		popAssert(t, hip, 12)
		popAssert(t, hip, 8)

		hip.insert(2)
		hip.insert(13)

		popAssert(t, hip, 12)
		popAssert(t, hip, 5)
		popAssert(t, hip, 2)
		popAssert(t, hip, 1)

		assert.Equal(t, 0, hip.size)
		_, ok := hip.delMax()
		assert.False(t, ok)
	})
}

// todo - make it generic
type heap interface {
	insert(int)
	max() (int, bool)
	delMax() (int, bool)
}

type arrayHeap struct {
	tab []int
	size int
}

func newArrayHeap() *arrayHeap {
	tab := []int{-900}// first element nil to help with index math

	for i := 0; i < 16; i++ {
		tab = append(tab, -1) // todo: think about resizing and dynamic size
	}

	return &arrayHeap{
		tab: tab,
		size: 0,
	}
}

func (a *arrayHeap) insert(v int) {
	a.size++
	a.tab[a.size] = v

	// swim up
	idx := a.size
	for {
		parent, ok := a.parentIdx(idx)
		if !ok {
			break
		}
		if a.tab[idx] > a.tab[parent] {
			swap(&a.tab[idx], &a.tab[parent])
		}
		idx = parent
	}
}

func swap(a,b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func (a *arrayHeap) max() (int, bool) {
	if a.size < 1 {
		return 0, false
	}
	return a.tab[1], true
}

func (a *arrayHeap) delMax() (int, bool) {
	if a.size < 1 {
		return 0, false
	}
	toRet := a.tab[1]

	// get the last element and push it down the heap
	a.tab[1] = a.tab[a.size]
	a.size--
	idx := 1
	for {
		leftIdx, lOk := a.leftChildIdx(idx)
		rightIdx, rOk := a.rightChildIdx(idx)

		if lOk && rOk {
			if a.tab[leftIdx] < a.tab[rightIdx] && a.tab[idx] < a.tab[leftIdx] {
				swap(&a.tab[idx], &a.tab[leftIdx])
				idx = leftIdx
			} else {
			}
		} else if lOk && !rOk {
		} else if !lOk && rOk {
		} else {
			break
		}
	}
	return toRet, true
}

func (a *arrayHeap) leftChildIdx(idx int) (int,bool) {
	return idx*2, (idx*2) <= a.size
}

func (a *arrayHeap) rightChildIdx(idx int) (int,bool) {
	return (idx*2 +1), (idx*2 +1) <= a.size
}

func (a *arrayHeap) parentIdx(idx int) (int, bool) {
	return idx/2, idx > 1
}
