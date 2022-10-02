package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSpreadsheet(t *testing.T) {
	input := `A,B,C,D,E
1,2,,4,asd
=A1+B1,,3,=A2+C2,`

	ss := spreadsheetParser{"\n", ","}.parse(input)

	t.Run("sizes", func(t *testing.T) {
		expectedColumnNumber := 5

		assert.Len(t, ss, expectedColumnNumber)
		assert.Len(t, ss["A"], 2)
		assert.Len(t, ss["B"], 1)
		assert.Len(t, ss["C"], 1)
		assert.Len(t, ss["D"], 2)
		assert.Len(t, ss["E"], 1)

		assert.Equal(t, expectedColumnNumber, ss.columns())
	})

	t.Run("elements", func(t *testing.T) {
		assertPresent := func(row string, col int, exp cell) {
			v, ok := ss.read(row, col)
			assert.True(t, ok, fmt.Sprintf("expected el on %v%d but not found", row, col))
			assert.Equal(t, exp, v, fmt.Sprintf("invalid value on %v%d", row, col))
		}

		absent := func(row string, col int) {
			_, ok := ss.read(row, col)
			assert.False(t, ok, fmt.Sprintf("el on %v%d not expected but found", row, col))
		}
		assertPresent("A", 1, numberCell(1))
		assertPresent("B", 1, numberCell(2))
		absent("C", 1)
		assertPresent("D", 1, numberCell(4))
		assertPresent("E", 1, stringCell("asd"))

		assertPresent("A", 2, expressionCell{"=A1+B1"})
		absent("B", 2)
		assertPresent("C", 2, numberCell(3))
		assertPresent("D", 2, expressionCell{"=A2+C2"})
		absent("E", 2)
	})
}

func TestParseCoords(t *testing.T) {
	testCases := []struct {
		in  string
		row string
		col int
	}{
		{"A2", "A", 2},
		{"B32", "B", 32},
		{"AA32", "AA", 32},
		{"ABCDEFG123456", "ABCDEFG", 123456},
	}
	for _, tC := range testCases {
		t.Run(tC.in, func(t *testing.T) {
			id, ok := parseCoords(tC.in)
			require.True(t, ok)
			assert.Equal(t, tC.col, id.col)
			assert.Equal(t, tC.row, id.row)
		})
	}
}

func TestParseCoordsInvalid(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		_, ok := parseCoords("")
		assert.False(t, ok)
	})

	t.Run("no column", func(t *testing.T) {
		_, ok := parseCoords("B")
		assert.False(t, ok)
	})

	t.Run("no row", func(t *testing.T) {
		_, ok := parseCoords("123")
		assert.False(t, ok)
	})

	t.Run("col with characters", func(t *testing.T) {
		_, ok := parseCoords("B12A3")
		assert.False(t, ok)
	})
}
