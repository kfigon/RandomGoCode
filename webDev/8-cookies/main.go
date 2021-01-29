package main

import (
	"io"
	"net/http"
)

// cookie - small file that server can write to clients machine!
// used for adding state to http (session)
// cookies are per domain
// browser checks if we got cookie for that domain and just sends that every time, on every request
func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi!")
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)
	return mux
}

func main() {
	http.ListenAndServe(":8080", createMux())
}