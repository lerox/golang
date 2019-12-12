package sorting

import (
	"reflect"
	"testing"
)

func TestSelectionSortWithSuccess(t *testing.T) {
	expected := []int{1, 7, 77, 90, 9999, 88888}

	unordered := []int{90, 7, 1, 9999, 88888, 77}
	result := SelectionSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("MergeSort is returning %v and we expect %v", result, expected)
	}

	expected2 := []int{1, 100}

	unordered2 := []int{100, 1}
	result2 := SelectionSort(unordered2)

	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("MergeSort is returning %v and we expect %v", result2, expected2)
	}
}

func TestSelectionSortWithEmptySlice(t *testing.T) {
	var expected []int

	var unordered []int
	result := SelectionSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("We expected %v but we got %v", expected, result)
	}
}

func TestSelectionSortWithJustOneElement(t *testing.T) {
	expected := []int{1}

	unordered := []int{1}
	result := SelectionSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("We expected %v but we got %v", expected, result)
	}
}

func TestSelectionSortWithAnAlreadySortedSlice(t *testing.T) {
	expected := []int{1, 7, 77, 90, 9999, 88888}

	sorted := []int{1, 7, 77, 90, 9999, 88888}
	result := SelectionSort(sorted)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("SelectionSort is returning %v and we expect %v", result, expected)
	}
}

func TestSelectionSortWithSomeRepeatedElements(t *testing.T) {
	expected := []int{1, 7, 77, 77, 90, 123, 123, 9999, 88888}

	unordered := []int{90, 7, 1, 77, 123, 9999, 123, 88888, 77}
	result := SelectionSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("SelectionSort is returning %v and we expect %v", result, expected)
	}
}
