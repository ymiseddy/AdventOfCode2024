package main

import (
	"fmt"
	"github.com/ymiseddy/AdventOfCode2024/shared"
	"os"
	"slices"
)

var title string = "Advent of Code 2024, Day 1"

func puzzle01(entries [][]int64) int64 {
	var listL []int64
	var listR []int64
	for _, fields := range entries {
		listL = append(listL, fields[0])
		listR = append(listR, fields[1])
	}

	// Sort the lists
	slices.Sort(listL)
	slices.Sort(listR)

	// Calculate the sum of the differences
	var sum int64 = 0
	for i := 0; i < len(listL); i++ {
		distance := listL[i] - listR[i]
		distance = max(distance, -distance)
		sum += distance
	}

	return sum
}

func puzzle02(entries [][]int64) int64 {
	var listL []int64
	mapR := make(map[int64]int64)
	for _, fields := range entries {
		listL = append(listL, fields[0])
		numR := fields[1]
		if _, ok := mapR[numR]; ok {
			mapR[numR]++
		} else {
			mapR[numR] = 1
		}
	}
	sum := int64(0)
	for _, numL := range listL {
		if count, ok := mapR[numL]; ok {
			sum += numL * count
		}
	}
	return sum
}

func main() {
	fmt.Println(title)
	inputs, err := shared.ReadIntsFromStream(os.Stdin)
	if err != nil {
		panic(err)
	}
	code := puzzle01(inputs)
	fmt.Printf("Puzzle 1 result %d\n", code)

	code = puzzle02(inputs)
	fmt.Printf("Puzzle 2 result %d\n", code)
}
