package main

import (
	"fmt"
)

func linear_search(arr1 []int, search_value int) (int, error) {
	for i, v := range arr1 {
		if v == search_value {
			return i, nil
		}
	}

	return -1, fmt.Errorf("Value is not found within the array")
}

func main() {
	arr1 := []int{3, 17, 75, 80, 202}
	searchValue := 22

	index, foundError := linear_search(arr1, searchValue)

	if foundError != nil {
		fmt.Println(foundError)
		return
	}

	fmt.Printf("%d is found at Index %d\n", searchValue, index)
}
