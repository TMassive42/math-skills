package mathskills

import (
	"testing"
)

func TestMedian(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := float64(5)
	got := CalculateMedian(input)

	if got != expected {
		t.Fatalf("expected: %v and got %v", expected, got)
	}
}