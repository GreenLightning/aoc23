package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

func northOf(p Position) Position {
	return Position{p.x, p.y - 1}
}

func southOf(p Position) Position {
	return Position{p.x, p.y + 1}
}

func westOf(p Position) Position {
	return Position{p.x - 1, p.y}
}

func eastOf(p Position) Position {
	return Position{p.x + 1, p.y}
}

// Lists of tiles that connect in a given direction:
const (
	north = "|LJ"
	south = "|7F"
	west  = "-J7"
	east  = "-LF"
)

func connects(list string, tile byte) bool {
	return strings.ContainsRune(list, rune(tile))
}

func main() {
	lines := readLines("input.txt")

	// Convert to array, so that we can update tiles.
	grid := make([][]byte, len(lines))
	for y, line := range lines {
		grid[y] = []byte(line)
	}

	at := func(p Position) byte {
		if p.y >= 0 && p.y < len(grid) {
			if p.x >= 0 && p.x < len(grid[p.y]) {
				return grid[p.y][p.x]
			}
		}
		return '.'
	}

	printGrid := func() {
		for _, line := range grid {
			fmt.Printf("%s\n", string(line))
		}
	}

	_ = printGrid

	var start Position

	// Find starting position.
	for y, line := range grid {
		for x, tile := range line {
			if tile == 'S' {
				start = Position{x, y}
			}
		}
	}

	// Determine the correct tile symbol for the starting position.
	{
		hasNorth := connects(south, at(northOf(start)))
		hasSouth := connects(north, at(southOf(start)))
		hasWest := connects(east, at(westOf(start)))
		hasEast := connects(west, at(eastOf(start)))
		if hasNorth && hasSouth {
			grid[start.y][start.x] = '|'
		} else if hasEast && hasWest {
			grid[start.y][start.x] = '-'
		} else if hasNorth && hasEast {
			grid[start.y][start.x] = 'L'
		} else if hasNorth && hasWest {
			grid[start.y][start.x] = 'J'
		} else if hasSouth && hasWest {
			grid[start.y][start.x] = '7'
		} else if hasSouth && hasEast {
			grid[start.y][start.x] = 'F'
		} else {
			panic("invalid input")
		}
	}

	// Flood-fill the loop, keeping track of the last distance,
	// because the algorithm will visit the furthest tile last.
	var distance int

	distances := make(map[Position]int)
	queue := make([]Position, 0)
	distances[start] = 0
	queue = append(queue, start)

	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		tile := at(p)
		distance = distances[p]
		if connects(north, tile) {
			next := northOf(p)
			if _, ok := distances[next]; !ok {
				distances[next] = distance + 1
				queue = append(queue, next)
			}
		}
		if connects(south, tile) {
			next := southOf(p)
			if _, ok := distances[next]; !ok {
				distances[next] = distance + 1
				queue = append(queue, next)
			}
		}
		if connects(west, tile) {
			next := westOf(p)
			if _, ok := distances[next]; !ok {
				distances[next] = distance + 1
				queue = append(queue, next)
			}
		}
		if connects(east, tile) {
			next := eastOf(p)
			if _, ok := distances[next]; !ok {
				distances[next] = distance + 1
				queue = append(queue, next)
			}
		}
	}

	// Remove junk pipes.
	for y, line := range grid {
		for x := range line {
			if _, ok := distances[Position{x, y}]; !ok {
				grid[y][x] = '.'
			}
		}
	}

	{
		// printGrid()
		fmt.Println("--- Part One ---")
		fmt.Println(distance)
	}

	// Mark and count enclosed tiles.
	var count int
	for _, line := range grid {
		for x, tile := range line {
			if tile != '.' {
				continue
			}
			// Walk to the edge of the map and count the number of crossings
			// with the pipe. If that number is even, we must be outside,
			// otherwise we must be inside the pipe.
			var crossings int
			for xx := x - 1; xx >= 0; xx-- {
				if strings.ContainsRune("|F7", rune(line[xx])) {
					crossings++
				}
			}
			if crossings%2 != 0 {
				line[x] = 'I'
				count++
			}
		}
	}

	{
		// printGrid()
		fmt.Println("--- Part Two ---")
		fmt.Println(count)
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
