package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func main() {
	app := makeApp(makeDb())
	view := &view{app}
	log.Fatal(http.ListenAndServe(":8080", createMux(view)))
}

func makeApp(db dataProvider) *app {
	return &app{db}
}

func makeDb() *mapDb {
	return &mapDb{
		data: []TodoListItem {
			TodoListItem{Title: "first task", IsDone: false, Date: "20-02-2021"},
			TodoListItem{Title: "second task", IsDone: true, Date: "15-01-2020"},
		},
	}
}
type mapDb struct {
	data []TodoListItem
}
func (m *mapDb) readList() []TodoListItem{
	return m.data
}
func (m *mapDb) insert(entry TodoListItem) error{
	m.data = append(m.data, entry)
	return nil
}
func (m *mapDb) update(TodoListItem) error{
	return nil
}

type app struct {
	db dataProvider
}

type dataProvider interface {
	readList() []TodoListItem
	insert(TodoListItem) error
	update(TodoListItem) error
}

type TodoListItem struct {
	IsDone bool
	Title string
	Date string
}

func (a *app) readList() []TodoListItem {
	return a.db.readList()
}

func (a *app) createNewEntry(entry TodoListItem) error {
	return a.db.insert(entry)
}

func (a *app) update(entry TodoListItem) error {
	return a.db.update(entry)
}


type view struct{
	app *app
}

func createMux(v *view) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", v.handleIndex)
	mux.HandleFunc("/list", v.handleList)
	mux.HandleFunc("/addNew", v.handleAddNew)
	return mux
}

func (v *view) handleIndex(w http.ResponseWriter, req* http.Request) {
	tpl := template.Must(template.ParseFiles("base.html", "landingPage.html"))
	tpl.Execute(w, "ziomx")
}

func (v *view) handleList(w http.ResponseWriter, req* http.Request) {
	tpl := template.Must(template.ParseFiles("base.html", "list.html"))
	tpl.Execute(w, v.app.readList())
}

func (v *view) handleAddNew(w http.ResponseWriter, req* http.Request) {
	defer http.Redirect(w,req,"/list",http.StatusSeeOther)
	
	if req.Method != "POST" {
		return
	}

	newTodoEntry, err := v.parseForm(req)
	if err != nil {
		return
	}

	v.app.createNewEntry(newTodoEntry)
}

func (v *view) parseForm(req *http.Request) (TodoListItem, error) {
	err := req.ParseForm()
	if err != nil {
		return TodoListItem{}, err
	}

	isDone := false
	if value := req.FormValue("isDone"); value == "on" {
		isDone = true
	}
	newTodoEntry := TodoListItem{
		Title: req.FormValue("title"),
		Date: req.FormValue("date"),
		IsDone: isDone,
	}
	if newTodoEntry.Title == "" || newTodoEntry.Date == "" {
		return TodoListItem{}, fmt.Errorf("Invalid data provided")
	}
	return newTodoEntry, nil
}