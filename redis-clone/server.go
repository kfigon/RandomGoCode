package main

import (
	"fmt"
	"io"
	"net"
)

const defaultPort int = 6379

func startServer(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println("Error during creating connection", err)
		return
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("error when accepting connection", err)
		return
	}
	defer conn.Close()

	// bufio scanner will listen up to specific token.
	// it's better to handle full message and parse manualy
	
	// todo: add timeout and grow the buffer if needed. For now 64k will do
	for {
		data := make([]byte, 64*1024)
		b, err := conn.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("error when reading buffer", err)
				return 
			}
		}
		fmt.Println("read", b ,"bytes")
		handleCommand(data[0:b])
	}
	fmt.Println("that's all folks")
}

func handleCommand(data []byte) {
	fmt.Println(string(data))
	cmd := command(data)

	if err := cmd.validate(); err != nil {
		fmt.Println("Invalid command:", err)
		return
	}

	if cmd.isStringCmd() {
		fmt.Println("got STRING", cmd.simpleString())
	} else if cmd.isBulk() {
		fmt.Println("got BULK")
	}else if cmd.isArray() {
		fmt.Println("got ARRAY")
	}
}
