package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSpreadsheet(t *testing.T) {
	input := `A,B,C,D,E
1,2,,4,asd
=A1+B1,,3,=A2+C2,`

	ss := newSpreadsheet(input)

	t.Run("sizes", func(t *testing.T) {
		expectedColumnNumber := 5
		expectedDataRowNumber := 2
		
		assert.Len(t, ss, expectedColumnNumber)
		for _, v := range ss {
			assert.Len(t, v, expectedDataRowNumber)
		}
		assert.Equal(t, expectedColumnNumber, ss.columns())
		assert.Equal(t, expectedDataRowNumber, ss.rows())
	})

	t.Run("elements", func(t *testing.T) {
		f := func(row string, col int, exp cell) {
			v, ok := ss.read(row,col)
			assert.True(t, ok, fmt.Sprintf("expected el on %v%d but not found", row,col))
			assert.Equal(t, exp, v, fmt.Sprintf("invalid value on %v%d", row,col))
		}
		f("A",1, numberCell(1))
		f("B",1, numberCell(2))
		f("C",1, emptyCell{})
		f("D",1, numberCell(4))
		f("E",1, stringCell("asd"))

		f("A",2, expressionCell{"=A1+B1"})
		f("B",2, emptyCell{})
		f("C",2, numberCell(3))
		f("D",2, expressionCell{"=A2+C2"})
		f("E",2, emptyCell{})
	})
}