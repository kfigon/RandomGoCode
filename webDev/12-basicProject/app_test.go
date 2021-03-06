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
	insertFun func(entry TodoListItem) error
	updateFun func() error
}
func (m mockDb) readList() []TodoListItem {
	if m.readListFun == nil {
		return []TodoListItem{}
	}
	return m.readListFun()
}
func (m mockDb) insert(entry TodoListItem) error {
	if m.insertFun == nil {
		return nil
	}
	return m.insertFun(entry)
}
func (m mockDb) update(entry TodoListItem) error {
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

func TestCreateNewEntry(t *testing.T) {
	// given
	app := makeApp(mockDb{})
	// when
	err := app.createNewEntry(TodoListItem{})
	// then
	if err != nil {
		t.Error("Error during creation not expected, got", err)
	}
}

func TestCreateNewEntryWhenNotSucceed(t *testing.T) {
	// given
	mock := mockDb{insertFun: func(TodoListItem) error {return fmt.Errorf("got error")}}
	app := makeApp(mock)
	// when
	err := app.createNewEntry(TodoListItem{})
	// then
	if err == nil {
		t.Error("Expected error during insert, not received")
	}
}

func TestUpdateEntry(t *testing.T) {
	// given
	app := makeApp(mockDb{})
	// when
	err := app.update(TodoListItem{})
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
	err := app.update(TodoListItem{})
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
	if !strings.Contains(responseBody, "hi there ziomx") {
		t.Error("Invalid response")
	}
}

func TestRoutingGetEmptyList(t *testing.T) {
	srv := createServer(&mockDb{})
	defer srv.Close()

	resp, _ := http.Get(srv.URL +"/list")
	assertStatus(t, resp.StatusCode, http.StatusOK)

	responseBody := getStringBody(t, resp)
	if !strings.Contains(responseBody, "To do list") {
		t.Error("Invalid response")
	}
}

func TestRoutingGetList(t *testing.T) {
	readMock := func() []TodoListItem {
		return []TodoListItem{
			TodoListItem{Title: "do testing", IsDone: true, Date: "2021-02-20"},
		}
	}
	srv := createServer(&mockDb{readListFun: readMock})
	defer srv.Close()

	resp, _ := http.Get(srv.URL +"/list")
	assertStatus(t, resp.StatusCode, http.StatusOK)

	responseBody := getStringBody(t, resp)
	if !strings.Contains(responseBody, "To do list") {
		t.Error("Invalid response")
	}
	if !strings.Contains(responseBody, "do testing") {
		t.Error("Invalid response")
	}
	if !strings.Contains(responseBody, "true") {
		t.Error("Invalid response")
	}
	if !strings.Contains(responseBody, "2021-02-20") {
		t.Error("Invalid response")
	}
}

func TestRoutingPostToAddNew(t *testing.T) {
	assertOnInsertFun := func(newEntry TodoListItem) error {
		if newEntry.Title != "My new title" {
			t.Error("Invalid entity saved", newEntry)
		}
		return nil
	}

	db := &mockDb{
		insertFun: assertOnInsertFun,
	}
	srv := createServer(db)
	defer srv.Close()

	data := map[string][]string {
		"title": []string{"My new title"},
		"date": []string{"2021-02-20"},
		"isDone": []string{"on"},
	}
	resp, _ := http.PostForm(srv.URL +"/addNew", data)
	assertStatus(t, resp.StatusCode, http.StatusOK)

	responseBody := getStringBody(t, resp)
	if !strings.Contains(responseBody, "To do list") {
		t.Error("Invalid response")
	}
}

func TestParseAddNewForm(t *testing.T) {
	v := view{}
	
	body := strings.NewReader("title=My new Title&date=2021-02-20&isDone=on")
	request := httptest.NewRequest("POST", "/", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	parsed, err := v.parseForm(request)
	if err != nil {
		t.Error("Got error during parsing", err)
	}
	if parsed.Title != "My new Title"{
		t.Error("Invalid title parsed", parsed.Title)
	}
	if parsed.Date != "2021-02-20" {
		t.Error("Invalid date parsed", parsed.Date)
	}
	if parsed.IsDone != true {
		t.Error("Invalid isDone parsed", parsed.IsDone)
	}
}