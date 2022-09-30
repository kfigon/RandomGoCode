package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

type config struct {
	MainPort int `json:"main_port"`
	Backends []int `json:"backends"`
}

func validateConfig(c config) error {
	if c.MainPort == 0 {
		return fmt.Errorf("no main port provided")
	} else if len(c.Backends) == 0 {
		return fmt.Errorf("no backends provided")
	}

	return nil
}

func main() {
	path := flag.String("path", "", "path to configuration file in json format")
	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		fmt.Println("file not found:", path)
		return
	}
	var c config
	if err = json.NewDecoder(file).Decode(&c); err != nil {
		fmt.Println("invalid json file:", err)
		return
	}
	file.Close()

	if err = validateConfig(c); err != nil {
		fmt.Println("invalid config provided:", err)
		return
	}

	f := newForwarder(port(c.MainPort))
	for _, v := range c.Backends {
		go serv(port(v))
		f.add(port(v))
	}

	fmt.Printf("running forwarder on %d for ports: %v\n", c.MainPort, c.Backends)
	f.run()
}

func serv(p port) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("Served by port %d\n", p))
	})
	http.ListenAndServe(p.toAddr(), mux)
}


type port int
func (p port) toAddr() string {
	return ":"+ strconv.Itoa(int(p))
}

type forwarder struct {
	backends []port
	cnt int32
	mainPort port
}

func newForwarder(p port) *forwarder {
	return &forwarder{
		backends: []port{},
		mainPort: p,
	}	
}

func (f *forwarder) run() {
	ln, err := net.Listen("tcp", f.mainPort.toAddr())
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

		go f.forwardSingleConnection(conn, f.nextPort())
	}
}

func (f *forwarder) add(p port) {
	f.backends = append(f.backends, p)
}

func (f *forwarder) nextPort() port {
	next := atomic.AddInt32(&f.cnt, 1)
	return f.backends[int(next) % len(f.backends)]
}

func (f *forwarder) forwardSingleConnection(request net.Conn, destinationPort port) {
	dst, err := net.Dial("tcp", destinationPort.toAddr())
	if err != nil {
		fmt.Println("error connecting to backend", destinationPort)
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
	go copy(dst, request)
	go copy(request, dst)

	wg.Wait()
}