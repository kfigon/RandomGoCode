package main

import (
	"strconv"
	"strings"
	"unicode"
)

// lack of sum types...
type cell interface{
	dummy()
	String() string
}

type stringCell string
func(_ stringCell) dummy(){}
func(s stringCell) String() string {
	return string(s)
}

type numberCell int
func(_ numberCell) dummy(){}
func(n numberCell) String() string{
	return strconv.Itoa(int(n))
}

type expressionCell struct{
	exp string
}
func(_ expressionCell) dummy(){}
func(e expressionCell) String() string{
	return e.exp
}

const lineSeparator = "\n"
const columnSeparator = ","

type spreadsheet map[string]map[int]cell
func newSpreadsheet(input string) spreadsheet {
	out := map[string]map[int]cell{}
	lines := strings.Split(input, lineSeparator)

	headers := map[int]string{}
	headerLine := strings.Split(lines[0], columnSeparator)
	for i, h := range  headerLine {
		headers[i] = h
		out[h]=map[int]cell{}
	}
	
	for i, line := range lines[1:] {
		columns := strings.Split(line, columnSeparator)
		for c, col := range columns {
			x := out[headers[c]]

			if intV, err := strconv.Atoi(col); err == nil {
				x[i] = numberCell(intV)
			} else if strings.HasPrefix(col, "=") {
				x[i] = expressionCell{col}
			} else if col != "" {
				x[i] = stringCell(col)
			}
			out[headers[c]] = x
		}
	}

	return out
}

func (s spreadsheet) read(row string, col int) (cell, bool) {
	r, ok := s[row]
	if !ok {
		return nil, false
	}
	v, ok := r[col-1]
	return v, ok
}

func (s spreadsheet) columns() int {
	return len(s)
}

type coordinate struct{
	row string
	col int
}

func parseId(in string) (coordinate, bool) {
	if len(in) < 2 {
		return coordinate{}, false
	}

	splittingPoint := 0
	for i,c := range in {
		if unicode.IsDigit(c) {
			splittingPoint = i
			break
		}
	}

	str := in[:splittingPoint]
	if len(str) == 0 {
		return coordinate{}, false
	}
	
	numStr := in[splittingPoint:]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return coordinate{}, false
	}
	return coordinate{row: str, col: num}, true
}