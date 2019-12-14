package sorting

import (
	"reflect"
	"testing"
)

type TestValue struct {
	input    int
	previous []int
	expected []int
}

func TestHeap_Add(t *testing.T) {
	testValues := []TestValue{
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

// TODO

//func TestHeapSortWithSuccess(t *testing.T) {
//	expected := []int{1, 7, 77, 90, 9999, 88888}
//
//	unordered := []int{90, 7, 1, 9999, 88888, 77}
//	result := HeapSort(unordered)
//
//	if !reflect.DeepEqual(expected, result) {
//		t.Errorf("HeapSort is returning %v and we expect %v", result, expected)
//	}
//}

//func TestHeapSortWithAnAlreadySortedSlice(t *testing.T) {
//	expected := []int{1, 7, 77, 90, 9999, 88888}
//
//	sorted := []int{1, 7, 77, 90, 9999, 88888}
//	result := HeapSort(sorted)
//
//	if !reflect.DeepEqual(expected, result) {
//		t.Errorf("HeapSort is returning %v and we expect %v", result, expected)
//	}
//}
//
//func TestHeapSortWithSomeRepeatedElements(t *testing.T) {
//	expected := []int{1, 7, 77, 77, 90, 123, 123, 9999, 88888}
//
//	unordered := []int{90, 7, 1, 77, 123, 9999, 123, 88888, 77}
//	result := HeapSort(unordered)
//
//	if !reflect.DeepEqual(expected, result) {
//		t.Errorf("HeapSort is returning %v and we expect %v", result, expected)
//	}
//}
//
//func TestHeapSortWithEmptySlice(t *testing.T) {
//	var expected []int
//
//	var unordered []int
//	result := HeapSort(unordered)
//
//	if !reflect.DeepEqual(expected, result) {
//		t.Errorf("We expected %v but we got %v", expected, result)
//	}
//}
//
//func TestHeapSortWithJustOneElement(t *testing.T) {
//	expected := []int{1}
//
//	unordered := []int{1}
//	result := HeapSort(unordered)
//
//	if !reflect.DeepEqual(expected, result) {
//		t.Errorf("We expected %v but we got %v", expected, result)
//	}
//}
