package main

import (
	"fmt"
)

func main() {
	//Slice expression is used to create a slice from another slice

	//Note: When a slice is created from another slice, its not making a copy but instead
	// the destination slice is sharing storage with the source slice

	x := []int {1,2,3,4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	a := []int{1,2,3,4}
	b := a[:2]
	c := a[1:]

	a[1] = 20
	b[0] = 10
	c[1] = 30

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}