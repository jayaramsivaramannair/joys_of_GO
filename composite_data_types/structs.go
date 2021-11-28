package main

import (
	"fmt"
)

type person struct {
	name string
	age int
	pet string
}

func main() {
	fred := person{}
	julia := person {
		name: "Julia",
		age: 40,
		pet : "cat",
	}

	fred.name = "Fred"
	fred.age = 55
	fred.pet = "I Hate Pets!"

	fmt.Println(fred)
	fmt.Println(julia)
}