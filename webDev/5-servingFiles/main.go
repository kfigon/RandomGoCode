package main

import (
	"io"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi!")
}

func getPicture(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/pic.jpg">`) // this will call to /pic.jpg
	// pic.jpg also works. / - absolute path
}

func servePicture(w http.ResponseWriter, req *http.Request) {
	// f, err := os.Open("pic.jpg")
	// if err != nil {
	// 	http.Error(w, "file not found", http.StatusNotFound)
	// 	return
	// }
	// defer f.Close()
	// io.Copy(w, f)

	http.ServeFile(w, req, "pic.jpg")
}


func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)
	mux.HandleFunc("/picture", getPicture)
	mux.HandleFunc("/pic.jpg", servePicture) // serve picture by exposing resource not working without it!
	

	return mux
}

func main() {
	http.ListenAndServe(":8080", createMux())
}