package main

import (
	"io"
	"net/http"
	"strconv"
	"log"
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
	mux.HandleFunc("/dupa", http.NotFound) // common pattern for /favicon.ico

	// serve files without exposing our file system
	// http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	return mux
}

func main() {
	err := http.ListenAndServe(":8080", createMux())
	log.Fatal(err)
	
	// we can serve static webpages with go server:
	// http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	
}