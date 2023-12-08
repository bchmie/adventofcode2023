package main

import (
	"bufio"
	"log"
	"os"
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

func parseNodes(input []string) map[string][2]string {
	nodes := make(map[string][2]string)
	for _, line := range input {
		nodes[line[0:3]] = [2]string{line[7:10], line[12:15]}
	}
	return nodes
}

func getMove(moves string, step int) int {
	move := string(moves[step%len(moves)])
	switch move {
	case "L":
		return 0
	case "R":
		return 1
	default:
		return -1
	}
}

func traverse(moves string, first string, nodes map[string][2]string) int {
	current := first
	step := 0
	for current != "ZZZ" {
		move := getMove(moves, step)
		current = nodes[current][move]
		step += 1
	}
	return step
}

func part1(input []string) {
	moves := input[0]
	nodes := parseNodes(input[2:])
	steps := traverse(moves, "AAA", nodes)
	log.Printf("PART 1: %v", steps)

}

// GCD and LCM are stolen from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func traverseAll(moves string, nodes map[string][2]string) int {
	var startingNodes []string
	for k, _ := range nodes {
		if k[2] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}

	var toFinish []int
	for _, node := range startingNodes {
		current := node
		step := 0
		for current[2] != 'Z' {
			move := getMove(moves, step)
			current = nodes[current][move]
			step += 1
		}
		toFinish = append(toFinish, step)
	}
	return LCM(toFinish[0], toFinish[1], toFinish[2:]...)
}

func part2(input []string) {
	moves := input[0]
	nodes := parseNodes(input[2:])
	steps := traverseAll(moves, nodes)
	log.Printf("PART 2: %v", steps)
}

func main() {
	input := readInput("8-haunted-wasteland/input.txt")
	part1(input)
	part2(input)
}
