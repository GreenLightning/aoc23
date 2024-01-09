package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Workflow struct {
	Rules []Rule
}

type Rule struct {
	Property   string
	Comparison string
	Threshold  int
	Target     string
}

type Item map[string]int

type ItemRange map[string]Range

func (item ItemRange) Clone() ItemRange {
	clone := ItemRange{}
	for property, r := range item {
		clone[property] = r
	}
	return clone
}

func (item ItemRange) Count() int {
	result := 1
	for _, r := range item {
		result *= r.Count()
	}
	return result
}

type Range struct {
	Min int
	Max int
}

func (r *Range) Valid() bool {
	return r.Min <= r.Max
}

func (r *Range) Count() int {
	return r.Max - r.Min + 1
}

func main() {
	lines := readLines("input.txt")

	workflows := make(map[string]Workflow)
	items := make([]Item, 0)

	workflowRegex := regexp.MustCompile(`^(.+)\{(.+)\}$`)
	ruleRegex := regexp.MustCompile(`^(\w+)(<|>)(\d+):(\w+)$`)
	itemRegex := regexp.MustCompile(`^{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$`)

	index := 0
	for ; index < len(lines); index++ {
		line := lines[index]
		if line == "" {
			break
		}

		matches := workflowRegex.FindStringSubmatch(line)
		name := matches[1]
		ruleList := strings.Split(matches[2], ",")

		var workflow Workflow
		for _, ruleString := range ruleList {
			if strings.Contains(ruleString, ":") {
				matches := ruleRegex.FindStringSubmatch(ruleString)
				workflow.Rules = append(workflow.Rules, Rule{
					Property:   matches[1],
					Comparison: matches[2],
					Threshold:  toInt(matches[3]),
					Target:     matches[4],
				})
			} else {
				workflow.Rules = append(workflow.Rules, Rule{
					Target: ruleString,
				})
			}
		}

		workflows[name] = workflow
	}

	index++
	for ; index < len(lines); index++ {
		line := lines[index]

		matches := itemRegex.FindStringSubmatch(line)
		item := make(Item)
		item["x"] = toInt(matches[1])
		item["m"] = toInt(matches[2])
		item["a"] = toInt(matches[3])
		item["s"] = toInt(matches[4])

		items = append(items, item)
	}

	{
		fmt.Println("--- Part One ---")

		var sum int
		for _, item := range items {
			target := "in"
			for target != "A" && target != "R" {
				workflow := workflows[target]
				for _, rule := range workflow.Rules {
					value := item[rule.Property]
					if rule.Comparison == "<" {
						if !(value < rule.Threshold) {
							continue
						}
					} else if rule.Comparison == ">" {
						if !(value > rule.Threshold) {
							continue
						}
					}
					target = rule.Target
					break
				}
			}
			if target == "A" {
				for _, value := range item {
					sum += value
				}
			}
		}

		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")

		type Entry struct {
			Item   ItemRange
			Target string
		}

		var queue []Entry
		queue = append(queue, Entry{
			Item: ItemRange{
				"x": Range{1, 4000},
				"m": Range{1, 4000},
				"a": Range{1, 4000},
				"s": Range{1, 4000},
			},
			Target: "in",
		})

		var total int
		for len(queue) != 0 {
			entry := queue[0]
			queue = queue[1:]

			if entry.Target == "R" {
				continue
			}

			if entry.Target == "A" {
				total += entry.Item.Count()
				continue
			}

			workflow := workflows[entry.Target]
			for _, rule := range workflow.Rules {
				value := entry.Item[rule.Property]
				if rule.Comparison == "<" {
					if value.Min < rule.Threshold {
						item := entry.Item.Clone()
						item[rule.Property] = Range{
							Min: value.Min,
							Max: min(value.Max, rule.Threshold-1),
						}
						queue = append(queue, Entry{item, rule.Target})

						value.Min = rule.Threshold
						if !value.Valid() {
							break
						}
						entry.Item[rule.Property] = value
					}
				} else if rule.Comparison == ">" {
					if value.Max > rule.Threshold {
						item := entry.Item.Clone()
						item[rule.Property] = Range{
							Min: max(value.Min, rule.Threshold+1),
							Max: value.Max,
						}
						queue = append(queue, Entry{item, rule.Target})

						value.Max = rule.Threshold
						if !value.Valid() {
							break
						}
						entry.Item[rule.Property] = value
					}
				} else {
					queue = append(queue, Entry{entry.Item, rule.Target})
					break
				}
			}
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
