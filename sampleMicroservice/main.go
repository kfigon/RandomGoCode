package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func main() {
	server := newServer(newLogin())

	const port = 8080
	log.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), server))
}

func newServer(loginServ *login) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthEndpoint)
	mux.HandleFunc("/login", withLog(loginServ.login))
	mux.HandleFunc("/auth", withLog(loginServ.authHandler))
	mux.HandleFunc("/data", withLog(withSecurity(loginServ,data)))
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

func withSecurity(loginServ *login, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !loginServ.auth(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w,r)
	}
}

func data(w http.ResponseWriter, r *http.Request) {
	response := &struct {
		Data string `json:"data"`
	}{ "foobar" }
	json.NewEncoder(w).Encode(response)
}

func healthEndpoint(w http.ResponseWriter, r *http.Request) {
	response := &struct {
		Status string `json:"status"`
	}{ "up" }
	json.NewEncoder(w).Encode(response)
}