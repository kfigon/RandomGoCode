package basicalgos

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {

	tdt := []struct {
		desc string
		fun func([]int, int) int
	}{
		{"iterative ", binarySearch},
		{"recursive ", binarySearchRec},
	}
	for _, tc := range tdt {
		t.Run(tc.desc+"empty", func(t *testing.T) {
			assert.Equal(t, -1, tc.fun([]int{}, 0))
			assert.Equal(t, -1, tc.fun([]int{}, 12))
			assert.Equal(t, -1, tc.fun([]int{}, 159))
		})

		t.Run(tc.desc+"not found", func(t *testing.T) {
			assert.Equal(t, -1, tc.fun([]int{1,2,3,4,5}, 0))
			assert.Equal(t, -1, tc.fun([]int{1,2,3,4}, 12))
			assert.Equal(t, -1, tc.fun([]int{1,2,3,5}, 4))
			assert.Equal(t, -1, tc.fun([]int{1,2,3,5,7}, 4))
			assert.Equal(t, -1, tc.fun([]int{1,3,4,5}, 2))
			assert.Equal(t, -1, tc.fun([]int{1,3,4,5,7}, 2))
			assert.Equal(t, -1, tc.fun([]int{1,3,4,5}, -1))
			assert.Equal(t, -1, tc.fun([]int{1,3,4,5,7}, -1))
			assert.Equal(t, -1, tc.fun([]int{1,3,4,5}, 6))
			assert.Equal(t, -1, tc.fun([]int{1,3,4,5,7}, 6))
		})

		t.Run(tc.desc+"found", func(t *testing.T) {
			assert.Equal(t, 0, tc.fun([]int{1,2,3,4,5}, 1))
			assert.Equal(t, 1, tc.fun([]int{1,2,3,4}, 2))
			assert.Equal(t, 2, tc.fun([]int{1,2,3,5}, 3))
			assert.Equal(t, 4, tc.fun([]int{1,2,3,5,7}, 7))
			assert.Equal(t, 2, tc.fun([]int{1,3,4,5}, 4))
			assert.Equal(t, 2, tc.fun([]int{1,3,4,5,7},4))
			
			assert.Equal(t, 0, tc.fun([]int{1,3,4,5,7}, 1))
			assert.Equal(t, 1, tc.fun([]int{1,3,4,5,7}, 3))
			assert.Equal(t, 2, tc.fun([]int{1,3,4,5,7}, 4))
			assert.Equal(t, 3, tc.fun([]int{1,3,4,5,7}, 5))
			assert.Equal(t, 4, tc.fun([]int{1,3,4,5,7}, 7))

			assert.Equal(t, 0, tc.fun([]int{1,2,3,4,5}, 1))
			assert.Equal(t, 1, tc.fun([]int{1,2,3,4,5}, 2))
			assert.Equal(t, 2, tc.fun([]int{1,2,3,4,5}, 3))
			assert.Equal(t, 3, tc.fun([]int{1,2,3,4,5}, 4))
			assert.Equal(t, 4, tc.fun([]int{1,2,3,4,5}, 5))
		})
	}
}

func binarySearch(tab []int, val int) int {
	left := 0
	right := len(tab)-1
	for left <= right {
		mid := left + (right-left)/2
		candidate := tab[mid]
		if val < candidate {
			right = mid-1
		} else if val > candidate {
			left = mid+1
		} else {
			return mid
		}
	}

	return -1
}

func binarySearchRec(tab []int, val int) int {
	var foo func(int,int) int
	foo = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := left + (right-left)/2
		if val < tab[mid] {
			return foo(0, mid-1)
		} else if val > tab[mid] {
			return foo(mid+1, right)
		}
		return mid
	}
	return foo(0, len(tab)-1)
}