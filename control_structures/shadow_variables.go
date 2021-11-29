package main

import (
	"fmt"
)

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		//A new variable shadowing x is created here
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}