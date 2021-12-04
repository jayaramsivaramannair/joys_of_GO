package main

import (
	"fmt"
)

//Note: An & cannot be used before a primitive literal (numbers, booleans, and strings) or a constant before they don't have memory addresses.
//Thus in order to create a pointer to a primitive type, declare a variable and point to it.

func stringp(s string) *string {
	return &s
}

func main() {

	type person struct {
		FirstName string
		MiddleName *string
		LastName string
	}
		
	p := person {
		FirstName: "Pat",
		MiddleName: stringp("Perry"),
		LastName: "Peterson",
	}

	fmt.Println(p.FirstName)
	fmt.Println(*p.MiddleName)
	fmt.Println(p.LastName)

}