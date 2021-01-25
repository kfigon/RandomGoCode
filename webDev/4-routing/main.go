package main

import (
	"io"
	"net/http"
)

type hotdog int
func (h hotdog) ServeHTTP (response http.ResponseWriter, req* http.Request) {
	io.WriteString(response, "Hello World!")
}

// mux == server
func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	var h hotdog
	mux.Handle("/", h)

	return mux
}

func main() {
	http.ListenAndServe(":8080", createMux())
}