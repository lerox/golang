package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	expected := []int{1, 7, 8, 10, 45, 123, 666, 999, 88888}

	unordered := []int{666, 123, 999, 1, 10, 45, 88888, 7, 8}
	result := BubbleSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("We expected %v but we got %v", expected, result)
	}
}

func TestBubbleSortWithEmptySlice(t *testing.T) {
	var expected []int

	var unordered []int
	result := BubbleSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("We expected %v but we got %v", expected, result)
	}
}

func TestBubbleSortWithJustOneElement(t *testing.T) {
	expected := []int{1}

	unordered := []int{1}
	result := BubbleSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("We expected %v but we got %v", expected, result)
	}
}

func TestBubbleSortWithAnAlreadySortedSlice(t *testing.T) {
	expected := []int{1, 7, 77, 90, 9999, 88888}

	sorted := []int{1, 7, 77, 90, 9999, 88888}
	result := BubbleSort(sorted)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("BubbleSort is returning %v and we expect %v", result, expected)
	}
}

func TestBubbleSortWithSomeRepeatedElements(t *testing.T) {
	expected := []int{1, 7, 77, 77, 90, 123, 123, 9999, 88888}

	unordered := []int{90, 7, 1, 77, 123, 9999, 123, 88888, 77}
	result := BubbleSort(unordered)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("BubbleSort is returning %v and we expect %v", result, expected)
	}
}
