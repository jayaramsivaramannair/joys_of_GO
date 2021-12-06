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


//Note : Pointer Receivers vs Value Receivers
//1. When to use a pointer receiver vs a value receiver for a method depends on other methods declared on the type
//2. If the method modifies the receiver, then use a pointer receiver
//3. If the method does not modify the receiver, then use a value receiver
//4. If a type has any pointer receiver methods, then for the sake of consistency use pointer receivers for all methods
// Contd...4 even the ones which which do not modify the receiver.

func main() {
	c := Counter{}

	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}