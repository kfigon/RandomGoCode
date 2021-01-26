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

func TestImageHtml(t *testing.T) {
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
	if !strings.Contains(responseString, `<img src="/pic.jpg">`) {
		t.Error("not contains img tag: ", responseString)
	}
}

func TestImage(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	response, err := http.Get(server.URL+"/pic.jpg")
	if err != nil {
		t.Fatal("Got error during request: ", err)
	}

	header := response.Header.Get("Content-Type")
	if header != "image/jpeg" {
		t.Error("Got invalid header: ", header)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("Got error during file read:", err)
	}
	if len(data) < 10000 {
		t.Error("Data not found len: ", len(data))
	}
}

func TestDownloadFile(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	response, err := http.Get(server.URL+"/downloadFile")
	if err != nil {
		t.Fatal("Got error during request: ", err)
	}

	header := response.Header.Get("Content-Type")
	if header != "application/octet-stream" {
		t.Fatal("Got invalid header: ", header)
	}
	data := readBodyAsString(t, response)
	if !strings.HasPrefix(data, "asdbar") {
		t.Error("Invalid beginning of data: ", data[:100])
	}
}

func TestNotFound(t *testing.T) {
	server := httptest.NewServer(createMux())
	defer server.Close()

	response, err := http.Get(server.URL+"/dupa")
	if err != nil {
		t.Fatal("Got error during request: ", err)
	}
	if response.StatusCode != http.StatusNotFound {
		t.Fatal("Invalid status, expected 404, got:", response.StatusCode)
	}
}