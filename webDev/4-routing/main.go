package main

import (
	"io"
	"net/http"
)

func helloFunc(response http.ResponseWriter, req* http.Request) {
	io.WriteString(response, "Hello World!")
}

func cat(response http.ResponseWriter, req* http.Request) {
	io.WriteString(response, "Hello Cat!")
}

// mux == server
func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloFunc)
	mux.HandleFunc("/cat", cat) // /cat/ - match everything after cat

	return mux
}

func main() {
	http.ListenAndServe(":8080", createMux())
}