package main

import (
	"io/ioutil"
	"net/http"
	"log"

)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil || len(data) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", upload)

	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", createMux()))
}