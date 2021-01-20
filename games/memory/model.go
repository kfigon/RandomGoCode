package memory

import (
	"math/rand"
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
	
}

func newField() *game {
	fields := make([]field, numberOfFields)
	fieldValues := generatePairsAndShuffle(numberOfFields)

	for i := 0; i < len(fields); i++ {
		fields[i] = field{revealed: false, val: fieldValues[i]}
	}

	return &game{
		fields: fields
	}
}

func (g *game) size() int {
	return len(g.fields)
}

func (g* game) isRevealed(idx int) bool {
	return g.fields[idx].revealed
}