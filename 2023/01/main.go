package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Problem 1")
	fmt.Println(Problem1())
	fmt.Println("Problem 2")
	fmt.Println(Problem2())
}

func Problem1() int {
	var calibrationValue int
	var totalCalibration int

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		regex, _ := regexp.Compile(`(\d){1}`)
		found := regex.FindAllString(line, -1)
		switch {
		case len(found) == 1:
			calibrationValue, _ = strconv.Atoi(found[0] + found[0])
		case len(found) >= 2:
			calibrationValue, _ = strconv.Atoi(found[0] + found[len(found)-1])
		}
		totalCalibration += calibrationValue
	}
	return totalCalibration
}

func Problem2() int {
	var calibrationValue int
	var totalCalibration int

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		regex, _ := regexp.Compile(`(one|two|three|four|five|six|seven|eight|nine)|(\d){1}`)
		found := regex.FindAllString(line, -1)
		switch {
		case len(found) == 1:
			calibrationValue, _ = strconv.Atoi(Convert(found[0]) + Convert(found[0]))
		case len(found) >= 2:
			calibrationValue, _ = strconv.Atoi(Convert(found[0]) + Convert(found[len(found)-1]))
		}
		totalCalibration += calibrationValue
	}
	return totalCalibration
}

func Convert(input string) string {
	switch input {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return input
	}
}
