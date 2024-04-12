package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template 

func main() {
	port := 3000

	tmpl = template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"style.css")	
	})
	http.HandleFunc("/favicon.ico", noop)

	http.HandleFunc("/", index)
	http.HandleFunc("GET /get-data", dataEndpoint)

	fmt.Println("started on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func noop(w http.ResponseWriter, r *http.Request) {}

func index(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}

func dataEndpoint(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "data", "fun fact #1234")
	_ = tmpl.ExecuteTemplate(w, "modal-response", "Success")
}