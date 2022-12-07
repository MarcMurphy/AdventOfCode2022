package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) PushItems(item []string) {
	s.items = append(s.items, item...)
}

func (s *Stack) Pop() string {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) PopItems(i int) []string {
	item := s.items[len(s.items)-i : len(s.items)]
	s.items = s.items[:len(s.items)-i]
	return item
}

func (s *Stack) Peek() string {
	return s.items[len(s.items)-1]
}

func part1() {
	/*file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]+")
	for scanner.Scan() {
		line := scanner.Text()
		numbers := re.FindAllString(line, -1)

		amountToMove, _ := strconv.Atoi(numbers[0])
		fromStack, _ := strconv.Atoi(numbers[1])
		destStack, _ := strconv.Atoi(numbers[2])

		for i := 0; i < amountToMove; i++ {
			item := stacks[fromStack].Pop()
			stacks[destStack].Push(item)
		}
	}

	for i := 1; i < 10; i++ {
		fmt.Println(stacks[i].Peek())
	}*/
}

func main() {
	firstStack := Stack{}
	firstStack.PushItems([]string{"D", "B", "J", "V"})

	secondStack := Stack{}
	secondStack.PushItems([]string{"P", "V", "B", "W", "R", "D", "F"})

	thirdStack := Stack{}
	thirdStack.PushItems([]string{"R", "G", "F", "L", "D", "C", "W", "Q"})

	fourthStack := Stack{}
	fourthStack.PushItems([]string{"W", "J", "P", "M", "L", "N", "D", "B"})

	fifthStack := Stack{}
	fifthStack.PushItems([]string{"H", "N", "B", "P", "C", "S", "Q"})

	sixthStack := Stack{}
	sixthStack.PushItems([]string{"R", "D", "B", "S", "N", "G"})

	seventhStack := Stack{}
	seventhStack.PushItems([]string{"Z", "B", "P", "M", "Q", "F", "S", "H"})

	eigthStack := Stack{}
	eigthStack.PushItems([]string{"W", "L", "F"})

	ninthStack := Stack{}
	ninthStack.PushItems([]string{"S", "V", "F", "M", "R"})

	stacks := []Stack{Stack{}, firstStack, secondStack, thirdStack, fourthStack, fifthStack, sixthStack, seventhStack, eigthStack, ninthStack}
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]+")
	for scanner.Scan() {
		line := scanner.Text()
		numbers := re.FindAllString(line, -1)

		amountToMove, _ := strconv.Atoi(numbers[0])
		fromStack, _ := strconv.Atoi(numbers[1])
		destStack, _ := strconv.Atoi(numbers[2])

		items := stacks[fromStack].PopItems(amountToMove)
		stacks[destStack].PushItems(items)
	}

	for i := 1; i < 10; i++ {
		fmt.Println(stacks[i].Peek())
	}
}
