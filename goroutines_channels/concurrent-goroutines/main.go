package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main() {
	//Allocate two logical processors for the scheduler to use
	//Allocating two logical processors will make the goroutines run simultaenously and concurrently
	runtime.GOMAXPROCS(2)

	//wg is used for wait for the program to finish
	//Add a count of two, one for each goroutine

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		//Schedule the call to Done to tell main we are done.
		defer wg.Done()
	
		//Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		//Schedule the call to Done to tell main we are done.
		defer wg.Done()
	
		//Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	

	//Wait for the goroutines to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")

}