package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID      int
	Samples []Sample
}

type Sample struct {
	Red, Green, Blue int
}

func main() {
	lines := readLines("input.txt")

	var games []Game
	for _, line := range lines {
		prefix, sampleList, ok := strings.Cut(line, ": ")
		if !ok {
			panic("missing :")
		}
		var game Game
		game.ID = toInt(strings.TrimPrefix(prefix, "Game "))
		for _, sampleString := range strings.Split(sampleList, "; ") {
			var sample Sample
			for _, entry := range strings.Split(sampleString, ", ") {
				count, color, ok := strings.Cut(entry, " ")
				if !ok {
					panic("invalid syntax")
				}
				switch color {
				case "red":
					sample.Red = toInt(count)
				case "green":
					sample.Green = toInt(count)
				case "blue":
					sample.Blue = toInt(count)
				default:
					panic("invalid syntax")
				}
			}
			game.Samples = append(game.Samples, sample)
		}
		games = append(games, game)
	}

	{
		fmt.Println("--- Part One ---")
		sum := 0
		for _, game := range games {
			ok := true
			for _, sample := range game.Samples {
				if sample.Red > 12 || sample.Green > 13 || sample.Blue > 14 {
					ok = false
					break
				}
			}
			if ok {
				sum += game.ID
			}
		}
		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")
		sum := 0
		for _, game := range games {
			var result Sample
			for _, sample := range game.Samples {
				result.Red = max(result.Red, sample.Red)
				result.Green = max(result.Green, sample.Green)
				result.Blue = max(result.Blue, sample.Blue)
			}
			sum += result.Red * result.Green * result.Blue
		}
		fmt.Println(sum)
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
