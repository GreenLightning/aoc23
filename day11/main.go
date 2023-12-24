package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

func main() {
	lines := readLines("input.txt")

	var galaxies []Position
	for y, line := range lines {
		for x, pixel := range line {
			if pixel == '#' {
				galaxies = append(galaxies, Position{x, y})
			}
		}
	}

	var emptyRows []int
	for y, line := range lines {
		found := false
		for _, pixel := range line {
			if pixel == '#' {
				found = true
				break
			}
		}
		if !found {
			emptyRows = append(emptyRows, y)
		}
	}

	var emptyColumns []int
	for x := 0; x < len(lines[0]); x++ {
		found := false
		for _, line := range lines {
			if line[x] == '#' {
				found = true
				break
			}
		}
		if !found {
			emptyColumns = append(emptyColumns, x)
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(sumOfShortestPaths(galaxies, emptyRows, emptyColumns, 2))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(sumOfShortestPaths(galaxies, emptyRows, emptyColumns, 1e6))
	}
}

func sumOfShortestPaths(galaxies []Position, emptyRows, emptyColumns []int, expansionFactor int) int {
	var sum int
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			a, b := galaxies[i], galaxies[j]
			path := abs(a.x-b.x) + abs(a.y-b.y)
			for _, y := range emptyRows {
				if (a.y < y && y < b.y) || (a.y > y && y > b.y) {
					path += expansionFactor - 1
				}
			}
			for _, x := range emptyColumns {
				if (a.x < x && x < b.x) || (a.x > x && x > b.x) {
					path += expansionFactor - 1
				}
			}
			sum += path
		}
	}
	return sum
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
