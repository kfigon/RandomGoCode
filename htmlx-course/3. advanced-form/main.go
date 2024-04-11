package main

import (
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"strconv"
)

var tmpl *template.Template 

type Goal struct {
	Name string
	Done bool
	Id int
}

type Page struct {
	Goals []Goal
}
var pendingId int
var page Page

func main() {
	port := 3000

	tmpl = template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"style.css")	
	})
	http.HandleFunc("/favicon.ico", noop)

	http.HandleFunc("/", index)
	http.HandleFunc("POST /store", store)
	http.HandleFunc("POST /done/{id}", done)
	http.HandleFunc("DELETE /remove/{id}", delete)

	fmt.Println("started on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func noop(w http.ResponseWriter, r *http.Request) {}

func index(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "index", page)
}

func store(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	v := r.FormValue("goal")
	if v == "" {
		return
	}
	
	out := Goal{v, false, pendingId}
	pendingId++
	page.Goals = append(page.Goals, out)
	_ = tmpl.ExecuteTemplate(w, "goal", out)
}

func done(w http.ResponseWriter, r *http.Request) {
	idS := r.PathValue("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		return
	}

	var out *Goal
	for i := 0; i < len(page.Goals); i++ {
		if page.Goals[i].Id == id {
			page.Goals[i].Done = !page.Goals[i].Done
			out = &page.Goals[i]
		}
	}
	
	_ = tmpl.ExecuteTemplate(w, "goal", out)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	
	page.Goals = slices.DeleteFunc(page.Goals, func(g Goal) bool {
		return g.Id == i
	})
	// empty response
}