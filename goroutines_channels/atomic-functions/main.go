//Atomic package is being used to provide safe access to numeric types
//and avoiding race conditions by writing and reading from a shared resource
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	//counter is a variable incremented by all goroutines
	counter int64

	//wg is used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {
	//Add a count of two, one for each goroutine
	wg.Add(2)

	//Create two goroutines
	go incCounter(1)
	go incCounter(2)

	//Wait for the goroutines to finish
	wg.Wait()

	//Display the final value
	fmt.Println("Final Counter:", counter)
}

//incCounter increments the package level counter variable.
func incCounter(id int) {
	//Schedule the call to Done to tell main we are done
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//Safely Add One To Counter.
		//AddInt64 function from the atomic package synchronizes the addition of integer values
		//to ensure that only one goroutine can perform and complete this add operation at a time.
		atomic.AddInt64(&counter, 1)
	}
}