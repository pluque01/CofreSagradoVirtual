package main

import (
	"regexp"
)

type ClientFile struct {
	fileContent [][]string
	fileTypes   ClientTypes
}

type ClientTypes struct {
	types map[string]regexp.Regexp
}
