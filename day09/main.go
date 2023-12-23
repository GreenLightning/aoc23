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

	var nextTotal int
	var previousTotal int
	for _, line := range lines {
		numbers := arrayToInt(strings.Split(line, " "))
		sequences := append([][]int{}, numbers)
		for {
			diff := make([]int, len(numbers)-1)
			allZero := true
			for i := range diff {
				diff[i] = numbers[i+1] - numbers[i]
				if diff[i] != 0 {
					allZero = false
				}
			}
			if allZero {
				break
			}
			sequences = append(sequences, diff)
			numbers = diff
		}

		next := 0
		previous := 0
		for i := len(sequences) - 1; i >= 0; i-- {
			seq := sequences[i]
			next = seq[len(seq)-1] + next
			previous = seq[0] - previous
		}
		nextTotal += next
		previousTotal += previous
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(nextTotal)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(previousTotal)
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
