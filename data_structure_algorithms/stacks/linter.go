package main

import (
	"fmt"
)

func lint(funnyCharacters string, braceStack Stack) {

	for _, v := range funnyCharacters {
		var last_brace rune
		if is_opening_brace(v) {
			braceStack.Push(v)
		} else if is_closing_brace(v) {
			if len(braceStack.items) > 0 {
				last_brace = braceStack.Pop()
			} else {
				fmt.Printf("%c doesn't have an opening brace\n", v)
				return
			}

			if !is_a_match(last_brace, v) {
				fmt.Printf("%c has mismatched opening brace\n", v)
				return
			}
		}
	}

	if len(braceStack.items) > 0 {
		fmt.Printf("%c does not have a closing brace\n", braceStack.Read())
		return
	}

	fmt.Printf("All the opening braces are properly closed!\n")
	return
}

func main() {

	braceCollection2 := "{(["
	braceCollection3 := "{([])"
	braceCollection := "({[]})"

	newStack := Stack{}.New()
	newStack2 := Stack{}.New()
	newStack3 := Stack{}.New()

	lint(braceCollection2, newStack2)
	lint(braceCollection3, newStack3)
	lint(braceCollection, newStack)

}

type Stack struct {
	items []rune
}

func (s Stack) New() Stack {
	s.items = []rune{}
	return s
}

func (s *Stack) Push(item rune) {
	(*s).items = append((*s).items, item)
	return
}

func (s Stack) Read() rune {
	//Reads the last item from the stack
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() rune {
	lastItem := (*s).items[len(s.items)-1]
	(*s).items = (*s).items[0 : len(s.items)-1]
	return lastItem
}

func is_opening_brace(brace rune) bool {
	opening_braces := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
	}

	_, ok := opening_braces[brace]

	if !ok {
		return false
	}

	return true
}

func is_closing_brace(brace rune) bool {
	closing_braces := map[rune]bool{
		')': true,
		']': true,
		'}': true,
	}

	_, ok := closing_braces[brace]

	if !ok {
		return false
	}

	return true
}

func is_a_match(opening_brace, closing_brace rune) bool {
	braces_collection := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	return braces_collection[opening_brace] == closing_brace
}
