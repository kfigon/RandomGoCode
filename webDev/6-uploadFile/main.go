package main

import (
	"net/http"
	"log"
	"io"
)

func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello!")
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)

	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", createMux()))
}