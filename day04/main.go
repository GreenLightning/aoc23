package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	var sum int
	copies := make([]int, len(lines))
	for index, line := range lines {
		_, line, ok := strings.Cut(line, ":")
		if !ok {
			panic("invalid input")
		}
		first, second, ok := strings.Cut(line, "|")
		if !ok {
			panic("invalid input")
		}
		winning := slices.DeleteFunc(strings.Split(first, " "), func(s string) bool { return s == "" })
		actual := slices.DeleteFunc(strings.Split(second, " "), func(s string) bool { return s == "" })

		var count int
		for _, number := range actual {
			if slices.Contains(winning, number) {
				count++
			}
		}

		if count != 0 {
			score := 1
			for i := 1; i < count; i++ {
				score *= 2
			}
			sum += score
		}

		copies[index]++
		for i := 1; i <= count; i++ {
			copies[index+i] += copies[index]
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")
		var total int
		for _, count := range copies {
			total += count
		}
		fmt.Println(total)
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

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
