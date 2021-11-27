package main

import (
	"fmt"
)


func main() {
	//:= is used for assigning a initial value whereas assignment operator "=" is used for reassignment
	x := 5
	for {
		fmt.Println("Do stuff", x)
		x += 3
		if x > 25 {
			break
		}
	}
}