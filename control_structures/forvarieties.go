package main

import (
	"fmt"
)

func main() {
	//There are no parentheses in the for statement similar to if else statements
	// := must be used to initialize and declare a variable in for statements as var is not legal.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Second variety which looks like the while statement in other languages
	j := 1
	for j < 100 {
		fmt.Println(j)
		j = j * 2
	}
}