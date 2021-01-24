package main

import (
	"io"
	"fmt"
	"html/template"
	"net/http"
)

type hotdog int
func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Got method", r.Method)
	// providedData := r.FormValue("myName")

	r.ParseForm() // data in query na in form (request data)
	values := r.Form // map[string][]string
	fmt.Println("Provided data:", values)
	providedData := values.Get("myName")
	getHelloTemplate(w, providedData)
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