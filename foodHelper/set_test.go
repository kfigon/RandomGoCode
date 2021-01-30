package main

import (
	"testing"
)

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
	set.add(10)

	assertContainsAll(t, set, []int{3,4,12,10})
	if set.has(1) {
		t.Error("4 shouldnt be in set")
	}
}

func TestRemove(t *testing.T) {
	set := newSet(5,6,7)
	set.remove(5)
	assertContainsAll(t, set, []int{6,7})
}

func TestRemoveNotExisting(t *testing.T) {
	set := newSet(5,6,7)
	set.remove(18)
	assertContainsAll(t, set, []int{6,7,5})
}

func TestIterateEmpty(t *testing.T) {
	set := newSet()
	elements := set.els()
	if gotElements := len(elements); gotElements != 0 {
		t.Errorf("Expected empty set, got: %v", gotElements)
	}
}
func TestIterateNotEmpty(t *testing.T) {
	set := newSet(5,6,7)
	elements := set.els()
	expLen := 3
	if gotElements := len(elements); gotElements != expLen {
		t.Errorf("Expected empty set, got: %v, exp :%v", gotElements, expLen)
	}
	contains := func(el int) bool {
		for _, v := range elements {
			if v == el{
				return true
			}
		}
		return false
	}
	if !contains(5) || !contains(6) || !contains(7) {
		t.Error("Required elements not contained")
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		desc	string
		first  []int	
		second []int
		exp []int
	}{
		{
			desc: "bothEmpty",
			first: []int{},
			second: []int{},
			exp: []int{},
		},
		{
			desc: "firstEmpty_secondNot",
			first: []int{},
			second: []int{5,6,7},
			exp: []int{5,6,7},
		},
		{
			desc: "firstNotEmpty_secondEmpty",
			first: []int{5,6,7},
			second: []int{},
			exp: []int{5,6,7},
		},
		{
			desc: "firstNotEmpty_secondNotEmpty",
			first: []int{5,6,7},
			second: []int{8,9},
			exp: []int{5,6,7,8,9},
		},
		{
			desc: "firstNotEmpty_secondNotEmpty2",
			first: []int{8,9},
			second: []int{5,6,7},
			exp: []int{5,6,7,8,9},
		},
		{
			desc: "withDuplicates",
			first: []int{8,9},
			second: []int{8,9,1},
			exp: []int{8,9,1},
		},
		{
			desc: "theSame",
			first: []int{8,9,1},
			second: []int{8,9,1},
			exp: []int{8,9,1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			set1 := newSet(tc.first...)
			set2 := newSet(tc.second...)

			result := set1.sum(set2)
			assertContainsAll(t, result, tc.exp)
		})
	}
}

func TestIntersection(t *testing.T) {
	testCases := []struct {
		desc	string
		first  []int	
		second []int
		exp []int
	}{
		{
			desc: "bothEmpty",
			first: []int{},
			second: []int{},
			exp: []int{},
		},
		{
			desc: "firstEmpty_secondNot",
			first: []int{},
			second: []int{5,6,7},
			exp: []int{},
		},
		{
			desc: "firstNotEmpty_secondEmpty",
			first: []int{5,6,7},
			second: []int{},
			exp: []int{},
		},
		{
			desc: "firstNotEmpty_secondNotEmpty",
			first: []int{5,6,7},
			second: []int{8,9},
			exp: []int{},
		},
		{
			desc: "firstNotEmpty_secondNotEmpty2",
			first: []int{7,9},
			second: []int{5,6,7},
			exp: []int{7},
		},
		{
			desc: "firstNotEmpty_secondNotEmpty2",
			first: []int{5,6,7},
			second: []int{7,9},
			exp: []int{7},
		},
		{
			desc: "withDuplicates",
			first: []int{8,9},
			second: []int{8,9,1},
			exp: []int{8,9},
		},
		{
			desc: "theSame",
			first: []int{8,9,1},
			second: []int{8,9,1},
			exp: []int{8,9,1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			set1 := newSet(tc.first...)
			set2 := newSet(tc.second...)

			result := set1.intersection(set2)
			assertContainsAll(t, result, tc.exp)
		})
	}
}


func TestImmutabilityOfSummed(t *testing.T) {
	firstSetData := []int{5,6,7}
	secondSetData := []int{8,9,2}
	set1 := newSet(firstSetData...)
	set2 := newSet(secondSetData...)

	set1.sum(set2)
	assertContainsAll(t, set1, firstSetData)
	assertContainsAll(t, set2, secondSetData)
}

func TestImmutabilityOfSummedWhenModifyingBoth(t *testing.T) {
	firstSetData := []int{5,6,7}
	secondSetData := []int{8,9,2}
	set1 := newSet(firstSetData...)
	set2 := newSet(secondSetData...)

	sum := set1.sum(set2)
	assertContainsAll(t, set1, firstSetData)
	assertContainsAll(t, set2, secondSetData)
	set1.add(100)
	set2.add(102)

	expectedResult := []int{5,6,7,8,9,2}
	assertContainsAll(t, sum, expectedResult)
	assertContainsAll(t, set1, []int{5,6,7,100})
	assertContainsAll(t, set2, []int{8,9,2,102})
}

