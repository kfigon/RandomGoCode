package main

import (
	"fmt"
	"io"
	"net"
)

func readSocket(conn net.Conn) ([]byte, int, error) {
	// bufio scanner will listen up to specific token.
	// it's better to handle full message and parse manualy
	
	// todo: add timeout and grow the buffer if needed. For now 8k will do
	data := make([]byte, 8*1024)
	b, err := conn.Read(data)
	if err != nil && err != io.EOF {
		return nil, 0, err
	}
	fmt.Println("read", b ,"bytes")
	return data, b, nil
}