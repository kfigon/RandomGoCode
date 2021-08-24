package lexer

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	input := `if (i==j) els = 0;
	else els=1;`

	got := Tokenize(input)

	exp := []Token {
		{Keyword, "if"},
		{Whitespace, " "},
		{OpenParam, "("},
		{Identifier, "i"},
		{Operator, "=="},
		{Identifier, "j"},
		{Whitespace, " "},
		{Identifier, "els"},
		{Whitespace, " "},
		{Assignment, "="},
		{Whitespace, " "},
		{Number, "0"},
		{Semicolon, ";"},
		{Whitespace, "\n\t"},
		{Keyword, "else"},
		{Whitespace, " "},
		{Identifier, "els"},
		{Assignment, "="},
		{Number, "1"},
		{Semicolon, ";"},
	}

	assertTokens(t, exp, got)
}

func assertTokens(t *testing.T, exp, got []Token) {
	if len(exp) != len(got) {
		t.Fatalf("Invalid array lengths, exp %v, got %v", len(exp), len(got))
	}

	for i := 0; i < len(got); i++ {
		g := got[i]
		e := exp[i]

		if g != e {
			t.Errorf("Invalid element on position %v, exp %v, got %v", i, e, g)
		}
	}
}

type mach string 
func (m mach) firstChar() rune {
	return rune(m[0])
}

func (m mach) find(input string, startingIdx int) (int, bool) {
	idx := startingIdx
	tokenIdx := 0
	for idx < len(input) && tokenIdx < len(m) {
		char := input[idx]
		toFind := m[tokenIdx]

		if char != toFind {
			return 0, false
		}

		idx++
		tokenIdx++
	}

	// found it
	if tokenIdx != len(m) {
		if idx == len(input) { // at the end
			return idx+1, true
		} else if input[idx] == ' ' {
			return idx, true
		}
	}
	
	return 0, false
}


func TestStateMachine(t *testing.T) {
	testCases := []struct {
		desc	string
		startingIdx int
		expectedIdx int
		inputStr string
	}{
		{
			desc: "Starting from beginning",
			startingIdx: 0,
			expectedIdx: 6,
			inputStr: `abc if abb`,
		},
		{
			desc: "Starting from near",
			startingIdx: 3,
			expectedIdx: 6,
			inputStr: `abc if abb`,
		},
		{
			desc: "Starting from same indx",
			startingIdx: 4,
			expectedIdx: 6,
			inputStr: `abc if abb`,
		},
		{
			desc: "Not found",
			startingIdx: 6,
			expectedIdx: -1,
			inputStr: `abc if abb`,
		},
		{
			desc: "Not found at beginning",
			startingIdx: 0,
			expectedIdx: -1,
			inputStr: `ifx abb`,
		},
		{
			desc: "Not found at beginning2",
			startingIdx: 1,
			expectedIdx: -1,
			inputStr: `ifx abb`,
		},
		{
			desc: "Starting exceeds len",
			startingIdx: 88,
			expectedIdx: -1,
			inputStr: `abc if abb`,
		},
		{
			desc: "Token at the beginning",
			startingIdx: 0,
			expectedIdx: 2,
			inputStr: `if abb`,
		},
		{
			desc: "Token at the beginning, start exceeds",
			startingIdx: 3,
			expectedIdx: -1,
			inputStr: `if abb`,
		},
		{
			desc: "Token at the end",
			startingIdx: 4,
			expectedIdx: 6,
			inputStr: `abb if`,
		},
		{
			desc: "Token not found at the end",
			startingIdx: 4,
			expectedIdx: -1,
			inputStr: `abb ifx`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			machine := mach("if")
			gotIdx, found := machine.find(tC.inputStr, tC.startingIdx)

			expectedNotFound := tC.expectedIdx == -1
			if expectedNotFound && found {
				t.Fatal("Not expected to find the element, but got it", gotIdx)
			}
			
			if !expectedNotFound && (!found || tC.expectedIdx != gotIdx) {
				t.Fatal("Expected idx", tC.expectedIdx, "got", gotIdx, "found:",found)
			}
		})
	}
}