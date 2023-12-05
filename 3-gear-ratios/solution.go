package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

const SYMBOLS = "!@#$%^&*-+=/"

func readInput(input string) []string {
	data, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

func getAdjCoords(i, j int) [][2]int {
	return [][2]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}
}

func isSymbol(c string) bool {
	return strings.Contains(SYMBOLS, c)
}

type Number struct {
	startX int
	endX   int
	y      int
	value  int
}

func (n Number) isAt(coord [2]int) bool {
	return coord[0] >= n.startX && coord[0] < n.endX && coord[1] == n.y
}

type Gear struct {
	x int
	y int
}

func part1(input []string) {
	maxLine := len(input) - 1
	maxChar := len(input[0]) - 1
	var sum int
	var curNum string
	isPartNum := false
	for i, line := range input {
		for j, c := range line {
			if !unicode.IsDigit(c) {
				if curNum != "" {
					if isPartNum {
						curNumInt, err := strconv.Atoi(curNum)
						if err != nil {
							log.Fatal(err)
						}
						sum += curNumInt
					}

					curNum = ""
					isPartNum = false
				}
			} else {
				curNum += string(input[i][j])
				if isPartNum {
					continue
				}
				coordsToCheck := getAdjCoords(i, j)
				for _, coord := range coordsToCheck {
					k := max(min(coord[0], maxLine), 0)
					l := max(min(coord[1], maxChar), 0)
					if isSymbol(string(input[k][l])) {
						isPartNum = true
						break
					}
				}
			}
		}
		if curNum != "" {
			if isPartNum {
				curNumInt, err := strconv.Atoi(curNum)
				if err != nil {
					log.Fatal(err)
				}
				sum += curNumInt
			}
			curNum = ""
		}
		isPartNum = false
	}
	log.Printf("PART 1: Sum of the symbols is: %d", sum)
}

func isGearSymbol(c string) bool {
	return c == "*"
}

func part2(input []string) {
	var nums []Number
	var gears []Gear
	var curNum string
	for i, line := range input {
		for j, c := range line {
			if c == '*' {
				gears = append(gears, Gear{j, i})
			}
			if !unicode.IsDigit(c) {
				if curNum != "" {
					curNumInt, err := strconv.Atoi(curNum)
					if err != nil {
						log.Fatal(err)
					}
					nums = append(nums, Number{j - len(curNum), j, i, curNumInt})
					curNum = ""
				}
			} else {
				curNum += string(input[i][j])
			}
		}
		if curNum != "" {
			curNumInt, err := strconv.Atoi(curNum)
			if err != nil {
				log.Fatal(err)
			}
			num := Number{len(line) - len(curNum), len(line), i, curNumInt}
			nums = append(nums, num)
			curNum = ""
		}
	}
	var sum int
	for _, gear := range gears {
		coordsToCheck := getAdjCoords(gear.x, gear.y)
		var numbersForGear []Number
		for _, coord := range coordsToCheck {
			for _, num := range nums {
				if num.isAt(coord) && !slices.Contains(numbersForGear, num) {
					numbersForGear = append(numbersForGear, num)
				}
			}
		}
		if len(numbersForGear) == 2 {
			sum += (numbersForGear[0].value * numbersForGear[1].value)
		}
	}
	log.Printf("PART 2: Sum of gear ratios is: %v", sum)
}

func main() {
	input := readInput("input.txt")
	part1(input)
	part2(input)
}
