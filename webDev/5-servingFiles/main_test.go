package main

import(
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
)

func TestHello(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	response, err := http.Get(server.URL+"/")
	if err != nil {
		t.Fatal("Got error during request: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Got error when reading response: ", err)
	}
	strBody := string(body)
	if strBody != "asd" {
		t.Errorf("Invalid body. Wanted %q, got %q", "asd", strBody)
	}
}