package sorting

// BubbleSort algorithm implementation
func BubbleSort(data []int) []int {
	loop := true

	for loop {
		loop = false

		for i := 0; i < len(data) - 1; i++ {
			if data[i+1] < data[i] {
				bkp := data[i]
				data[i] = data[i+1]
				data[i+1] = bkp

				loop = true
			}
		}
	}

	return data
}
