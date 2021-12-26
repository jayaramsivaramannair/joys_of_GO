package main

import (
	"sync"
	"fmt"
)

//Using Wait Groups to process requests to API in batches when the API is rate-limited

type request func()

func main() {
	requests := map[int]request {}

	for i := 1; i <= 100; i++ {
		f := func(n int) request {
			return func() {
				fmt.Println("request ", n)
			}
		}

		requests[i] = f(i)
	}

	var wg sync.WaitGroup

	//Batch number
	max := 10

	for i := 0; i < len(requests); i += max {
		//Process each request by a batch of 10
		for j := i; j< i + max; j++ {
			wg.Add(1)
			go func(r request) {
				defer wg.Done()
				r()
			}(requests[j+1])
		}
		wg.Wait()
		fmt.Println(max, "requests processed")
	}
}