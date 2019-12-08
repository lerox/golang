package sorting

// BubbleSort algorithm implementation
func BubbleSort(data []int) []int {
	for j := len(data) - 1; j > 0; j-- {

		alreadySorted := true

		for i := 0; i < j; i++ {
			if data[i] > data[i+1] {
				alreadySorted = false // I swapped, so...
				data[i+1], data[i] = data[i], data[i+1]
			}
		}

		if alreadySorted {
			break
		}
	}

	return data
}
