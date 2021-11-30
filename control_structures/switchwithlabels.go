package main

import (
	"fmt"
)


//Note: Use a loop to break out of a loop when a specific case of a swtich statement is met inside the loop.
func main () {

loop:
	for i := 0; i < 10; i++ {
		switch {
		 case i % 2 == 0:
				fmt.Println(i, "is even")
		 case i % 3 == 0:
				fmt.Println(i, "is divisible by 3 but not 2")
		 case i%7 == 0:
				fmt.Println("exit the loop!")
				break loop
		 default:
				fmt.Println(i, "is boring")
		}
	}
}