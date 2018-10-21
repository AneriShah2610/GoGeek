package calculator

import (
	"fmt"
)

func mergeSort(arrayForMergeSort []int) []int {
	if len(arrayForMergeSort) <= 1 {
		return arrayForMergeSort
	}
	mid := (len(arrayForMergeSort)) / 2
	return sortArray(mergeSort(arrayForMergeSort[:mid]), mergeSort(arrayForMergeSort[mid:]))
}
func sortArray(left, right []int) []int {
	splitArray := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(splitArray, right...)
		}
		if len(right) == 0 {
			return append(splitArray, left...)
		}
		if left[0] <= right[0] {
			splitArray = append(splitArray, left[0])
			left = left[1:]
		} else {
			splitArray = append(splitArray, right[0])
			right = right[1:]
		}
	}
	return splitArray
}
func MergeSort() {

	// generate random size array

	arrayForMergeSort := RandomArrayGeneration()
	fmt.Println("Array before MergerSort: ", arrayForMergeSort)
	fmt.Println("Array after MergeSort: ", mergeSort(arrayForMergeSort))
}
