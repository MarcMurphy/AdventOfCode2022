package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	const positionEqualsPriority = "-abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	totalPriority := 0
	for scanner.Scan() {
		line := scanner.Text()
		part1 := line[0 : len(line)/2]
		part2 := line[len(line)/2:]

		matchesFound := ""

		for _, c := range part1 {
			if strings.Contains(part2, string(c)) && !strings.Contains(matchesFound, string(c)) {
				matchesFound += string(c)
			}
		}

		for _, c := range matchesFound {
			totalPriority += strings.Index(positionEqualsPriority, string(c))
		}
	}
	fmt.Println(totalPriority)
}

func main() {
	const positionEqualsPriority = "-abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	totalPriority := 0
	var groups []string
	for scanner.Scan() {
		groups = append(groups, scanner.Text())
		if len(groups) != 3 {
			continue
		}

		for _, c := range groups[0] {
			if strings.Contains(groups[1], string(c)) && strings.Contains(groups[2], string(c)) {
				totalPriority += strings.Index(positionEqualsPriority, string(c))
				break
			}
		}
		groups = groups[:0]
	}
	fmt.Println(totalPriority)
}
