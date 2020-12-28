package dto

import "testing"

func TestParseGet(t *testing.T) {
	body:=`GET / HTTP/1.1\r\nHost: localhost:8080\r\nUser-Agent: curl/7.68.0\r\nAccept: */*\r\n\r\n\r\n`
	request, err := ParseResponse([]byte(body))
	if err != nil {
		t.Error("got error")
	}
	if request.Body != nil {
		t.Error("body should be null, got:", request.Body)
	}
	if request.Method != "GET" {
		t.Error("Method should be GET, got:", request.Method)
	}
	if request.Url != "/" {
		t.Error("URL should be /, got:", request.Url)
	}
}

func TestEmptyBody(t *testing.T) {
	_, err := ParseResponse([]byte(``))
	if err == nil {
		t.Error("Error should be present")
	}
}

func TestInvalidBody(t *testing.T) {
	_, err := ParseResponse([]byte(` `))
	if err == nil {
		t.Error("Error should be present when blank data")
	}

	_, err = ParseResponse([]byte(``))
	if err == nil {
		t.Error("Error should be present when empty data")
	}
}

func TestInvalidBody2(t *testing.T) {
	_, err := ParseResponse([]byte(`GET asd`))
	if err == nil {
		t.Error("Error should be present")
	}
}
func TestInvalidFirstLine(t *testing.T) {
	body:=`GET / HTTP/1.0`

	_, err := ParseResponse([]byte(body))
	if err == nil {
		t.Error("Error should be present")
	}
}