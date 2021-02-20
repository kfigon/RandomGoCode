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
	readListFun func() []TodoListItem
	readEntryFun func() *TodoEntry
	insertFun func() error
	updateFun func() error
}
func (m mockDb) readList() []TodoListItem {
	if m.readListFun == nil {
		return []TodoListItem{}
	}
	return m.readListFun()
}
func (m mockDb) readEntry(id int) *TodoEntry {
	if m.readEntryFun == nil {
		return nil
	}
	return m.readEntryFun()
}
func (m mockDb) insert(entry TodoEntry) error {
	if m.insertFun == nil {
		return nil
	}
	return m.insertFun()
}
func (m mockDb) update(entry TodoEntry) error {
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
	mock := mockDb{readListFun: func() []TodoListItem { return []TodoListItem { 
		TodoListItem{Title:"first task"},
		TodoListItem{Title:"second task"},
	}}}
	app := makeApp(mock)
	// when
	todos := app.readList()
	// then
	if size := len(todos); size != 2 {
		t.Error("Empty todo list expected, got", size)
	}
	if title := todos[0].Title; title != "first task" {
		t.Error("Invalid title fetched, got:", title)
	}
	if title := todos[1].Title; title != "second task" {
		t.Error("Invalid title fetched, got:", title)
	}
}

func TestReadSingleTodo(t *testing.T) {
	// given
	mockEntry := TodoEntry{ TodoListItem: TodoListItem{Title:"first task"}}
	mock := mockDb{readEntryFun: func() *TodoEntry { return &mockEntry}}
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
	if gotTitle := todo.Title; gotTitle != "first task" {
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
	todoEntry := TodoEntry{}
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
	todoEntry := TodoEntry{}
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
	todoEntry := TodoEntry{}
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
	todoEntry := TodoEntry{}
	err := app.update(todoEntry)
	// then
	if err == nil {
		t.Error("Expected error during insert, not received")
	}
}

func createServer(db *mockDb) *httptest.Server {
	application := makeApp(db)
	return httptest.NewServer(createMux(&view{app:application}))
}

func assertStatus(t *testing.T, got int, exp int) {
	if got != exp {
		t.Errorf("Wrong status, got %v, exp %v", got, exp)
	} 
}

func getStringBody(t *testing.T, resp *http.Response) string {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error during reading response:", err)
	}
	return string(data)
}

func TestBasicWebRouting(t *testing.T) {
	srv := createServer(&mockDb{})
	defer srv.Close()

	resp, _ := http.Get(srv.URL +"/")
	assertStatus(t, resp.StatusCode, http.StatusOK)
	responseBody := getStringBody(t, resp)
	if !strings.Contains(responseBody, "hi there") {
		t.Error("Invalid response")
	}
}

func TestRoutingGetList(t *testing.T) {
	srv := createServer(&mockDb{})
	defer srv.Close()

	resp, _ := http.Get(srv.URL +"/list")
	assertStatus(t, resp.StatusCode, http.StatusOK)

	responseBody := getStringBody(t, resp)
	if !strings.Contains(responseBody, "To do list") {
		t.Error("Invalid response")
	}
}