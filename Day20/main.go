package main

import (
	"fmt"
	"os"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day "

func ParseInput(lines []string) (shared.Coord, shared.Coord, [][]rune) {
	grid := make([][]rune, len(lines))
	start := shared.Coord{X: -1, Y: -1}
	end := shared.Coord{X: -1, Y: -1}
	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		for x, char := range line {
			if char == 'S' {
				start = shared.Coord{X: x, Y: y}
			}
			if char == 'E' {
				end = shared.Coord{X: x, Y: y}
			}
			grid[y][x] = char
		}
	}

	return start, end, grid
}

// This is waay overkill for this problem - themaze doesn't branch.
func LeastCostPath(start shared.Coord, end shared.Coord, grid [][]rune, debug bool) []shared.Coord {
	var total int = 0
	visited := make(map[shared.Coord]struct{})
	visitedList := []shared.Coord{start}
	queue := []shared.Coord{start}
	parent := make(map[shared.Coord]shared.Coord)
	parent[start] = shared.Coord{X: -1, Y: -1}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			total++
			continue
		}

		for _, adjacent := range current.Adjacencies() {
			if adjacent.X < 0 || adjacent.Y < 0 || adjacent.X >= len(grid) || adjacent.Y >= len(grid[0]) {
				continue
			}
			if adjacent == end {
				parent[end] = current
				path := []shared.Coord{end}
				path = append(path, current)
				walk := current
				for {
					walk = parent[walk]
					path = append(path, walk)
					if walk == start {
						break
					}
				}
				return path
			}

			if grid[adjacent.Y][adjacent.X] == '#' {
				continue
			}

			if _, ok := visited[adjacent]; !ok {
				visited[adjacent] = struct{}{}
				visitedList = append(visitedList, adjacent)
				parent[adjacent] = current
				queue = append(queue, adjacent)
			}
		}
	}

	shared.ShowGridStep(grid, debug, nil, 0)
	return nil
}

func Puzzle1(lines []string) int {
	total := 0
	start, end, grid := ParseInput(lines)
	lcp := LeastCostPath(start, end, grid, true)
	lcpDistance := map[shared.Coord]int{}
	reversedPath := []shared.Coord{}
	for n, p := range lcp {
		lcpDistance[p] = n
		reversedPath = append(reversedPath, p)
	}
	total = FindCheats(reversedPath, grid, lcpDistance, 2)
	return total
}

type StartEnd struct {
	Start shared.Coord
	End   shared.Coord
}

func FindCheats(path []shared.Coord, grid [][]rune, distances map[shared.Coord]int, cheatCount int) int {

	cheatMap := make(map[StartEnd]int)
	for _, p := range path {
		pEndDistance := distances[p]
		for _, k := range path {
			kEndDistance := distances[k]
			if p == k {
				continue
			}
			distance := shared.ManhattanDistance(p, k)
			if distance > cheatCount {
				continue
			}

			delta := (kEndDistance - (pEndDistance + distance))
			if delta >= 100 {
				cheatMap[StartEnd{p, k}] = delta
			}
		}
	}

	/*
		// Just so we can print this out to compare with samples.
		distancesSavedCount := map[int]int{}
		distancesSavedList := []int{}
		for _, v := range cheatMap {
			distance := v
			if _, ok := distancesSavedCount[distance]; !ok {
				distancesSavedList = append(distancesSavedList, distance)
				distancesSavedCount[distance] = 1
			} else {
				distancesSavedCount[distance]++
			}
		}

		// Sort distancesSavedList
		sort.Ints(distancesSavedList)

		for _, s := range distancesSavedList {
			fmt.Printf("Cheat %d: %d\n", s, distancesSavedCount[s])
		}
	*/

	return len(cheatMap)
}

func Puzzle2(lines []string) int {
	total := 0
	start, end, grid := ParseInput(lines)
	lcp := LeastCostPath(start, end, grid, true)
	lcpDistance := map[shared.Coord]int{}
	reversedPath := []shared.Coord{}
	for n, p := range lcp {
		lcpDistance[p] = n
		reversedPath = append(reversedPath, p)
	}
	total = FindCheats(reversedPath, grid, lcpDistance, 20)
	return total
}

func main() {
	fmt.Println(title)
	// Read all text from stdin

	lines := shared.ReadLinesFromStream(os.Stdin)

	res1 := Puzzle1(lines)
	fmt.Println("Puzzle 1 result: ", res1)
	res2 := Puzzle2(lines)
	fmt.Println("Puzzle 2 result: ", res2)
}
