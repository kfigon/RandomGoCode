package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"log"
)

func main() {
	li,_ := net.Listen("tcp", ":8080")
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error during accept: ", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn)  {
	parseReq(conn)
	response(conn, `<html>hi!</html>`)
	conn.Close()
}

func parseReq(conn net.Conn) {
	log.Println("==============New req:")
	scan := bufio.NewScanner(conn)
	i := 0
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break
		}
		parse(line, i)
		i++
	}
}

func parse(line string, lineNum int)  {
	if lineNum == 0 {
		elements := strings.Fields(line)
		if len(elements) != 3 {
			log.Printf("Invalid staring line: %q\n", elements)
			return
		}
		log.Printf("Method: %q, URI: %q, protocolVer: %q\n", elements[0], elements[1], elements[2])
		return
	}

	h := strings.Split(line, ": ")
	if len(h) != 2 {
		log.Printf("Invalid header %v\n", h)
		return
	}
	log.Printf("Header: %q : %q'n", h[0], h[1])
}

func response(conn net.Conn, responseBody string)  {
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %v\r\n", len(responseBody))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "%v", responseBody)
}