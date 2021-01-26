package main

import (
	"io"
	"net/http"
	"strconv"
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
	// manual work
	// f, err := os.Open("pic.jpg")
	// if err != nil {
	// 	http.Error(w, "file not found", http.StatusNotFound)
	// 	return
	// }
	// defer f.Close()
	// io.Copy(w, f)

	http.ServeFile(w, req, "pic.jpg")

	// http.FileServer() - we can serve whole directories
}

func downloadFile(w http.ResponseWriter, request *http.Request) {
	fileName := "aFile.txt"
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(fileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, request, fileName)
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)
	mux.HandleFunc("/picture", getPicture)
	mux.HandleFunc("/pic.jpg", servePicture) // serve picture by exposing resource not working without it!
	mux.HandleFunc("/downloadFile", downloadFile)
	

	return mux
}

func main() {
	http.ListenAndServe(":8080", createMux())
}