package mathskills

import (
	"sort"
    "math"
)
func CalculateMedian(data []int) float64 {
    if data == nil {
        return 0
    }
    sort.Ints(data)
    mid := len(data) / 2
    if len(data)%2 == 0 {
        return math.Round((float64(data[mid-1]) + float64(data[mid])) / 2)
    }
    return float64(data[mid])
}