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
	
	if res := reader.readOption("company_name"); res != nil {
		reader.csv_data = res
	}
	if res := reader.readOption("city"); res != nil {
		reader.csv_data = res
	}
	if res := reader.readOption("state"); res != nil {
		reader.csv_data = res
	}
	if res := reader.readOption("round"); res != nil {
		reader.csv_data = res
	}

	output := []map[string]string{}
	for i := 0; i < len(reader.csv_data); i++ {
		mapped := make(map[string]string)
		readProperties(mapped, reader.csv_data[i])
		output = append(output, mapped)
	}

	return output
}

func readProperties(mapped map[string]string, csv_data []string) {
	mapped["permalink"] = csv_data[0]
	mapped["company_name"] = csv_data[1]
	mapped["number_employees"] = csv_data[2]
	mapped["category"] = csv_data[3]
	mapped["city"] = csv_data[4]
	mapped["state"] = csv_data[5]
	mapped["funded_date"] = csv_data[6]
	mapped["raised_amount"] = csv_data[7]
	mapped["raised_currency"] = csv_data[8]
	mapped["round"] = csv_data[9]
}

func (reader *Reader) readPropertyInLoop(mapped map[string] string, keyName string, i int) bool {
	id := reader.columnIndexMapping[keyName]
	_, ok := reader.options[keyName]
	if ok == true {
		if reader.csv_data[i][id] == reader.options[keyName] {
			readProperties(mapped, reader.csv_data[i])
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
