package core

import "testing"

func TestHello(t *testing.T)  {
	res := Hello()
	expected := "asd"
	if res != expected {
		t.Errorf("Error, got %q, instead of %q", res, expected)
	}
}