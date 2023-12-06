package utils

import (
	"bufio"
	"os"
)

func LoopInput(path string, fn func(string)) {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fn(scanner.Text())
	}
}
