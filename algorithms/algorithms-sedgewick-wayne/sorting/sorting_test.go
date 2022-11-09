package sorting

import (
	"fmt"
	"sort"
	"testing"
	"math"

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
		{5,3,7,3,1,45,56,8,9,45,2,0,4,5,7,4,3,2,1,54,76},
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
		{"mergeSort", mergeSort},
		{"quickSort", quickSort},
		// {"shell", shellSort},
		{"radix", radixSort},
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

// todo: presort array to minimize number of swaps (by comparing elements with larger intervals)
// then finish with insertion sort
func shellSort(tab []int) []int {
	return tab
}

// O(nlogn)
func mergeSort(tab []int) []int{
	merge := func(a []int, b []int) []int {
		out := make([]int, len(a)+len(b))
		aI := 0
		bI := 0
		i := 0
		for aI < len(a) && bI < len(b) {
			if a[aI] < b[bI] {
				out[i] = a[aI]	
				aI++
				i++
			} else {
				out[i] = b[bI]	
				bI++
				i++
			}
		}

		for aI < len(a) {
			out[i] = a[aI]
			aI++
			i++
		}
		for bI < len(b) {
			out[i] = b[bI]
			bI++
			i++
		}
		return out
	}

	// non splitting recursive solution:
	// middle = lo + (hi - lo)/2
	if len(tab) < 2 {
		return tab
	}
	middle := len(tab)/2
	left := mergeSort(tab[:middle])
	right := mergeSort(tab[middle:])

	return merge(left, right)
}

// O(nlogn)
func quickSort(tab []int) []int{
	if len(tab) < 2 {
		return tab
	}
	pivot := 0
	left := []int{}
	right := []int{}
	for i, v := range tab {
		if i == pivot {
			continue
		}
		if v < tab[pivot] {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)
	out := []int{}
	out = append(out, sortedLeft...)
	out = append(out, tab[pivot])
	out = append(out, sortedRight...)
	return out
}


func radixSort(tab []int) []int {
	// idx from 0
	extractDigit := func (num int, digitNum int) int {
		return (num/int(math.Pow10(digitNum)) % 10)
	}

	out := copyArr(tab)

	singleIteration := func(idx int) bool {
		buckets := make([][]int,100)
		for i := 0; i < len(buckets); i++ {
			buckets[i] = make([]int,0)
		}
		
		canContinue := false
		for _, v := range out {
			digit := extractDigit(v,idx)
			if digit != 0 {
				canContinue = true
			}
			list := buckets[digit]
			list = append(list, v)
			buckets[digit] = list
		}
		// short circuit
		if !canContinue {
			return canContinue
		}

		outIdx := 0
		for _,subList := range buckets {
			for _,el := range subList {
				out[outIdx] = el
				outIdx++
			}
		}
		return canContinue
	}

	idx := 0
	for singleIteration(idx) {
		idx++
	}

	return out
}