package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	task1()
	task2()
	task3()
	task4()
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