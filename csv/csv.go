package csv

import (
	"encoding/csv"
	"os"
)

type CSVFileSaver struct {
	file   *os.File
	writer *csv.Writer
}

func (s *CSVFileSaver) Save(data string) error {
	return s.writer.Write([]string{data})
}

func NewCSVFileSaver(filePath string) (*CSVFileSaver, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)
	return &CSVFileSaver{file: file, writer: writer}, nil
}
