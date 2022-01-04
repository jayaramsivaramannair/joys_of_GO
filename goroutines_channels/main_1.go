package main

import (
	"fmt"
	"time"
)

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x - 1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}


func main() {
	//Function call is made for the animation and the program continues ahead without waiting
	go spinner(100 * time.Millisecond)
	const n = 45
	//The program will not finish executing until the result returns from the function call
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}