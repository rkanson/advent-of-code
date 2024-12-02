package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSafe := 0

	for scanner.Scan() {
		report := scanner.Text()
		var columns []int
		// Split the report into columns by whitespace
		levels := strings.Fields(report)
		for _, level := range levels {
			num, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			columns = append(columns, num)
		}

		// Check each item in columns to see if they're all increasing
		totalIncs := 0
		allIncreasing := false
		for i := 0; i < len(columns)-1; i++ {
			if columns[i+1] > columns[i] {
				totalIncs++
			}
		}
		if totalIncs == len(columns)-1 {
			allIncreasing = true
		}

		totalDecs := 0
		allDecreasing := false
		for i := 0; i < len(columns)-1; i++ {
			if columns[i+1] < columns[i] {
				totalDecs++
			}
		}
		if totalDecs == len(columns)-1 {
			allDecreasing = true
		}

		allSameOperations := allIncreasing == allDecreasing

		// Check the columns, getting the delta between each column.
		// If the delta between any two columns is greater than or equal to  1, but less than or equal to 3,
		// then the columns are considered deltaSafe.
		deltaSafe := false
		for i := 0; i < len(columns)-1; i++ {
			delta := columns[i+1] - columns[i]
			// Get the absolute value
			if delta < 0 {
				delta = -delta
			}
			if delta >= 1 && delta <= 3 {
				deltaSafe = true
			}
		}

		if !allSameOperations && deltaSafe {
			totalSafe++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalSafe)
}
