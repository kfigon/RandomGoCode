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
	http.HandleFunc("GET /get-data", load)
	http.HandleFunc("POST /send", sendData)

	fmt.Println("started on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func noop(w http.ResponseWriter, r *http.Request) {}

func index(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}

func load(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "get-data", "my data")
}

func sendData(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	w.Write([]byte(r.FormValue("the-data")))
}