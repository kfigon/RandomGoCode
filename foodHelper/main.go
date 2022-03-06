package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 8080

	http.HandleFunc("/api/healthcheck", healthcheck)

	log.Println("Starting on port", port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	out := struct {
		Status string `json:"status"`
	}{"ok"}
	json.NewEncoder(w).Encode(&out)
}