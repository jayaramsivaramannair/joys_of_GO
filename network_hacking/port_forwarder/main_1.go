package main

import (
	"net"
	"log"
	"io"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "joescatcam.website:80")

	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}

	defer dst.Close()

	go func() {
		//Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	//copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}


func main() {
	//Listen on local port 80
	listener, err := net.Listen("tcp", ":80")

	if err != nil {
		log.Fatalln("Unable to bind to the port")
	}

	log.Println("Listening on localhost:80")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		log.Println("Connection Received")
		go handle(conn)
	}

}