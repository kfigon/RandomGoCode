package core

import (
	"fmt"
	"net"
	"log"
	"phigo/core/dto"
)
const bufferSize = 4096



type HandlerFunction (func() string)

// HttpServer - server class
type HttpServer struct {
	Port int
	Running bool
	AllowedMethods map[dto.HttpEndpointId]HandlerFunction
}

func DefaultAllowedMethods() map[dto.HttpEndpointId]HandlerFunction {
	methods := make(map[dto.HttpEndpointId]HandlerFunction)
	methods[dto.HttpEndpointId{Method: "GET", Url: "/health"}] = func() string{return prepareBasicResponse("200", "Ok")}
	methods[dto.HttpEndpointId{Method: "GET", Url: "/"}] = func() string{return prepareBasicResponse("200", "Ok")}
	
	return methods
}

func (server *HttpServer) isMethodAllowed(request dto.HttpRequest) (HandlerFunction, bool) {
	toFind := request.HttpEndpointId
	fun, ok := server.AllowedMethods[toFind]
	return fun, ok
}

// NewServer - contructor for HttpServer
func NewServer(port int, allowedMethods map[dto.HttpEndpointId]HandlerFunction) *HttpServer {
	return &HttpServer{Port: port, 
		Running: true,
		AllowedMethods: allowedMethods,
	}
}

// Run - creates connection and handles request
func (server *HttpServer) Run() {
	ln, err := net.Listen("tcp", "localhost:"+ fmt.Sprintf("%d",server.Port))
	if err != nil {
		log.Fatal("can't create server", err)
	}
	defer ln.Close()

	for server.Running == true {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("error during listening to requests", err)
			continue
		}
		defer conn.Close()

		go server.handleConnection(conn)
	}
}

func (server *HttpServer) handleConnection(conn net.Conn)  {
	var data = make([]byte, bufferSize)
	byteLen, err := conn.Read(data)
	if err != nil {
		log.Println("error when reading data", err)
		sendResponse(prepareBasicResponse("500", "Internal Server Error"), conn)
		return
	}
	
	log.Println("response size", byteLen)
	request, err := dto.ParseResponse(data)
	if err != nil {
		log.Println("error when parsing request", err)
		sendResponse(prepareBasicResponse("400", "Bad Request"), conn)
		return
	}

	handler, allowed := server.isMethodAllowed(request)
	if allowed == false {
		log.Println("Not allowed method", err)
		sendResponse(prepareBasicResponse("404", "Not Found"), conn)
		return
	}

	log.Println("Allowed method", request.HttpEndpointId)
	sendResponse(handler(), conn)
}

func sendResponse(resp string, conn net.Conn) {
	byteLen, err := conn.Write([]byte(resp))
	if err != nil {
		log.Println("error during responding", err)
	}
	log.Println("sent response with bytes:", byteLen)
}

func prepareBasicResponse(code, status string) string {
	return fmt.Sprintf("HTTP/1.1 %s %s\r\n"+
		"Server: Phigo\r\n"+
		"Date: Mon, 28 Dec 2020 11:48:50 CET\r\n"+
		"Content-Type: text/html\r\n"+
		"Content-Length: 0\r\n\r\n", code, status)
}