package main

import (
	"io"
	"net/http"
)

func helloFunc (response http.ResponseWriter, req* http.Request) {
	io.WriteString(response, "Hello World!")
}

// mux == server
func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloFunc)

	return mux
}

func main() {
	http.ListenAndServe(":8080", createMux())
}