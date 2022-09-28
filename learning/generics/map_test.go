package main

import (
	"strconv"
	"testing"
)

func mapFn[T any, K any](arr []T, fn func(T)K) []K {
	var out []K
	for _, v := range arr {
		out = append(out, fn(v))
	}
	return out
}

func filterFn[T any](arr []T, fn func(T)bool) []T {
	var out []T
	for _, v := range arr {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func TestFilterMap(t *testing.T) {
	in := []int{1,2,3,4}
	even := filterFn(in, func(t int) bool {return t % 2 == 0})
	double := mapFn(even, func(t int) int {return t * 2})
	strs := mapFn(double, strconv.Itoa)

	assertEqualArr(t, []string{"4", "8"}, strs)
}