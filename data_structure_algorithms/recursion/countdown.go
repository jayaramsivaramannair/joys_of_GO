package main

import (
	"fmt"
)

func countdown(number int) {
	fmt.Println(number)
	if number == 0 {
		return
	}
	countdown(number - 1)
}

func main() {
	countdown(20)
}
