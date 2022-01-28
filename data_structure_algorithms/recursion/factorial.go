package main

import (
	"fmt"
)

func factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	fmt.Println("Factorial of 6 is, ", factorial(6))
	fmt.Println("Factorial of 7 is, ", factorial(7))
	fmt.Println("Factorial of 8 is, ", factorial(8))
}
