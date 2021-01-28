package main

import (
	"io"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi")
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.NotFound)
	mux.HandleFunc("/resource", greet)

	return mux
}

func main() {

	http.ListenAndServe(":8080", createMux())
}