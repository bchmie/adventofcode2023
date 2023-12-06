package main

import (
	"bufio"
	"log"
	"os"
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

type Race struct {
	duration       int
	recordDistance int
}

func stringToIntSlice(input string) []int {
	var result []int
	strSlice := strings.Split(input, " ")
	for _, str := range strSlice {
		if str == "" {
			continue
		}
		integer, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, integer)
	}
	return result
}

func readRaces(input []string) []Race {
	times := stringToIntSlice(strings.Trim(strings.Trim(input[0], "Time:"), " "))
	distances := stringToIntSlice(strings.Trim(strings.Trim(input[1], "Distance:"), " "))
	numRaces := len(times)

	var races []Race
	for i := 0; i < numRaces; i++ {
		races = append(races, Race{times[i], distances[i]})
	}
	return races
}

func part1(input []string) {
	races := readRaces(input)
	total := 1
	for _, race := range races {
		var waysToBeat int
		for i := 1; i < race.duration; i++ {
			distance := i * (race.duration - i)
			if distance > race.recordDistance {
				waysToBeat += 1
			}
		}
		total = total * waysToBeat
	}

	log.Printf("PART 1: Product of ways to win: %v", total)
}

func main() {
	input := readInput("input.txt")
	part1(input)
	//part2(input) - I've just edited the input.txt file
}
