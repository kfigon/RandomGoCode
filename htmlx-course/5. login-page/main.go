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
	http.HandleFunc("POST /login", login)
	http.HandleFunc("/authenticated", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<h1>hello mr president</h1>`))
	})

	fmt.Println("started on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func noop(w http.ResponseWriter, r *http.Request) {}

func index(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "index", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	usrError := ""
	passError := ""

	usr := r.FormValue("user")
	if usr == "" || usr == "asdf" {
		usrError = "invalid user name"
	}

	pass := r.FormValue("password")
	if pass == "" || pass == "asdf" {
		passError = "invalid password"
	} 

	if usrError == "" && passError == "" {
		w.Header().Set("HX-Redirect", "/authenticated")
		return
	} 
	_,_ = w.Write([]byte(fmt.Sprintf(`<div hx-swap-oob="true:#user-error">%s</div>`, usrError)))
	_,_ = w.Write([]byte(fmt.Sprintf(`<div hx-swap-oob="true:#pass-error">%s</div>`, passError)))

}