package main

import (
	"net/http"
	"log"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleReq)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleReq(w http.ResponseWriter, req *http.Request) {
	if !authenticated(req) {
		w.WriteHeader(http.StatusForbidden)
		return
	} 
	
	w.Write([]byte(`hello mr president`))
}

func authenticated(req *http.Request) bool {
	return false
}