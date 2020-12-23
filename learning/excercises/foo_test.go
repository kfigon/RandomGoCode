package main

import (
	"testing"
)

func assertFibo(t *testing.T, arg, exp int)  {
	if r := fibo(arg); r != exp {
		t.Errorf("fibo(%d) != %v, got: %v\n", arg, exp, r)
	}
}

func assertFizzbuzz(t *testing.T, arg int, exp string) {
	if r := fizzbuzz(arg); r != exp {
		t.Errorf("fizzbuzz(%d) != %v, got: %v\n", arg, exp, r)
	}
}

func assertGenerateEven(t *testing.T, arg int, exp []int) {
	compare := func(a []int, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i,v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	if r := generateEven(arg); !compare(r, exp) {
		t.Errorf("generateEvent(%d) != %v, got: %v\n", arg, exp, r)
	}
}

func assertCharCount(t *testing.T, in string, exp map[string]int) {
	r := charCount(in)
	if len(r) != len(exp) {
		t.Errorf("different lengths of maps - got %v, exp %v. input %v, res %v, exp %v", len(r), len(exp), in, r, exp)
	}

	for key := range exp {
		if exp[key] != r[key] {
			t.Errorf("values on key %v - got %v, exp %v. input %v, res %v, exp %v", key, r[key], exp[key], in, r , exp)
		}
	}
}

func TestFibo(t *testing.T)  {
	assertFibo(t, 0, 1)
	assertFibo(t, 1, 1)
	assertFibo(t, 2, 2)
	assertFibo(t, 3, 3)
	assertFibo(t, 4, 5)
	assertFibo(t, 5, 8)
	assertFibo(t, 6, 13)
}

func TestFizzbuzz(t *testing.T)  {
	assertFizzbuzz(t, 0, "fizzbuzz")
	assertFizzbuzz(t, 1, "1")
	assertFizzbuzz(t, 2, "2")
	assertFizzbuzz(t, 3, "fizz")
	assertFizzbuzz(t, 4, "4")
	assertFizzbuzz(t, 5, "buzz")
	assertFizzbuzz(t, 6, "fizz")
	assertFizzbuzz(t, 10, "buzz")
	assertFizzbuzz(t, 15, "fizzbuzz")
	assertFizzbuzz(t, 16, "16")
}

func TestGenerateEven(t *testing.T)  {
	assertGenerateEven(t, 1, []int{0})
	assertGenerateEven(t, 2, []int{0})
	assertGenerateEven(t, 3, []int{0,2})
	assertGenerateEven(t, 5, []int{0,2,4})
	assertGenerateEven(t, 10, []int{0,2,4,6,8})
}

func TestCharCount(t *testing.T)  {
	assertCharCount(t, "kamil", map[string]int{"k":1, "a":1, "m":1,"i":1,"l":1})
	assertCharCount(t, "golang", map[string]int{"g":2, "o":1, "l":1,"a":1,"n":1})
	assertCharCount(t, "missisipi", map[string]int{"m":1, "i":4, "s":3,"p":1})
}