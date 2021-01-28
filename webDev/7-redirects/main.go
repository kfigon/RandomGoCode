package main

import (
	"io"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w,r, "/resource", http.StatusSeeOther)
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.NotFound)
	mux.HandleFunc("/resource", greet)
	mux.HandleFunc("/redirect", redirect)

	return mux
}

func main() {

	http.ListenAndServe(":8080", createMux())
}