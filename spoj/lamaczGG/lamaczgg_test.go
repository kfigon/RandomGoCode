package main

import (
	"strings"
	"testing"
)

func TestLamaczGG(t *testing.T) {
	testCases := []struct {
		input	string
		expected string
	}{
		{"BGCGDGEGFGGGHGIGJGKG", "abcdefghij"},
		{"LGBGEHBGDHEHCHPGGGBG", "katastrofa"},
		{"PGCGPGKHPGHHJGDHLGPG", "obozowisko"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			got := decode(tC.input)

			if tC.expected != got {
				t.Errorf("Got %q, exp %q", got, tC.expected)
			}
		})
	}
}

func TestMapChar(t *testing.T) {
	testCases := []struct {
		input	string
		expected string
	}{
		{"BG", "a"},
		{"CG", "b"},
		{"DG", "c"},
		{"KH", "z"},
		{"PG", "o"},
		{"HH", "w"},
		{"BH", "q"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			got := mapChar(tC.input)

			if tC.expected != got {
				t.Errorf("Got %q, exp %q", got, tC.expected)
			}
		})
	}
}

func decode(input string) string {
	out := ""
	for i := 0; i < len(input); i+=2 {
		out += mapChar(input[i:i+2])
	}
	return out
}

func mapChar(input string) string {
	if len(input) != 2 {
		return ""
	}

	first,second := input[0],input[1]

	modifier := second-byte('G')
	charCode := first - 1 + modifier*16
	return strings.ToLower(string(charCode))
}