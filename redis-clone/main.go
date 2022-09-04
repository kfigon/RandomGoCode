package main

import (
	"flag"
	"fmt"
	srv "redis-clone/server"
	"redis-clone/client"
	"redis-clone/command"
)

func main() {
	conf, err := parseCliConfig()
	if err != nil {
		fmt.Println("invalid config:", err)
		return
	}
	conf.run()
}

type cliMode int
const (
	server cliMode = iota

	tcpClient
	getCmd
	setCmd
	deleteCmd
)

func parseCliConfig() (cliCommand, error) {
	data := flag.String("data", "", "data you want to send. It'll add termination in tcp mode\\r\\n. In SET - send data in key=value format")
	port := flag.Int("port", srv.DefaultPort, "Port you want to use")
	modeInt := flag.Int("mode", 0, "application mode. 0 - server; 1 - tcpClient; 2 - GET; 3 - SET; 4 - DELETE")
	flag.Parse()

	mode := cliMode(*modeInt)
	if mode != server && *data == "" {
		return nil, fmt.Errorf("no data provided in client mode")
	}

	switch mode {
	case server: return &serverCliCommand{port: *port}, nil
	case tcpClient: return &clientCliCommand{port: *port, data: *data+"\r\n"}, nil
	case getCmd: return &clientCliCommand{port: *port, data: command.BuildGetCommand(*data)}, nil
	case setCmd: return &clientCliCommand{port: *port, data: command.BuildSetCommand(*data)}, nil
	case deleteCmd: return &clientCliCommand{port: *port, data: command.BuildDeleteCommand(*data)}, nil
	default: return nil, fmt.Errorf("invalid mode provided: %v", mode)
	}
}

type cliCommand interface {
	run()
}

type serverCliCommand struct {
	port int
}
func (s *serverCliCommand) run() {
	fmt.Println("starting a server on port", s.port)
	srv.StartServer(s.port)
}

type clientCliCommand struct {
	port int
	data string
}
func (c *clientCliCommand) run() {
	fmt.Println("client mode - sending data to port", c.port)
    fmt.Printf("data:\t%+0x \n", []byte(c.data))

	res, err := client.SendData(c.port, []byte(c.data))
	if err != nil {
		fmt.Println("client error:", err)
		return
	}
	fmt.Println("Got response:")
	fmt.Printf("%q\n", string(res))
}