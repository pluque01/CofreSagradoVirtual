package stringmetrics

import (
	"errors"
)

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
	matchDistance := (Max(len(s1), len(s2)) / 2) - 1

	// Loop over the first string
	for i := 0; i < len(s1); i++ {
		// Define the start and end index for comparison
		start := Max(0, i-matchDistance)
		end := Min(i+matchDistance+1, len(s2))

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
	length := Min(len(s1), len(s2))
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
