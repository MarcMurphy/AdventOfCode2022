package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	file, _ := os.Open("./input.txt")

	scoreMap := make(map[string]int)
	scoreMap["A X"] = 1 + 3 // Rock + draw
	scoreMap["A Y"] = 2 + 6 // Paper + win
	scoreMap["A Z"] = 3 + 0 // Scissors + lose
	scoreMap["B X"] = 1 + 0 // Rock + lose
	scoreMap["B Y"] = 2 + 3 // Paper + draw
	scoreMap["B Z"] = 3 + 6 // Scissors + win
	scoreMap["C X"] = 1 + 6 // Rock + Win
	scoreMap["C Y"] = 2 + 0 // Paper + lose
	scoreMap["C Z"] = 3 + 3 // Scissors + Draw
	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalScore += scoreMap[line]
	}
	fmt.Println(totalScore)
}

func main() {
	file, _ := os.Open("./input.txt")

	scoreMap := make(map[string]int)
	scoreMap["A Y"] = 1 + 3 // Rock + draw
	scoreMap["B Y"] = 2 + 3 // Paper + draw
	scoreMap["C Y"] = 3 + 3 // Scissors + Draw

	scoreMap["A X"] = 3 + 0 // Scissors + lose
	scoreMap["B X"] = 1 + 0 // Rock + lose
	scoreMap["C X"] = 2 + 0 // Paper + lose

	scoreMap["A Z"] = 2 + 6 // Paper + win
	scoreMap["B Z"] = 3 + 6 // Scissors + win
	scoreMap["C Z"] = 1 + 6 // Rock + win

	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalScore += scoreMap[line]
	}
	fmt.Println(totalScore)
}
