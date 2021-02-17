package sorting

func copyTab(tab []int) []int {
	out := make([]int, len(tab))
	for i := 0; i < len(tab); i++ {
		out[i] = tab[i]
	}
	return out
}

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

func selectionSort(tab []int) []int {
	out := copyTab(tab)

	findMinIdx := func(startIdx int) int {
		minIdx := startIdx
		for j := startIdx; j < len(out); j++ {
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