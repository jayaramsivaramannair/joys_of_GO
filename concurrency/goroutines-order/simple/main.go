package main

import (
	"fmt"
	"sync"
)


func main() {
	//Step: 1 Declare the variable wg to use WaitGroup from sync package
	var wg sync.WaitGroup
	//Step: 2: Add the # of go routines to run
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
				//Step:3 : Run this before exiting the function
				defer wg.Done()
				fmt.Println("go routine", i + 1)

		}(i)
	}

	//Step:4 : Wait for all the go routines to finish running
	wg.Wait()
}