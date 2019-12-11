package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSortWithSuccess(t *testing.T) {
	expected := []int{1, 7, 77, 90, 9999, 88888}

	unordered := []int{90, 7, 1, 9999, 88888, 77}
	result := QuickSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("MergeSort is returning %v and we expect %v", result, expected)
	}
}

func TestQuickSortWithEmptySlice(t *testing.T) {
	var expected []int

	var unordered []int
	result := QuickSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("We expected %v but we got %v", expected, result)
	}
}

func TestQuickSortWithJustOneElement(t *testing.T) {
	expected := []int{1}

	unordered := []int{1}
	result := QuickSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("We expected %v but we got %v", expected, result)
	}
}

func TestQuickSortWithAnAlreadySortedSlice(t *testing.T) {
	expected := []int{1, 7, 77, 90, 9999, 88888}

	sorted := []int{1, 7, 77, 90, 9999, 88888}
	result := QuickSort(sorted)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("QuickSort is returning %v and we expect %v", result, expected)
	}
}

func TestQuickSortWithSomeRepeatedElements(t *testing.T) {
	expected := []int{1, 7, 77, 77, 90, 123, 123, 9999, 88888}

	unordered := []int{90, 7, 1, 77, 123, 9999, 123, 88888, 77}
	result := QuickSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("QuickSort is returning %v and we expect %v", result, expected)
	}
}
