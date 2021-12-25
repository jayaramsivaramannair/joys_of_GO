package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//Once the go keyword is added, the functions will be called but there will be no waiting and the main function will not wait for these functions to return values
	go task1()
	go task2()
	go task3()
	go task4()
	fmt.Println("elapsed:", time.Since(now))

}

func task1() {
	//Wait a second before printing
	time.Sleep(100*time.Millisecond)
	fmt.Println("task1")
}

func task2() {
	//Wait for two seconds before printing
	time.Sleep(200*time.Millisecond)
	fmt.Println("task2")
}

func task3() {
	fmt.Println("task3")
}

func task4() {
	//Wait a second before printing
	time.Sleep(100*time.Millisecond)
	fmt.Println("task4")
}