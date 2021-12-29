package main

import (
	"sync"
	"fmt"
	"sync/atomic"
)

func main() {

	var count int32
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		//Increment the count variable by 10
		atomic.StoreInt32(&count, 10)
	}()

	go func() {
		defer wg.Done()
		//Increment the count variable by -15
		atomic.StoreInt32(&count, -15)
	}()

	go func() {
		defer wg.Done()
		//Increment the count variable by 1
		atomic.StoreInt32(&count, 1)
	}()

	go func() {
		defer wg.Done()
		//Increment the count variable by 0
		atomic.StoreInt32(&count, 0)
	}()


	go func() {
		defer wg.Done()
		//Set the count variable to 100
		atomic.StoreInt32(&count, 100)
	}()

	wg.Wait()
	fmt.Println("count", count)
}