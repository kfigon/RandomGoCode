package main

import (
	"flag"
	"fmt"
)

func main() {
	conf := parseCliConfig()
	fmt.Println(conf)
	if err := conf.validate(); err != nil {
		fmt.Println("invalid config", err)
		return
	}

	if conf.mode == server {
		fmt.Println("starting a server on port", conf.port)
		startServer()
	} else if conf.mode == client {
		fmt.Println("sending data to port", conf.port)
		fmt.Printf("%x\n", conf.data)
		sendData(conf.port, []byte(conf.data))
	}
}

type cliMode int
const (
	server cliMode = iota
	client
)

type config struct {
	mode cliMode
	port int
	data string
}

func parseCliConfig() *config {
	conf := config{}
	flag.StringVar(&conf.data, "data", "", "data you want to send")
	flag.IntVar(&conf.port, "port", 6379, "Port you want to use")
	
	var mode int
	flag.IntVar(&mode, "mode", 0, "application mode. 0 - server; 1 - client")
	
	flag.Parse()

	conf.mode = cliMode(mode)
	return &conf
}

func (c *config) validate() error {
	if c.mode != server && c.mode != client {
		return fmt.Errorf("invalid mode provided: %v", c.mode)
	}

	return nil
}