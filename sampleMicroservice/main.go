package main

import (
	"log"
	"net/http"
)

func main() {
	server := newServer(newLogin())
	log.Fatal(http.ListenAndServe(":8080", server))
}

func newServer(loginServ *login) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginServ.login)
	mux.HandleFunc("/auth", loginServ.auth)
	return mux
}

