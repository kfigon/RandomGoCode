package main

import (
	"fmt"
	"net"
	"log"
)

func listner()  {
	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		log.Fatal("error during opening connection", err)
	}
	defer conn.Close()
}

func createServer()  {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("can't create server", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("error during listening to requests", err)
			continue
		}
		defer conn.Close()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn)  {
	var data = make([]byte, 512)
	byteLen, err := conn.Read(data)
	if err != nil {
		log.Println("error when reading data", err)
	}
	log.Println("response size", byteLen)
	fmt.Println(string(data))

	responseBody := fmt.Sprintf("HTTP/1.1 200 OK\r\n"+
					"Server: Phigo\r\n"+
					"Date: Mon, 28 Dec 2020 11:48:50 CET\r\n"+
					"Content-Type: text/html\r\n"+
					"Content-Length: 0\r\n\r\n")

	byteLen, err = conn.Write([]byte(responseBody))
	if err != nil {
		log.Println("error during responding", err)
	}
	log.Println("sent response with bytes:", byteLen)
}

func main() {
	createServer()
}	