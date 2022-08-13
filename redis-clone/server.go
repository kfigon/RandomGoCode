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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	data, b, err := readSocket(conn)
	if err != nil {
		fmt.Println("error reading socket", err)
		return
	}
	handleCommand(data[0:b])
	conn.Write([]byte("+ok\r\n"))
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
