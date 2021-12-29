package main

import (
	"sync"
	"time"
	"fmt"
	"sync/atomic"
)

func main() {
	var count int64
	var wg sync.WaitGroup

	wg.Add(1)

	// Reader - This will run randomnly at any point during execution of the operation
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)

		fmt.Println("count in go routine", atomic.LoadInt64(&count))

	}()

	wg.Add(50)

	// Writers
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			//Second argument in atomic function specifies that the value at provided lcoation needs to be incremented by 1
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	fmt.Println("count in main", count)
}