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
		grid := makeGrid(lines)
		north(grid)
		fmt.Println(load(grid))
	}

	{
		fmt.Println("--- Part Two ---")

		const limit = 1000000000

		grid := makeGrid(lines)

		cycles := 0
		cache := make(map[string]int)
		cache[key(grid)] = cycles

		for {
			cycle(grid)
			cycles++

			k := key(grid)
			if old, ok := cache[k]; ok {
				length := old - cycles
				cycles += (limit - cycles) / length * length
				break
			}
			cache[k] = cycles
		}

		for ; cycles < limit; cycles++ {
			cycle(grid)
		}

		fmt.Println(load(grid))
	}
}

func makeGrid(lines []string) (grid [][]byte) {
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	return
}

func key(grid [][]byte) string {
	var builder strings.Builder
	builder.Grow(len(grid) * len(grid[0]))
	for _, line := range grid {
		builder.Write(line)
	}
	return builder.String()
}

func load(grid [][]byte) (result int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'O' {
				result += len(grid) - y
			}
		}
	}
	return
}

func cycle(grid [][]byte) {
	north(grid)
	west(grid)
	south(grid)
	east(grid)
}

func north(grid [][]byte) {
	for {
		moved := false

		for y := 0; y+1 < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == '.' && grid[y+1][x] == 'O' {
					grid[y][x] = 'O'
					grid[y+1][x] = '.'
					moved = true
				}
			}
		}

		if !moved {
			break
		}
	}
}

func west(grid [][]byte) {
	for {
		moved := false

		for x := 0; x+1 < len(grid[0]); x++ {
			for y := 0; y < len(grid); y++ {
				if grid[y][x] == '.' && grid[y][x+1] == 'O' {
					grid[y][x] = 'O'
					grid[y][x+1] = '.'
					moved = true
				}
			}
		}

		if !moved {
			break
		}
	}
}

func south(grid [][]byte) {
	for {
		moved := false

		for y := len(grid) - 1; y-1 >= 0; y-- {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == '.' && grid[y-1][x] == 'O' {
					grid[y][x] = 'O'
					grid[y-1][x] = '.'
					moved = true
				}
			}
		}

		if !moved {
			break
		}
	}
}

func east(grid [][]byte) {
	for {
		moved := false

		for x := len(grid[0]) - 1; x-1 >= 0; x-- {
			for y := 0; y < len(grid); y++ {
				if grid[y][x] == '.' && grid[y][x-1] == 'O' {
					grid[y][x] = 'O'
					grid[y][x-1] = '.'
					moved = true
				}
			}
		}

		if !moved {
			break
		}
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
