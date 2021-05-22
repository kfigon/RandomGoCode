package main

import (
	"log"
	"time"
	"net/http"
	"github.com/google/uuid"
)

func main() {
	server := newServer(newLogin())
	log.Fatal(http.ListenAndServe(":8080", server))
}

func newServer(loginServ *login) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", withLog(loginServ.login))
	mux.HandleFunc("/auth", withLog(loginServ.auth))
	return mux
}

func withLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New().String()
		start := time.Now()
		log.Printf("Starting [%v] req %v to %v\n", uuid, r.Method, r.URL)
		next(w,r)
		log.Printf("Ending [%v] req %v to %v, time: %v\n", uuid, r.Method, r.URL, time.Since(start))
	}
}
