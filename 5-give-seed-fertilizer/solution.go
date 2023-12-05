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

type Mapping struct {
	destinationStart int
	sourceStart      int
	mapRange         int
}

func (m Mapping) CanMap(source int) bool {
	return source >= m.sourceStart && source <= m.sourceStart+m.mapRange
}

func (m Mapping) Map(source int) int {
	if !m.CanMap(source) {
		log.Fatal("Invalid Map call to mapping")
	}
	offset := source - m.sourceStart
	return m.destinationStart + offset
}

type RangeMap struct {
	mappings []Mapping
}

func (r RangeMap) Map(source int) int {
	for _, mapping := range r.mappings {
		if mapping.CanMap(source) {
			return mapping.Map(source)
		}
	}
	return source
}

type ChainMap struct {
	maps []RangeMap
}

func (c ChainMap) Map(source int) int {
	for _, rangeMap := range c.maps {
		source = rangeMap.Map(source)
	}
	return source
}

func pluckSeeds(input string) []int {
	input = strings.Trim(input, "seeds: ")
	var seeds []int
	seedsStr := strings.Split(input, " ")
	for _, seedStr := range seedsStr {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func lineToIntSlice(line string) []int {
	splitted := strings.Split(line, " ")
	var nums []int
	for _, numStr := range splitted {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func readMap(mapName string, input []string) RangeMap {
	var mappingsStart int
	for i, line := range input {
		if strings.HasPrefix(line, mapName) {
			mappingsStart = i + 1
			break
		}
	}
	if mappingsStart == 0 {
		log.Fatal("Could not find map start")
	}

	var mappings []Mapping
	for _, line := range input[mappingsStart:] {
		if line == "" {
			break
		}
		mapping := lineToIntSlice(line)
		mappings = append(mappings, Mapping{mapping[0], mapping[1], mapping[2]})
	}
	return RangeMap{mappings}
}

func readAlmanac(input []string) ChainMap {
	seedsToSoil := readMap("seed-to-soil", input)
	soilToFertilizer := readMap("soil-to-fertilizer", input)
	fertilizerToWater := readMap("fertilizer-to-water", input)
	waterToLight := readMap("water-to-light", input)
	lightToTemperature := readMap("light-to-temperature", input)
	temperatureToHumidity := readMap("temperature-to-humidity", input)
	humidityToLocation := readMap("humidity-to-location", input)
	almanac := ChainMap{
		[]RangeMap{
			seedsToSoil,
			soilToFertilizer,
			fertilizerToWater,
			waterToLight,
			lightToTemperature,
			temperatureToHumidity,
			humidityToLocation,
		},
	}
	return almanac
}

func readMinLocation(seeds []int, almanac ChainMap) int {
	var minLocation int
	for _, seed := range seeds {
		location := almanac.Map(seed)
		if minLocation == 0 {
			minLocation = location
		} else {
			minLocation = min(minLocation, location)
		}
	}
	return minLocation
}

func part1(input []string) {
	seeds := pluckSeeds(input[0])
	almanac := readAlmanac(input)
	minLocation := readMinLocation(seeds, almanac)
	log.Printf("PART 1: Min location is: %v", minLocation)
}

func normalizeSeedRanges(seedRanges []int) []SeedRange {
	var seeds []SeedRange
	for i := 0; i < len(seedRanges); i += 2 {
		start := seedRanges[i]
		length := seedRanges[i+1]
		seeds = append(seeds, SeedRange{start, length})
	}
	return seeds
}

type SeedRange struct {
	start  int
	length int
}

func part2(input []string) {
	seedRanges := pluckSeeds(input[0])
	almanac := readAlmanac(input)
	seeds := normalizeSeedRanges(seedRanges)
	log.Printf("Len of seeds: %d", len(seeds))
	var minLocation int
	for _, seed := range seeds {
		for i := seed.start; i < seed.start+seed.length; i++ {
			location := almanac.Map(i)
			if minLocation == 0 {
				minLocation = location
			} else {
				minLocation = min(minLocation, location)
			}
		}
	}
	log.Printf("PART 2: Min location is: %v", minLocation)
}

func main() {
	input := readInput("test_input.txt")
	part1(input)
	part2(input)
}
