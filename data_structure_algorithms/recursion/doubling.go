package main

import (
	"fmt"
)

func doubling(arr []int, index int) {
	if index >= len(arr) {
		return
	}
	arr[index] *= 2
	doubling(arr, index+1)
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Array before doubling: ", arr1)
	doubling(arr1, 0)
	fmt.Println("Array after doubling: ", arr1)
}
