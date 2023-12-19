package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card int

const (
	CardNone  Card = 0
	CardJoker Card = 1
	CardJack  Card = 11
	CardQueen Card = 12
	CardKing  Card = 13
	CardAce   Card = 14
)

func (c Card) String() string {
	return string(c.Rune())
}

func (c Card) Rune() rune {
	switch {
	case c == CardNone:
		return '_'
	case c == CardJoker:
		return 'J'
	case c >= 2 && c <= 9:
		return '0' + rune(c)
	case c == 10:
		return 'T'
	case c == CardJack:
		return 'J'
	case c == CardQueen:
		return 'Q'
	case c == CardKing:
		return 'K'
	case c == CardAce:
		return 'A'
	default:
		return '#'
	}
}

func parseCard(r rune) Card {
	if r >= '2' && r <= '9' {
		return Card(r - '0')
	}
	switch r {
	case 'T':
		return Card(10)
	case 'J':
		return CardJack
	case 'Q':
		return CardQueen
	case 'K':
		return CardKing
	case 'A':
		return CardAce
	default:
		return CardNone
	}
}

type Hand [5]Card

func (hand Hand) String() string {
	runes := make([]rune, len(hand))
	for i, c := range hand {
		runes[i] = c.Rune()
	}
	return string(runes)
}

func parseHand(s string) (hand Hand, err error) {
	if len(s) != len(hand) {
		err = errors.New("invalid input")
		return
	}
	for i, r := range s {
		hand[i] = parseCard(r)
		if hand[i] == CardNone {
			err = errors.New("invalid input")
		}
	}
	return
}

type Rank int

const (
	RankHighCard Rank = iota
	RankOnePair
	RankTwoPair
	RankThreeOfAKind
	RankFullHouse
	RankFourOfAKind
	RankFiveOfAKind
)

func rankNormal(hand Hand) Rank {
	var counts [15]int
	for _, card := range hand {
		counts[card]++
	}
	slices.Sort(counts[:])
	slices.Reverse(counts[:])
	return rankCounts(counts[0], counts[1])
}

func rankJoker(hand Hand) Rank {
	counts := make([]int, 15)
	for _, card := range hand {
		counts[card]++
	}
	jokers := counts[CardJoker]
	counts = counts[2:]
	slices.Sort(counts)
	slices.Reverse(counts)
	a := rankCounts(counts[0]+jokers, counts[1])
	first := max(counts[0], counts[1]+jokers)
	second := min(counts[0], counts[1]+jokers)
	b := rankCounts(first, second)
	return max(a, b)
}

func rankCounts(first, second int) Rank {
	if first == 5 {
		return RankFiveOfAKind
	} else if first == 4 {
		return RankFourOfAKind
	} else if first == 3 && second == 2 {
		return RankFullHouse
	} else if first == 3 {
		return RankThreeOfAKind
	} else if first == 2 && second == 2 {
		return RankTwoPair
	} else if first == 2 {
		return RankOnePair
	} else {
		return RankHighCard
	}
}

type Entry struct {
	Hand Hand
	Rank Rank
	Bid  int
}

func compare(a, b Entry) int {
	if a.Rank != b.Rank {
		return int(a.Rank - b.Rank)
	}
	for i := range a.Hand {
		if a.Hand[i] != b.Hand[i] {
			return int(a.Hand[i] - b.Hand[i])
		}
	}
	return 0
}

func (e Entry) String() string {
	return fmt.Sprintf("%v %v", e.Hand, e.Bid)
}

func main() {
	lines := readLines("input.txt")

	var entries []Entry
	for _, line := range lines {
		before, after, _ := strings.Cut(line, " ")
		hand, err := parseHand(before)
		check(err)
		bid := toInt(after)
		entries = append(entries, Entry{Hand: hand, Bid: bid})
	}

	{
		fmt.Println("--- Part One ---")

		for i, entry := range entries {
			entries[i].Rank = rankNormal(entry.Hand)
		}

		slices.SortFunc(entries, compare)

		fmt.Println(score(entries))
	}

	{
		fmt.Println("--- Part Two ---")

		for i := range entries {
			for j, card := range entries[i].Hand {
				if card == CardJack {
					entries[i].Hand[j] = CardJoker
				}
			}
			entries[i].Rank = rankJoker(entries[i].Hand)
		}

		slices.SortFunc(entries, compare)

		fmt.Println(score(entries))
	}
}

func score(entries []Entry) int {
	var total int
	for i, entry := range entries {
		total += (i + 1) * entry.Bid
	}
	return total
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
