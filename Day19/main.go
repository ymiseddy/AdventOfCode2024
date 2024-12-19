package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day "

func ParseInput(lines []string) ([]string, []string) {
	// split line 0 on comma spaces
	towelTypes := strings.Split(lines[0], ", ")
	wantedDesigns := lines[2:]
	return towelTypes, wantedDesigns
}

var memoDesigns = map[string]bool{}

func matchDesign(wantedDesign string, towelTypes []string) bool {
	result, ok := memoDesigns[wantedDesign]
	if ok {
		return result
	}

	// See if we can construct wantedDesign from towelTypes
	if len(wantedDesign) == 0 {
		return true
	}
	for _, towelType := range towelTypes {
		if strings.HasPrefix(wantedDesign, towelType) {
			res := matchDesign(wantedDesign[len(towelType):], towelTypes)
			if res {
				memoDesigns[wantedDesign] = true
				return true
			} else {
				memoDesigns[wantedDesign] = false
			}
		}
	}
	return false
}

func Puzzle1(lines []string) int {
	total := 0
	towelTypes, wantedDesigs := ParseInput(lines)
	for _, wantedDesign := range wantedDesigs {
		if matchDesign(wantedDesign, towelTypes) {
			total++
		}
	}
	return total
}

var designWays = map[string]int{}

func matchDesign2(wantedDesign string, towelTypes []string) int {
	result, ok := designWays[wantedDesign]
	if ok {
		return result
	}

	// See if we can construct wantedDesign from towelTypes
	if len(wantedDesign) == 0 {
		return 1
	}

	countWays := 0
	for _, towelType := range towelTypes {
		if strings.HasPrefix(wantedDesign, towelType) {
			res := matchDesign2(wantedDesign[len(towelType):], towelTypes)
			countWays += res
		}
	}
	designWays[wantedDesign] = countWays
	return countWays
}

func Puzzle2(lines []string) int {
	total := 0
	towelTypes, wantedDesigs := ParseInput(lines)
	for _, wantedDesign := range wantedDesigs {
		//fmt.Printf("wantedDesign: '%v'", wantedDesign)
		countWays := matchDesign2(wantedDesign, towelTypes)
		//fmt.Printf("%d ways\n", countWays)
		total += countWays
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
