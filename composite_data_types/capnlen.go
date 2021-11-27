package main

import (
	"fmt"
)

func main() {
	var x []int16
	//len prints out the elements currently held by the slice
	// cap prints out the elements which the slice can hold
	//Note: if the slice capacity needs to be increased for accomodating an additional element, it is doubled until the capacity is less than 1024
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

}
