package main

import (
	"fmt"
)

func main() {
	x := []int {1,2,3,4}
	y := make([]int, 4)

	//To create a slice which is independent of the original use the copy function
	num := copy(y, x)
	fmt.Println(y, num)
}