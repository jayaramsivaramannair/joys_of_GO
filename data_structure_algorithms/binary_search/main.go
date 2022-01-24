package main

import (
	"fmt"
)

func binary_search(arr []int, search_value int) (int, error) {
	lower_bound := 0
	upper_bound := len(arr) - 1

	for lower_bound <= upper_bound {
		midpoint := (upper_bound + lower_bound) / 2

		midpoint_value := arr[midpoint]

		if search_value == midpoint_value {
			return midpoint, nil
		} else if search_value < midpoint_value {
			upper_bound = midpoint - 1
		} else if search_value > midpoint_value {
			lower_bound = midpoint + 1
		}

	}
	return -1, fmt.Errorf("Value is not found within the array")
}

func main() {
	arr1 := []int{3, 17, 75, 80, 202}
	searchValue := 30

	index, foundError := binary_search(arr1, searchValue)

	if foundError != nil {
		fmt.Println(foundError)
		return
	}

	fmt.Printf("%d is found at Index %d\n", searchValue, index)

}
