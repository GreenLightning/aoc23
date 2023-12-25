package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")

		var sum int
		for _, line := range lines {
			first, second, ok := strings.Cut(line, " ")
			if !ok {
				panic("invalid input")
			}

			sum += arrangements(first, second)
		}

		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println()
	}
}

func arrangements(first, second string) int {
	spring := []byte(first)
	counts := arrayToInt(strings.Split(second, ","))

	var openDefects, openChoices int
	for _, count := range counts {
		openDefects += count
	}
	for _, part := range spring {
		if part == '#' {
			openDefects--
		} else if part == '?' {
			openChoices++
		}
	}

	return arrangementsImpl(spring, counts, 0, openDefects, openChoices)
}

func arrangementsImpl(spring []byte, counts []int, index int, openDefects, openChoices int) int {
	if index >= len(spring) {
		return 1
	}
	if spring[index] == '.' || spring[index] == '#' {
		return arrangementsImpl(spring, counts, index+1, openDefects, openChoices)
	}
	var result int
	if openChoices > openDefects {
		spring[index] = '.'
		if consistent(spring, counts, index) {
			result += arrangementsImpl(spring, counts, index+1, openDefects, openChoices-1)
		}
	}
	if openDefects > 0 {
		spring[index] = '#'
		if consistent(spring, counts, index) {
			result += arrangementsImpl(spring, counts, index+1, openDefects-1, openChoices-1)
		}
	}
	spring[index] = '?'
	return result
}

func consistent(spring []byte, counts []int, index int) (result bool) {
	countIndex := 0
	segment := 0
	for _, part := range spring {
		if part == '#' {
			if countIndex >= len(counts) {
				return false
			}
			segment++
		} else if part == '.' {
			if segment > 0 {
				if segment != counts[countIndex] {
					return false
				}
				countIndex++
				segment = 0
			}
		} else {
			if segment > 0 {
				if segment > counts[countIndex] {
					return false
				}
			}
			return true
		}
	}
	if segment > 0 {
		if segment != counts[countIndex] {
			return false
		}
		countIndex++
		segment = 0
	}
	return countIndex == len(counts)
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func arrayToInt(input []string) (output []int) {
	output = make([]int, len(input))
	for i, text := range input {
		output[i] = toInt(text)
	}
	return output
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
