package main

import (
	"fmt"
)

func reversal(str string) string {
	if len(str) == 1 {
		return string(str[0])
	}
	return reversal(str[1:len(str)]) + string(str[0])
}

func main() {
	str1 := "abcdefg"

	fmt.Println("String before reversal: ", str1)
	fmt.Println("String after reversal: ", reversal(str1))
}
