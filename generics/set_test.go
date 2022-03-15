package main

import "testing"

type key interface {
	comparable
}

type void struct{}

type set[T key] map[T]void

func newSet[T key]() set[T] {
	return set[T]{}
}

func (s set[T]) add(val T) {
	s[val] = void{}
}

func (s set[T]) contains(val T) bool {
	_, ok := s[val]
	return ok
}

func (s set[T]) len() int {
	return len(s)
}

func perform[T key](t *testing.T, val T) {
	s := newSet[T]()
	assertEqual(t, 0, s.len())
	assertEqual(t, false, s.contains(val))

	s.add(val)
	assertEqual(t, 1, s.len())
	assertEqual(t, true, s.contains(val))
}

func TestNewStringSet(t *testing.T) {
	t.Run("string set", func(t *testing.T) {
		perform(t, "foobar")
	})

	t.Run("int set", func(t *testing.T) {
		perform(t, 5)
	})
}