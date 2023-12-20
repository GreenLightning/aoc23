package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Left, Right string
}

func main() {
	lines := readLines("input.txt")

	instructions := lines[0]
	nodes := make(map[string]Node)

	nodeRegex := regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		matches := nodeRegex.FindStringSubmatch(line)
		nodes[matches[1]] = Node{
			Left:  matches[2],
			Right: matches[3],
		}
	}

	{
		fmt.Println("--- Part One ---")

		current := "AAA"
		steps := 0
		for ; current != "ZZZ"; steps++ {
			if instructions[steps%len(instructions)] == 'L' {
				current = nodes[current].Left
			} else {
				current = nodes[current].Right
			}
		}

		fmt.Println(steps)
	}

	{
		fmt.Println("--- Part Two ---")

		// This solution requires that each start leads to a repeating
		// sequence with exactly one end node and the end node must appear at
		// the position that is equal to the period of the repetition.
		// E.g. A -> P -> Q -> R -> S -> Z -> T -> U
		//                     ^                   |
		//                     +-------------------+
		// Here Z appears at position 5 and then again at 10, 15, etc.
		// (Ignoring LR instructions in this example.)

		result := 1

		for current := range nodes {
			if strings.HasSuffix(current, "A") {
				steps := 0
				for ; !strings.HasSuffix(current, "Z"); steps++ {
					if instructions[steps%len(instructions)] == 'L' {
						current = nodes[current].Left
					} else {
						current = nodes[current].Right
					}
				}
				result = lcm(result, steps)
			}
		}

		fmt.Println(result)
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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
