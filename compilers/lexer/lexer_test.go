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

// accept only digit "1"
// 2 states - init, end
func simpleState(input string) func() (string, bool) {
	state := "init"
	indx := 0

	return func() (string, bool) {
		switch state {
		case "init":
			char := input[indx]
			if char == '1' {
				state = "accept"
				indx++
				return state, true
			} 
			// stuck, reject
			state = "reject"
			return state, false

		case "accept":
			if indx >= len(input) {
				return "accept",false
			}
		}
		return "reject",false
	}
}

// go test ./lexer -run TestSimpleAutomaton
func TestSimpleAutomaton(t *testing.T) {
	type stateMachine func()(string,bool)
	drain := func(machine stateMachine) string {
		state, more := machine()
		for more{
			state, more = machine()
		}
		return state
	}

	t.Run("Simple run", func(t *testing.T) {
		machine := simpleState("1")
		state := drain(machine)
		if state != "accept" {
			t.Error("Invalid state, expected accept, got",state)
		}
	})

	t.Run("Simple run - invalid", func(t *testing.T) {
		machine := simpleState("0")
		state := drain(machine)
		if state != "reject" {
			t.Error("Invalid state, expected reject, got",state)
		}
	})
	
	t.Run("Simple run - too long string", func(t *testing.T) {
		machine := simpleState("10")
		state := drain(machine)
		if state != "reject" {
			t.Error("Invalid state, expected reject, got",state)
		}
	})
}

// any number of 1s followed by a single 0
func anyNumberOfOnes(input string) func()(string,bool) {
	state := "A"
	idx := 0
	return func() (string, bool) {
		switch state {
		case "A":
			if idx >= len(input) {
				return "reject", false
			} else if input[idx] == '1' {
				state = "A"
			} else if input[idx] == '0' {
				state = "B"
			}
			idx++
			return state, true
		case "B":
			if idx >= len(input) {
				return "accept", false
			} 
		}
		return "reject", false
	}
}
func TestLessSimpleAutomaton(t *testing.T) {
	type stateMachine func()(string,bool)
	drain := func(machine stateMachine) string {
		state, more := machine()
		for more{
			state, more = machine()
		}
		return state
	}

	t.Run("Simple run, invalid", func(t *testing.T) {
		machine := anyNumberOfOnes("11")
		state := drain(machine)
		if state != "reject" {
			t.Error("Invalid state, expected reject, got",state)
		}
	})

	t.Run("Simple run, no ones, valid", func(t *testing.T) {
		machine := anyNumberOfOnes("0")
		state := drain(machine)
		if state != "accept" {
			t.Error("Invalid state, expected accept, got",state)
		}
	})

	t.Run("Simple run, invalid2", func(t *testing.T) {
		machine := anyNumberOfOnes("100")
		state := drain(machine)
		if state != "reject" {
			t.Error("Invalid state, expected reject, got",state)
		}
	})
	
	t.Run("Simple run - ok", func(t *testing.T) {
		machine := anyNumberOfOnes("110")
		state := drain(machine)
		if state != "accept" {
			t.Error("Invalid state, expected accept, got",state)
		}
	})
}