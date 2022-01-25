package main

import (
	"fmt"
)

func bubble_sort(arr []int) {
	unsorted_until_index := len(arr) - 1
	sorted := false

	for sorted == false {
		sorted = true
		for i := 0; i < unsorted_until_index; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		unsorted_until_index -= 1
	}
}

func main() {
	arr1 := []int{65, 55, 45, 35, 25, 15, 10}
	fmt.Println("Unsorted Array: ", arr1)
	bubble_sort(arr1)
	fmt.Println("Sorted Array: ", arr1)
}
