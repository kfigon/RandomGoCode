package main

import (
	"net"
	"log"
	"phigo/core"
)

func listner()  {
	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		log.Fatal("error during opening connection", err)
	}
	defer conn.Close()
}

func main() {
	core.NewServer(8080, core.DefaultAllowedMethods()).Run()
}	