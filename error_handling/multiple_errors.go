package main

import (
	"fmt"
	"errors"
)

//Using errors.New to catch errors
func doubleEven(i int) (int, error) {
	if i % 2 != 0 {
		return 0, errors.New("Only even numbers are processed")
	}

	return i * 2, nil
}

//Using fmt.Errorf to catch errors
func doubleEven2(i int) (int, error) {
	if i % 2 != 0 {
		return 0, fmt.Errorf("%d is not an even number", i)
	}

	return i * 2, nil
}

func main() {
	result, err1 := doubleEven(5)
	result2, err2 := doubleEven2(11)

	fmt.Println(result, err1)
	fmt.Println(result2, err2)
}