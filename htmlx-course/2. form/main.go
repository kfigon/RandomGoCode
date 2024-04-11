package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template 

var goals []string

func main() {
	port := 3000

	tmpl = template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"style.css")	
	})
	http.HandleFunc("/favicon.ico", noop)

	http.HandleFunc("/", index)
	http.HandleFunc("/store", sendData)

	fmt.Println("started on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func noop(w http.ResponseWriter, r *http.Request) {}

func index(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "index", goals)
}

func sendData(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	v := r.FormValue("goal")
	
	if v != "" {
		goals = append(goals, v)
		// alternatively just return the single element
		// hx-swap="beforeend"
		// w.Write([]byte(fmt.Sprintf(`<li>%s</li>`, v)))
	}
	_ = tmpl.ExecuteTemplate(w, "goal-list", goals)
}