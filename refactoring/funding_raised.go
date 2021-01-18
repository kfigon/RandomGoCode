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

func readOption(options map[string]string, keyName string, id int, csv_data [][]string) [][]string {
	_, ok := options[keyName]
	if ok == true {
		results := [][]string{}
		for i := 0; i < len(csv_data); i++ {
			if csv_data[i][id] == options[keyName] {
				results = append(results, csv_data[i])
			}
		}
		return results
	}
	return nil
}



func Where(options map[string]string) []map[string]string {
	csv_data := readFile("startup_funding.csv")
	if res := readOption(options, "company_name", 1, csv_data); res != nil {
		csv_data = res
	}
	if res := readOption(options, "city", 4, csv_data); res != nil {
		csv_data = res
	}
	if res := readOption(options, "state", 5, csv_data); res != nil {
		csv_data = res
	}
	if res := readOption(options, "round", 9, csv_data); res != nil {
		csv_data = res
	}


	output := []map[string]string{}
	for i := 0; i < len(csv_data); i++ {
		mapped := make(map[string]string)
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
		output = append(output, mapped)
	}

	return output
}

func FindBy(options map[string]string) (map[string]string, error) {
	f, _ := os.Open("startup_funding.csv")
	reader := csv.NewReader(bufio.NewReader(f))
	csv_data := [][]string{}

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		csv_data = append(csv_data, row)
	}

	for i := 0; i < len(csv_data); i++ {
		var ok bool
		mapped := make(map[string]string)

		_, ok = options["company_name"]
		if ok == true {
			if csv_data[i][1] == options["company_name"] {
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
			} else {
				continue
			}
		}

		_, ok = options["city"]
		if ok == true {
			if csv_data[i][4] == options["city"] {
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
			} else {
				continue
			}
		}

		_, ok = options["state"]
		if ok == true {
			if csv_data[i][5] == options["state"] {
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
			} else {
				continue
			}
		}

		_, ok = options["round"]
		if ok == true {
			if csv_data[i][9] == options["round"] {
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
			} else {
				continue
			}
		}

		return mapped, nil
	}

	return make(map[string]string), errors.New("Record Not Found")
}
