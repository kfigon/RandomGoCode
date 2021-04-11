package main

import (
	"io"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/login", handleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl -s localhost:8080/login -i
func handleLogin(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hi there")
	// todo: get github auth url
	// redirect then
	// handle their redirect with separate handler (specified on oauth config)
}