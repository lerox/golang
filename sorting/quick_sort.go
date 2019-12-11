package sorting

// QuickSort algorithm implementation
func QuickSort(data []int) []int {
	return quickSortAux(data, 0, len(data)-1)
}

func quickSortAux(data []int, l int, h int) []int {
	if l < h {
		p := partition(data, l, h)
		quickSortAux(data, l, p-1)
		quickSortAux(data, p+1, h)
	}

	return data
}

func partition(data []int, l int, h int) int {
	// j controls the loop
	// i controls what is `<=` in comparison with the pivot

	pivot := data[h]
	i := l

	for j := l; j < h; j++ {
		if data[j] <= pivot {
			data[j], data[i] = data[i], data[j]
			i++
		}
	}

	data[i], data[h] = data[h], data[i]

	return i
}
