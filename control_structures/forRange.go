package main

import (
	"fmt"
)

func main() {
	//for-range loop is used to iterate over slices, maps, arrays and strings
	evenVals := []int{2, 4, 6, 8, 10, 12}

	//i and v are idiomatic variables used for iterating over a collection 
	// i indicates index/position of the value and v indicates the value itself
	// When iterating over a map, k is instead instead of i since the elements represent key-value pairs
	// if the index is not required in the body of the loop, then use an underscore _ instead.
	for i, v := range evenVals {
		fmt.Println(i, v)
	}


	oddVals := [] int{1, 3, 5, 7, 9, 11}

	for _, v := range oddVals {
		fmt.Println(v)
	}


	//We can choose to ignore the second value of the range if we do not intend to use the value inside the loop
	uniqueNames := map[string] bool {"Fred": true, "Raul" : true, "Wilma" : true}

	for k := range uniqueNames {
		fmt.Println(k)
	}

	//Note that for-range makes a copy and assigns each value to the v variable thus any changs made to v
	// does not affect these values in the compount type itself

	noVals := []int {2,4,6,8,10,12}
	for _, v := range noVals {
		v *= 2
	}

	fmt.Println(noVals)
}