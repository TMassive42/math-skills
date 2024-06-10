package mathskills

import (
	"math"
)

func CalculateStandardDeviation(variance float64) float64 {
    return math.Round(math.Sqrt(variance))
}