package core

import "testing"

func TestHello(t *testing.T)  {
	res := Hello()
	expected := "asd"
	if res != expected {
		t.Errorf("Error, got %q, instead of %q", res, expected)
	}
}

func TestDto(t *testing.T) {
	res := asd()
	exp := "{asd 123}"
	if res != exp {
		t.Errorf("Expected %q, got %q", exp, res)
	}
}