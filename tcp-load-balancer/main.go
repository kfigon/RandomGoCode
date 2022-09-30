package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	mainPort := 8000
	backendPort := 4000
	go serv(backendPort)

	ln, err := net.Listen("tcp", toAddr(mainPort))
	if err != nil {
		fmt.Println("error when opening port:", err)
		return
	}
	defer ln.Close()

	for {
		conn,err := ln.Accept()
		if err != nil {
			fmt.Println("error when reading connection:", err)
			return
		}

		go func (conn net.Conn)  {
			dst, err := net.Dial("tcp", toAddr(backendPort))
			if err != nil {
				fmt.Println("error connecting to backend", backendPort)
				return
			}

			var wg sync.WaitGroup

			copy := func(first, second net.Conn) {
				io.Copy(first, second)
				first.Close()
				second.Close()
				
				wg.Done()
			}

			wg.Add(2)
			// order doesnt matter - response will be empty anyway, till we copy all data from source
			go copy(dst, conn)
			go copy(conn, dst)

			wg.Wait()
		}(conn)
	}
}

func serv(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("reached backend")
		fmt.Fprintf(w, fmt.Sprintf("Served by port %d\n", port))
	})

	http.ListenAndServe(toAddr(port), mux)
}

func toAddr(port int) string {
	return ":"+ strconv.Itoa(port)
}