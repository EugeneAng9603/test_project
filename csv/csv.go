package csv

import (
	"encoding/csv"
	"os"
)

type CSVFileSaver struct {
	File   *os.File
	Writer *csv.Writer
}

func (s *CSVFileSaver) Save(data string) error {
	return s.Writer.Write([]string{data})
}

func NewCSVFileSaver(filePath string) (*CSVFileSaver, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)
	return &CSVFileSaver{File: file, Writer: writer}, nil
}
