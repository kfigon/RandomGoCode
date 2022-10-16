package bagqueuestack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeListWithElements(vals []int) *singleLinkedList[int] {
	s := &singleLinkedList[int]{}
	for _, v := range vals {
		s.add(v)
	}
	return s
}

func TestSinglyLinkedList(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		s := &singleLinkedList[int]{}
		assert.Empty(t, s.getElements())
	})

	t.Run("single", func(t *testing.T) {
		s := &singleLinkedList[int]{}
		s.add(2)
		assert.Equal(t, []int{2}, s.getElements())
	})

	t.Run("many", func(t *testing.T) {
		s := makeListWithElements([]int{2,3,4,5})
		assert.Equal(t, []int{2,3,4,5}, s.getElements())
	})
}

func TestRemoveWithIndex(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		s := &singleLinkedList[int]{}
		for i := -10; i < 10; i++ {
			s.removeIdx(i)
		}
	})

	t.Run("non existing", func(t *testing.T) {
		s := makeListWithElements([]int{1,2,3,4})

		s.removeIdx(-1)
		s.removeIdx(4)
		s.removeIdx(5)
		s.removeIdx(6)

		assert.Equal(t, []int{1,2,3,4},s.getElements())
	})

	t.Run("first", func(t *testing.T) {
		s := makeListWithElements([]int{1,2,3,4})

		s.removeIdx(0)

		assert.Equal(t, []int{2,3,4}, s.getElements())
	})

	t.Run("last", func(t *testing.T) {
		s := makeListWithElements([]int{1,2,3,4})

		s.removeIdx(3)

		assert.Equal(t, []int{1,2,3}, s.getElements())
	})

	t.Run("middle", func(t *testing.T) {
		s := makeListWithElements([]int{1,2,3,4})

		s.removeIdx(1)
		s.removeIdx(1)

		assert.Equal(t, []int{1,4}, s.getElements())
	})

	t.Run("pre last", func(t *testing.T) {
		s := makeListWithElements([]int{1,2,3,4})

		s.removeIdx(2)

		assert.Equal(t, []int{1,2,4}, s.getElements())
	})

	t.Run("last with short list", func(t *testing.T) {
		s := makeListWithElements([]int{1,2})

		s.removeIdx(1)

		assert.Equal(t, []int{1}, s.getElements())
	})

	t.Run("pre last with short list", func(t *testing.T) {
		s := makeListWithElements([]int{1,2})

		s.removeIdx(0)

		assert.Equal(t, []int{2}, s.getElements())
	})
}

func TestReverseLinkedList(t *testing.T) {
	t.Fatal("1.3.30")
	t.Run("empty", func(t *testing.T) {
		s := &singleLinkedList[int]{}
		s.reverse()
		assert.Empty(t, s.getElements())
	})	

	t.Run("single element", func(t *testing.T) {
		s := makeListWithElements([]int{5})
		s.reverse()
		assert.Equal(t, []int{5},s.getElements())
	})

	t.Run("2 elements", func(t *testing.T) {
		s := makeListWithElements([]int{5,6})
		s.reverse()
		assert.Equal(t, []int{6,5},s.getElements())
	})

	t.Run("many elements", func(t *testing.T) {
		s := makeListWithElements([]int{1,2,3,4,5})
		s.reverse()
		assert.Equal(t, []int{5,4,3,2,1},s.getElements())
	})
}

type listNode[T any] struct {
	val T
	next *listNode[T]
}

func newListNode[T any](val T) *listNode[T] {
	return &listNode[T]{
		val: val,
	}
}

type singleLinkedList[T any] struct {
	root *listNode[T]
}

func (s *singleLinkedList[T]) add(val T) {
	newNode := newListNode(val)
	if s.root == nil {
		s.root = newNode
		return
	}

	last := s.root
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

func (s *singleLinkedList[T]) getElements() []T {
	out := []T{}

	ptr := s.root
	for ptr != nil {
		out = append(out, ptr.val)
		ptr = ptr.next
	}

	return out
}

func (s *singleLinkedList[T]) removeIdx(idx int) {
	if idx < 0 || s.root == nil {
		return
	} else if idx == 0 {
		s.root = s.root.next
		return
	}

	currentIdx := 0
	preEl := s.root
	for currentIdx != (idx-1) && preEl.next != nil {
		preEl = preEl.next
		currentIdx++
	}

	toDelete := preEl.next
	if toDelete == nil {
		preEl.next = nil
	} else {
		preEl.next = toDelete.next
	}
}

func (s *singleLinkedList[T]) reverse() {
	if s.root == nil || s.root.next == nil {
		return
	}

	prev := s.root
	cur := s.root.next
	for cur != nil {
		tmp := cur

		cur.next = prev
		prev = cur
		cur = tmp.next
	}
	s.root = prev
}