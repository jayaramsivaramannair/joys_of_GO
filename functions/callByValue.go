package main

import (
	"fmt"
)

type person struct {
	age int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

//Note:
//string, integer and struct passed as parameter into a function are passed by value.
//This means that they are copied into the function.
//Thus any change made to the copies will affect the original values.


//Note:
//However if a map or slice are passed into a function as parameters then any changes made to them will modify and replace the original values.
//But the if the change involves lenghtening the slice then it will not stick.

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
 	}

	 s = append(s, 10)
}

func main() {
	p := person {}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)

	fmt.Println(i, s, p);

	m := map[int]string {
		1 : "first",
		2 : "second",
	}

	modMap(m)
	fmt.Println(m)

	sl := []int {1, 2, 3}
	modSlice(sl)
	fmt.Println(sl)

}