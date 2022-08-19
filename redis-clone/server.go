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
	cmd := command(data)

	if err := cmd.basicValidation(); err != nil {
		fmt.Println("Invalid command:", err)
		return nil
	}

	if cmd.isStringCmd() {
		fmt.Println("got STRING")
	} else if cmd.isBulk() {
		fmt.Println("got BULK")
	} else if cmd.isArray() {
		fmt.Println("got ARRAY")
	}
	return []byte("+ok\r\n")
}
