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

func TestHeap_Add(t *testing.T) {

	testValues := []AddDataProvider{
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

	for i := range testValues {
		maxHeap := MaxHeap{}
		maxHeap.data = testValues[i].previous

		result := maxHeap.Add(testValues[i].input)

		if !reflect.DeepEqual(testValues[i].expected, result) {
			t.Errorf(
				"MaxHeap on execution %d number failed because we expected %v and we got %v",
				i,
				testValues[i].expected,
				result,
			)
		}
	}
}

type RemoveDataProvider struct {
	previous []int
	expected []int
}

func TestMaxHeap_Remove(t *testing.T) {

	expectedResults := []RemoveDataProvider{

		{previous: []int{}, expected: []int{}},
		//  100
		//  /
		// 50
		{previous: []int{100, 50}, expected: []int{50, 0}},

		//  100
		//  / \
		// 20  60
		{previous: []int{100, 20, 60}, expected: []int{60, 20, 0}},

		//  100
		//  /  \
		// 60  20
		{previous: []int{100, 60, 20}, expected: []int{60, 20, 0}},

		//     100
		//     /  \
		//   60   20
		//  / \
		// 1  30
		{previous: []int{100, 60, 20, 1, 30}, expected: []int{60, 30, 20, 1, 0}},

		//     100
		//     / \
		//   60   20
		//  / \  /  \
		// 1  30 5  10
		{previous: []int{100, 60, 20, 1, 30, 5, 10}, expected: []int{60, 30, 20, 1, 10, 5, 0}},
	}

	for k, i := range expectedResults {
		maxHeap := MaxHeap{}
		maxHeap.data = i.previous

		result := maxHeap.Remove()

		if !reflect.DeepEqual(i.expected, result) {
			t.Errorf(
				fmt.Sprintf(
					"We got an error in execution %d we expected %v and we got %v", k, i.expected, result,
				),
			)
		}
	}
}
