package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type State struct {
	x, y   int
	area   int
	border int
}

// See: https://stackoverflow.com/a/67115906
func (s *State) Update(dir string, steps int) {
	switch dir {
	case "U", "3":
		s.y -= steps
	case "D", "1":
		s.y += steps
	case "L", "2":
		s.x -= steps
		s.area -= steps * s.y
	case "R", "0":
		s.x += steps
		s.area += steps * s.y
	}
	s.border += steps
}

func (s *State) Area() int {
	// The area only accounts for the border on one side,
	// so we have to add the other half of the border.
	// Also, we start with one square already dug out.
	return abs(s.area) + s.border/2 + 1
}

func main() {
	lines := readLines("input.txt")

	var part1, part2 State

	regex := regexp.MustCompile(`^([UDLR]) (\d+) \(#([0-9a-f]{5})([0-3])\)$`)
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)

		{
			dir := matches[1]
			steps := toInt(matches[2])
			part1.Update(dir, steps)
		}

		{
			steps, err := strconv.ParseInt(matches[3], 16, 0)
			check(err)
			dir := matches[4]
			part2.Update(dir, int(steps))
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(part1.Area())
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(part2.Area())
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
