package main

import (
	"fmt"
)

func main() {
	x := []int{1,2,3,4}
	//Creating a smaller slice from a bigger slice makes the smaller slice still retain capacity of the bigger slice
	y := x[:2]

	fmt.Println(cap(x), cap(y))
	y = append(y, 30)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}