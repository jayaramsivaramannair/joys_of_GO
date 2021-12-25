package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup

	//Wait for one go routine to finish executing
	wg.Add(1)

	//This is the fork point. work function is invoked inside an anonymous function and invoked at the same time
	go func() {
		//wg.Done tells that the work is done and you can go ahead and exit the function 
		defer wg.Done()
		work()
	}()

	//This is the join point. The main function returns before the work() function gets a chance to finish executing
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exits")
}

func work() {
	//Wait for 5 second before printing
	time.Sleep(500*time.Millisecond)
	fmt.Println("printing some stuff")
}