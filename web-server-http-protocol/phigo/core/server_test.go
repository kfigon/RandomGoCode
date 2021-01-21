package core

import (
	"testing"
	"math"
)

type foo struct{}

//todo: add assert (closure can help)
func (f foo) Write(data []byte) (int, error) {
	return 1, nil
}
func (f foo) Read(data []byte) (int, error) {
	output := []byte(`GET / HTTP/1.1\r\nHost: localhost:8080\r\nUser-Agent: curl/7.68.0\r\nAccept: */*\r\n\r\n`)

	for i := 0; i < int(math.Min(float64(len(data)), float64(len(output)))); i++ {
		data[i] = output[i]
	}
	return 1, nil
}

func TestHandleConnectionAllOk(t *testing.T)  {
	server := NewServer(8080, DefaultAllowedMethods())
	connectionMock := foo{}
	server.handleConnection(connectionMock)
}