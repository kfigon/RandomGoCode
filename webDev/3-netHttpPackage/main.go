package main

import (
	"io"
	"fmt"
	"html/template"
	"net/http"
)

type hotdog int
func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	if r.Method == "GET" && path == "/" {
		getHelloTemplate(w, nil)
	} else if r.Method == "POST" && path == "/" {
		// providedData := r.FormValue("myName")

		r.ParseForm()
		values := r.Form // get collection, we can iterate in template
		providedData := values.Get("myName")
		getHelloTemplate(w, providedData)
	}
}

func getHelloTemplate(w io.Writer, data interface{}){
	tp := template.Must(template.ParseFiles("template.html"))
	tp.Execute(w, data)
}

func main() {
	fmt.Println("Starting")
	var d hotdog
	http.ListenAndServe(":8080", d)
}