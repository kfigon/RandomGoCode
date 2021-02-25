package main

import (
	"net/http"
	"log"
	"mywebapp/getlist"
)

func createMux(v *getlist.GetListView) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req* http.Request) {
		v.Render(w)
	})
	return mux
}

func main() {
	controller := getlist.CreateGetListController(getlist.MakeDb())
	view := getlist.CreateView(controller)
	log.Fatal(http.ListenAndServe(":8080", createMux(view)))
}

