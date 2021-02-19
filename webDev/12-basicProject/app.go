package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", createMux()))
}

func makeApp(db dataProvider) *app {
	return &app{db}
}

type app struct {
	db dataProvider
}

type dataProvider interface {
	readList() []todoListItem
	readEntry(int) *todoEntry
	insert(todoEntry) error
	update(todoEntry) error
}

type todoListItem struct {
	isDone bool
	title string
	date string
}

type todoEntry struct {
	todoListItem
	description string
}

func (a *app) readList() []todoListItem {
	return a.db.readList()
}

func (a *app) readEntry(id int) (*todoEntry,error) {
	entry := a.db.readEntry(id)
	if entry == nil {
		return entry, fmt.Errorf("Entity not found, id %v", id)
	}
	return entry, nil
}

func (a *app) createNewEntry(entry todoEntry) error {
	return a.db.insert(entry)
}

func (a *app) update(entry todoEntry) error {
	return a.db.update(entry)
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	return mux
}

func index(w http.ResponseWriter, req* http.Request) {
	tpl := template.Must(template.ParseFiles("base.html"))

	tpl.Execute(w, "ziomx")
}