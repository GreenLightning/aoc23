package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")

		sum := 0
		for _, line := range lines {
			first := -1
			last := 0
			for _, char := range line {
				if char >= '0' && char <= '9' {
					value := int(char - '0')
					if first == -1 {
						first = value
					}
					last = value
				}
			}
			sum += 10*first + last
		}

		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")

		digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

		sum := 0
		for _, line := range lines {
			first := -1
			last := 0

			for pos, char := range line {
				value := -1

				if char >= '0' && char <= '9' {
					value = int(char - '0')
				} else {
					for i, digit := range digits {
						if strings.HasPrefix(line[pos:], digit) {
							value = i + 1
						}
					}
				}

				if value != -1 {
					if first == -1 {
						first = value
					}
					last = value
				}
			}

			sum += 10*first + last
		}

		fmt.Println(sum)
	}
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
