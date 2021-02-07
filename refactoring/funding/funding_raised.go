package fundingraised

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
)

func readFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}
	}
	reader := csv.NewReader(bufio.NewReader(f))
	csvData := [][]string{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			return csvData
		}
		csvData = append(csvData, row)
	}
}

type optionsWrapper struct {
	options map[string]string
}

func (o optionsWrapper) isPresentInRow(row []string, propertyName string) bool {
	valueToFind := o.options[propertyName]
	return row[columnNameToIdx[propertyName]] == valueToFind
}

func (o optionsWrapper) fieldProvidedForSearch(propertyName string) bool {
	_, ok := o.options[propertyName]
	return ok
}

func (o optionsWrapper) filterData(csvData [][]string, propertyName string) [][]string {
	if !o.fieldProvidedForSearch(propertyName) {
		return csvData
	}

	results := [][]string{}
	for _, row := range csvData {
		if o.isPresentInRow(row, propertyName) {
			results = append(results, row)
		}
	}
	return results
}


var columnNameToIdx = map[string]int {
	"permalink": 0,
	"company_name": 1,
	"number_employees": 2,
	"category": 3,
	"city": 4,
	"state": 5,
	"funded_date": 6,
	"raised_amount": 7,
	"raised_currency": 8,
	"round": 9,
}

var fieldsForSearch = []string {"company_name","city","state","round"}

func collectDataFromRow(aggregatedData map[string]string, row []string) {
	for key := range columnNameToIdx {
		aggregatedData[key] = row[columnNameToIdx[key]]
	}
}

func Where(options map[string]string) []map[string]string {
	csvData := readFile("startup_funding.csv")
	opts := optionsWrapper{options}
	narrowedData := csvData
	
	for _,field := range fieldsForSearch {
		narrowedData = opts.filterData(narrowedData, field)
	}

	output := []map[string]string{}
	for _, row := range narrowedData {
		aggregatedData := make(map[string]string)
		collectDataFromRow(aggregatedData, row)
		output = append(output, aggregatedData)
	}
	return output	
}

func collectDataAndStopSearching(opts optionsWrapper, propertyName string, row []string, aggregatedData map[string]string) bool {
	if !opts.fieldProvidedForSearch(propertyName) {
		return false
	}
	propertyPresentInRow := opts.isPresentInRow(row, propertyName)
	if !propertyPresentInRow {
		return true // skip row - provided property not found in row
	}
	// all good - append data and proceed with next property
	collectDataFromRow(aggregatedData, row)
	return false
}

func FindBy(options map[string]string) (map[string]string, error) {
	csvData := readFile("startup_funding.csv")

	opts := optionsWrapper{options}

	for _, row := range csvData {
		aggregatedRowData := make(map[string]string)

		goToNextRow := false
		for _, field := range fieldsForSearch {
			if collectDataAndStopSearching(opts, field, row, aggregatedRowData) {
				goToNextRow = true
				break
			}
		}
		if goToNextRow {
			continue
		}

		// all found in this row
		return aggregatedRowData, nil
	}

	return make(map[string]string), errors.New("Record Not Found")
}