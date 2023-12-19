package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")
	timeString := strings.TrimPrefix(lines[0], "Time:")
	recordString := strings.TrimPrefix(lines[1], "Distance:")

	{
		fmt.Println("--- Part One ---")

		times := arrayToInt(deleteEmptyStrings(strings.Split(timeString, " ")))
		records := arrayToInt(deleteEmptyStrings(strings.Split(recordString, " ")))

		if len(times) != len(records) {
			panic("invalid input")
		}

		total := 1
		for index := range times {
			total *= computeWaysToWin(times[index], records[index])
		}

		fmt.Println(total)
	}

	{
		fmt.Println("--- Part Two ---")
		time := toInt(strings.ReplaceAll(timeString, " ", ""))
		record := toInt(strings.ReplaceAll(recordString, " ", ""))
		fmt.Println(computeWaysToWin(time, record))
	}
}

func computeWaysToWin(time, record int) int {
	count := 0
	distance := 0
	for pressed := 1; pressed < time; pressed++ {
		distance += (time - pressed) - (pressed - 1)
		if distance > record {
			count++
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

func deleteEmptyStrings(input []string) []string {
	return slices.DeleteFunc(input, func(s string) bool { return s == "" })
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
