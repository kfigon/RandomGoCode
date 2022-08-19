package main

import (
	"fmt"
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

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error when accepting connection", err)
			return
		}
		handleConnection(conn)
	}
	fmt.Println("that's all folks")
}

// todo: accept conn io.ReadWriteCloser for better testability
func handleConnection(conn net.Conn) {
	defer conn.Close()

	data, b, err := readSocket(conn)
	if err != nil {
		fmt.Println("error reading socket", err)
		return
	}
	resp := handleCommand(data[0:b])
	conn.Write(resp)
}

func handleCommand(data []byte) []byte {
	fmt.Println(string(data))
	
	c, err := parseCommand(data)
	if err != nil {
		fmt.Println("invalid command:", err)
		return []byte("-ok\r\n")
	}

	switch e := c.(type) {
	case *simpleStringCommand:
		fmt.Println("got string", e.simpleString())
	case *bulkCommand:
		fmt.Println("got bulk", e.bulkString())
	case *arrayCommand:
		fmt.Println("got array", e.commands())
	default:
		fmt.Println("invalid cmd")
	}
	return []byte("+ok\r\n")
}
