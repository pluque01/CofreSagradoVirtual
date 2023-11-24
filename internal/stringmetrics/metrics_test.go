package stringmetrics

import (
  "math"
	"testing"
)

var jaroWinklerTests = []struct {
	s1       string
	s2       string
	expected float64
}{
	{"carlos", "carlps", 0.933},
	{"pablo", "pavlo", 0.893},
	{"rebeca", "rebecca", 0.971},
	{"maría", "maria", 0.875},
}

func TestJaroWinkler(t *testing.T) {
	for _, tt := range jaroWinklerTests {
		t.Run(tt.s1, func(t *testing.T) {
			ans, err := JaroWinkler(tt.s1, tt.s2)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			if math.Abs(ans-tt.expected) > 0.001 {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}
