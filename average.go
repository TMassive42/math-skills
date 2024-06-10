package mathskills

import "math"

func CalculateAverage(data []int) float64 {
    sum := 0
    if data == nil {
        return 0.0
    }
    for _, value := range data {
        sum += value
    }
    return math.Round(float64(sum) / float64(len(data)))
}