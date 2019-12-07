package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSortWithSuccess(t *testing.T) {
	expected := []int{1, 7, 77, 90, 9999, 88888}

	unordered := []int{90, 7, 1, 9999, 88888, 77}
	result := MergeSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("MergeSort is returning %v and we expect %v", result, expected)
	}
}

func TestMergeSortWithAnAlreadySortedSlice(t *testing.T) {
	expected := []int{1, 7, 77, 90, 9999, 88888}

	sorted := []int{1, 7, 77, 90, 9999, 88888}
	result := MergeSort(sorted)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("MergeSort is returning %v and we expect %v", result, expected)
	}
}

func TestMergeSortWithSomeRepeatedElements(t *testing.T) {
	expected := []int{1, 7, 77, 77, 90, 123, 123, 9999, 88888}

	unordered := []int{90, 7, 1, 77, 123, 9999, 123, 88888, 77}
	result := MergeSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("MergeSort is returning %v and we expect %v", result, expected)
	}
}
