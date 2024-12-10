package main

import (
	"fmt"
	"github.com/ymiseddy/AdventOfCode2024/shared"
	"os"
)

var title string = "Advent of Code 2024, Day 10"

var cardinalDirections [][]int = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type PointStruct struct {
	X int
	Y int
}

type PointSet map[PointStruct]struct{}

func Map(lines []string, value int, pathX, pathY int) (PointSet, int) {
	if value == 9 {
		set := PointSet{
			{X: pathX, Y: pathY}: {},
		}
		return set, 1
	}
	pointSet := PointSet{}
	raiting := int(0)
	for _, dir := range cardinalDirections {
		x := pathX + dir[0]
		y := pathY + dir[1]
		if x < 0 || x >= len(lines) {
			continue
		}
		if y < 0 || y >= len(lines[x]) {
			continue
		}

		val := int(lines[y][x] - '0')
		if val < 0 || val > 9 {
			continue
		}
		// fmt.Printf("Invalid value at (%d, %d): %c\n", x, y, lines[x][y])
		if val == value+1 {
			newPoints, subRaiting := Map(lines, value+1, x, y)
			raiting += subRaiting
			for point := range newPoints {
				pointSet[point] = struct{}{}
			}
		} else {
		}
	}
	return pointSet, raiting
}

func Puzzle1(lines []string) int64 {
	score := int64(0)
	for y, line := range lines {
		for x, char := range line {
			if char == '0' {
				pointSet, _ := Map(lines, 0, x, y)
				// fmt.Printf("Point set at (%d, %d): %v\n", x, y, len(pointSet))
				score += int64(len(pointSet))
			}
		}
	}

	return score
}

func Puzzle2(lines []string) int64 {
	score := int64(0)
	for y, line := range lines {
		for x, char := range line {
			if char == '0' {
				_, raiting := Map(lines, 0, x, y)
				// fmt.Printf("Point set at (%d, %d): %v\n", x, y, raiting)
				score += int64(raiting)
			}
		}
	}

	return score
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
