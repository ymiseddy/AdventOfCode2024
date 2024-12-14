package main

import (
	"fmt"
	"os"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 14"

func Puzzle1(lines []string) int64 {
	var total int64 = 0
	return total
}

func Puzzle2(lines []string) int64 {
	var total int64 = 0
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
