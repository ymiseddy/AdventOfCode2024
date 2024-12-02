package main

import (
	"fmt"
	"github.com/ymiseddy/AdventOfCode2024/shared"
	"os"
)

func Puzzle1(values [][]int64) {
	fmt.Println("Puzzle 1")
	for _, value := range values {
		fmt.Println(value)
	}
}

func Puzzle2(values [][]int64) {
	fmt.Println("Puzzle 1")
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	fmt.Println("Advent of Code 2024, Day 3")
	values, err := shared.ReadIntsFromStream(os.Stdin)
	if err != nil {
		panic(err)
	}
	Puzzle1(values)
	Puzzle2(values)
}
