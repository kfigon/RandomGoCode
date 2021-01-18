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
	csv_data [][]string
	columnIndexMapping map[string]int
}

func newReader(path string, opts map[string]string) *Reader {
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

	return &Reader{
		options: opts,
		csv_data: readFile(path),
		columnIndexMapping: colIdxMap,
	}
}

func (reader *Reader) readOption(columnName string) [][]string {
	_, ok := reader.options[columnName]
	if ok == true {
		id := reader.columnIndexMapping[columnName]
		results := [][]string{}
		for i := 0; i < len(reader.csv_data); i++ {
			if reader.csv_data[i][id] == reader.options[columnName] {
				results = append(results, reader.csv_data[i])
			}
		}
		return results
	}
	return nil
}

func Where(options map[string]string) []map[string]string {
	reader := newReader("startup_funding.csv", options)
	
	allowedColumns := []string{"company_name", "city", "state", "round"}
	for _, col := range allowedColumns {
		if res := reader.readOption(col); res != nil {
			reader.csv_data = res
		}
	}

	output := []map[string]string{}
	for i := 0; i < len(reader.csv_data); i++ {
		mapped := make(map[string]string)
		reader.readProperties(mapped, i)
		output = append(output, mapped)
	}

	return output
}

func (reader *Reader) readProperties(mapped map[string]string, i int) {
	writeToMap := func(colName string) {
		id := reader.columnIndexMapping[colName]
		mapped[colName] = reader.csv_data[i][id]
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

func (reader *Reader) readPropertyInLoop(mapped map[string] string, keyName string, i int) bool {
	id := reader.columnIndexMapping[keyName]
	_, ok := reader.options[keyName]
	if ok == true {
		if reader.csv_data[i][id] == reader.options[keyName] {
			reader.readProperties(mapped, i)
		} else {
			return true
		}
	}
	return false
}

func FindBy(options map[string]string) (map[string]string, error) {
	reader := newReader("startup_funding.csv", options)

	for i := 0; i < len(reader.csv_data); i++ {
		mapped := make(map[string]string)
		
		if reader.readPropertyInLoop(mapped, "company_name", i) || 
			reader.readPropertyInLoop(mapped, "city", i) ||
			reader.readPropertyInLoop(mapped, "state", i) || 
			reader.readPropertyInLoop(mapped, "round", i) {
			continue
		}

		return mapped, nil
	}

	return make(map[string]string), errors.New("Record Not Found")
}
