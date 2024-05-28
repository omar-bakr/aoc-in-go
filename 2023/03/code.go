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
	// when you're ready to do part 2, remove this "not implemented" block
	lines := strings.Split(input, "\n")
	sum := 0

	if part2 {
		for lineIndex, line := range lines {
			i := 0
			for i < len(line) {
				if line[i] == '*' {
					sum += getPartNumbers(lineIndex, lines, i)
				}
				i++
			}
		}
		return sum
	}
	for lineIndex, line := range lines {
		i := 0
		for i < len(line) {
			if isDigit(line[i]) {
				start := i
				for i < len(line) && isDigit(line[i]) {
					i++
				}
				end := i
				// Convert the substring to an integer
				number, _ := strconv.Atoi(line[start:end])
				// fmt.Println(number)
				// fmt.Println(lineIndex, line, start, end)
				// fmt.Println(start-1 >= 0 && line[start-1] != '.')
				// fmt.Println(end < len(line) && line[end] != '.')
				// fmt.Println(checkDiagonal(lineIndex, lines, start, end))

				if start-1 >= 0 && line[start-1] != '.' || end < len(line) && line[end] != '.' || checkDiagonal(lineIndex, lines, start, end) {
					sum += number

				}

			} else {
				i++
			}
		}
	}
	return sum
}
func getPartnumber(line string, index int) int {
	start := index
	end := index
	for start >= 0 {
		if start-1 >= 0 && isDigit(line[start-1]) {
			start--
		} else {
			break
		}
	}

	for end < len(line) {
		if end+1 < len(line) && isDigit(line[end]) {
			end++
		} else {
			break
		}
	}
	if end < len(line) && isDigit(line[end]) {
		end++
	}
	num, _ := strconv.Atoi(line[start:end])
	return num
}

func getPartNumbers(lineIndex int, lines []string, symbolIndex int) int {
	multp := 1
	count := 0

	muliplyWithPartNumber := func(str string, index int) {
		multp *= getPartnumber(str, index)
		count++
	}

	processAdjacent := func(str string, index int) {
		if index >= 0 && isDigit(str[index]) && index < len(str) {
			muliplyWithPartNumber(str, index)
		}
	}

	if lineIndex-1 >= 0 {
		prev := lines[lineIndex-1]
		cur := isDigit(prev[symbolIndex])
		if cur {
			muliplyWithPartNumber(prev, symbolIndex)
		} else {
			processAdjacent(prev, symbolIndex-1)
			processAdjacent(prev, symbolIndex+1)
		}
	}

	if lineIndex+1 < len(lines) {
		next := lines[lineIndex+1]
		cur := isDigit(next[symbolIndex])
		if cur {
			muliplyWithPartNumber(next, symbolIndex)
		} else {
			processAdjacent(next, symbolIndex-1)
			processAdjacent(next, symbolIndex+1)
		}
	}

	line := lines[lineIndex]
	processAdjacent(line, symbolIndex-1)
	processAdjacent(line, symbolIndex+1)

	if count == 2 {
		return multp
	}
	return 0
}

func isDigit(char byte) bool {
	// return char >= '0' && char <= '9'
	return unicode.IsDigit(rune(char))
}

func checkDiagonal(lineIndex int, lines []string, start int, end int) bool {
	line := lines[lineIndex]
	if lineIndex-1 >= 0 {
		prev := lines[lineIndex-1]
		for i := start; i < end; i++ {
			if !isDigit(prev[i]) && prev[i] != '.' {
				return true
			}
		}
		if start-1 >= 0 && prev[start-1] != '.' || end < len(line) && prev[end] != '.' {
			return true
		}
	}

	if lineIndex+1 < len(lines) {
		next := lines[lineIndex+1]
		for i := start; i < end; i++ {
			if !isDigit(next[i]) && next[i] != '.' {
				return true
			}
		}

		if start-1 >= 0 && next[start-1] != '.' || end < len(line) && next[end] != '.' {
			return true
		}
	}
	return false
}
