package main


import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	//Since we are not specifiying the number of go routines to wait for using Add. The main function will execute and exit without waiting for the go routine
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("go routine: done")

	}()

	wg.Wait()
	fmt.Println("executed immediately")
}