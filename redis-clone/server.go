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
		data := scanner.Text()
		fmt.Println(data)
	}
	fmt.Println("that's all folks")
}