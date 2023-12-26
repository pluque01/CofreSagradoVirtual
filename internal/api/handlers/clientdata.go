package handlers

import (
	"log"
	"net/http"
	"regexp"
)

type ClientData struct {
	l *log.Logger
}

func NewClientData(l *log.Logger) *ClientData {
	return &ClientData{l}
}

var (
	postDataRe    = regexp.MustCompile(`^/clientdata/$`)
	getDataRowsRe = regexp.MustCompile(`^/clientdata/([a-z0-9]+)/rows$`)
	getDataRowRe  = regexp.MustCompile(`^/clientdata/([a-z0-9]+)/([0-9]+)$`)
)

func (c *ClientData) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	switch {
	case h.Method == http.MethodPost && postDataRe.MatchString(h.URL.Path):
		// define logic for POST /clientdata/
		return
	case h.Method == http.MethodGet && getDataRowsRe.MatchString(h.URL.Path):
		// define logic for GET /clientdata/hash/rows
		return
	case h.Method == http.MethodGet && getDataRowRe.MatchString(h.URL.Path):
		// define logic for GET /clientdata/hash/row
		return
	}
}
