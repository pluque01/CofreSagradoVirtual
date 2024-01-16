package validcsv

import (
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/pluque01/CofreSagradoVirtual/internal/projectpath"
)

var readFileTests = []struct {
	expectedErr string
	filePath    string
	expected    [][]string
	separator   rune
}{
	{"", string(projectpath.Root + "/test/data/default.csv"), [][]string{{"Nombre", "Apellido", "telefono"}, {"juan", "Perez", "123456789"}, {"paco", "Gomez", "987654321"}}, ';'},
	{"", string(projectpath.Root + "/test/data/compound_record.csv"), [][]string{{"Nombre", "Apellido", "telefono"}, {"Juan Alberto", "Perez", "123456789"}, {"Francisco Javier", "Gomez", "987654321"}}, ';'},
	{"", string(projectpath.Root + "/test/data/missing_record.csv"), [][]string{{"Nombre", "Apellido", "telefono"}, {"juan", "", "123456789"}, {"paco", "Gomez", ""}}, ','},
	{"", string(projectpath.Root + "/test/data/missing_type.csv"), [][]string{{"Nombre", "", "telefono"}, {"juan", "Perez", "123456789"}, {"paco", "Gomez", "987654321"}}, ';'},
	{"invalid separator", string(projectpath.Root + "/test/data/default.csv"), [][]string{{"Nombre", "Apellido", "telefono"}, {"juan", "Perez", "123456789"}, {"paco", "Gomez", "987654321"}}, '\r'},
}

func TestReadFile(t *testing.T) {
	for _, tt := range readFileTests {
		t.Run(tt.filePath, func(t *testing.T) {
			// Open the file and read it
			csvFile, _ := os.Open(tt.filePath)
			ans, err := readFile(csvFile, tt.separator)
			defer csvFile.Close()
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

var inferTypesTests = []struct {
	values   []string
	expected []string
}{
	{[]string{"nombre", "apellidos", "telefono"}, []string{"name", "surname", "telephone"}},
	{[]string{"NOMBRE", "ApEllidos", "Móvil"}, []string{"name", "surname", "telephone"}},
	{[]string{"Nombre 1", "Apellido 2", "Número 3"}, []string{"name", "surname", "telephone"}},
	{[]string{"Ciudad", "Calle", ""}, []string{"unknown", "unknown", "unknown"}},
}

func TestInferTypes(t *testing.T) {
	for _, tt := range inferTypesTests {
		t.Run(tt.values[0], func(t *testing.T) {
			ans := inferTypes(tt.values)
			if !reflect.DeepEqual(ans, tt.expected) {
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

func TestValidateFileContent(t *testing.T) {
	// Create a sample ClientFile instance with test data
	clientFile := &ClientFile{
		fileTypes: ClientTypes{
			types: map[int]regexp.Regexp{
				0: *regexp.MustCompile(`[A-Z][a-z]+`),
				1: *regexp.MustCompile(`\d{3}-\d{3}-\d{4}`),
			},
		},
		fileContent: [][]string{
			{"John23", "123-456-7890"},
			{"Jane", "abc-def-ghij"},
			{"Jane", "abc def-ghij"},
		},
	}

	// Define the expected results
	expectedResults := [][]string{
		{"23", ""},
		{"", "abc-def-ghij"},
		{"", "abc def-ghij"},
	}

	// Call the method being tested
	results := clientFile.ValidateFileContent()

	// Compare the actual results with the expected results
	if !reflect.DeepEqual(*results, expectedResults) {
		t.Errorf("Validation failed. Expected: %v, but got: %v", expectedResults, *results)
	}
}

func TestRecommendStrings(t *testing.T) {
	names := []string{"John", "Jane", "Jack", "Jill", "Jim"}

	t.Run("Matching name found", func(t *testing.T) {
		name := "John"
		expected := []string{"John"}
		result, err := RecommendStrings(name, &names)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RecommendStrings(%s, %v) = %v, expected %v", name, names, result, expected)
		}
	})

	t.Run("No matching name found", func(t *testing.T) {
		name := "Alex"
		expected := []string{}
		result, err := RecommendStrings(name, &names)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RecommendStrings(%s, %v) = %v, expected %v", name, names, result, expected)
		}
	})

	t.Run("Multiple matching names found", func(t *testing.T) {
		name := "Ji"
		expected := []string{"Jim", "Jill"}
		result, err := RecommendStrings(name, &names)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RecommendStrings(%s, %v) = %v, expected %v", name, names, result, expected)
		}
	})

	t.Run("Empty names list", func(t *testing.T) {
		name := "John"
		names := []string{}
		expected := []string{}
		result, err := RecommendStrings(name, &names)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RecommendStrings(%s, %v) = %v, expected %v", name, names, result, expected)
		}
	})
}

func TestGetNonMatchingPattern(t *testing.T) {
	originalStr := "Hello, World!"
	matchingStr := "Hello"
	expectedResult := ", World!"

	result, err := getNonMatchingPattern(originalStr, matchingStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expectedResult {
		t.Errorf("Expected %q, but got %q", expectedResult, result)
	}
}

func TestGetMatchingRuneIndex(t *testing.T) {
	a := "Hello, World!"
	r := 'o'
	expectedResult := 4

	result, err := getMatchingRuneIndex(a, r)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expectedResult {
		t.Errorf("Expected %d, but got %d", expectedResult, result)
	}
}
