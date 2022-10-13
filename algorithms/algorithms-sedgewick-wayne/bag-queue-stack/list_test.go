package bagqueuestack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeListWithElements(vals []int) *singleLinkedList {
	s := &singleLinkedList{}
	for _, v := range vals {
		s.add(v)
	}
	return s
}

func TestSinglyLinkedList(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		s := &singleLinkedList{}
		assert.Empty(t, s.getElements())
	})

	t.Run("single", func(t *testing.T) {
		s := &singleLinkedList{}
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
		s := &singleLinkedList{}
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

type listNode struct {
	val int
	next *listNode
}

func newListNode(val int) *listNode {
	return &listNode{
		val: val,
	}
}

type singleLinkedList struct {
	root *listNode
}

func (s *singleLinkedList) add(val int) {
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

func (s *singleLinkedList) getElements() []int {
	out := []int{}

	ptr := s.root
	for ptr != nil {
		out = append(out, ptr.val)
		ptr = ptr.next
	}

	return out
}

func (s *singleLinkedList) removeIdx(idx int) {
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