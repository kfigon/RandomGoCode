package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func main() {
	app := makeApp(makeDb())
	view := view{app}
	log.Fatal(http.ListenAndServe(":8080", createMux(view)))
}

func makeApp(db dataProvider) *app {
	return &app{db}
}

func makeDb() mapDb {
	return mapDb{
		data: []TodoListItem {
			TodoListItem{Title: "first task", IsDone: false, Date: "20-02-2021"},
			TodoListItem{Title: "second task", IsDone: true, Date: "15-01-2020"},
		},
	}
}
type mapDb struct {
	data []TodoListItem
}
func (m mapDb) readList() []TodoListItem{
	return m.data
}
func (m mapDb) readEntry(int) *TodoEntry{
	return nil
}
func (m mapDb) insert(TodoEntry) error{
	return nil
}
func (m mapDb) update(TodoEntry) error{
	return nil
}

type app struct {
	db dataProvider
}

type dataProvider interface {
	readList() []TodoListItem
	readEntry(int) *TodoEntry
	insert(TodoEntry) error
	update(TodoEntry) error
}

type TodoListItem struct {
	IsDone bool
	Title string
	Date string
}

type TodoEntry struct {
	TodoListItem
	Description string
}

func (a *app) readList() []TodoListItem {
	return a.db.readList()
}

func (a *app) readEntry(id int) (*TodoEntry,error) {
	entry := a.db.readEntry(id)
	if entry == nil {
		return entry, fmt.Errorf("Entity not found, id %v", id)
	}
	return entry, nil
}

func (a *app) createNewEntry(entry TodoEntry) error {
	return a.db.insert(entry)
}

func (a *app) update(entry TodoEntry) error {
	return a.db.update(entry)
}


type view struct{
	app *app
}
type basicView interface {
	handleIndex(w http.ResponseWriter, req* http.Request)
	handleList(w http.ResponseWriter, req* http.Request)
}

func createMux(v basicView) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", v.handleIndex)
	mux.HandleFunc("/list", v.handleList)
	return mux
}

func (v view) handleIndex(w http.ResponseWriter, req* http.Request) {
	tpl := template.Must(template.ParseFiles("base.html", "landingPage.html"))
	tpl.Execute(w, "ziomx")
}

func (v view) handleList(w http.ResponseWriter, req* http.Request) {
	tpl := template.Must(template.ParseFiles("base.html", "list.html"))
	tpl.Execute(w, v.app.readList())
}