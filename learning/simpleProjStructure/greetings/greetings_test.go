package greetings

import "testing"

func TestHelloGut(t *testing.T) {
	v, err := Hello("Adam")
	if err != nil {
		t.Error("error should be nil in simple case")
	}
	if v != "Hello from Adam" {
		t.Error("Wrong response", v)
	}
}

func TestHelloEmpty(t *testing.T)  {
	v,err := Hello("")
	if v != "" || err == nil {
		t.Error("on error we should get empty response")
	}
}