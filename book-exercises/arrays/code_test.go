package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrays(t *testing.T) {
	modArr := func (a [3]int) { // full copy
		a[0] = 123
	}

	anArray := [...]int{1,2,3}
	modArr(anArray)
	assert.Equal(t, anArray, [3]int{1,2,3})
}

func TestSlice(t *testing.T) {
	modSlice := func (a []int) { // copy just the header, underlying can be change. Append wont work though
		a[0] = 123
	}

	slice := []int{1,2,3}
	modSlice(slice)
	assert.Equal(t, slice, []int{123,2,3})
}

// also - use copy() to make sure we're not modifying the original array when slicing

func TestIndexing(t *testing.T) {
	arr := []int{1,2,3,4}

	assert.Equal(t, []int{2,3,4}, arr[1:])
	assert.Equal(t, []int{1,2}, arr[:2])
	assert.Panics(t, func() {
		_ = arr[:80]
	})
}