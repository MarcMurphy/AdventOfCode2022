package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		elfAStart, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
		elfAEnd, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])
		elfBStart, _ := strconv.Atoi(strings.Split(parts[1], "-")[0])
		elfBEnd, _ := strconv.Atoi(strings.Split(parts[1], "-")[1])

		if (elfAStart <= elfBStart && elfAEnd >= elfBEnd) || (elfBStart <= elfAStart && elfBEnd >= elfAEnd) {
			counter++
		}
	}
	fmt.Println(counter)
}

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		elfAStart, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
		elfAEnd, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])
		elfBStart, _ := strconv.Atoi(strings.Split(parts[1], "-")[0])
		elfBEnd, _ := strconv.Atoi(strings.Split(parts[1], "-")[1])

		if elfAStart <= elfBEnd && elfBStart <= elfAEnd {
			counter++
		}
	}
	fmt.Println(counter)
}
