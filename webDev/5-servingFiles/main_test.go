package main

import(
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
)

func TestHello(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	response, err := http.Get(server.URL+"/")
	if err != nil {
		t.Fatal("Got error during request: ", err)
	}

	strBody := readBodyAsString(t, response)
	if strBody != "Hi!" {
		t.Errorf("Invalid body. Wanted %q, got %q", "Hi!", strBody)
	}
}

func readBodyAsString(t *testing.T, response *http.Response) string{
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Got error when reading response: ", err)
	}
	return string(body)
}

func TestImage(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	response, err := http.Get(server.URL+"/picture")
	if err != nil {
		t.Fatal("Got error during request: ", err)
	}

	header := response.Header.Get("Content-Type")
	if header != "text/html; charset=utf-8" {
		t.Error("Got invalid header: ", header)
	}

	responseString := readBodyAsString(t, response)
	if !strings.Contains(responseString, `<img src="pic.jpg">`) {
		t.Error("not contains img tag: ", responseString)
	}
}