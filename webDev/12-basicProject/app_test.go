package main

import (
	"testing"
	// "fmt"
)
type mockDb struct {
	list []todoListItem
	entry *todoEntry
}
func (m mockDb) readList() []todoListItem {
	return m.list
}
func (m mockDb) readEntry(id int) *todoEntry {
	return m.entry
}

func createMock(list []todoListItem, entry *todoEntry) mockDb {
	return mockDb{
		list:list,
		entry: entry,
	}
}

func TestReadTodoListWhenEmpty(t *testing.T) {
	// given
	app := makeApp(createMock([]todoListItem{}, nil))
	// when
	todos := app.readList()
	// then
	if size := len(todos); size != 0 {
		t.Error("Empty todo list expected, got", size)
	}
}

func TestReadTodoList(t *testing.T) {
	// given
	mockTodos := []todoListItem { 
		todoListItem{title:"first task"},
		todoListItem{title:"second task"},
	}

	app := makeApp(createMock(mockTodos, nil))
	// when
	todos := app.readList()
	// then
	if size := len(todos); size != 2 {
		t.Error("Empty todo list expected, got", size)
	}
	if title := todos[0].title; title != "first task" {
		t.Error("Invalid title fetched, got:", title)
	}
	if title := todos[1].title; title != "second task" {
		t.Error("Invalid title fetched, got:", title)
	}
}

func TestReadSingleTodo(t *testing.T) {
	// given
	singleEntry := todoEntry{ todoListItem: todoListItem{title:"first task"}, }

	app := makeApp(createMock(nil, &singleEntry))
	// when
	todo,err := app.readEntry(456)
	// then
	if err != nil {
		t.Error("Error not expected, got", err)
	}
	if todo == nil {
		t.Error("Entry not found")
	}
	if gotTitle := todo.title; gotTitle != "first task" {
		t.Error("Invalid title read, got", gotTitle)
	}
}

func TestReadTodoWhenNotFound(t *testing.T) {
	// given
	app := makeApp(createMock(nil, nil))
	// when
	todo,err := app.readEntry(456)
	// then
	if err == nil {
		t.Error("Error expected, not received")
	}
	if todo != nil {
		t.Error("Entry found, but shouldn't. Got:", todo)
	}
}