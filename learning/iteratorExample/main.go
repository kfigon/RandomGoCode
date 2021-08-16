package main

import (
	"fmt"
)

func main() {
	fmt.Println("foo")
	s := newStack()
	s.add(2)
	s.add(3)
	s.add(4)
	s.add(5)

	fmt.Println("collect:",s.collect())

	var els []int
	it := s.iter()
	for it.hasNext() {
		els = append(els, it.next())
	}
	fmt.Println("collected iterator:",els)
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