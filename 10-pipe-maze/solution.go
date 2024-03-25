package main

import (
	"adventofcode2023/input"
	"errors"
	"log"
	"os"
	"slices"
)

func findStart(maze []string) ([2]int, error) {
	for i, row := range maze {
		for j, symbol := range row {
			if symbol == 'S' {
				return [2]int{i, j}, nil
			}
		}
	}
	return [2]int{0, 0}, errors.New("could not find starting point")
}

func outOfBound(point [2]int, maze []string) bool {
	i := point[0]
	j := point[1]

	if i < 0 || i >= len(maze) {
		return true
	} else if j < 0 || j >= len(maze[0]) {
		return true
	} else {
		return false
	}
}

func neightbourNodes(point [2]int, maze []string) [2][2]int {
	symbol := maze[point[0]][point[1]]
	i, j := point[0], point[1]

	switch symbol {
	case '|':
		return [2][2]int{{i - 1, j}, {i + 1, j}}
	case '-':
		return [2][2]int{{i, j - 1}, {i, j + 1}}
	case 'L':
		return [2][2]int{{i - 1, j}, {i, j + 1}}
	case 'J':
		return [2][2]int{{i, j - 1}, {i - 1, j}}
	case '7':
		return [2][2]int{{i, j - 1}, {i + 1, j}}
	case 'F':
		return [2][2]int{{i, j + 1}, {i + 1, j}}
	case '.':
		return [2][2]int{{-1, -1}, {-1, -1}}
	default:
		return [2][2]int{{-1, -1}, {-1, -1}}
	}
}

func part1(maze []string) {
	start, err := findStart(maze)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start is at i=%d j=%d", start[0], start[1])

	i := start[0]
	j := start[1]
	adjecents := [4][2]int{
		{i - 1, j},
		{i, j + 1},
		{i + 1, j},
		{i, j - 1},
	}
	found := false
	for _, point := range adjecents {
		if found {
			break
		}
		if outOfBound(point, maze) {
			continue
		}
		path := make([][2]int, 0)
		path = append(path, [][2]int{start, point}...)

		steps := 1
		for !outOfBound(point, maze) {
			points := neightbourNodes(point, maze)
			if slices.Contains(path, points[0]) && slices.Contains(path, points[1]) {
				log.Printf("After %v steps we have a loop at %v ", steps, points)
				found = true
				break
			} else if slices.Contains(path, points[0]) {
				point = points[1]
			} else {
				point = points[0]
			}
			steps++
			path = append(path, point)
		}
		furthest := steps / 2
		if steps % 2 != 0 {
			furthest += 1
		}
		log.Printf("Furthest point is %v steps away", furthest)
	}
}

func main() {
	here, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	puzzleInput := input.ReadInput(here + "/input.txt")
	part1(puzzleInput)
}
