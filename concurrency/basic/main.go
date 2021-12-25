package main

import (
	"fmt"
	"sync"
)


func main() {
	//1. Create a wg variable
	var wg sync.WaitGroup

	//2. Add the number of go routines that you need to wait for
	wg.Add(3)

	go func() {
		//3. Exit the function once done executing
		defer wg.Done()
		fmt.Println("1")
	}()

	go func() {
		//3. Exit the function once done executing
		defer wg.Done()
		fmt.Println("2")

	}()

	go func() {
		//3. Exit the function once done executing
		defer wg.Done()
		fmt.Println("3")
	}()

	//4. Wait for all the go routines to finish executing
	//However, all the go routines will end up executing in random order
	wg.Wait()

}