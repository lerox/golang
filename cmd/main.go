package main

import (
	"fmt"
	"github.com/lerox/golang/sorting"
)

func main() {
	unordered := []int{90, 7, 1, 9999, 88888, 77}

	result := sorting.MergeSort(unordered)

	fmt.Println(result)
}
