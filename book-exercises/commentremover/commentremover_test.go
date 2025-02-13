package commentremover

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Exercise 1-23. Write a program to remove all comments from a C program.
// Don’t forget to handle quoted strings and character constants properly. C comments do not nest.

func TestRemoveComments(t *testing.T) {
	testCases := []struct {
		desc	string
		input 	string
		expected string
	}{
		{
			desc: "no comments",
			input: `#include <stdio.h>

int main() {
    return 0;
}`,
			expected: `#include <stdio.h>

int main() {
    return 0;
}`,
		},
		{
			desc: "inline",
			input: 
`#include <stdio.h>
// foobar
int main() {
	int a = 1 / 3;
	int b = 1 * 3;

	return 0; // return 1; // nested!
} //dont forget me`,
			expected: 
`#include <stdio.h>

int main() {
	int a = 1 / 3;
	int b = 1 * 3;

	return 0; 
} `,
		},
		{
			desc: "terminated comments",
			input: 
`#include <stdio.h>
// foobar
int main() { /* this is fine */
	/* dont forget to 
	do this */
	/* sdfg *
	asdf
	*/

	// oopsie /* asd */
	// oopsie */ asd 
    return 0; // return 1;
}`,
			expected: 
`#include <stdio.h>

int main() { 
	
	

	
	
    return 0; 
}`,
	},
	{
		desc: "comments in string",
		input: 
`#include <stdio.h>
int main() {
	printf("//please dont remove me");
    return 0; // return 1;
}`,
		expected: 
`#include <stdio.h>
int main() {
	printf("//please dont remove me");
    return 0; 
}`,
	},
}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expected, removeComments(tC.input))
		})
	}
}

func removeComments(s string) string {
	out := strings.Builder{}
	const (
		passthrough int = iota
		slashFound
		singleLine
		multiLine
		multilineStarFound
		stringFound
	)

	state := passthrough
	for _, c := range s {
		switch state {
		case passthrough:
			if c == '/' {
				state = slashFound
			} else if c == '"' {
				state = stringFound
				out.WriteRune(c)
			} else {
				out.WriteRune(c)
			}
		case slashFound:
			if c == '/' {
				state = singleLine
			} else if c == '*' {
				state = multiLine
			} else {
				state = passthrough
				out.WriteRune('/')
				out.WriteRune(c)
			}
		case singleLine:
			if c == '\n' {
				state = passthrough
				out.WriteRune(c)
			}
		case multiLine:
			if c == '*' {
				state = multilineStarFound
			}
		case multilineStarFound:
			if c == '/' {
				state = passthrough
			} else {
				state = multiLine
			}
		case stringFound:
			if c == '"' {
				state = passthrough
				out.WriteRune(c)
			} else {
				out.WriteRune(c)
			}
		}
		// * " / \n default
	}
	return out.String()
}