package state

import (
	"encoding/csv"
	"os"
)

type CSVData struct {
	Headers []string
	Data    [][]string
}

func (c *CSVData) ParseCSV(filePath string) error {
	rawData, err := readCSVFile(filePath)
	if err != nil {
		return err
	}

	c.Headers = rawData[0]
	if len(rawData[0]) <= 1 {
		return nil
	}
	rows := [][]string{}
	for _, row := range rawData {
		if row[0] == "name" {
			continue
		}
		rows = append(rows, row)
	}
	c.Data = rows
	return nil
}

func readCSVFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
