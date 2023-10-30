package main

import (
	"regexp"
)

type TipoDatosCliente struct {
	diccionarioTipos map[string]regexp.Regexp
}
