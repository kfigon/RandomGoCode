package compiler

import "monkey-lang/objects"

type Stack[T any] struct {
	s []T
	pointer int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		s: make([]T, 0, 512),
		pointer: 0,
	}
}

func (s *Stack[T]) Push(v T) {
	s.s = append(s.s, v)
	s.pointer++
}

func (s *Stack[T]) Pop() T {
	s.pointer--
	out := s.s[s.pointer]
	return out
}

func (s *Stack[T]) Empty() bool {
	return s.pointer <= 0
}

type VM struct {
	instructions Instructions
	constants []objects.Object

	stack Stack[objects.Object]
}