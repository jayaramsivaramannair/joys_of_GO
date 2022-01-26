package main

import (
	"fmt"
)

func main() {

	newStruct := Stack{}.New()

	newStruct.Push("1")
	newStruct.Push("2")
	newStruct.Push("3")

	fmt.Println(newStruct.items)
	fmt.Println(newStruct.Read())

	fmt.Println(newStruct.Pop())
	fmt.Println(newStruct.items)

}

type Stack struct {
	items []string
}

func (s Stack) New() Stack {
	s.items = []string{}
	return s
}

func (s *Stack) Push(item string) {
	(*s).items = append((*s).items, item)
	return
}

func (s Stack) Read() string {
	//Reads the last item from the stack
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() string {
	lastItem := (*s).items[len(s.items)-1]
	(*s).items = (*s).items[0 : len(s.items)-1]
	return lastItem
}
