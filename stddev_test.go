package mathskills

import (
	"math"
	"testing"
)

func TestStddev(t *testing.T) {
	input := float64(7)
	expected := float64(3)
	got := CalculateStandardDeviation(input)

	if math.Round(got) != expected {
		t.Fatalf("expected: %v and got %v", expected, math.Round(got))
	}
}
