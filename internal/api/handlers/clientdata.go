package handlers

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"regexp"
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
		return
	case r.Method == http.MethodGet && getDataRowRe.MatchString(r.URL.Path):
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
		return
	case h.Method == http.MethodGet && getDataRowRe.MatchString(h.URL.Path):
		// define logic for GET /clientdata/hash/row
		return
	}
}
