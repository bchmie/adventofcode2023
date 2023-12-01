package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const PUZZLE_INPUT = "input.txt"
const NUMERIC = "one|two|three|four|five|six|seven|eight|nine|[0-9]"

var numeralToDigit = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func findFirst(line string, r *regexp.Regexp) string {
	return r.FindString(line)
}

func findLast(line string, r *regexp.Regexp) (matched string) {
	for i := len(line) - 1; i >= 0; i-- {
		matched = r.FindString(line[i:])
		if matched != "" {
			return
		}
	}
	return
}

func asDigit(numeral string) string {
	matched, err := regexp.Match(`\d`, []byte(numeral))
	if err != nil {
		panic(err)
	}
	if matched {
		return numeral
	} else {
		return numeralToDigit[numeral]
	}
}

func main() {
	f, err := os.Open(PUZZLE_INPUT)
	if err != nil {
		log.Fatal("Unable to open puzzle input")
	}
	defer closeFile(f)

	r, err := regexp.Compile(NUMERIC)
	if err != nil {
		log.Fatal(err)
	}

	var calibrationValues []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		first := asDigit(findFirst(line, r))
		last := asDigit(findLast(line, r))
		calibrationValues = append(calibrationValues, first+last)
	}

	var sum int
	for _, value := range calibrationValues {
		i, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		sum += i
	}
	log.Printf("Sum of all calibration values is %d", sum)
}
