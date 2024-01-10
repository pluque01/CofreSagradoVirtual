package validcsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"

	log "github.com/pluque01/CofreSagradoVirtual/internal/logger"
	"github.com/pluque01/CofreSagradoVirtual/internal/stringmetrics"
)

type ClientFile struct {
	fileTypes   ClientTypes
	fileContent [][]string
}

type ClientTypes struct {
	types map[int]regexp.Regexp
}

type DataType struct {
	KeyRegex  regexp.Regexp
	TypeRegex regexp.Regexp
}

var knownDataTypes = map[string]DataType{
	"name":      {*regexp.MustCompile(`(?i)(.*name.*|.*nombre.*)`), *regexp.MustCompile(`[A-ZÁÉÍÓÚÜÑ][a-záéíóúüÜñÑ]+(\s[A-ZÁÉÍÓÚÜÑ][a-záéíóúüÜñÑ]+)?(\s(?:(de|del)\s(?:las?\s)?[A-ZÁÉÍÓÚÜÑ][a-záéíóúüÜñÑ]+))?`)},
	"surname":   {*regexp.MustCompile(`(?i)(.*surname.*|.*apellido.*)`), *regexp.MustCompile(`[A-ZÁÉÍÓÚÜÑ][a-záéíóúüñ]+(?:[-\s][A-ZÁÉÍÓÚÜÑ][a-záéíóúüñ]+)*(?:\s[A-ZÁÉÍÓÚÜÑ][a-záéíóúüñ]+(?:[-\s][A-ZÁÉÍÓÚÜÑ][a-záéíóúüñ]+)*)?`)},
	"telephone": {*regexp.MustCompile(`(?i)(.*phone.*|.*tel(e|é)fono.*|.*m(o|ó)vil.*|.*n(u|ú)mero.*|.*number.*)`), *regexp.MustCompile(`\d{3}(?:[-\s]?\d{2,3}){2}`)},
}

func NewClientTypes() *ClientTypes {
	return &ClientTypes{types: map[int]regexp.Regexp{}}
}

func IsValidSeparator(separator rune) bool {
	return separator != '\r' && separator != '\n' && separator != '\uFFFD'
}

func readFile(file io.Reader, separator rune) ([][]string, error) {
	if !IsValidSeparator(separator) {
		return nil, errors.New("invalid separator")
	}
	reader := csv.NewReader(file)
	reader.Comma = separator

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read records: %w", err)
	}
	return records, nil
}

func inferTypes(values []string) []string {
	types := []string{}
	for _, value := range values {
		found := false
		for key, element := range knownDataTypes {
			// Matches only to the first known dataType
			// If no match found, then it is unknown
			if element.KeyRegex.MatchString(value) {
				types = append(types, key)
				found = true
				break
			}
		}
		if !found {
			log.Default.Logger.Debug().Msgf("Unknown type: %s", value)
			types = append(types, "unknown")
		}
	}
	return types
}

func (c *ClientFile) ValidateFileContent() *[][]string {
	results := [][]string{}
	for i := range c.fileContent {
		results = append(results, *c.ValidateRow(i))
	}
	return &results
}

func (c ClientFile) ValidateRow(row int) *[]string {
	records := c.fileContent[row]
	recordsValidated := []string{}
	// Iterate over each value of the row
	for i, value := range records {
		var nonMatches string
		var err error
		pattern := c.fileTypes.types[i]
		matches := pattern.FindStringSubmatch(value)
		if len(matches) == 0 {
			nonMatches = value
		} else if value != matches[0] {
			nonMatches, err = getNonMatchingPattern(value, matches[0])
			if err != nil {
				log.Default.Logger.Warn().Msgf("Error getting non-matching pattern: %s", err)
				return nil
			}
		}
		recordsValidated = append(recordsValidated, nonMatches)
	}
	return &recordsValidated
}

func RecommendStrings(name string, names *[]string) ([]string, error) {
	const MaxRecommendations uint8 = 10
	const MatchThreshold float64 = 0.8
	matches := []struct {
		name  string
		value float64
	}{}
	for _, n := range *names {
		score, err := stringmetrics.JaroWinkler(name, n)
		if err != nil {
			continue
		}
		if score > MatchThreshold {
			matches = append(matches, struct {
				name  string
				value float64
			}{n, score})
		}
	}

	if len(matches) == 0 {
		return []string{}, errors.New("no matches found")
	}
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].value > matches[j].value
	})
	matchesSplit := matches[:stringmetrics.Min(int(MaxRecommendations), len(matches))]
	recommendations := []string{}

	for _, recommendation := range matchesSplit {
		recommendations = append(recommendations, recommendation.name)
	}
	return recommendations, nil
}

func getMatchingRuneIndex(a string, r rune) (int, error) {
	for i, c := range a {
		if c == r {
			return i, nil
		}
	}
	return 0, errors.New("no matching character")
}

func getNonMatchingPattern(originalStr, matchingStr string) (string, error) {
	// Find the non-matching part
	var nonMatchingPart string
	firstMatchIndex, err := getMatchingRuneIndex(originalStr, rune(matchingStr[0]))
	if err != nil {
		return "", err
	}
	startRange := firstMatchIndex
	endRange := firstMatchIndex + len(matchingStr)
	for i, c := range originalStr {
		if i < startRange || i >= endRange {
			nonMatchingPart += string(c)
		}
	}
	return nonMatchingPart, nil
}

func NewClientFile(records [][]string) *ClientFile {
	clientTypes := NewClientTypes()
	types := inferTypes(records[0])
	for i := range records[0] {
		clientTypes.types[i] = knownDataTypes[types[i]].TypeRegex
	}
	return &ClientFile{fileTypes: *clientTypes, fileContent: records[1:]}
}

func (c *ClientFile) Print() {
	for _, record := range c.fileContent {
		for _, value := range record {
			fmt.Printf("%s\t", value)
		}
		fmt.Println()
	}
}
