package patterns

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type generalAlgorithm struct {}

func (g *generalAlgorithm) extractData(s string) (int, error) {
	// long algorithm to be shared
	return strconv.Atoi(s)
}

func (g *generalAlgorithm) handleInput(i int) int {
	// another long alg to be shared
	return i * 2
}

type theDoer struct {
	*generalAlgorithm
	input string
}

// use the algorithm
func (d *theDoer) execute() (int, error) {
	v, err := d.extractData(d.input)
	if err != nil {
		return 0, err
	}
	out := d.handleInput(v)
	return out, nil
}

// another doer that reuses stuff
// type anotherDoer struct {
// 	*generalAlgorithm
// }

func TestComposition(t *testing.T) {
	init := func(in string) *theDoer {
		return &theDoer{&generalAlgorithm{}, in}
	}

	t.Run("Invalid case", func(t *testing.T) {
		d := init("foo")
		_, err := d.execute()

		assert.Error(t, err)
	})

	t.Run("Valid case", func(t *testing.T) {
		d := init("123")
		v, err := d.execute()
		
		assert.NoError(t, err)
		assert.Equal(t, 246, v)
	})
}