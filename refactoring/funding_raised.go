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
	csvData [][]string
	columnIndexMapping map[string]int
}

func newReader(opts map[string]string) *Reader {
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
		csvData: readFile("startup_funding.csv"),
		columnIndexMapping: colIdxMap,
	}
}

func (reader *Reader) readOptionAndReplaceCsvData(columnName string) {
	value, ok := reader.options[columnName]
	if ok != true {
		return
	}

	id := reader.columnIndexMapping[columnName]
	results := [][]string{}
	for i := 0; i < len(reader.csvData); i++ {
		if reader.csvData[i][id] == value {
			results = append(results, reader.csvData[i])
		}
	}
	reader.csvData = results
}

func Where(options map[string]string) []map[string]string {
	reader := newReader(options)
	
	allowedColumns := []string{"company_name", "city", "state", "round"}
	for _, col := range allowedColumns {
		reader.readOptionAndReplaceCsvData(col)
	}

	output := []map[string]string{}
	for i := 0; i < len(reader.csvData); i++ {
		mapped := make(map[string]string)
		reader.readProperties(mapped, i)
		output = append(output, mapped)
	}

	return output
}

func (reader *Reader) readProperties(mapped map[string]string, i int) {
	writeToMap := func(colName string) {
		id := reader.columnIndexMapping[colName]
		mapped[colName] = reader.csvData[i][id]
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
func (reader *Reader) readSingleProperty(mapped map[string] string, keyName string, i int) bool {
	value, ok := reader.options[keyName]
	if ok != true {
		return false
	}
	id := reader.columnIndexMapping[keyName]
	if reader.csvData[i][id] != value {
		return true
	} 
	reader.readProperties(mapped, i)
	return false
}

func FindBy(options map[string]string) (map[string]string, error) {
	reader := newReader(options)

	for i := 0; i < len(reader.csvData); i++ {
		mapped := make(map[string]string)
		
		if !reader.readSingleProperty(mapped, "company_name", i) &&
			!reader.readSingleProperty(mapped, "city", i) &&
			!reader.readSingleProperty(mapped, "state", i) && 
			!reader.readSingleProperty(mapped, "round", i) {
			return mapped, nil
		}
	}

	return make(map[string]string), errors.New("Record Not Found")
}
