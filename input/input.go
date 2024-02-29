package input

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

func ReadInput(filename string) []string {
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

func StringToIntSlice(input string) []int {
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