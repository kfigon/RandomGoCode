package main

import (
	"time"
	"strings"
	"bufio"
	"net"
	"log"
	"fmt"
)

// call through telnet or browser
func main() {
	li, _ := net.Listen("tcp", ":8080")
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			log.Print("Got error during accept: ", err)
		}
		go handle(con) // multiple connections
	}
}

func handle(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(10*time.Second))

	skaner := bufio.NewScanner(conn)
	for skaner.Scan() {
		t := skaner.Text()

		log.Println("read line: ", t)
		if strings.Contains(t, "exit") {
			break
		}
		fmt.Fprintf(conn, "\nEcho: %q\n", t)
	}
	conn.Close()
}