package dto

import (
	"fmt"
	"errors"
	"strings"
)

type HttpEndpointId struct {
	Method string
	Url string
}

type HttpRequest struct {
	HttpEndpointId
	Body []byte
}

func ParseResponse(body []byte) (HttpRequest, error) {
	req := HttpRequest{}
	if len(body) == 0 {
		return req, errors.New("Empty body provided")
	}
	stringBasedBody := string(body)
	lines := strings.Split(stringBasedBody, `\r\n`)
	if len(lines) == 0 {
		return req, errors.New("no newline found in request")
	}
	firstLine := lines[0]
	parsedFirstLine := strings.Fields(firstLine)
	if len(parsedFirstLine) < 3 {
		return req, errors.New(fmt.Sprint("Invalid first line, got:", firstLine))
	}

	if methodAllowed(parsedFirstLine[0]) == false {
		return req, errors.New(fmt.Sprint("Not allowed method:", parsedFirstLine[0]))
	}

	if validUrl(parsedFirstLine[1]) == false {
		return req, errors.New(fmt.Sprint("Invalid url:", parsedFirstLine[1]))
	}

	if parsedFirstLine[2] != "HTTP/1.1" {
		return req, errors.New(fmt.Sprint("Invalid protocol version, got:", parsedFirstLine[2]))
	}

	req.Method = parsedFirstLine[0]
	req.Url = parsedFirstLine[1]
	
	return req, nil
}

func methodAllowed(method string) bool {
	allowed := [...]string{"GET", "POST"}
	for _,al := range(allowed) {
		if al == method {
			return true
		}
	}

	return false
}

func validUrl(url string) bool {
	if len(url) <=0 {
		return false
	}
	return url[0] == '/'
}