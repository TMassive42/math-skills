package mathskills

import (
	"testing"
	"math"
)

func TestVariance(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := float64(7)
	got := CalculateVariance(input)

	if math.Round(got) != expected {
		t.Fatalf("expected: %v and got %v", expected, math.Round(got))
	}
}