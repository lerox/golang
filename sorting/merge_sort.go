package sorting

// sort by merging
// divide & conquer
// give me a small problem and I can solve it!

import (
	"math"
)

// MergeSort algorithm implementation
func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	middle := int(math.Floor(float64(len(data) / 2)))

	left := make([]int, middle)
	right := make([]int, len(data) - middle)

	for i := 0; i < middle; i++ {
		left[i] = data[i]
	}

	for j := 0; j < len(data) - middle; j++ {
		right[j] = data[j + middle]
	}

	lSorted := MergeSort(left)
	rSorted := MergeSort(right)

	return merge(data, lSorted, rSorted)
}

func merge(data []int, left []int, right []int) []int {
	var i,j,k int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		data[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		data[k] = right[j]
		j++
		k++
	}

	return data
}
