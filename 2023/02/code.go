package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	sum := 0
	lines := strings.Split(input, "\n")
	colorMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	if part2 {
		for _, line := range lines {
			_, sets := parseLine(line)
			sum += getMinColorsPower(sets)
		}
		return sum
	}

	for _, line := range lines {
		game, sets := parseLine(line)
		if isValid(sets, colorMap) {
			sum += extractGameID(game)
		}
	}
	return sum
}

func parseLine(line string) (string, string) {
	parts := strings.SplitN(line, ":", 2)
	return parts[0], parts[1]
}

func extractGameID(game string) int {
	gameIDStr := strings.Fields(game)
	gameID, _ := strconv.Atoi(gameIDStr[len(gameIDStr)-1])
	return gameID
}

func isValid(sets string, colorMap map[string]int) bool {
	for _, set := range strings.Split(sets, ";") {
		for _, colorAndNum := range strings.Split(set, ",") {
			color, num := parseColorAndNum(colorAndNum)
			if value, exists := colorMap[color]; exists && num > value {
				return false
			}
		}
	}
	return true
}

func parseColorAndNum(colorAndNum string) (string, int) {
	parts := strings.Fields(colorAndNum)
	num, _ := strconv.Atoi(parts[0])
	color := parts[1]
	return color, num
}

func getMinColorsPower(sets string) int {
	maxMap := map[string]int{
		"red":   -1,
		"green": -1,
		"blue":  -1,
	}
	for _, set := range strings.Split(sets, ";") {
		for _, colorAndNum := range strings.Split(set, ",") {
			color, num := parseColorAndNum(colorAndNum)
			maxMap[color] = int(math.Max(float64(num), float64(maxMap[color])))
		}
	}

	product := 1

	for _, v := range maxMap {
		product *= v
	}
	return product
}
