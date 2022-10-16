package bagqueuestack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// usage of stacks - Dijkstra's two-stack algorithm for expression eval

func TestEvaluator(t *testing.T) {
	testCases := []struct {
		input	string
		exp 	int		
	}{
		{
			input: "(1+((2+3)*(4*5)))",
			exp: 101,
			
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			got := evaluateExpression(tC.input)
			assert.Equal(t, tC.exp, got)
		})
	}
}

func evaluateExpression(input string) int {
	// operatorStack := &stack{}
	// operandsStack := &stack{}

	// for _, char := range input {
	// 	switch char {
	// 	case '+', '*','-': operatorStack.push(int(char))
	// 	}
	// }

	return -1
}