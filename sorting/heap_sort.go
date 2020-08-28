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

	// context: it adds at the end of the heap, in order to continue being a complete binary tree
	// and after it it ensures that it stills a max heap

	// it can be: O(1) ~ O(log n)

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

// Remove an element from MaxHeap
func (h *MaxHeap) Remove() []int {

	if len(h.data) < 1 {
		return []int{}
	}

	// context: it removes from top & picks last element on the heap (array|slice) to replace it
	// in order to continue to be a complete binary tree (there's no gaps)
	// and after it we must ensure that it stills a max heap

	// it can be: O(1) ~ O(log n)

	lastPos := len(h.data) - 1

	h.data[0] = h.data[lastPos]
	h.data[lastPos] = 0

	h.checkChildren(0)

	return h.data
}

func (h *MaxHeap) checkChildren(pos int) {

	if len(h.data) <= pos*2+1 {
		return
	}

	var higherChildrenPos int

	if len(h.data)-1 == pos*2+1 || h.data[pos*2+1] > h.data[pos*2+2] {
		higherChildrenPos = pos*2 + 1
	} else {
		higherChildrenPos = pos*2 + 2
	}

	if h.data[higherChildrenPos] > h.data[pos] {
		h.data[pos], h.data[higherChildrenPos] = h.data[higherChildrenPos], h.data[pos]

		h.checkChildren(higherChildrenPos)
	}
}

// MaxHeapSort algorithm implementation
func MaxHeapSort(data []int) []int {

	// TODO

	// data = [50, 15, 10, 30]

	return data
}
