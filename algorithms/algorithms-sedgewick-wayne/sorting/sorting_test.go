package sorting

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	testCases := [][]int{
		{6,5,4,3,2,1},
		{7,6,5,4,3,2,1},
		{5,3,7,3,1,45,56,8,9,45,2,4,5,7,4,3,2,1,54,76},
		{5,3,7,3,1,45,56,8,9,45,2,-1,4,5,7,4,3,2,1,54,76},
		{1,2,3,4,5,6},
		{1,2,3,4,5,6,7},
	}
	algos := []struct{
		desc string
		fn algoFn
	}{
		{"bubble", bubbleSort},
	}

	for _, algo := range algos {
		t.Run(algo.desc, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
					got := algo.fn(copyArr(tc))

					sorted := copyArr(tc)
					sort.Ints(sorted)

					assert.Equal(t, sorted, got)
				})
			}
		})
	}
}

func copyArr(tab []int) []int {
	out := []int{}
	for i := 0; i < len(tab); i++ {
		out = append(out, tab[i])
	}
	return out
}

type algoFn func([]int) []int

func bubbleSort(tab []int) []int {
	return tab
}