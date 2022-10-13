package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


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
		s := &singleLinkedList{}
		s.add(2)
		s.add(3)
		s.add(4)
		s.add(5)
		assert.Equal(t, []int{2,3,4,5}, s.getElements())
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