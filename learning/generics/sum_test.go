package main

import "testing"


func sumInt(vals []int) int {
	var out int
	for _, v := range vals {
		out += v
	}
	return out
}

// or use number type:
type Number interface {
    int | float64
}
func sum[T int | float32](vals []T) T {
	var out T
	for _, v := range vals {
		out += v
	}
	return out
}
func TestSum(t *testing.T) {
	t.Run("non generic", func(t *testing.T) {
		assertEqual(t, 10, sumInt([]int{1,2,3,4}))
	})

	t.Run("generic - int", func(t *testing.T) {
		assertEqual(t, 10, sum([]int{1,2,3,4}))
	})

	t.Run("generic float", func(t *testing.T) {
		assertEqual(t, 10, sum([]float32{1,2,3,4}))
	})	
}