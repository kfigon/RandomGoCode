package main

import (
	"bufio"
	"fmt"
	"net"
)

const port int = 6379

func startServer() {
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

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		data := scanner.Bytes()
		handleCommand(data)
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
