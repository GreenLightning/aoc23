package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("input.txt")

	var patterns [][][]byte
	var pattern [][]byte
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = nil
		} else {
			pattern = append(pattern, []byte(line))
		}
	}
	if len(pattern) != 0 {
		patterns = append(patterns, pattern)
	}

	{
		fmt.Println("--- Part One ---")
		var sum int
		for _, pattern := range patterns {
			pos, _ := reflections(pattern)
			sum += pos
		}
		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")
		var sum int
		for _, pattern := range patterns {
			original, _ := reflections(pattern)

		loop:
			for y := 0; y < len(pattern); y++ {
				for x := 0; x < len(pattern[y]); x++ {
					flip(pattern, x, y)
					new, count := reflections(pattern)
					flip(pattern, x, y)
					if new != 0 && new != original {
						sum += new
						if count > 1 {
							sum -= original
						}
						break loop
					}
				}
			}
		}
		fmt.Println(sum)
	}
}

func reflections(pattern [][]byte) (pos int, count int) {
ysearch:
	for y := 1; y < len(pattern); y++ {
		for i := 0; y-(i+1) >= 0 && y+i < len(pattern); i++ {
			for x := 0; x < len(pattern[0]); x++ {
				if pattern[y-(i+1)][x] != pattern[y+i][x] {
					continue ysearch
				}
			}
		}
		pos += 100 * y
		count++
	}
xsearch:
	for x := 1; x < len(pattern[0]); x++ {
		for i := 0; x-(i+1) >= 0 && x+i < len(pattern[0]); i++ {
			for y := 0; y < len(pattern); y++ {
				if pattern[y][x-(i+1)] != pattern[y][x+i] {
					continue xsearch
				}
			}
		}
		pos += x
		count++
	}
	return
}

func flip(pattern [][]byte, x, y int) {
	if pattern[y][x] == '.' {
		pattern[y][x] = '#'
	} else {
		pattern[y][x] = '.'
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
