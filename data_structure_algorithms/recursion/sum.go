package main

import (
	"fmt"
)

func array_sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + array_sum(arr[1:len(arr)])
}

func main() {
	arr1 := []int{2, 3, 4, 5, 6}

	fmt.Println("Sum of elements inside the array is ", array_sum(arr1))
}
