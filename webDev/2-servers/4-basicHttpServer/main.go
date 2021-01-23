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
	req := parseReq(conn)
	response(conn, req)
	conn.Close()
}

type req struct {
	method string
	url string
	headers map[string]string
	body []string
}

func parseReq(conn net.Conn) *req {
	log.Println("==============New req:")
	scan := bufio.NewScanner(conn)
	i := 0
	var r req
	r.headers = make(map[string]string)
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break
		}
		parse(line, i, &r)
		i++
	}
	return &r
}

func parse(line string, lineNum int, r *req)  {
	if lineNum == 0 {
		elements := strings.Fields(line)
		if len(elements) != 3 {
			log.Printf("Invalid staring line: %q\n", elements)
			return
		}
		log.Printf("Method: %q, URI: %q, protocolVer: %q\n", elements[0], elements[1], elements[2])
		r.method = elements[0]
		r.url = elements[1]
		return 
	}

	h := strings.Split(line, ": ")
	if len(h) != 2 {
		log.Printf("Invalid header %v\n", h)
		return
	}
	log.Printf("Header: %q : %q'n", h[0], h[1])
	r.headers[h[0]] = h[1]
}

func response(conn net.Conn, r *req)  {
	log.Println(r)
	switch r.url {
	case "/": handleIndex(conn)
	default: handleUnknown(conn)
	}
	
}

func handleIndex(conn net.Conn)  {
	responseBody := `<html>hi!</html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %v\r\n", len(responseBody))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "%v", responseBody)
}

func handleUnknown(conn net.Conn)  {
	fmt.Fprintf(conn, "HTTP/1.1 404 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %v\r\n", 0)
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
}