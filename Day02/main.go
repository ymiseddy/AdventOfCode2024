package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/ymiseddy/AdventOfCode2024/shared"

	"strconv"
	"strings"
)

var title string = "Advent of Code 2024, Day 2"

func fieldsAsInts(line string) []int64 {
	fields := strings.Fields(line)
	ints := make([]int64, len(fields))
	for i, field := range fields {
		val, err := strconv.ParseInt(field, 10, 64)
		if err != nil {
			panic(err)
		}
		ints[i] = val
	}
	return ints
}

func puzzle1(values [][]int64) int64 {
	safe := int64(0)
outer:
	for _, fields := range values {
		originalDir := int64(0)
		for n, field := range fields {
			if n == 0 {
				continue
			}
			delta := fields[n-1] - field
			absDelta := max(delta, -delta)
			if delta == 0 {
				continue outer
			}
			dir := delta / absDelta
			if originalDir == 0 {
				originalDir = dir
			}
			if dir != originalDir {
				continue outer
			}

			if absDelta < 1 || absDelta > 3 {
				continue outer
			}
		}
		safe += 1
	}

	// Iterate ofer the left list
	return safe
}

func checkField(dir int64, prev int64, curr int64) (bool, int64) {
	delta := prev - curr
	absDelta := max(delta, -delta)
	if absDelta < 1 || absDelta > 3 {
		return false, dir
	}
	newDir := delta / absDelta
	if dir != 0 && newDir != dir {
		return false, dir
	}
	return true, newDir
}

func checkFields(fields []int64, direction int64, missed bool) bool {

	// Recursive solution

	// Base case - we only have one element, we're done.
	if len(fields) == 1 {
		return true
	}

	// Compute difference and direction.
	delta := fields[1] - fields[0]
	absDelta := max(delta, -delta)
	var newDirection int64 = 0
	if absDelta > 0 {
		newDirection = delta / absDelta
	}

	// Check constraints.
	if absDelta > 0 && absDelta < 4 &&
		(direction == 0 || direction == newDirection) {

		// Check forward.
		if checkFields(fields[1:], newDirection, missed) {
			return true
		}
	}

	// If we have previously missed, we can't re-check.
	if missed {
		return false
	}

	// When directon is 0, we are the first element.
	// Check moving forward without the current element and invariant direction.
	if direction == 0 && checkFields(fields[1:], 0, true) {
		res := checkFields(fields[1:], direction, true)
		return res
	}

	// Check without the next element.
	if checkFields(slices.Concat(fields[:1], fields[2:]), direction, true) {
		return true
	}

	return false
}

func puzzle2(values [][]int64) int64 {
	number := 0
	safe := int64(0)
	for _, fields := range values {
		number++
		if checkFields(fields, 0, false) {
			safe++
			continue
		}
	}

	// Iterate ofer the left list
	return safe
}

func main() {
	values, err := shared.ReadIntsFromStream(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(title)
	result1 := puzzle1(values)
	fmt.Printf("Puzzle 1 result: %d\n", result1)
	result2 := puzzle2(values)
	fmt.Printf("Puzzle 2 result: %d\n", result2)
}
