package main

import (
	"fmt"
)


func main() {
	//comman ok idiom is used to determine if a key exists inside a map or not
	//Usually if we try to access a value by providing a non-existent key, the value returned will be zero
	//However this does not tell us if the key actually exists inside the map or not
	//Thus, comma ok idiom helps us find this out

	m := map[string] int {
		"hello" : 5,
		"world" : 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}