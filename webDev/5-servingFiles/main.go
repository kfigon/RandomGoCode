package main

import (
	"io"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi!")
}

func servePicture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/pic.jpg">`)
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)
	mux.HandleFunc("/picture", servePicture)
	return mux
}

func main() {
	http.ListenAndServe(":8080", createMux())
}