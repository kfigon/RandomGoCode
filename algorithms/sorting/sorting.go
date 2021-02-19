package main

func copyTab(tab []int) []int {
	out := make([]int, len(tab))
	for i := 0; i < len(tab); i++ {
		out[i] = tab[i]
	}
	return out
}

// O(n^2)
func bubbleSort(tab []int) []int {
	out := copyTab(tab)
	
	numberOfSortedElements := 0
	arrayLen := len(out)

	for i := 0; i < arrayLen; i++ {
		for j := 0; j < arrayLen-1-numberOfSortedElements; j++ {
			currentIdx := j
			nextIdx := j+1

			if out[currentIdx] > out[nextIdx] {
				tmp := out[currentIdx]
				out[currentIdx] = out[nextIdx]
				out[nextIdx] = tmp
			}
		}

		numberOfSortedElements++
	}
	return out
}

// O(n^2)
func insertionSort(tab []int) []int {
	out := copyTab(tab)
	
	for i := 1; i < len(out); i++ {
		curElement := out[i]
		j := i-1
		for j >= 0 && curElement < out[j] {
			out[j+1] = out[j]
			j--
		}
		out[j+1] = curElement
	}
	return out
}

// O(n^2)
func selectionSort(tab []int) []int {
	out := copyTab(tab)

	findMinIdx := func(startIdx int) int {
		minIdx := startIdx
		for j := startIdx+1; j < len(out); j++ {
			if out[j] < out[minIdx] {
				minIdx = j
			}
		}
		return minIdx
	}

	for i := 0; i < len(out); i++ {
		minIdx := findMinIdx(i)
		tmp := out[i]
		out[i] = out[minIdx]
		out[minIdx] = tmp
	}
	return out
}

// O(n+m)
func mergeTabs(a []int, b []int) []int {
	out := make([]int, len(a)+len(b))

	aIdx := 0
	bIdx := 0
	outIdx := 0
	for aIdx < len(a) && bIdx < len(b) {
		if a[aIdx] < b[bIdx] {
			out[outIdx] = a[aIdx]
			aIdx++
		} else {
			out[outIdx] = b[bIdx]
			bIdx++
		}
		outIdx++
	}
	appendRest := func(tab[]int, tabIdx *int) {
		for *tabIdx < len(tab) {
			out[outIdx] = tab[*tabIdx]
			*tabIdx++
			outIdx++
		}	
	}
	appendRest(a, &aIdx)
	appendRest(b, &bIdx)

	return out
}

// go test -v -run "TestSort/merge->*"
// O(nlogn)
func mergeSort(tab []int) []int {
	if len(tab) <= 1 {
		return tab
	}
	splittingPoint := len(tab)/2
	left := tab[:splittingPoint]
	right := tab[splittingPoint:]

	return mergeTabs(mergeSort(left), mergeSort(right))
}

func mergeSortParallel(tab []int) []int {
	out := copyTab(tab)
	return out
}

func splitThroughPivot(tab []int) ([]int, []int, int) {
	left := make([]int,0)
	right := make([]int,0)
	pivotIdx := 0
	pivot := tab[pivotIdx]
	for i := 1; i < len(tab); i++ {
		el := tab[i]
		if el < pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}
	return left, right, pivot
}

// O(nlogn)
// in worst case - O(n^2)
// better to use idx, not slices
func quickSort(tab []int) []int {
	if len(tab) <= 1 {
		return tab
	}
	left,right,pivot := splitThroughPivot(tab)

	out := make([]int,0)
	newLeft := quickSort(left)
	for i := 0; i < len(newLeft); i++ { out = append(out, newLeft[i]) }
	
	out = append(out, pivot)
	
	newRight := quickSort(right)
	for i := 0; i < len(newRight); i++ { out = append(out, newRight[i]) }
	return out
}

func raidxSort(tab []int) []int {
	out := copyTab(tab)
	return out
}