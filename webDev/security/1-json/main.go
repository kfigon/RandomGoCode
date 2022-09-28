package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/encode", handleGet)
	mux.HandleFunc("/decode", handlePost)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

type person struct {
	Name string
	Age int		`json:"theAge"`
}

// curl -i localhost:8080/encode
func handleGet(w http.ResponseWriter, req *http.Request) {
	data := person{"Michael", 25}
	
	encoded, _ := json.Marshal(data)
	header := w.Header()
	header.Add("Content-Type", "application/json")
	log.Println("Responding with", string(encoded))
	w.Write(encoded)

	// or just
	// json.NewEncoder(w).Encode(data)
}

//  curl -i -d '{ "Name": "ziomx", "theAge": 123 }' 'localhost:8080/decode'
func handlePost(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var p person
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(p)
}