package main

import (
	"fmt"
	"net"
)

func sendData(port int, data []byte) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println("error connecting to the server", err)
		return
	}
	defer conn.Close()

	n, err := conn.Write(data)
	if err != nil {
		fmt.Println("error sending data",err)
		return
	}
	fmt.Println("send", n, "bytes")
}