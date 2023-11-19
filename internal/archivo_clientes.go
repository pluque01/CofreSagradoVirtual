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
	types map[int]regexp.Regexp
}

type DataType struct {
	KeyRegex  regexp.Regexp
	TypeRegex regexp.Regexp
}

var knownDataTypes = map[string]DataType{
	"name":      {*regexp.MustCompile(`(?i)(.*name.*|.*nombre.*)`), *regexp.MustCompile(`^[A-Za-z ]+$`)},
	"surname":   {*regexp.MustCompile(`(?i)(.*surname.*|.*apellido.*)`), *regexp.MustCompile(`^[A-Za-z ]+$`)},
	"telephone": {*regexp.MustCompile(`(?i)(.*phone.*|.*tel(e|é)fono.*|.*m(o|ó)vil.*|.*n(u|ú)mero.*|.*number.*)`), *regexp.MustCompile(`^[0-9]+$`)},
}

func NewClientTypes() *ClientTypes {
	return &ClientTypes{types: map[int]regexp.Regexp{}}
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

func inferTypes(values []string) []string {
	types := []string{}
	for _, value := range values {
		found := false
		for key, element := range knownDataTypes {
			// Matches only to the first known dataType
			// If no match found, then it is unknown
			if element.KeyRegex.MatchString(value) {
				types = append(types, key)
				found = true
				break
			}
		}
		if !found {
			types = append(types, "unknown")
		}
	}
	return types
}

func NewClientFile(filePath string, separator rune) *ClientFile {
	clientTypes := NewClientTypes()
	records, err := readFile(filePath, separator)
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	types := inferTypes(records[0])
	for i := range records[0] {
		clientTypes.types[i] = knownDataTypes[types[i]].TypeRegex
	}
	return &ClientFile{fileTypes: *clientTypes, fileContent: records[1:]}
}

