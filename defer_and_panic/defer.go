package main

import (
	"fmt"
)

func main() {
	//When a function called is deferred using defer keyword, it gets added to the stack thus the last function call (on top of the stack ) will be executed first
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}