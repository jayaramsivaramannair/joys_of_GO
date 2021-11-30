package main

import (
	"fmt"
	"errors"
)

//Go supports multiple return values unlike other programming languages
//However, the types of return values need to be specified in parenthesis
//Return values inside the function must be separated by commas and not to be put inside parenthesis
func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return numerator / denominator, numerator % denominator, nil
}

func main() {
	fmt.Println(divAndRemainder(5,2))
	fmt.Println(divAndRemainder(5,0))
}