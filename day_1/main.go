package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func part1() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	currentSum := 0
	biggestSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentSum > biggestSum {
				biggestSum = currentSum
			}
			currentSum = 0
		}
		intval, _ := strconv.Atoi(line)
		currentSum += intval
	}
	fmt.Println(biggestSum)
}

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	currentSum := 0
	var allSums []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			allSums = append(allSums, currentSum)
			currentSum = 0
		}
		intval, _ := strconv.Atoi(line)
		currentSum += intval
	}
	sort.Slice(allSums, func(i, j int) bool {
		return allSums[j] < allSums[i]
	})
	totalCals := allSums[0] + allSums[1] + allSums[2]
	fmt.Println(totalCals)
}
