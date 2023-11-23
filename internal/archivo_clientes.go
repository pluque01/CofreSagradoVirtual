package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"regexp"
	"sort"
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

func readFile(filePath string, separator rune) ([][]string, error) {
	if separator == '\r' || separator == '\n' || separator == '\uFFFD' {
		return nil, errors.New("invalid separator")
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while reading file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = separator

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records", err)
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
			types = append(types, "unknown")
		}
	}
	return types
}

func (c *ClientFile) ValidateFileContent() *[][]string {
	results := [][]string{}
	for _, records := range c.fileContent {
		columnResults := []string{}
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
					log.Println(err)
				}
			}
			columnResults = append(columnResults, nonMatches)
		}
		results = append(results, columnResults)
	}
	return &results
}

// JaroWinkler function calculates the Jaro-Winkler similarity between two strings
func JaroWinkler(s1, s2 string) (float64, error) {
	if len(s1) == 0 || len(s2) == 0 {
		return 0, errors.New("input strings must not be empty")
	}
	// Create two match arrays
	s1Matches := make([]bool, len(s1))
	s2Matches := make([]bool, len(s2))

	// Initialize variables
	matches := 0
	transpositions := 0
	halfTranspositions := 0

	// Define the match distance
	matchDistance := (max(len(s1), len(s2)) / 2) - 1

	// Loop over the first string
	for i := 0; i < len(s1); i++ {
		// Define the start and end index for comparison
		start := max(0, i-matchDistance)
		end := min(i+matchDistance+1, len(s2))

		// Loop over the second string within the defined range
		for j := start; j < end; j++ {
			// If the character in the second string is already matched, continue
			if s2Matches[j] {
				continue
			}
			// If the characters do not match, continue
			if s1[i] != s2[j] {
				continue
			}
			// If the characters match, mark them as matched and increment the match count
			s1Matches[i] = true
			s2Matches[j] = true
			matches++
			break
		}
	}
	// If there are no matches, return 0
	if matches == 0 {
		return 0, nil
	}
	// Calculate transpositions
	k := 0
	for i := 0; i < len(s1); i++ {
		if !s1Matches[i] {
			continue
		}
		for !s2Matches[k] {
			k++
		}
		if s1[i] != s2[k] {
			halfTranspositions++
		}
		k++
	}
	// Calculate the number of transpositions
	transpositions = halfTranspositions / 2
	// Calculate the Jaro similarity
	jaro := ((float64(matches) / float64(len(s1))) +
		(float64(matches) / float64(len(s2))) +
		((float64(matches) - float64(transpositions)) / float64(matches))) / 3.0
	// Calculate the prefix length
	prefix := 0
	length := min(len(s1), len(s2))
	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			break
		}
		prefix++
	}
	cl := 0.1
	if prefix > 4 {
		prefix = 4
	}
	// Return the Jaro-Winkler similarity
	return jaro + (float64(prefix) * cl * (1 - jaro)), nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RecommendStrings(name string, names *[]string) ([]string, error) {
	const MaxRecommendations uint8 = 10
	const MatchThreshold float64 = 0.8
	matches := []struct {
		name  string
		value float64
	}{}
	for _, n := range *names {
		score, err := JaroWinkler(name, n)
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
	matchesSplit := matches[:min(int(MaxRecommendations), len(matches))]
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

func NewClientFile(filePath string, separator rune) *ClientFile {
	clientTypes := NewClientTypes()
	records, err := readFile(filePath, separator)
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	types := inferTypes(records[0])
	for i := range records[0] {
		clientTypes.types[i] = knownDataTypes[types[i]].TypeRegex
	}
	return &ClientFile{fileTypes: *clientTypes, fileContent: records[1:]}
}

