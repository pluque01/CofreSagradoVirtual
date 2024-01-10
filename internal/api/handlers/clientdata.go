package handlers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pluque01/CofreSagradoVirtual/internal/api/utils"
	"github.com/pluque01/CofreSagradoVirtual/internal/logger"
	"github.com/pluque01/CofreSagradoVirtual/internal/validcsv"
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

func (c *ClientData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	logger.Default.Logger.Info().Msgf("Request received: %s %s", r.Method, r.URL.Path)
	switch {
	case r.Method == http.MethodPost && postDataRe.MatchString(r.URL.Path):
		c.handlePostData(rw, r)
		return
	case r.Method == http.MethodGet && getDataRowsRe.MatchString(r.URL.Path):
		c.handleGetRows(rw, r)
		return
	case r.Method == http.MethodGet && getDataRowRe.MatchString(r.URL.Path):
		c.handleGetRow(rw, r)
		return
	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *ClientData) handlePostData(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(rw, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(rw, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the content of the file into a buffer
	var fileBuffer bytes.Buffer
	if _, err := io.Copy(&fileBuffer, file); err != nil {
		http.Error(rw, "Error reading the file", http.StatusInternalServerError)
		return
	}

	// Calculate MD5 hash
	md5Hash := utils.GetFileMd5(bytes.NewReader(fileBuffer.Bytes()))

	// Read the CSV file from the buffer
	reader := csv.NewReader(&fileBuffer)
	delimiter := ','
	if r.FormValue("delimiter") != "" {
		d, err := strconv.Atoi(r.FormValue("delimiter"))
		if err != nil {
			http.Error(rw, "Error converting delimiter to rune", http.StatusBadRequest)
		}
		if validcsv.IsValidSeparator(rune(d)) {
			delimiter = rune(d)
		}
	}
	reader.Comma = delimiter

	records, err := reader.ReadAll()
	if err != nil {
		http.Error(rw, "Error reading the file", http.StatusBadRequest)
	}

	clientfile := validcsv.NewClientFile(records)
	validcsv.SaveClientFile(md5Hash, clientfile)

	logger.Default.Logger.Info().Msgf("Saved file with hash: %s", md5Hash)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "%s", md5Hash)
}

func (c *ClientData) handleGetRows(rw http.ResponseWriter, r *http.Request) {
	hash := getDataRowsRe.FindStringSubmatch(r.URL.Path)[1]
	rows, err := validcsv.GetClientFileRows(hash)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error getting the number of rows: %s", err), http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "%d", rows)
}

func (c *ClientData) handleGetRow(rw http.ResponseWriter, r *http.Request) {
	hash := getDataRowRe.FindStringSubmatch(r.URL.Path)[1]
	row := getDataRowRe.FindStringSubmatch(r.URL.Path)[2]
	convertedRow, err := strconv.Atoi(row)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error converting row to int: %s", err), http.StatusBadRequest)
		return
	}
	rowValidated, err := validcsv.GetValidatedRow(hash, convertedRow)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error getting the row: %s", err), http.StatusNotFound)
		return
	}
	json.NewEncoder(rw).Encode(rowValidated)
	rw.WriteHeader(http.StatusOK)
}
