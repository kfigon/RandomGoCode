package main

import (
	"fmt"
)

func main() {

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