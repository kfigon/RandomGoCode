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
		readProperties(mapped, reader.csv_data, i)
		output = append(output, mapped)
	}

	return output
}

func readProperties(mapped map[string]string, csv_data [][]string, i int) {
	mapped["permalink"] = csv_data[i][0]
	mapped["company_name"] = csv_data[i][1]
	mapped["number_employees"] = csv_data[i][2]
	mapped["category"] = csv_data[i][3]
	mapped["city"] = csv_data[i][4]
	mapped["state"] = csv_data[i][5]
	mapped["funded_date"] = csv_data[i][6]
	mapped["raised_amount"] = csv_data[i][7]
	mapped["raised_currency"] = csv_data[i][8]
	mapped["round"] = csv_data[i][9]
}

func FindBy(options map[string]string) (map[string]string, error) {
	csv_data := readFile("startup_funding.csv")

	for i := 0; i < len(csv_data); i++ {
		var ok bool
		mapped := make(map[string]string)
		
		readProperty := func(keyName string, id int) bool {
			_, ok = options[keyName]
			if ok == true {
				if csv_data[i][id] == options[keyName] {
					readProperties(mapped, csv_data, i)
				} else {
					return true
				}
			}
			return false
		}
		
		if res := readProperty("company_name", 1); res {
			continue
		}
		if res := readProperty("city", 4); res {
			continue
		}
		if res := readProperty("state", 5); res {
			continue
		}
		if res := readProperty("round", 9); res {
			continue
		}
		
		return mapped, nil
	}

	return make(map[string]string), errors.New("Record Not Found")
}
