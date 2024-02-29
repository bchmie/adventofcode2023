package main

import (
	"adventofcode2023/input"
	"log"
)

func derivative(original []int) []int {
	var derivative []int
	for i := 0; i < len(original)-1; i++ {
		derivative = append(derivative, original[i+1]-original[i])
	}
	return derivative
}

func flat(series []int) bool {
	for _, v := range series {
		if v != 0 {
			return false
		}
	}
	return true
}

func part1(in []string) {
	var seriesSlice [][]int
	for _, line := range in {
		series := input.StringToIntSlice(line)
		seriesSlice = append(seriesSlice, series)
	}
	var sumNextValues int
	for _, series := range seriesSlice {
		nexts := [][]int{series}
		next := derivative(series)
		for !flat(next) {
			nexts = append(nexts, next)
			next = derivative(next)
		}
		nexts = append(nexts, next)

		sum := 0
		for i := len(nexts) - 1; i > 0; i-- {
			current := nexts[i]
			lastIndex := len(current) - 1
			log.Printf("Current: %v. lastIndex: %v, adding %v", current, lastIndex, nexts[i-1][lastIndex+1])
			sum += nexts[i-1][lastIndex+1]
		}
		sumNextValues += sum
	}
	log.Print(sumNextValues)
}

func part2(in []string) {
	var seriesSlice [][]int
	for _, line := range in {
		series := input.StringToIntSlice(line)
		seriesSlice = append(seriesSlice, series)
	}
	var sumOfFirstValues int
	for _, series := range seriesSlice {
		nexts := [][]int{series}
		next := derivative(series)
		for !flat(next) {
			nexts = append(nexts, next)
			next = derivative(next)
		}
		nexts = append(nexts, next)

		for i := len(nexts) - 2; i >= 0; i-- {
			from := nexts[i+1][0]
			to := nexts[i][0]
			step := to - from
			if i == 0 {
				sumOfFirstValues += step
			} else {
				nexts[i] = append([]int{step}, nexts[i]...)
			}
		}
	}
	log.Print(sumOfFirstValues)
}

func main() {
	puzzleInput := input.ReadInput("9-mirage-maintenance/input.txt")
	//part1(input)
	part2(puzzleInput)
}
