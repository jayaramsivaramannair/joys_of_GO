package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	//If panic is encountered in code then any code after the panic will not execute
	panic("This is a panic")
	fmt.Println("End")
}