package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 7"

func CanProduceResult(result int64, aggr int64, operands []int64) bool {
	if len(operands) == 0 {
		return result == aggr
	}

	if CanProduceResult(result, aggr+operands[0], operands[1:]) {
		return true
	} else if CanProduceResult(result, aggr*operands[0], operands[1:]) {
		return true
	}
	return false
}

func IntConcat(a, b int64) int64 {
	return int64(float64(a)*math.Pow(float64(10), 1+math.Floor(math.Log10(float64(b)))) + float64(b))
}

func CanProduceResult2(result int64, aggr int64, operands []int64) bool {
	if len(operands) == 0 {
		return result == aggr
	}

	if CanProduceResult2(result, IntConcat(aggr, operands[0]), operands[1:]) {
		return true
	} else if CanProduceResult2(result, aggr+operands[0], operands[1:]) {
		return true
	} else if CanProduceResult2(result, aggr*operands[0], operands[1:]) {
		return true
	}
	return false
}

func Puzzle1(lines []string) int64 {
	total := int64(0)
	re := regexp.MustCompile(`:?\s+`)
	for _, line := range lines {
		parts := re.Split(line, -1)

		// Parse parts to ints
		intParts := make([]int64, len(parts))
		for n, part := range parts {
			result, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				panic(err)
			}
			intParts[n] = result
		}
		result := intParts[0]
		operands := intParts[1:]
		if CanProduceResult(result, operands[0], operands[1:]) {
			total += result
		}
	}

	return total
}

func Puzzle2(lines []string) int64 {
	total := int64(0)
	re := regexp.MustCompile(`:?\s+`)
	for _, line := range lines {
		parts := re.Split(line, -1)
		// Parse parts to ints
		intParts := make([]int64, len(parts))
		for n, part := range parts {
			result, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				panic(err)
			}
			intParts[n] = result
		}
		result := intParts[0]
		operands := intParts[1:]
		if CanProduceResult2(result, operands[0], operands[1:]) {
			total += result
		}
	}

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
