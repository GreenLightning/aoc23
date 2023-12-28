package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	grid := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(solve(grid, 0, 3))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(solve(grid, 4, 10))
	}
}

func solve(grid []string, minSteps, maxSteps int) int {
	height, width := len(grid), len(grid[0])

	seen := make(map[PriorityItem]bool)

	var queue PriorityQueue
	queue.Push(&PriorityItem{})

	for !queue.Empty() {
		item := queue.Pop()

		if item.X == width-1 && item.Y == height-1 && item.StraightSteps >= minSteps {
			return item.Loss
		}

		if item.X-1 >= 0 && item.LastX != item.X-1 && (item.StraightSteps >= minSteps || item.LastY == item.Y) && (item.StraightSteps < maxSteps || item.LastX == item.X) {
			newItem := &PriorityItem{
				X:             item.X - 1,
				Y:             item.Y,
				LastX:         item.X,
				LastY:         item.Y,
				StraightSteps: item.StraightSteps + 1,
			}
			if item.LastX == item.X {
				newItem.StraightSteps = 1
			}
			if !seen[*newItem] {
				seen[*newItem] = true
				newItem.Loss = item.Loss + int(grid[newItem.Y][newItem.X]-'0')
				newItem.Cost = newItem.Loss + (width - 1 - newItem.X) + (height - 1 - newItem.Y)
				queue.Push(newItem)
			}
		}

		if item.X+1 < width && item.LastX != item.X+1 && (item.StraightSteps >= minSteps || item.LastY == item.Y) && (item.StraightSteps < maxSteps || item.LastX == item.X) {
			newItem := &PriorityItem{
				X:             item.X + 1,
				Y:             item.Y,
				LastX:         item.X,
				LastY:         item.Y,
				StraightSteps: item.StraightSteps + 1,
			}
			if item.LastX == item.X {
				newItem.StraightSteps = 1
			}
			if !seen[*newItem] {
				seen[*newItem] = true
				newItem.Loss = item.Loss + int(grid[newItem.Y][newItem.X]-'0')
				newItem.Cost = newItem.Loss + (width - 1 - newItem.X) + (height - 1 - newItem.Y)
				queue.Push(newItem)
			}
		}

		if item.Y-1 >= 0 && item.LastY != item.Y-1 && (item.StraightSteps >= minSteps || item.LastX == item.X) && (item.StraightSteps < maxSteps || item.LastY == item.Y) {
			newItem := &PriorityItem{
				X:             item.X,
				Y:             item.Y - 1,
				LastX:         item.X,
				LastY:         item.Y,
				StraightSteps: item.StraightSteps + 1,
			}
			if item.LastY == item.Y {
				newItem.StraightSteps = 1
			}
			if !seen[*newItem] {
				seen[*newItem] = true
				newItem.Loss = item.Loss + int(grid[newItem.Y][newItem.X]-'0')
				newItem.Cost = newItem.Loss + (width - 1 - newItem.X) + (height - 1 - newItem.Y)
				queue.Push(newItem)
			}
		}

		if item.Y+1 < height && item.LastY != item.Y+1 && (item.StraightSteps >= minSteps || item.LastX == item.X) && (item.StraightSteps < maxSteps || item.LastY == item.Y) {
			newItem := &PriorityItem{
				X:             item.X,
				Y:             item.Y + 1,
				LastX:         item.X,
				LastY:         item.Y,
				StraightSteps: item.StraightSteps + 1,
			}
			if item.LastY == item.Y {
				newItem.StraightSteps = 1
			}
			if !seen[*newItem] {
				seen[*newItem] = true
				newItem.Loss = item.Loss + int(grid[newItem.Y][newItem.X]-'0')
				newItem.Cost = newItem.Loss + (width - 1 - newItem.X) + (height - 1 - newItem.Y)
				queue.Push(newItem)
			}
		}
	}

	return -1
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

type PriorityItem struct {
	X             int
	Y             int
	LastX         int
	LastY         int
	StraightSteps int
	Loss          int
	Cost          int
	Index         int
}

type PriorityStorage []*PriorityItem

func (s PriorityStorage) Len() int {
	return len(s)
}

func (s PriorityStorage) Less(i, j int) bool {
	return s[i].Cost < s[j].Cost
}

func (s PriorityStorage) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].Index, s[j].Index = i, j
}

func (s *PriorityStorage) Push(x interface{}) {
	item := x.(*PriorityItem)
	item.Index = len(*s)
	*s = append(*s, item)
}

func (s *PriorityStorage) Pop() interface{} {
	len := len(*s)
	item := (*s)[len-1]
	item.Index = -1
	*s = (*s)[:len-1]
	return item
}

type PriorityQueue struct {
	storage PriorityStorage
}

func (q *PriorityQueue) Len() int {
	return len(q.storage)
}

func (q *PriorityQueue) Empty() bool {
	return len(q.storage) == 0
}

func (q *PriorityQueue) Push(item *PriorityItem) {
	heap.Push(&q.storage, item)
}

func (q *PriorityQueue) Pop() *PriorityItem {
	return heap.Pop(&q.storage).(*PriorityItem)
}

func (q *PriorityQueue) Update(item *PriorityItem) {
	heap.Fix(&q.storage, item.Index)
}
