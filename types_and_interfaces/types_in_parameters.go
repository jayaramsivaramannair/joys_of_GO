package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total int
	lastUpdated time.Time
}

//Since the total field inside Counter type is modified, Counter is provided as a pointer receiver to the method.
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

//Since this method merely prints updated values of fields inside the Counter type, Counter is provided as a value receiver.
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

//Since the struct is passed directly, any changes made to the struct inside the function will not stick as struct is passed by value.
func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

//Since a pointer to the struct is passed as a parameter, any changes made to the struct inside the function will be permananet
func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}


func main() {
	c := Counter {}
	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())

}