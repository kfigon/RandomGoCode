package sorting

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	testCases := [][]int{
		{},
		{1},
		{3,4},
		{4,3},
		{6,5,4,3,2,1},
		{6,3,5,4,3,2,1,3},
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
		{"bubbleSort", bubbleSort},
		{"selectionSort", selectionSort},
		{"insertionSort", insertionSort},
		{"shellSort", shellSort},
		{"quickSort", quickSort},
		{"mergeSort", mergeSort},
	}

	for _, algo := range algos {
		t.Run(algo.desc, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
					got := algo.fn(copyArr(tc))

					sorted := copyArr(tc)
					sort.Ints(sorted)

					assert.Equal(t, sorted, got)
					assert.True(t, isSorted(got))
				})
			}
		})
	}
}

func isSorted(tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] > tab[i+1] {
			return false
		}
	}
	return true
}

func copyArr(tab []int) []int {
	out := []int{}
	for i := 0; i < len(tab); i++ {
		out = append(out, tab[i])
	}
	return out
}

type algoFn func([]int) []int

func swap(i, j *int) {
	tmp := *j
	*j = *i
	*i = tmp
}

func bubbleSort(tab []int) []int {
	for i := 0; i < len(tab)-1; i++ {
		for j := 0; j < len(tab)-1-i; j++ {
			if tab[j] > tab[j+1] {
				swap(&tab[j], &tab[j+1])
			}
		}
	}
	return tab
}

// find the minimum of unsorted subarray, add to sorted part
// select the min item
func selectionSort(tab []int) []int{
	for i := 0; i < len(tab); i++ {
		minIdx := i
		for j := i+1; j < len(tab); j++ {
			if tab[j] < tab[minIdx]	{
				minIdx = j
			}
		}
		swap(&tab[i], &tab[minIdx])
	}
	return tab
}

// deck of cards
func insertionSort(tab []int) []int {
	for i := 1; i < len(tab); i++ {
		prev := i-1
		cur := i
		for prev >= 0 && tab[prev] > tab[cur] {
			swap(&tab[prev], &tab[cur])
			cur--
			prev--
		}
	}
	return tab
}

func shellSort(tab []int) []int {
	return tab
}

func quickSort(tab []int) []int{
	return tab
}

func mergeSort(tab []int) []int{
	return tab
}