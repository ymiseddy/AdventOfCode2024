package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 11"

func intDigits(n int64) int64 {
	return int64(1 + math.Floor(math.Log10(float64(n))))
}

type cacheResult struct {
	stone int64
	depth int
}

// Cache previously calculated stone/depth combinations otherwise we will
// be here for a long time.
var stoneMap map[cacheResult]int64 = make(map[cacheResult]int64)

func BlinkStone(stone int64, blinks int) int64 {
	point := cacheResult{stone, blinks}
	if val, ok := stoneMap[point]; ok {
		return val
	}

	if blinks == 0 {
		stoneMap[point] = 1
		return 1
	}
	digits := len(fmt.Sprint(stone))
	if stone == 0 {
		val := BlinkStone(1, blinks-1)
		stoneMap[point] = val
		return val
	}

	if digits%2 == 0 {
		half := digits / 2
		first := stone / int64(math.Pow(10, float64(half)))
		second := stone % int64(math.Pow(10, float64(half)))
		firstVal := BlinkStone(first, blinks-1)
		secondVal := BlinkStone(second, blinks-1)

		val := firstVal + secondVal
		stoneMap[point] = val
		return val
	}
	val := BlinkStone(stone*2024, blinks-1)
	stoneMap[point] = val
	return val
}

func Exec(lines []string, numberOfBlinks int) int64 {
	// Split second line into parts and convert to ints
	parts := strings.Fields(lines[0])
	stones := make([]int64, len(parts))
	for n, part := range parts {
		result, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			panic(err)
		}
		stones[n] = result
	}

	totalStones := int64(0)
	for _, stone := range stones {
		newStones := BlinkStone(stone, int(numberOfBlinks))
		totalStones += newStones
	}

	return totalStones
}

func Puzzle1(lines []string) int64 {
	return Exec(lines, 25)
}

func Puzzle2(lines []string) int64 {
	return Exec(lines, 75)
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
