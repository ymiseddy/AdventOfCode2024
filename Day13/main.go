package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 13"

func scalarMultiply(a int, b []int) []int {
	res := make([]int, len(b))
	for i := range b {
		res[i] = a * b[i]
	}
	return res
}

func Add(a []int, b []int) []int {
	res := make([]int, len(a))
	for i := range a {
		res[i] = a[i] + b[i]
	}
	return res
}

type PuzzleInput struct {
	ButtonA []int
	ButtonB []int
	Prize   []int
}

func ReadInputs(lines []string) []PuzzleInput {
	buttonARe := regexp.MustCompile(`^(Button A): X\+(\d+), Y\+(\d+)$`)
	buttonBRe := regexp.MustCompile(`^(Button B): X\+(\d+), Y\+(\d+)$`)
	prizeRe := regexp.MustCompile(`^(Prize): X=(\d+), Y=(\d+)$`)
	var puzzleInputs []PuzzleInput
	puzzleInput := PuzzleInput{}
	for n, line := range lines {
		match := buttonARe.FindStringSubmatch(line)
		if match != nil {
			strx := match[2]
			stry := match[3]
			x, _ := strconv.Atoi(strx)
			y, _ := strconv.Atoi(stry)
			puzzleInput.ButtonA = []int{x, y}
			continue
		}
		match = buttonBRe.FindStringSubmatch(line)
		if match != nil {
			strx := match[2]
			stry := match[3]
			x, _ := strconv.Atoi(strx)
			y, _ := strconv.Atoi(stry)
			puzzleInput.ButtonB = []int{x, y}
			continue
		}

		match = prizeRe.FindStringSubmatch(line)
		if match != nil {
			strx := match[2]
			stry := match[3]
			x, _ := strconv.Atoi(strx)
			y, _ := strconv.Atoi(stry)
			puzzleInput.Prize = []int{x, y}
			if n != len(lines)-1 {
				continue
			}
		}

		if line == "" || n == len(lines)-1 {
			puzzleInputs = append(puzzleInputs, puzzleInput)
			puzzleInput = PuzzleInput{}
			continue
		}

		panic(fmt.Sprintf("Invalid line: '%s'", line))
	}
	return puzzleInputs
}

func FindMinPrize(puzzleInput PuzzleInput) (int64, []int) {
	aCost := 3
	bCost := 1

	xf := float64(puzzleInput.Prize[0])
	yf := float64(puzzleInput.Prize[1])
	x1 := float64(puzzleInput.ButtonA[0])
	y1 := float64(puzzleInput.ButtonA[1])
	x2 := float64(puzzleInput.ButtonB[0])
	y2 := float64(puzzleInput.ButtonB[1])

	aPress := (-x2*yf + xf*y2) / (x1*y2 - x2*y1)
	bPress := (x1*yf - xf*y1) / (x1*y2 - x2*y1)
	if math.Floor(aPress) == aPress && math.Floor(bPress) == bPress {
		cost := int64(aCost*int(aPress) + bCost*int(bPress))
		return cost, []int{int(aPress), int(bPress)}
	}

	return 0, []int{0, 0}
}

func Puzzle1(lines []string) int64 {
	var total int64 = 0
	puzzleInputs := ReadInputs(lines)
	for _, puzzleInput := range puzzleInputs {
		cost, _ := FindMinPrize(puzzleInput)
		total += cost
	}
	return total
}

func Puzzle2(lines []string) int64 {
	var total int64 = 0
	puzzleInputs := ReadInputs(lines)
	for _, puzzleInput := range puzzleInputs {
		puzzleInput.Prize[0] = 10000000000000 + puzzleInput.Prize[0]
		puzzleInput.Prize[1] = 10000000000000 + puzzleInput.Prize[1]
		cost, _ := FindMinPrize(puzzleInput)
		total += cost
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
