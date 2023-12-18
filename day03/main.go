package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	sum := 0
	symbolNumbers := make(map[int][]int)
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			length := 0
			for isDigit(at(lines, x+length, y)) {
				length++
			}
			if length == 0 {
				continue
			}
			number := toInt(line[x : x+length])
			symbols := 0
			for i := -1; i <= length; i++ {
				symbols += update(lines, symbolNumbers, number, x+i, y-1)
				symbols += update(lines, symbolNumbers, number, x+i, y)
				symbols += update(lines, symbolNumbers, number, x+i, y+1)
			}
			if symbols != 0 {
				sum += number
			}
			x += length
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")
		sum := 0
		for _, list := range symbolNumbers {
			if len(list) == 2 {
				sum += list[0] * list[1]
			}
		}
		fmt.Println(sum)
	}
}

func update(lines []string, symbolNumbers map[int][]int, number int, x, y int) int {
	c := at(lines, x, y)
	if c == '*' {
		index := y*len(lines) + x
		symbolNumbers[index] = append(symbolNumbers[index], number)
	}
	isSymbol := c != '.' && !isDigit(c)
	if isSymbol {
		return 1
	}
	return 0
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func at(lines []string, x, y int) byte {
	if y >= 0 && y < len(lines) {
		if x >= 0 && x < len(lines[y]) {
			return lines[y][x]
		}
	}
	return '.'
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
