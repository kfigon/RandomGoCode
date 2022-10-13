package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	popAssert := func(t *testing.T, s *stack, exp int) {
		v, ok := s.pop()
		assert.True(t, ok)
		assert.Equal(t, exp, v)
	}

	t.Run("empty", func(t *testing.T) {
		s := &stack{}
		assert.Equal(t, 0, s.len())

		for i := 0; i < 5; i++ {
			_, ok := s.pop()
			assert.False(t, ok)
		}
	})

	t.Run("push", func(t *testing.T) {
		s := &stack{}
		s.push(5)
		assert.Equal(t, 1, s.len())

		popAssert(t, s, 5)

		for i := 0; i < 5; i++ {
			_, ok := s.pop()
			assert.False(t, ok)
		}
	})

	t.Run("push multiple", func(t *testing.T) {
		s := &stack{}
		s.push(5)
		s.push(6)
		s.push(7)
		assert.Equal(t, 3, s.len())

		popAssert(t,s,7)
		popAssert(t,s,6)
		assert.Equal(t, 1, s.len())

		popAssert(t,s,5)
		assert.Equal(t, 0, s.len())

		for i := 0; i < 5; i++ {
			_, ok := s.pop()
			assert.False(t, ok)
		}
	})

	t.Run("push pop push", func(t *testing.T) {
		s := &stack{}
		s.push(5)
		s.push(6)

		popAssert(t, s, 6)
		s.push(123)
		
		popAssert(t, s, 123)
		popAssert(t, s, 5)

		for i := 0; i < 5; i++ {
			_, ok := s.pop()
			assert.False(t, ok)
		}
	})
}

type stack struct {
	top *listNode
}

func (s *stack) len() int{
	ln := 0
	ptr := s.top
	for ptr != nil {
		ptr = ptr.next
		ln++
	}
	return ln
}

func (s *stack) pop() (int, bool) {
	if s.top == nil {
		return -1, false
	}
	toRet := s.top.val
	s.top = s.top.next
	return toRet, true
}

func (s *stack) push(val int) {
	newNode := newListNode(val)
	if s.top == nil {
		s.top = newNode
		return
	}

	newNode.next = s.top
	s.top = newNode
}