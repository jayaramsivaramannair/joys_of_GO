package main

import (
	"fmt"
)

func main() {
	//Slice should always be created with an estimate of initial capacity so that it does not have to resize automatically
	//make function allows, in this case, to a slice of int with an initial capacity of 5
	//In this case all elements are initialized with 0 value
	x := make([]int, 5)

	//If we want to specify both the length and capacity while using the make function. capacity follows length
	y := make([]int, 5, 15)

	//We can also specify the capacity but set the length to zero
	z := make([]int, 0, 10)

	fmt.Println(x, len(x), cap(x))
	fmt.Println(y, len(y), cap(y))
	fmt.Println(z, len(z), cap(z))
	z = append(z, 4,5,6,7,8)
	fmt.Println(z, len(z), cap(z))
}