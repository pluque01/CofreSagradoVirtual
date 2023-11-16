package main

import (
	"reflect"
	"testing"
)

var readFileTests = []struct {
	expectedErr string
	filePath    string
	expected    [][]string
	separator   rune
}{
	{"", "../test/data/default.csv", [][]string{{"Nombre", "Apellido", "telefono"}, {"juan", "Perez", "123456789"}, {"paco", "Gomez", "987654321"}}, ';'},
	{"", "../test/data/compound_record.csv", [][]string{{"Nombre", "Apellido", "telefono"}, {"Juan Alberto", "Perez", "123456789"}, {"Francisco Javier", "Gomez", "987654321"}}, ';'},
	{"", "../test/data/missing_record.csv", [][]string{{"Nombre", "Apellido", "telefono"}, {"juan", "", "123456789"}, {"paco", "Gomez", ""}}, ','},
	{"", "../test/data/missing_type.csv", [][]string{{"Nombre", "", "telefono"}, {"juan", "Perez", "123456789"}, {"paco", "Gomez", "987654321"}}, ';'},
	{"invalid separator", "../test/data/default.csv", [][]string{{"Nombre", "Apellido", "telefono"}, {"juan", "Perez", "123456789"}, {"paco", "Gomez", "987654321"}}, '\r'},
}

func TestReadFile(t *testing.T) {
	for _, tt := range readFileTests {
		t.Run(tt.filePath, func(t *testing.T) {
			ans, err := readFile(tt.filePath, tt.separator)
			if tt.expectedErr != "" && err == nil {
				t.Errorf("got %v, want %v", err, tt.expectedErr)
			} else if tt.expectedErr == "" && err != nil {
				t.Errorf("got %v, want %v", err, tt.expectedErr)
			} else if tt.expectedErr == "" && err == nil && !reflect.DeepEqual(ans, tt.expected) {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
	if t.Failed() {
		t.Logf("FAIL - %s", t.Name())
	} else {
		t.Logf("OK - %s", t.Name())
	}
}
