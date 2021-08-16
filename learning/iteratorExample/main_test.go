package main

import (
	"testing"
)

func assertTab(t *testing.T, got []int, exp []int) {
	if len(got) != len(exp) {
		t.Fatalf("invalid len, got %v, exp %v", len(got), len(exp))
	}
	for i := 0; i < len(got); i++ {
		g := got[i]
		e := exp[i]
		if e != g {
			t.Errorf("Invalid el on %v, got %v, exp %v", i, g, e)
		}
	}
}

func buildInit() *linkedList {
	s := newStack()
	s.add(2)
	s.add(3)
	s.add(4)
	s.add(5)
	return s
}

func TestCollect(t *testing.T) {
	s := buildInit()
	assertTab(t, s.collect(), []int{2,3,4,5})
	assertTab(t, s.collect(), []int{2,3,4,5})
}

func TestIterator1(t *testing.T) {
	s := buildInit()
	var el []int
	it := s.iter()
	for it.hasNext() {
		el = append(el, it.next())
	}
	assertTab(t, el, []int{2,3,4,5})
	assertTab(t, s.collect(), []int{2,3,4,5})
}

func TestIterator2(t *testing.T) {
	s := buildInit()
	var el []int
	it := s.iterClosure()
	for {
		val, ok := it()
		if !ok {
			break
		}
		el = append(el, val)
	}
	assertTab(t, el, []int{2,3,4,5})
	assertTab(t, s.collect(), []int{2,3,4,5})
}

func TestIterator3(t *testing.T) {
	s := buildInit()
	var el []int
	it := s.iterClosure2()
	for valOpt := it(); valOpt.ok; valOpt = it(){
		el = append(el, valOpt.val)
	}
	assertTab(t, el, []int{2,3,4,5})
	assertTab(t, s.collect(), []int{2,3,4,5})
}