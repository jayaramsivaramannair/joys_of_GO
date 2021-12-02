package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		// Anonymous function declared inside of a for loop
		func (j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		} (i)
	}
}


