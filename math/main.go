package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"mathskills"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter file name:")
		return
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	// Read the data from the file
	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numStr := strings.TrimSpace(scanner.Text())
		if numStr == "" {
			continue
		}

		val, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			fmt.Println("Error parsing data:", err)
			return
		}

		data = append(data, int(val))
	}

	avg := mathskills.CalculateAverage(data)
	median := mathskills.CalculateMedian(data)
	variance := mathskills.CalculateVariance(data)
	stdDev := mathskills.CalculateStandardDeviation(variance)

	// Print the results
	fmt.Println("Average:", int(avg))
	fmt.Println("Median:", median)
	fmt.Println("Variance:", int(variance))
	fmt.Println("Standard Deviation:", int(stdDev))
}

