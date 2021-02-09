package main

import (
	"html/template"
	"log"
	"net/http"
	"io"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:8080", createMux()))
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/info", info)
	mux.HandleFunc("/data", data)
	return mux
}

func index(response http.ResponseWriter, request *http.Request) {
	indexTemplate(response)
}

func indexTemplate(response http.ResponseWriter) {
	htmlStr := `<h1>Welcome!</h1>
	<a href="/info">info</a><br>
	<a href="/data">data</a>`
	tpl := template.Must(template.New("").Parse(htmlStr))
	tpl.Execute(response, nil)
}

func info(response http.ResponseWriter, request *http.Request){
	io.WriteString(response, "info page!")
}

func data(response http.ResponseWriter, request *http.Request){
	io.WriteString(response, "data page!")
}