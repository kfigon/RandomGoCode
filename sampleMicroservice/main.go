package main

import (
	"encoding/json"
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