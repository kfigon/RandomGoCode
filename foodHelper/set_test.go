package main

import (
	"testing"
)

func TestCreateEmptySet(t *testing.T) {
	set := newSet()
	assertSize(t, set, 0)
	if set.has(1) {
		t.Error("Empty set does not have any element")
	}
}

func TestCreateWithElements(t *testing.T) {
	set := newSet(4,1,32)
	assertContainsAll(t, set, []int{1,4,32})
}

func TestInitWithDuplicates(t *testing.T) {
	set := newSet(1,1,3,1,2,3)
	assertContainsAll(t, set, []int{1,2,3})
	if set.has(4) {
		t.Error("4 shouldnt be in set")
	}
}
func assertSize(t *testing.T, s *set, expectedSize int) {
	if le := s.size(); le != expectedSize {
		t.Errorf("Length should be: %v, got: %v", expectedSize, le)
	}
}
func assertContainsAll(t *testing.T, s *set, expected []int) {
	assertSize(t, s, len(expected))
	for _, v := range expected {
		if !s.has(v) {
			t.Errorf("%v not found in set", v)
		}
	}
}

func TestAddWhenEmpty(t *testing.T) {
	set := newSet()
	expectedSize := 0
	assertSize(t, set, expectedSize)
	
	set.add(3)
	set.add(4)
	set.add(3)
	set.add(3)
	
	assertContainsAll(t, set, []int{4,3})
	if set.has(1) {
		t.Error("4 shouldnt be in set")
	}
}

func TestAddWithElements(t *testing.T) {
	set := newSet(12,10)
	assertSize(t, set, 2)
	set.add(3)
	set.add(4)
	set.add(3)
	set.add(3)
	if !set.has(3) || !set.has(4) || !set.has(12) || !set.has(10) {
		t.Error("Element should be present")
	}
	if set.has(1) {
		t.Error("4 shouldnt be in set")
	}
	assertSize(t, set, 4)
}

func TestRemove(t *testing.T) {
	t.Fail()
}

func TestRemoveNotExisting(t *testing.T) {
	t.Fail()
}

func TestIntersection(t *testing.T) {
	t.Fail()
}

