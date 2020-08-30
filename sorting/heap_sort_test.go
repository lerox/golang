package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

type AddDataProvider struct {
	input    int
	previous []int
	expected []int
}

func Test_Max_Heap_Add(t *testing.T) {

	testData := []AddDataProvider{
		// 50
		{input: 50, previous: []int{}, expected: []int{50}},

		//  50
		//  /
		// 1
		{input: 1, previous: []int{50}, expected: []int{50, 1}},

		//  100
		//  / \
		// 1  50
		{input: 100, previous: []int{50, 1}, expected: []int{100, 1, 50}},

		//    100
		//    / \
		//   25  50
		//  /
		// 1
		{input: 25, previous: []int{100, 1, 50}, expected: []int{100, 25, 50, 1}},

		//     100
		//    /  \
		//   25   50
		//  / \
		// 1   2
		{input: 2, previous: []int{100, 25, 50, 1}, expected: []int{100, 25, 50, 1, 2}},

		//     999
		//    /  \
		//   25   100
		//  / \   /
		// 1   2 50
		{input: 999, previous: []int{100, 25, 50, 1, 2}, expected: []int{999, 25, 100, 1, 2, 50}},
	}

	for k, v := range testData {
		maxHeap := NewMaxHeap(v.previous)

		maxHeap.Add(v.input)

		if !reflect.DeepEqual(v.expected, maxHeap.data) {
			t.Errorf(
				"We got an error in execution %d, we expected %v and we got %v",
				k,
				v.expected,
				maxHeap.data,
			)
		}
	}
}

type RemoveDataProvider struct {
	previous []int
	expected RemoveExpectedData
}

type RemoveExpectedData struct {
	removedElem   int
	result        []int
	containsError bool
}

func Test_Max_Heap_Remove(t *testing.T) {

	testData := []RemoveDataProvider{

		{previous: []int{}, expected: RemoveExpectedData{result: []int{}, removedElem: 0, containsError: true}},
		//  100
		//  /
		// 50
		{previous: []int{100, 50}, expected: RemoveExpectedData{result: []int{50, 0}, removedElem: 100}},

		//  100
		//  / \
		// 20  60
		{previous: []int{100, 20, 60}, expected: RemoveExpectedData{result: []int{60, 20, 0}, removedElem: 100}},

		//  100
		//  /  \
		// 60  20
		{previous: []int{100, 60, 20}, expected: RemoveExpectedData{result: []int{60, 20, 0}, removedElem: 100}},

		//     100
		//     /  \
		//   60   20
		//  / \
		// 1  30
		{
			previous: []int{100, 60, 20, 1, 30},
			expected: RemoveExpectedData{result: []int{60, 30, 20, 1, 0}, removedElem: 100},
		},

		//     100
		//     / \
		//   60   20
		//  / \  /  \
		// 1  30 5  10
		{
			previous: []int{100, 60, 20, 1, 30, 5, 10},
			expected: RemoveExpectedData{result: []int{60, 30, 20, 1, 10, 5, 0}, removedElem: 100},
		},
	}

	for k, v := range testData {
		maxHeap := NewMaxHeap(v.previous)

		removed, err := maxHeap.Remove()

		if !reflect.DeepEqual(v.expected.result, maxHeap.data) {
			t.Errorf(
				fmt.Sprintf(
					"We got an error in execution %d, we expected result %v and we got %v",
					k,
					v.expected.result,
					maxHeap.data,
				),
			)
		}

		if v.expected.containsError && err == nil {
			t.Error(fmt.Sprintf("We got an error in execution %d, we expected an error as result", k))
		}

		if removed != v.expected.removedElem {
			t.Error(fmt.Sprintf(
				"We got an error in execution %d, we expected removedElem %v and we got %v",
				k,
				v.expected.removedElem,
				removed,
			))
		}
	}
}

type HeapifyDataProvider struct {
	previous []int
	expected []int
}

func Test_Max_Heap_Heapify(t *testing.T) {

	testData := []HeapifyDataProvider{
		{previous: []int{}, expected: []int{}},

		{previous: []int{100, 50}, expected: []int{100, 50}},

		{previous: []int{100, 20, 60}, expected: []int{100, 20, 60}},

		{previous: []int{100, 60, 20}, expected: []int{100, 60, 20}},

		//       100
		//      /   \
		//     60    80
		//    / \    / \
		//   1   3  20 30
		{previous: []int{60, 3, 20, 1, 100, 80, 30}, expected: []int{100, 60, 80, 1, 3, 20, 30}},
	}

	for k, v := range testData {

		maxHeap := NewMaxHeap(v.previous)
		maxHeap.Heapify()

		if !reflect.DeepEqual(maxHeap.data, v.expected) {
			t.Errorf("We got an error in execution %d, we expected %v and we got %v", k, v.expected, maxHeap.data)
		}
	}
}
