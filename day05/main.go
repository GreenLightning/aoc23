package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Entry struct {
	Destination int
	Source      int
	Length      int
}

type Map struct {
	From    string
	To      string
	Entries []Entry
}

type Range struct {
	Start  int
	Length int
}

func (r *Range) Cut(wanted int) (before, after Range) {
	if r.Length < wanted {
		wanted = r.Length
	}
	before = Range{r.Start, wanted}
	after = Range{r.Start + wanted, r.Length - wanted}
	return
}

func main() {
	lines := readLines("input.txt")

	seedString, ok := strings.CutPrefix(lines[0], "seeds: ")
	checkOk(ok)
	seeds := arrayToInt(strings.Split(seedString, " "))

	var maps []*Map
	var currentMap *Map

	titleRegex := regexp.MustCompile(`^(\w+)-to-(\w+) map:$`)
	entryRegex := regexp.MustCompile(`^(\d+) (\d+) (\d+)$`)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		if matches := titleRegex.FindStringSubmatch(line); matches != nil {
			currentMap = new(Map)
			currentMap.From = matches[1]
			currentMap.To = matches[2]
			maps = append(maps, currentMap)
			continue
		}

		if matches := entryRegex.FindStringSubmatch(line); matches != nil {
			currentMap.Entries = append(currentMap.Entries, Entry{
				Destination: toInt(matches[1]),
				Source:      toInt(matches[2]),
				Length:      toInt(matches[3]),
			})
			continue
		}

		panic("invalid input")
	}

	for _, m := range maps {
		slices.SortFunc(m.Entries, func(a, b Entry) int {
			return a.Source - b.Source
		})
	}

	{
		fmt.Println("--- Part One ---")

		kind := "seed"
		values := slices.Clone(seeds)

		for kind != "location" {
			var currentMap *Map
			for _, m := range maps {
				if m.From == kind {
					currentMap = m
					break
				}
			}

			kind = currentMap.To

			for i, value := range values {
				for _, entry := range currentMap.Entries {
					if value >= entry.Source && value < entry.Source+entry.Length {
						value += entry.Destination - entry.Source
						break
					}
				}
				values[i] = value
			}

		}

		fmt.Println(slices.Min(values))
	}

	{
		fmt.Println("--- Part Two ---")

		kind := "seed"
		var values []Range
		for i := 0; i+1 < len(seeds); i += 2 {
			values = append(values, Range{
				Start:  seeds[i],
				Length: seeds[i+1],
			})
		}

		for kind != "location" {
			var currentMap *Map
			for _, m := range maps {
				if m.From == kind {
					currentMap = m
					break
				}
			}

			var newValues []Range
			for _, value := range values {
				for _, entry := range currentMap.Entries {
					if value.Start < entry.Source {
						var before Range
						before, value = value.Cut(entry.Source - value.Start)
						newValues = append(newValues, before)
					}

					if value.Length == 0 {
						break
					}

					if value.Start >= entry.Source && value.Start < entry.Source+entry.Length {
						var before Range
						before, value = value.Cut(entry.Source + entry.Length - value.Start)
						before.Start += entry.Destination - entry.Source
						newValues = append(newValues, before)
					}

					if value.Length == 0 {
						break
					}
				}

				if value.Length > 0 {
					newValues = append(newValues, value)
				}
			}

			kind = currentMap.To
			values = newValues
		}

		fmt.Println(slices.MinFunc(values, func(a, b Range) int {
			return a.Start - b.Start
		}).Start)
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

func checkOk(ok bool) {
	if !ok {
		panic("invalid input")
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
