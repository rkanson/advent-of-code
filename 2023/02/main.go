package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rkanson/advent-of-code/2023/utils"
)

var digits = regexp.MustCompile(`(\d{1,2})`)
var red = regexp.MustCompile(`(\d{1,2}).?(red)`)
var blue = regexp.MustCompile(`(\d{1,2}).?(blue)`)
var green = regexp.MustCompile(`(\d{1,2}).?(green)`)

func main() {
	fmt.Println("Problem 1")
	fmt.Println(Problem1())
	fmt.Println("Problem 2")
	fmt.Println(Problem2())
}

func Problem1() int {
	index := 1
	totalValue := 0
	utils.LoopInput("input.txt", func(line string) {
		okRed := GameCanExist(red.FindAllString(line, -1), "red")
		okBlue := GameCanExist(blue.FindAllString(line, -1), "blue")
		okGreen := GameCanExist(green.FindAllString(line, -1), "green")
		if okRed && okBlue && okGreen {
			totalValue += index
		}
		index += 1
	})
	return totalValue
}

func GameCanExist(found []string, pattern string) bool {
	maxes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, v := range found {
		digit := digits.FindAllString(v, -1)
		value, _ := strconv.Atoi(digit[0])
		if value > maxes[pattern] {
			return false
		}
	}
	return true
}

func Problem2() int {
	total := 0
	utils.LoopInput("input.txt", func(line string) {
		minRed := MinToExist(red.FindAllString(line, -1))
		minBlue := MinToExist(blue.FindAllString(line, -1))
		minGreen := MinToExist(green.FindAllString(line, -1))
		total += minRed * minBlue * minGreen
	})
	return total
}

func MinToExist(found []string) int {
	highest := 0
	for _, v := range found {
		digit := digits.FindAllString(v, -1)
		value, _ := strconv.Atoi(digit[0])
		if value > highest {
			highest = value
		}
	}
	return highest
}
