package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleAuth)
	mux.HandleFunc("/resource", handleSecure)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type Auth struct {
	User string 	`json:"user"`
	Password string `json:"password"`
}

// curl -i -d '{ "user": "ziomx", "password": "123" }' localhost:8080
func handleAuth(w http.ResponseWriter, r *http.Request) {
	log.Println("got request")
	var data Auth
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "error in decoding json", http.StatusBadRequest)
		return
	}

	log.Println("Json decoded", data)
	w.Write([]byte(`hello mr president`))
}

func handleSecure(w http.ResponseWriter, r *http.Request) {
	if !authorised(r) {
		http.Error(w, "Unauthenticated", http.StatusForbidden)
		return
	}

	w.Write([]byte(`hello mr president`))
}

func authorised(r *http.Request) bool {
	return false
}