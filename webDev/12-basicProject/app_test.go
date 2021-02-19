package main

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
)

type mockDb struct {
	readListFun func() []todoListItem
	readEntryFun func() *todoEntry
	insertFun func() error
	updateFun func() error
}
func (m mockDb) readList() []todoListItem {
	if m.readListFun == nil {
		return []todoListItem{}
	}
	return m.readListFun()
}
func (m mockDb) readEntry(id int) *todoEntry {
	if m.readEntryFun == nil {
		return nil
	}
	return m.readEntryFun()
}
func (m mockDb) insert(entry todoEntry) error {
	if m.insertFun == nil {
		return nil
	}
	return m.insertFun()
}
func (m mockDb) update(entry todoEntry) error {
	if m.updateFun == nil {
		return nil
	}
	return m.updateFun()
}


func TestReadTodoListWhenEmpty(t *testing.T) {
	// given
	app := makeApp(mockDb{})
	// when
	todos := app.readList()
	// then
	if size := len(todos); size != 0 {
		t.Error("Empty todo list expected, got", size)
	}
}

func TestReadTodoList(t *testing.T) {
	// given
	mock := mockDb{readListFun: func() []todoListItem { return []todoListItem { 
		todoListItem{title:"first task"},
		todoListItem{title:"second task"},
	}}}
	app := makeApp(mock)
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
	mockEntry := todoEntry{ todoListItem: todoListItem{title:"first task"}}
	mock := mockDb{readEntryFun: func() *todoEntry { return &mockEntry}}
	app := makeApp(mock)
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
	app := makeApp(mockDb{})
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

func TestCreateNewEntry(t *testing.T) {
	// given
	app := makeApp(mockDb{})
	// when
	todoEntry := todoEntry{}
	err := app.createNewEntry(todoEntry)
	// then
	if err != nil {
		t.Error("Error during creation not expected, got", err)
	}
}

func TestCreateNewEntryWhenNotSucceed(t *testing.T) {
	// given
	mock := mockDb{insertFun: func() error {return fmt.Errorf("got error")}}
	app := makeApp(mock)
	// when
	todoEntry := todoEntry{}
	err := app.createNewEntry(todoEntry)
	// then
	if err == nil {
		t.Error("Expected error during insert, not received")
	}
}

func TestUpdateEntry(t *testing.T) {
	// given
	app := makeApp(mockDb{})
	// when
	todoEntry := todoEntry{}
	err := app.update(todoEntry)
	// then
	if err != nil {
		t.Error("Error not expected, got:",err)
	}
}

func TestUpdateEntryButFailed(t *testing.T) {
	// given
	mock := mockDb{updateFun: func() error {return fmt.Errorf("error occured")}}
	app := makeApp(mock)
	// when
	todoEntry := todoEntry{}
	err := app.update(todoEntry)
	// then
	if err == nil {
		t.Error("Expected error during insert, not received")
	}
}

func createServer() *httptest.Server {
	return httptest.NewServer(createMux())
}

func TestBasicWeb(t *testing.T) {
	srv := createServer()
	defer srv.Close()

	resp, _ := http.Get(srv.URL +"/")
	expStatus := http.StatusOK
	if gotStatus := resp.StatusCode; gotStatus != expStatus {
		t.Errorf("Wrong status, got %v, exp %v", gotStatus, expStatus)
	} 
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error during reading response:", err)
	}
	if responseBody := string(data); !strings.Contains(responseBody, "hi there") {
		t.Errorf("Invalid response, got: %v", responseBody)
	}
}