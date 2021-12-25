package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	//Create a new channel to contain value of a struct type
	done := make(chan struct{})

	//This is the fork point. work function is invoked inside an anonymous function and invoked at the same time
	go func() { 
		work()
		type emptyStruct struct {}
		done <- emptyStruct {}
	}()


	//This is the join point. The main function returns before the work() function gets a chance to finish executing
	<- done //Wait till the done channel is populated with a value to read from
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("done waiting, main exits")
}

func work() {
	//Wait for 5 second before printing
	time.Sleep(500*time.Millisecond)
	fmt.Println("printing some stuff")
}