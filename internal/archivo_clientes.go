package main

import (
	"errors"
	"regexp"
)

type ClientFile struct {
	fileContent [][]string
	fileTypes   ClientTypes
}

type ClientTypes struct {
	types map[string]regexp.Regexp
}
func readFile(filePath string, separator rune) ([][]string, error) {
}
