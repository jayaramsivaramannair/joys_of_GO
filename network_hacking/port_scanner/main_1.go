package main

import (
	"fmt"
	"net"
	"log"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")

	if err != nil {
		log.Printf("Check the address string!")
	}

	if err == nil {
		//This confirms that a successful connection with the host was made. 
		fmt.Println("Connection successful")
	}

}