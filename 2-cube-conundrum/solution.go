package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT = "input.txt"


var gameLimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func readInput(input string) []string {
	f, err := os.Open(INPUT)
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

func part2(input []string) {
	var sum int
	for _, line := range input {
		game := strings.Split(line, ": ")
		samples := strings.Split(game[1], "; ")

		fewestNumbers := make(map[string]int)
		for _, sample := range samples {
			for _, cubeType := range strings.Split(sample, ", ") {
				amountColor := strings.Split(cubeType, " ")
				color := amountColor[1]
				amount, err := strconv.Atoi(amountColor[0])
				if err != nil {
					log.Fatal(err)
				}
				if fewestNumbers[color] < amount {
					fewestNumbers[color] = amount
				}

			}
		}
		power := 1
		for _, v := range fewestNumbers {
			power *= v
		}
		sum += power
		clear(fewestNumbers)
	}
	log.Printf("PART 2: Sum of the powers is: %v", sum)
}

func part1(input []string) {
	var sum int
GAME:
	for _, line := range input {
		game := strings.Split(line, ": ")
		gameID := game[0][5:]
		samples := strings.Split(game[1], "; ")

		for _, sample := range samples {
			for _, cubeType := range strings.Split(sample, ", ") {
				amountColor := strings.Split(cubeType, " ")
				amount, err := strconv.Atoi(amountColor[0])
				if err != nil {
					log.Fatal(err)
				}
				overLimit := amount > gameLimit[amountColor[1]]
				if overLimit {
					continue GAME
				}
			}

		}
		gameIdInt, err := strconv.Atoi(gameID)
		if err != nil {
			log.Fatal(err)
		}
		sum += gameIdInt

	}
	log.Printf("PART 1: Sum of valid Game IDs is: %v", sum)
}

func main() {
	input := readInput(INPUT)
	part1(input)
	part2(input)
}
