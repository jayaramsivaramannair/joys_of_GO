package main

import (
	"sync"
)

func main() {
	//Step: 1 First Step for WaitGroup
	var wg sync.WaitGroup

	//Step: 2 Second Step for WaitGroup
	wg.Add(1)


	//Step: 4 Fourth and Final Step for WaitGroup
	wg.Wait()

	//Since Add is called without being followed by Done. The code will error out with a deadlock message
}