package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIdx)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleIdx(w http.ResponseWriter, req *http.Request) {
	data := struct {
		Name string
		Age int		`json:"theAge"`
	} { "Michael", 25}
	
	encoded, _ := json.Marshal(data)
	header := w.Header()
	header.Add("Content-Type", "application/json")
	w.Write(encoded)
}