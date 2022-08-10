package main

import (
	"bufio"
	"fmt"
	"net"
)

const port int = 6379

func startServer() {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Println("Error during creating connection", err)
		return
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("error when accepting connection", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		data := scanner.Text()
		handleCommand(data)
	}
	fmt.Println("that's all folks")
}

func handleCommand(data string) {
	fmt.Println(data)
	cmd := newCmd(&data)

	if err := cmd.validate(); err != nil {
		fmt.Println("Invalid command:", err)
		return
	}

	if cmd.isStringCmd() {
		fmt.Println("got STRING")
	}
}

func newCmd(data *string) *command {
	c := command(*data)
	return &c
}

type command string
func (c *command) validate() error {
	if len(*c) < 3 {
		return fmt.Errorf("too short")
	} else if !c.isStringCmd() && (*c)[0] != '*' && (*c)[0] != '$' {
		return fmt.Errorf("invalid first character: %v", (*c)[0])
	} else if c.termination() != `\r\n` {
		return fmt.Errorf("invalid termination: %q", c.termination())
	}
	return nil
}

func (c *command) isStringCmd() bool {
	return (*c)[0] == '+'
}

func (c *command) termination() string {
	sub := string(*c)
	ln := len(sub)-2
	return sub[ln:]
}