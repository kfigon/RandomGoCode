package main

import "testing"

type ifoobar[T any] interface {
	do() T
}

type strFn func() string
func (s strFn) do() string {
	return s()
}

type intFn func() int
func (i intFn) do() int {
	return i()
}

func processor[X any, T ifoobar[X]](iface T) X {
	return iface.do()
}

func TestConstraint(t *testing.T) {
	sFn := func() string { return "hello" }
	iFn := func() int { return 5 }
	
	assertEqual(t, "hello", processor[string, strFn](sFn))
	assertEqual(t, 5, processor[int, intFn](iFn))
}