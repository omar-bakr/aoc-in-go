package main

import (
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
		return "not implemented"
	}

	for _, line := range lines {
		lineSplitted := strings.Split(line, ":")
		game := lineSplitted[0]
		gameIdStr := strings.Split(game, " ")
		sets := lineSplitted[1]
		if isValid(sets, colorMap) {
			gameIdInt, _ := strconv.Atoi(gameIdStr[len(gameIdStr)-1])
			sum += gameIdInt
		}

	}
	return sum
}

func isValid(sets string, colorMap map[string]int) bool {
	for _, set := range strings.Split(sets, ";") {
		for _, colorAndNum := range strings.Split(set, ",") {
			colorAndNum := strings.TrimSpace(colorAndNum)
			numStr := strings.Split(colorAndNum, " ")[0]
			num, _ := strconv.Atoi(numStr)
			color := strings.Split(colorAndNum, " ")[1]
			if value, exsists := colorMap[color]; exsists && num > value {
				return false
			}
		}
	}
	return true
}
