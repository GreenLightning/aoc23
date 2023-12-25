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

		var sum int
		for _, line := range lines {
			first, second, ok := strings.Cut(line, " ")
			if !ok {
				panic("invalid input")
			}

			sum += arrangements(repeat(first, "?", 5), repeat(second, ",", 5))
		}

		fmt.Println(sum)
	}
}

type Key struct {
	Description string
	List        string
}

var cache = make(map[Key]int)

func arrangements(description, list string) int {
	if list == "" {
		for _, part := range description {
			if part == '#' {
				return 0
			}
		}
		return 1
	}

	key := Key{description, list}
	if result, ok := cache[key]; ok {
		return result
	}

	first, restList, _ := strings.Cut(list, ",")
	number := toInt(first)

	var result int

	spring := []byte(description)
	for start := 0; start+number <= len(spring); start++ {
		ok := true
		for i := start; i < start+number; i++ {
			if spring[i] == '.' {
				ok = false
				break
			}
		}
		if start+number < len(spring) {
			if spring[start+number] == '#' {
				ok = false
			}
		}

		if ok {
			restSpring := spring[start+number:]
			if len(restSpring) != 0 {
				restSpring = restSpring[1:]
			}
			result += arrangements(string(restSpring), restList)
		}

		if spring[start] == '#' {
			break
		}
	}

	cache[key] = result
	return result
}

func repeat(s string, sep string, count int) string {
	return strings.TrimSuffix(strings.Repeat(s+sep, count), sep)
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
