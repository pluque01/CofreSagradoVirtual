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

type DataTypes struct {
	key       string
	keyRegex  regexp.Regexp
	typeRegex regexp.Regexp
}

var knownDataTypes = []DataTypes{
	{"name", *regexp.MustCompile(`(?i)(.*name.*|.*nombre.*)`), *regexp.MustCompile(`^[A-Za-z ]+$`)},
	{"surname", *regexp.MustCompile(`(?i)(.*surname.*|.*apellido.*)`), *regexp.MustCompile(`^[A-Za-z ]+$`)},
	{"telephone", *regexp.MustCompile(`(?i)(.*phone.*|.*tel(e|é)fono.*|.*m(o|ó)vil.*|.*n(u|ú)mero.*|.*number.*)`), *regexp.MustCompile(`^[0-9]+$`)},
	{"unknown", *regexp.MustCompile(`.*`), *regexp.MustCompile(`.*`)},
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
		for _, dataType := range knownDataTypes {
			// Matches only to the first known dataType
			// If no match found, then it is unknown
			if dataType.keyRegex.MatchString(value) {
				types = append(types, dataType.key)
				break
			}
		}
	}
	return types
}
