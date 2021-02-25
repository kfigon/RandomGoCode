package main

import (
	"net/http"
	"log"
	"mywebapp/getlist"
)

func createMux(c *getlist.GetController) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req* http.Request) {
		getlist.Render(w, c.GetList())
	})
	return mux
}

func main() {
	controller := getlist.CreateGetListController(getlist.MakeDb())
	log.Fatal(http.ListenAndServe(":8080", createMux(controller)))
}
