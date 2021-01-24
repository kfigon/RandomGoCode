package main

import (
	"io"
	"fmt"
	"html/template"
	"net/http"
)

type hotdog int
func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	// providedData := r.FormValue("myName")

	r.ParseForm() // data in query na in form (request data)
	values := r.Form // map[string][]string
	providedData := values.Get("myName")
	getHelloTemplate(w, providedData)
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // optional, go can fill it
}

func getHelloTemplate(w io.Writer, data interface{}){
	tp := template.Must(template.ParseFiles("baseTemplate.html", "template.html"))
	tp.Execute(w, data)
}

func main() {
	fmt.Println("Starting")
	var d hotdog
	http.ListenAndServe(":8080", d)
}