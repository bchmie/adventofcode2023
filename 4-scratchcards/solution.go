package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func readInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to open puzzle input")
	}
	defer closeFile(f)
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type Scratchcard struct {
	id             string
	numbers        []int
	winningNumbers []int
}

func (s Scratchcard) MatchingNumbers() int {
	var sum int
	for _, number := range s.numbers {
		if slices.Contains(s.winningNumbers, number) {
			sum += 1
		}
	}
	return sum
}

func (s Scratchcard) Score() float64 {
	matchingNumbers := s.MatchingNumbers()
	switch matchingNumbers {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return math.Pow(2, float64(matchingNumbers-1))
	}
}

func stringSliceToIntSlice(input []string) (output []int) {
	for _, numberStr := range input {
		if numberStr == "" {
			continue
		}
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, number)
	}
	return
}

func part1(input []string) {
	var cards []Scratchcard
	for _, line := range input {
		card := strings.Split(line, ": ")
		id := strings.Trim(card[0], "Card ")
		allNumbers := strings.Split(card[1], " | ")
		numbersStr := strings.Split(allNumbers[0], " ")
		winningNumbersStr := strings.Split(allNumbers[1], " ")
		numbers := stringSliceToIntSlice(numbersStr)
		winningNumbers := stringSliceToIntSlice(winningNumbersStr)
		cards = append(cards, Scratchcard{id, numbers, winningNumbers})
	}

	var sum float64
	for _, card := range cards {
		score := card.Score()
		sum += score
	}
	log.Printf("Sum of the points is: %f", sum)
}

func main() {
	input := readInput("input.txt")
	part1(input)
}
