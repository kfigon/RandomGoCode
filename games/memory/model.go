package memory

import (
	"math/rand"
	"time"
)

type field struct {
	revealed bool
	val int
}

type game struct {
	fields []field
}

const numberOfFields = 16


func generatePairsAndShuffle(numOfGameFields int) []int {
	rand.Seed(time.Now().UnixNano())
	vals := make([]int,numOfGameFields)
	counter := 0

	for i := 0; i < len(vals); i+=2 {
		vals[i] = counter
		vals[i+1]= counter
		counter++
	}

	// shuffle
	for i := 0; i < len(vals); i++ {
		src := rand.Intn(len(vals))
		dst := rand.Intn(len(vals))
		
		tmp := vals[src]
		vals[src] = vals[dst]
		vals[dst] = tmp
	}

	return vals
}

func newField() *game {
	fields := make([]field, numberOfFields)
	fieldValues := generatePairsAndShuffle(numberOfFields)

	for i := 0; i < len(fields); i++ {
		fields[i] = field {revealed: false, val: fieldValues[i]}
	}

	return &game{
		fields: fields,
	}
}

func (g *game) size() int {
	return len(g.fields)
}

func (g* game) isRevealed(idx int) bool {
	return g.fields[idx].revealed
}