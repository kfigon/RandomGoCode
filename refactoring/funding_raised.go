package fundingraised

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
)
func readFile(path string)[][]string {
	f, _ := os.Open(path)
	defer f.Close()
	reader := csv.NewReader(bufio.NewReader(f))
	csvData := [][]string{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		csvData = append(csvData, row)
	}
	return csvData
}

type Reader struct {
	options map[string]string
	columnIndexMapping map[string]int
}

func newReader(opts map[string]string) (*Reader, [][]string) {
	colIdxMap := make(map[string]int)
	colIdxMap["permalink"] = 0
	colIdxMap["company_name"] = 1
	colIdxMap["number_employees"] = 2
	colIdxMap["category"] = 3
	colIdxMap["city"] = 4
	colIdxMap["state"] = 5
	colIdxMap["funded_date"] = 6
	colIdxMap["raised_amount"] = 7
	colIdxMap["raised_currency"] = 8
	colIdxMap["round"] = 9

	return &Reader {
		options: opts,
		columnIndexMapping: colIdxMap,
	}, readFile("startup_funding.csv")
}

func (reader *Reader) findRows(key string, rows [][]string) [][]string {
	value, ok := reader.options[key]
	if ok != true {
		return nil
	}

	results := [][]string{}
	id := reader.columnIndexMapping[key]
	for _, row := range rows {
		if row[id] == value {
			results = append(results, row)
		}
	}
	return results
}

func Where(options map[string]string) []map[string]string {
	reader, rows := newReader(options)
	
	allowedColumns := []string{"company_name", "city", "state", "round"}
	for _, col := range allowedColumns {
		if res := reader.findRows(col, rows); res != nil {
			rows = res
		}
	}

	output := []map[string]string{}
	for _, row := range rows {
		mapped := make(map[string]string)
		reader.readProperties(mapped, row)
		output = append(output, mapped)
	}

	return output
}

func (reader *Reader) readProperties(mapped map[string]string, row []string) {
	writeToMap := func(colName string) {
		id := reader.columnIndexMapping[colName]
		mapped[colName] = row[id]
	}
	writeToMap("permalink")
	writeToMap("company_name")
	writeToMap("number_employees")
	writeToMap("category")
	writeToMap("city")
	writeToMap("state")
	writeToMap("funded_date")
	writeToMap("raised_amount")
	writeToMap("raised_currency")
	writeToMap("round")
}

// returns - should skip to next iteration
func (reader *Reader) readSingleProperty(mapped map[string] string, keyName string, row []string) bool {
	value, ok := reader.options[keyName]
	if ok != true {
		return false
	}
	id := reader.columnIndexMapping[keyName]
	if row[id] != value {
		return true
	} 
	reader.readProperties(mapped, row)
	return false
}

func FindBy(options map[string]string) (map[string]string, error) {
	reader, rows := newReader(options)

	for _,row := range rows {
		mapped := make(map[string]string)
		
		if !reader.readSingleProperty(mapped, "company_name", row) &&
			!reader.readSingleProperty(mapped, "city", row) &&
			!reader.readSingleProperty(mapped, "state", row) && 
			!reader.readSingleProperty(mapped, "round", row) {
			return mapped, nil
		}
	}

	return make(map[string]string), errors.New("Record Not Found")
}
