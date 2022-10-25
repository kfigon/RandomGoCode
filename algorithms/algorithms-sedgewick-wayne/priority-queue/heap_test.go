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
		assert.Len(t, hip.tab, 0)

		_, ok = hip.max()
		assert.False(t, ok)
	})

	t.Run("single", func(t *testing.T) {
		hip := newArrayHeap()

		hip.insert(5)
		assert.Len(t, hip.tab, 1)

		popAssert(t, hip, 5)
		assert.Len(t, hip.tab, 0)

		_, ok := hip.delMax()
		assert.False(t, ok)
		assert.Len(t, hip.tab, 0)
	})

	t.Run("multiple", func(t *testing.T) {
		hip := newArrayHeap()

		for i, v := range []int{4,3,6,8,32,2,1} {
			hip.insert(v)
			assert.Len(t, hip.tab, i+1)
		}
		assert.Len(t, hip.tab, 7)
		
		expectedSorted := []int{32,8,6,4,3,2,1}
		got := []int{}
		for {
			v, ok := hip.delMax()
			if !ok {
				break
			}
			got = append(got, v)
			assert.Len(t, hip.tab, len(expectedSorted) - len(got))
		}
		assert.Equal(t, expectedSorted, got)
		assert.Len(t, hip.tab, 0)

		_, ok := hip.max()
		assert.False(t, ok)
	})

	t.Run("inserts between pops", func(t *testing.T) {
		hip := newArrayHeap()
		hip.insert(5)
		hip.insert(8)
		hip.insert(1)
		hip.insert(13)

		popAssert(t, hip, 13)
		popAssert(t, hip, 8)

		hip.insert(2)
		hip.insert(12)

		popAssert(t, hip, 12)
		popAssert(t, hip, 5)
		popAssert(t, hip, 2)
		popAssert(t, hip, 1)

		assert.Len(t, hip.tab, 0)
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
}

func newArrayHeap() *arrayHeap {
	tab := []int{}// alternatively we can insert dummy first element to help with index math - parent idx/2, children: idx*2, idx*2 + 1

	return &arrayHeap{
		tab: tab,
	}
}

func (a *arrayHeap) insert(v int) {
	a.tab = append(a.tab, v)
	
	// swim up
	idx := len(a.tab)-1
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
	if len(a.tab) < 1 {
		return 0, false
	}
	return a.tab[0], true
}

func (a *arrayHeap) delMax() (int, bool) {
	if len(a.tab) < 1 {
		return 0, false
	}
	toRet := a.tab[0]

	// get the last element and push it down the heap
	a.tab[0] = a.tab[len(a.tab)-1]
	a.tab = a.tab[:len(a.tab)-1]
	idx := 0
	for {
		leftIdx, lOk := a.leftChildIdx(idx)
		rightIdx, rOk := a.rightChildIdx(idx)

		maxIdx := -1
		if lOk && rOk {
			if a.tab[leftIdx] > a.tab[rightIdx] {
				maxIdx = leftIdx
			} else {
				maxIdx = rightIdx
			}
		} else if lOk && !rOk {
			maxIdx = leftIdx
		} else if !lOk && rOk {
			maxIdx = rightIdx
		} else {
			break
		}

		if maxIdx != -1 && a.tab[maxIdx] > a.tab[idx] {
			swap(&a.tab[maxIdx], &a.tab[idx])
			idx = maxIdx
		} else {
			break
		}

	}
	return toRet, true
}

func (a *arrayHeap) leftChildIdx(idx int) (int,bool) {
	return idx*2 +1, (idx*2 +1) < len(a.tab)
}

func (a *arrayHeap) rightChildIdx(idx int) (int,bool) {
	return (idx*2 +2), (idx*2 +2) < len(a.tab)
}

func (a *arrayHeap) parentIdx(idx int) (int, bool) {
	return (idx-1)/2, (idx-1) >= 0
}
