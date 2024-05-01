package main

import (
	"strconv"
	"strings"
	"unicode"

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
	lines := strings.Split(input, "\n")
	sum := 0
	numsMap := getNumsMap()

	if part2 {
		for _, line := range lines {
			firstDigit := getFirstDigit(line, numsMap) * 10
			lastDigit := getLastDigit(line, numsMap)
			sum += (firstDigit + lastDigit)
		}
		return sum
	}

	for _, line := range lines {
		sum += getNum(line)
	}

	return sum
}

func getNumsMap() map[string]string {
	return map[string]string{
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
}

func getNum(word string) int {
	left := 0
	right := len(word) - 1
	for left < right {
		if !unicode.IsDigit(rune(word[left])) {
			left++
		}

		if !unicode.IsDigit(rune(word[right])) {
			right--
		}

		if unicode.IsDigit(rune(word[left])) && unicode.IsDigit(rune(word[right])) {
			break
		}
	}
	firstDigit := int(word[left] - '0')
	lastDigit := int(word[right] - '0')
	return firstDigit*10 + lastDigit
}

func getFirstDigit(word string, nums map[string]string) int {
	for i, char := range word {
		if unicode.IsDigit(rune(char)) {
			return int(char - '0')
		}

		for j := 3; j <= 5 && i+j <= len(word); j++ {
			digit := word[i : i+j]
			if value, exsists := nums[digit]; exsists {
				num, _ := strconv.Atoi(value)
				return num
			}

		}

	}
	return -1
}

func getLastDigit(word string, nums map[string]string) int {
	for i := len(word) - 1; i >= 0; i-- {
		char := word[i]

		if unicode.IsDigit(rune(char)) {
			return int(char - '0')
		}

		for j := 3; j <= 5 && i-j+1 >= 0; j++ {
			digit := word[i-j+1 : i+1]
			if value, exists := nums[digit]; exists {
				num, _ := strconv.Atoi(value)
				return num
			}
		}
	}
	return -1
}
