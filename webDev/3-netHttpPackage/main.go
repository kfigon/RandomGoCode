package main

import (
	"io"
	"fmt"
	"html/template"
	"net/http"
)

type hotdog int
func (h hotdog) ServeHTTP(response http.ResponseWriter, request *http.Request)  {
	// providedData := request.FormValue("myName")

	request.ParseForm() // data in query na in form (request data)
	values := request.Form // map[string][]string
	providedData := values.Get("myName")
	getHelloTemplate(response, providedData)
	response.Header().Set("Content-Type", "text/html; charset=utf-8") // optional, go can fill it
}

func getHelloTemplate(response io.Writer, data interface{}){
	tp := template.Must(template.ParseFiles("baseTemplate.html", "template.html"))
	tp.Execute(response, data)
}

func main() {
	fmt.Println("Starting")
	var d hotdog
	http.ListenAndServe(":8080", d)
}