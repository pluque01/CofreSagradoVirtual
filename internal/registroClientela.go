package main

import (
	"regexp"
)

type RegistroClientera struct {
	diccionarioTipos map[string]regexp.Regexp
}
