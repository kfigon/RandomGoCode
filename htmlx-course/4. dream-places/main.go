package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tmpl *template.Template 
var placesDb []Place

var favouritePlaces map[string]Place

type Place struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Image Img `json:"image"`
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type Img struct {
	Src string `json:"src"`
	Description string `json:"alt"`
}

func main() {
	port := 3000
	
	tmpl = template.Must(template.ParseFiles("index.html"))
	d, err := os.ReadFile("places.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(d, &placesDb); err != nil {
		panic(err)
	}
	favouritePlaces = map[string]Place{}

	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	http.HandleFunc("/style.css", style)
	http.HandleFunc("/favicon.ico", noop)

	http.HandleFunc("/", index)
	http.HandleFunc("POST /favourite/{id}", favourite)

	fmt.Println("started on port", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func noop(w http.ResponseWriter, r *http.Request) {}
func style(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"style.css")	
}

func index(w http.ResponseWriter, r *http.Request) {
	_ = tmpl.ExecuteTemplate(w, "index", remainingPlaces())
}

func remainingPlaces() []Place {
	out := []Place{}
	for _, v := range placesDb {
		if _, ok := favouritePlaces[v.Id]; !ok {
			out = append(out, v)
		}
	}
	return out
}

func favourite(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var f *Place
	for _, v := range placesDb {
		if v.Id == id {
			f = &v
			break
		}
	}

	// todo: rerender places
	if f == nil {
		return
	}

	favouritePlaces[f.Id] = *f
}