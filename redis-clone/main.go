package main

import (
	"flag"
	"fmt"
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
	client
)

func parseCliConfig() (cliCommand, error) {
	data := flag.String("data", "", "data you want to send. It'll add \\r\\n")
	port := flag.Int("port", defaultPort, "Port you want to use")
	modeInt := flag.Int("mode", 0, "application mode. 0 - server; 1 - client")
	flag.Parse()

	mode := cliMode(*modeInt)
	if mode == client && *data == "" {
		return nil, fmt.Errorf("no data provided in client mode")
	}

	switch mode {
	case server: return &serverCliCommand{port: *port}, nil
	case client: return &clientCliCommand{port: *port, data: *data+"\r\n"}, nil
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
	startServer(s.port)
}

type clientCliCommand struct {
	port int
	data string
}
func (c *clientCliCommand) run() {
	fmt.Println("client mode - sending data to port", c.port)
    fmt.Printf("data:\t%+0x \n", []byte(c.data))

	res, err := sendData(c.port, []byte(c.data))
	if err != nil {
		fmt.Println("client error:", err)
		return
	}
	fmt.Println("Got response:")
	fmt.Printf("%q\n", string(res))
}