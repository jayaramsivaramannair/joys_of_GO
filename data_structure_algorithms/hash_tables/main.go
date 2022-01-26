package main

import (
	"fmt"
)

func is_subset(arr1, arr2 []string) bool {
	var largerArray []string
	var smallerArray []string
	var hashTable = map[string]bool{}

	if len(arr1) > len(arr2) {
		largerArray = append(largerArray, arr1...)
		smallerArray = append(smallerArray, arr2...)
	} else {
		largerArray = append(largerArray, arr2...)
		smallerArray = append(largerArray, arr1...)
	}

	//Store all items from the largerArray inside a hashTable
	for _, v := range largerArray {
		hashTable[v] = true
	}

	//Iterate through each item in smallerArray and return false
	//if we encounter an item not inside hashtable
	for _, v := range smallerArray {
		//Check if the value exists in the hashtable or not
		_, ok := hashTable[v]

		//If the value does not exist
		if !ok {
			return false
		}
	}

	return true

}

func main() {

	arr1 := []string{"a", "b", "c", "d", "e", "f"}
	arr2 := []string{"b", "d", "f"}

	if is_subset(arr1, arr2) {
		fmt.Println("Array 2 is a subset of Array 1")
	} else {
		fmt.Println("Array 2 is NOT a subset of Array 1")
	}

	arr3 := []string{"a", "b", "c", "d", "e", "f"}
	arr4 := []string{"b", "d", "f", "h"}

	if is_subset(arr3, arr4) {
		fmt.Println("Array 4 is a subset of Array 3")
	} else {
		fmt.Println("Array 4 is NOT a subset of Array 3")
	}

	return

}
