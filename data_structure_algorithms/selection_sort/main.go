package main

import (
	"fmt"
)

func selection_sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		//When starting off, think that it is the lowest value
		lowestNumberIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}

		//If it is not the same value that we started off the passthrough with
		if lowestNumberIndex != i {
			temp := arr[i]
			arr[i] = arr[lowestNumberIndex]
			arr[lowestNumberIndex] = temp
		}
	}
}

func main() {
	arr1 := []int{65, 55, 45, 35, 25, 15, 10}

	fmt.Println("Unsorted Array: ", arr1)
	selection_sort(arr1)
	fmt.Println("Sorted Array: ", arr1)
}
