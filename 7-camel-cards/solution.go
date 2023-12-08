package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

const CARDS = "AKQT98765432J"

type HandBid struct {
	hand string
	bid  int
}

func (h HandBid) Rank() int {
	kinds := make(map[string]int)
	for i := 0; i < len(h.hand); i++ {
		c := string(h.hand[i])
		kinds[c]++
	}

	switch {
	case fiveOfKind(kinds):
		return 7
	case fourOfKind(kinds):
		return 6
	case fullHouse(kinds):
		return 5
	case threeOfKind(kinds):
		return 4
	case twoPair(kinds):
		return 3
	case onePair(kinds):
		return 2
	case highCard(kinds):
		return 1
	}
	return 0
}

func hasAtLeast(n int, kinds map[string]int) bool {
	for k, v := range kinds {
		if k == "J" && n == v {
			//log.Print(kinds)
			return true
		} else if k != "J" && (v >= n-kinds["J"]) {
			if n == 5 {
				log.Print(kinds)
			}
			return true
		}
	}
	return false
}

func fiveOfKind(kinds map[string]int) bool {
	return hasAtLeast(5, kinds)
}

func fourOfKind(kinds map[string]int) bool {
	return hasAtLeast(4, kinds)
}

func fullHouse(kinds map[string]int) bool {
	hasThree := false
	hasPair := false
	jokersToBorrow := kinds["J"]
	for k, v := range kinds {
		if k == "J" {
			continue
		}
		if v+jokersToBorrow >= 3 {
			hasThree = true
			jokersToBorrow -= max(3-v, 0)
			continue
		}
		if v+jokersToBorrow >= 2 {
			hasPair = true
			jokersToBorrow -= max(2-v, 0)
			continue
		}
	}
	return hasThree && hasPair
}

func threeOfKind(kinds map[string]int) bool {
	return hasAtLeast(3, kinds)
}

func twoPair(kinds map[string]int) bool {
	var pairs int
	jokersToBorrow := kinds["J"]
	for k, v := range kinds {
		if k == "J" {
			continue
		}
		if v+jokersToBorrow >= 2 {
			pairs++
			jokersToBorrow -= max(2-v, 0)
		}
	}
	return pairs == 2
}

func onePair(kinds map[string]int) bool {
	return hasAtLeast(2, kinds)
}

func highCard(kinds map[string]int) bool {
	return len(kinds) == 5
}

func parseHands(input []string) []HandBid {
	var hands []HandBid
	for _, line := range input {
		split := strings.Split(line, " ")
		hand := split[0]
		bidStr := split[1]
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			log.Fatal(err)
		}
		hands = append(hands, HandBid{hand, bid})
	}
	return hands
}

func sortHands(hands []HandBid) {
	sort.SliceStable(hands, func(i, j int) bool {
		rankI := hands[i].Rank()
		rankJ := hands[j].Rank()
		if rankI == rankJ {
			for k := 0; k < 5; k++ {
				indexI := strings.Index(CARDS, string(hands[i].hand[k]))
				indexJ := strings.Index(CARDS, string(hands[j].hand[k]))
				if indexI > indexJ {
					return true
				} else if indexI < indexJ {
					return false
				}
			}
		}
		return rankI < rankJ
	})
}

func solve(input []string) {
	hands := parseHands(input)
	sortHands(hands)

	var sum int
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}
	log.Printf("PART 2: Sum of winnings is %v", sum)
}

func main() {
	input := readInput("7-camel-cards/input.txt")
	solve(input)
}
