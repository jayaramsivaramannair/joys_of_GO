package main
import (
	"fmt"
)

func main() {

	x := 15
	a := &x
	
	//a stores the memory address for the location where x is stored
	fmt.Println(a)
	//Prints the value stored at the memory address when asterisk is used
	fmt.Println(*a)

	*a = 5
	fmt.Println(*a)
	fmt.Println(*a**a)
}