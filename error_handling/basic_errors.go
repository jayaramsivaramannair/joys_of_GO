package main

import (
	"fmt"
	"errors"
	"os"
)

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Denominator is 0")
	}
	remainder := numerator / denominator
	mod := numerator % denominator
	return remainder, mod, nil
}

//In order to return an error, use a simple if statement to check that the error variable is non-nil

func main() {
	numerator := 20
	denominator := 0

	remainder, mod, err := calcRemainderAndMod(numerator, denominator)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(remainder, mod)
}