package main

import (
	"fmt"
	"github.com/ymiseddy/AdventOfCode2024/shared"
	"os"
)

var title string = "Advent of Code 2024, Day 4"

func Puzzle1(lines []string) int64 {
	total := int64(0)
	return total
}

func Puzzle2(lines []string) int64 {
	total := int64(0)
	return total
}

func main() {
	fmt.Println(title)
	// Read all text from stdin
	lines := shared.ReadLinesFromStream(os.Stdin)
	fmt.Println(lines)

	res1 := Puzzle1(lines)
	fmt.Println("Puzzle 1 result: ", res1)
	res2 := Puzzle2(lines)
	fmt.Println("Puzzle 2 result: ", res2)
}
