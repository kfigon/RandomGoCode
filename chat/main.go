package main

import (
	"bufio"
	"fmt"
	"net"
)

// nc localhost 8000
func main() {
    fmt.Println("hello")
    fmt.Println(newServer().run(8000))
}

type clientChan chan<- string

type server struct {
    text chan string
    userEnter chan clientChan // channel of channels
    userLeft chan clientChan
}

func newServer() *server {
    return &server{
        text: make(chan string),
        userEnter: make(chan clientChan),
        userLeft: make(chan clientChan),
    }
}

func (s *server) run(port int) error {
    conn, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
    if err != nil {
        return err
    }
    defer conn.Close()

    go s.broadcaster()

    for {
        c, err := conn.Accept()
        if err != nil {
            fmt.Println("error accepting connection", err)
            continue
        }
        go s.handleConn(c)
    }
}

func (s *server) broadcaster() {
    clients := make(map[clientChan]struct{}) // all connected clients
    for {
        select {
        case msg := <-s.text:
            // Broadcast incoming message to all clients channels
            for client := range clients {
                client <- msg
            }

        case cli := <-s.userEnter:
            clients[cli] = struct{}{}

        case cli := <-s.userLeft:
            delete(clients, cli)
            close(cli)
        }
    }
}

func (s *server) handleConn(con net.Conn) {
    userChan := make(chan string)
    go func() {
        // write all other msgs to this chan
        for msg := range userChan {
            fmt.Fprintln(con, msg)
        }
    }()

    s.userEnter <- userChan
    addr := con.RemoteAddr()
    
    s.text <- fmt.Sprintf("%s joined", addr)
    
    defer con.Close()
    scanner := bufio.NewScanner(con)
    for scanner.Scan() {
        s.text <- fmt.Sprintf("%s says %s", addr, scanner.Text())
    }
    s.text <- fmt.Sprintf("%s left", addr)
    
    s.userLeft <- userChan
}