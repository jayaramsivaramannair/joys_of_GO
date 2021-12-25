package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("number of cores", runtime.NumCPU())
	//Step: 1 for WaitGroup
	var wg sync.WaitGroup
	//Step: 2 for WaitGroup
	wg.Add(10)

	now := time.Now()

	for i := 0; i < 10; i++ {
		go work(&wg, i + 1)
	}

	//Step: 4 for WaitGroup
	wg.Wait()
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("main is done")

}

func work(wg *sync.WaitGroup, id int) {
	//Step: 3 for WaitGroup
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task ", id, " is done")
}