package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

//Method provided on a user defined type. In this case, a String() method is provided on the Person type.
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}


func main() {
	fmt.Println("Hello, World!")

	fred := Person{"Fred", "Smith", 42}
	fmt.Println(fred.String())
}