package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Lens struct {
	Label string
	Power int
}

func main() {
	input := readFile("input.txt")
	steps := strings.Split(input, ",")

	{
		fmt.Println("--- Part One ---")
		var sum int
		for _, step := range steps {
			sum += int(hash(step))
		}
		fmt.Println(sum)
	}

	{
		fmt.Println("--- Part Two ---")
		var boxes [256][]Lens

		for _, step := range steps {
			if label, ok := strings.CutSuffix(step, "-"); ok {
				bi := hash(label)
				box := boxes[bi]
				for li, lens := range box {
					if lens.Label == label {
						copy(box[li:], box[li+1:])
						box = box[:len(box)-1]
						boxes[bi] = box
						break
					}
				}
			} else if label, power, ok := strings.Cut(step, "="); ok {
				bi := hash(label)
				box := boxes[bi]
				found := false
				for li, lens := range box {
					if lens.Label == label {
						box[li].Power = toInt(power)
						found = true
						break
					}
				}
				if !found {
					boxes[bi] = append(box, Lens{
						Label: label,
						Power: toInt(power),
					})
				}
			}
		}

		var total int
		for bi, box := range boxes {
			for li, lens := range box {
				total += (bi + 1) * (li + 1) * lens.Power
			}
		}
		fmt.Println(total)
	}
}

func hash(s string) (result byte) {
	for _, b := range []byte(s) {
		result += b
		result *= 17
	}
	return
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
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
