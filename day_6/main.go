package main

import (
	"bufio"
	"fmt"
	"os"
)

func isUnique(s string) bool {
	duplicateMap := make(map[rune]bool)
	for _, c := range s {
		if duplicateMap[c] {
			return false
		}
		duplicateMap[c] = true
	}
	return true
}

func part1() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataStream := scanner.Text()
		for i := 0; i < len(dataStream)-4; i++ {
			characters := dataStream[i : i+4]
			if isUnique(characters) {
				fmt.Println(i + 4)
				break
			}
		}
	}
}

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataStream := scanner.Text()
		for i := 0; i < len(dataStream)-14; i++ {
			characters := dataStream[i : i+14]
			if isUnique(characters) {
				fmt.Println(i + 14)
				break
			}
		}
	}
}
