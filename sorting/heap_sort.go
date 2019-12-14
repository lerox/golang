package sorting

import "math"

// Heap struct
type MaxHeap struct {
	data []int
}

// Add an element to MaxHeap
func (h *MaxHeap) Add(i int) []int {
	newElemPos := len(h.data)

	h.data = append(h.data, i)

	parentPos := getParentPos(newElemPos)

	for h.data[newElemPos] > h.data[parentPos] {
		h.data[parentPos], h.data[newElemPos] = h.data[newElemPos], h.data[parentPos]
		newElemPos = parentPos

		parentPos = getParentPos(newElemPos)
	}

	return h.data
}

func getParentPos(newElemPos int) int {
	return int(math.Floor(float64((newElemPos - 1) / 2)))
}

// Remove an element to heap
// Returns the removed element and also the MaxHeap
func (h *MaxHeap) Remove() (int, []int) {
	// TODO

	return 1, h.data
}

// MaxHeapSort algorithm implementation
func MaxHeapSort(data []int) []int {

	// TODO

	// data = [50, 15, 10, 30]

	return data
}
