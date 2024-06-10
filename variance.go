package mathskills

import (
	"math"
)
func CalculateVariance(data []int) float64 {
    sum := 0.0
    if data == nil {
        return 0.0
    }
    avg := CalculateAverage(data)
    for _, value := range data {
        sum += math.Pow(float64(value)-avg, 2)
    }
    return math.Round(sum/float64(len(data)))
}
