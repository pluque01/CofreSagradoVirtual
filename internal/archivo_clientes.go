package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"regexp"
)

type ClientFile struct {
	fileTypes   ClientTypes
	fileContent [][]string
}

type ClientTypes struct {
	types map[string]regexp.Regexp
}
func readFile(filePath string, separator rune) ([][]string, error) {
	if separator == '\r' || separator == '\n' || separator == '\uFFFD' {
		return nil, errors.New("invalid separator")
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while reading file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = separator

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records", err)
	}
	return records, nil
}

