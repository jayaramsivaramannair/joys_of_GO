package main

import (
	"fmt"
	"time"
)

func main() {
	go work() //fork point is created where the function is invoked
	time.Sleep(100*time.Millisecond)
	fmt.Println("done waiting, main exits")

	//join points. The main function returns before the work() function gets a chance to finish executing
}

func work() {
	//Wait for 5 second before printing
	time.Sleep(500*time.Millisecond)
	fmt.Println("printing some stuff")
}