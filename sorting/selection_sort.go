package sorting

// SelectionSort algorithm implementation
func SelectionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		lowestPos := i

		for j := i + 1; j < len(data); j++ {
			if data[j] < data[lowestPos] {
				lowestPos = j
			}
		}

		data[i], data[lowestPos] = data[lowestPos], data[i]
	}

	return data
}
