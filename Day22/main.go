package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day "

func ParseInput(lines []string) []int {
	// Parse input
	output := []int{}
	for _, line := range lines {
		intLine, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		output = append(output, intLine)
	}

	return output
}

func GenerateNext(secretNumber int) int {

	// Part 1
	result1 := secretNumber * 64
	result2 := secretNumber ^ result1
	result3 := result2 % 16777216

	// Part 2
	result4 := result3 / 32
	result5 := result3 ^ result4
	result6 := result5 % 16777216

	// Part 3
	result7 := result6 * 2048
	result8 := result6 ^ result7
	result9 := result8 % 16777216

	return result9

}

func Puzzle1(lines []string) int {
	total := 0
	data := ParseInput(lines)
	for _, line := range data {
		a := line
		for i := 0; i < 2000; i++ {
			a = GenerateNext(a)
		}
		total += a
	}
	return total
}

type sequenceKey struct {
	v1 int
	v2 int
	v3 int
	v4 int
}
type sequenceMap map[sequenceKey]int
type monkeySequenceMap map[int]sequenceMap
type sequenceSet map[sequenceKey]struct{}

func Puzzle2(lines []string) int {
	total := 0
	data := ParseInput(lines)
	monkeySequences := monkeySequenceMap{}
	allSequences := sequenceSet{}
	var sequenceSetBananas map[sequenceKey]int = map[sequenceKey]int{}

	// Determine the best sequence
	var bestCumulativeSequence sequenceKey
	maxBananas := 0

	for monkeyNumber, line := range data {
		a := line
		currentSequence := make([]int, 0, 5)
		currentBananas := a % 10
		monkeySequences[monkeyNumber] = make(sequenceMap)
		for i := 0; i < 2000; i++ {
			a = GenerateNext(a)

			newCurrentBananas := a % 10
			diff := newCurrentBananas - currentBananas
			currentBananas = newCurrentBananas

			currentSequence = append(currentSequence, diff)
			// Can't count until we have 4 elements
			if len(currentSequence) < 4 {
				continue
			}
			if len(currentSequence) > 4 {
				currentSequence = currentSequence[1:]
			}
			sequenceKey := sequenceKey{currentSequence[0], currentSequence[1], currentSequence[2], currentSequence[3]}
			if _, ok := monkeySequences[monkeyNumber][sequenceKey]; !ok {
				monkeySequences[monkeyNumber][sequenceKey] = currentBananas
				if _, ok := sequenceSetBananas[sequenceKey]; !ok {
					sequenceSetBananas[sequenceKey] = 0
				}
				sequenceSetBananas[sequenceKey] += currentBananas
				if sequenceSetBananas[sequenceKey] > maxBananas {
					maxBananas = sequenceSetBananas[sequenceKey]
					bestCumulativeSequence = sequenceKey
				}
				allSequences[sequenceKey] = struct{}{}
			}
		}
	}

	fmt.Printf("Best Cumulative Sequence: %v\n", bestCumulativeSequence)
	fmt.Printf("Best Cumulative Bananas: %d\n", maxBananas)

	total = maxBananas
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
