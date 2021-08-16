package main

import (
	"fmt"
)

func main() {
	fmt.Println("foo")
}

type node struct {
	val int
	next *node
}

type linkedList struct {
	head *node
}

func newStack() *linkedList {
	return &linkedList{}
}

func (s *linkedList) add(val int) {
	newnode := &node{val:val}
	if s.head == nil {
		s.head = newnode
		return
	}

	prelast := s.head
	for prelast.next != nil {
		prelast = prelast.next
	}

	prelast.next = newnode
}

func (s *linkedList) collect() []int {
	var out []int
	ptr := s.head
	for ptr != nil {
		out = append(out, ptr.val)
		ptr = ptr.next
	}
	return out
}

type iterator struct {
	current *node
}

func (s *linkedList) iter() *iterator {
	return &iterator{s.head}
}

func (i *iterator) hasNext() bool {
	return i.current != nil
}

func (i *iterator) next() int {
	val := i.current.val
	i.current = i.current.next
	return val
}

func (l *linkedList) iterClosure() func()(int,bool) {
	node := l.head
	return func() (int, bool) {
		if node == nil {
			return 0, false
		}
		val := node.val
		node = node.next
		return val, true
	}
}

type optionInt struct {
	val int
	ok bool
}

func (l *linkedList) iterClosure2() func() optionInt {
	node := l.head
	return func() optionInt {
		if node == nil {
			return optionInt{}
		}
		val := node.val
		node = node.next
		return optionInt{val, true}
	}
}