package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct {})
	now := time.Now()
	//Once the go keyword is added, the functions will be called but there will be no waiting and the main function will not wait for these functions to return values
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	//Waiting for value to be read from done four times since values are pushed four times into done variable inside each of the tasks
	<- done
	<- done
	<- done
	<- done
	fmt.Println("elapsed:", time.Since(now))

}

func task1(done chan struct {}) {
	//Wait a second before printing
	time.Sleep(100*time.Millisecond)
	fmt.Println("task1")
	done <- struct{}{}
}

func task2(done chan struct {}) {
	//Wait for two seconds before printing
	time.Sleep(200*time.Millisecond)
	fmt.Println("task2")
	done <- struct{}{}
}

func task3(done chan struct {}) {
	fmt.Println("task3")
	done <- struct{}{}

}

func task4(done chan struct {}) {
	//Wait a second before printing
	time.Sleep(100*time.Millisecond)
	fmt.Println("task4")
	done <- struct{}{}
}