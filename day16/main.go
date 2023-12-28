package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

type Direction int

const (
	DirEast Direction = (1 << iota)
	DirNorth
	DirWest
	DirSouth
)

func main() {
	grid := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(test(grid, 0, 0, DirEast))
	}

	{
		fmt.Println("--- Part Two ---")
		var best int
		height, width := len(grid), len(grid[0])
		for y := 0; y < height; y++ {
			best = max(best, test(grid, 0, y, DirEast))
			best = max(best, test(grid, width-1, y, DirWest))
		}
		for x := 0; x < width; x++ {
			best = max(best, test(grid, x, 0, DirSouth))
			best = max(best, test(grid, x, height-1, DirNorth))
		}
		fmt.Println(best)
	}
}

func test(grid []string, startX, startY int, startDir Direction) int {
	beams := make([][]Direction, len(grid))
	for y, line := range grid {
		beams[y] = make([]Direction, len(line))
	}

	type Item struct {
		pos Position
		dir Direction
	}

	var queue []Item
	queue = append(queue, Item{Position{startX, startY}, startDir})

	push := func(pos Position, dir Direction) {
		switch dir {
		case DirEast:
			pos.x++
		case DirNorth:
			pos.y--
		case DirWest:
			pos.x--
		case DirSouth:
			pos.y++
		}
		queue = append(queue, Item{pos, dir})
	}

	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]

		if item.pos.y < 0 || item.pos.y >= len(grid) || item.pos.x < 0 || item.pos.x >= len(grid[0]) {
			continue
		}

		if beams[item.pos.y][item.pos.x]&item.dir != 0 {
			continue
		}

		beams[item.pos.y][item.pos.x] |= item.dir

		switch grid[item.pos.y][item.pos.x] {
		case '.':
			push(item.pos, item.dir)
		case '/':
			switch item.dir {
			case DirEast:
				push(item.pos, DirNorth)
			case DirNorth:
				push(item.pos, DirEast)
			case DirWest:
				push(item.pos, DirSouth)
			case DirSouth:
				push(item.pos, DirWest)
			}
		case '\\':
			switch item.dir {
			case DirEast:
				push(item.pos, DirSouth)
			case DirSouth:
				push(item.pos, DirEast)
			case DirWest:
				push(item.pos, DirNorth)
			case DirNorth:
				push(item.pos, DirWest)
			}
		case '|':
			switch item.dir {
			case DirEast, DirWest:
				push(item.pos, DirNorth)
				push(item.pos, DirSouth)
			case DirNorth, DirSouth:
				push(item.pos, item.dir)
			}
		case '-':
			switch item.dir {
			case DirNorth, DirSouth:
				push(item.pos, DirWest)
				push(item.pos, DirEast)
			case DirWest, DirEast:
				push(item.pos, item.dir)
			}
		}
	}

	var count int
	for _, line := range beams {
		for _, dirs := range line {
			if dirs != 0 {
				count++
			}
		}
	}
	return count
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
