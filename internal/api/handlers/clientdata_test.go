package handlers

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pluque01/CofreSagradoVirtual/internal/api/utils"
	"github.com/pluque01/CofreSagradoVirtual/internal/validcsv"
)

func TestHandleGetRows(t *testing.T) {
	// Create a new instance of ClientData
	clientData := NewClientData(nil)

	clientrecords := [][]string{
		{"Nombre", "Apellido", "Telefono"},
		{"juan", "Perez", "123456789"},
		{"paco", "gomez", "987654321"},
	}
	clientfile := validcsv.NewClientFile(clientrecords)
	validcsv.SaveClientFile("abc123", clientfile)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/clientdata/abc123/rows", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handleGetRows method
	handler := http.HandlerFunc(clientData.handleGetRows)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "2"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}

func TestHandleGetRow(t *testing.T) {
	// Create a new instance of ClientData
	clientData := NewClientData(nil)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/clientdata/abc123/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handleGetRow method
	handler := http.HandlerFunc(clientData.handleGetRow)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestHandlePostData(t *testing.T) {
	// Create a new instance of ClientData
	clientData := NewClientData(nil)

	// Create a new HTTP request with a sample file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the CSV file to the multipart form
	fileContent := `Nombre;Apellido;telefono
juan;Perez;123456789
paco;Gomez;987654321`
	part, err := writer.CreateFormFile("file", "data.csv")
	if err != nil {
		t.Fatal(err)
	}
	part.Write([]byte(fileContent))

	// Close the multipart writer
	writer.Close()

	req, err := http.NewRequest("POST", "/clientdata/", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a new ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handlePostData method
	handler := http.HandlerFunc(clientData.handlePostData)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := utils.GetFileMd5(bytes.NewReader([]byte(fileContent)))
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}
