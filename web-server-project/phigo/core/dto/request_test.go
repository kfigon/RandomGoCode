package dto

import (
	"fmt"
	"testing"
)

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

func TestInvalidBody(t *testing.T) {
	testCases := []struct {in string} {
		{in: ` `},
		{in: ``},
		{in: `GET asd`},
		{in: `GET / HTTP/1.0`},
	}

	for _,testCase := range testCases {
		t.Run(fmt.Sprintf("%q", testCase.in), func(t *testing.T) {
			if _, err := ParseResponse([]byte(testCase.in)); err == nil {
				t.Error("Error should be present when blank data")
			}
		})
	}
}