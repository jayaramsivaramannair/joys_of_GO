package main

import (
	"fmt"
	"math/rand"
)

//Like if statements in other languages, the conditions specified in if else part does not require surrounding parenthesis

func main() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

}