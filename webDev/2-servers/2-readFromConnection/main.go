package main

import (
	"bufio"
	"net"
	"log"
)

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
	skaner := bufio.NewScanner(conn)
	for skaner.Scan() {
		t := skaner.Text()

		log.Println("read line: ", t)
		if t == "exit" {
			break
		}
	}
	defer conn.Close()
}