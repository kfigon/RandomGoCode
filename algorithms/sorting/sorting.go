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
	
	for i := 0; i < len(out); i++ {
		curIdx := i
		prevIdx := i-1
		for prevIdx >= 0 && out[curIdx] < out[prevIdx] {
			tmp := out[curIdx]
			out[curIdx] = out[prevIdx]
			out[prevIdx] = tmp
			curIdx--
			prevIdx--
		}
	}
	return out
}