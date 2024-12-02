package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	Problem1()
	Problem2()
}

func Problem1() {
	// Parse input.txt, a space delimited file of integers
	// Split the integers into two slices, one for each column
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1 []int
	var col2 []int

	for {
		var a, b int
		_, err := fmt.Fscanf(file, "%d %d\n", &a, &b)
		if err != nil {
			break
		}
		col1 = append(col1, a)
		col2 = append(col2, b)
	}

	// Sort the columns
	sort.Ints(col1)
	sort.Ints(col2)

	var totalDistance int

	// Iterate through each slice, finding the delta between the two values
	for i := 0; i < len(col1); i++ {
		// Get the absolute value of the difference between the two values
		diff := col1[i] - col2[i]
		if diff < 0 {
			diff = -diff
		}
		totalDistance += diff
	}

	fmt.Println(totalDistance)
}

func Problem2() {
	// Parse input.txt, a space delimited file of integers
	// Split the integers into two slices, one for each column
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1 []int
	var col2 []int

	for {
		var a, b int
		_, err := fmt.Fscanf(file, "%d %d\n", &a, &b)
		if err != nil {
			break
		}
		col1 = append(col1, a)
		col2 = append(col2, b)
	}

	// Sort the columns
	sort.Ints(col1)
	sort.Ints(col2)

	// Find the number of times the item in the 1st column is present in the 2nd column
	var similarityScore int

	for i := 0; i < len(col1); i++ {
		var simCount int
		// Count the number of times the item in the 1st column is present in the 2nd column
		for j := 0; j < len(col2); j++ {
			if col1[i] == col2[j] {
				simCount++
			}
		}
		// Multiply the simCount by the value in the 1st column and add it to the similarityScore
		similarityScore += simCount * col1[i]
	}

	fmt.Println(similarityScore)
}
