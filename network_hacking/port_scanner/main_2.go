package main

import (
	"fmt"
	"net"
)


func main() {
	//Filter 1 to 1024 ports
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			//Error signifies that the ports are either closed or filtered
			continue //Skips the iteration for the closed port
		}

		//Close the port connection once it has been successfully established
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}