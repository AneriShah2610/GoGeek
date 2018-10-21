package calculator

import (
	"fmt"
)

func RandomArrayGeneration() []int {

	// generate random size array
	var arraySize int
	fmt.Print("Define array size:")
	fmt.Scanln(&arraySize)
	arrayForMergeSort := make([]int, arraySize)
	fmt.Println("Add elements in array:")
	for i := range arrayForMergeSort {
		fmt.Scanln(&arrayForMergeSort[i])
	}
	return arrayForMergeSort
}
