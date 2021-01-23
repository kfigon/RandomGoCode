package main

import (
	"io"
	"net"
	"log"
)
// https://tools.ietf.org/html/rfc7230 - standard http/1.1
// well work on TCP layer

// we can use telnet (putty) to dial to this connection
func main() {
	log.Print("starting server")
	li, _ := net.Listen("tcp", ":8080")
	defer li.Close()

	for {
		log.Println("Waiting for connection...")
		con, err := li.Accept()
		if err != nil {
			log.Print("Got error during accept: ", err)
		}
		log.Println("Got it!")
		io.WriteString(con, "Hello form my http server")
		con.Close()
	}
}