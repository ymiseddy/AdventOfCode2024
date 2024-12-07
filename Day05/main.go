package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 5"

func Puzzle1(lines []string) int64 {
	total := int64(0)
	// Map of int64 to int64
	forward := make(map[int64][]int64)
	backward := make(map[int64][]int64)
	next := 0
	for c, line := range lines {
		if line == "" {
			next = c
			break
		}
		// split on bars
		split := strings.Split(line, "|")
		if len(split) != 2 {
			continue
		}
		// convert to int64
		l1, _ := strconv.ParseInt(split[0], 10, 64)
		l2, _ := strconv.ParseInt(split[1], 10, 64)
		// Check if l1 is in map
		_, ok := forward[l1]
		if !ok {
			forward[l1] = []int64{l2}
		} else {
			forward[l1] = append(forward[l1], l2)
		}

		_, ok = backward[l2]
		if !ok {
			backward[l2] = []int64{l1}
		} else {
			backward[l2] = append(backward[l2], l1)
		}
	}
	remaining := lines[next+1:]
	for _, line := range remaining {
		split := strings.Split(line, ",")
		seenSet := make(map[int64]struct{})

		// Create a slice of int64
		arr := make([]int64, len(split))
		inOrder := true
		for n, s := range split {
			i, _ := strconv.ParseInt(s, 10, 64)
			arr[n] = i
			rules := forward[i]
			/*
				if rules == nil {
					continue
				}
			*/
			seenSet[i] = struct{}{}
			for _, r := range rules {
				if _, ok := seenSet[r]; ok {
					inOrder = false
					break
				}
			}
		}
		if inOrder {
			// Get middle element
			midpoint := (len(arr) / 2)
			mid := arr[midpoint]
			total += mid
		} else {
		}
	}
	return total
}

func Puzzle2(lines []string) int64 {
	total := int64(0)
	// Map of int64 to int64
	forward := make(map[int64][]int64)
	backward := make(map[int64][]int64)
	next := 0
	for c, line := range lines {
		if line == "" {
			next = c
			break
		}
		// split on bars
		split := strings.Split(line, "|")
		if len(split) != 2 {
			continue
		}
		// convert to int64
		l1, _ := strconv.ParseInt(split[0], 10, 64)
		l2, _ := strconv.ParseInt(split[1], 10, 64)
		// Check if l1 is in map
		_, ok := forward[l1]
		if !ok {
			forward[l1] = []int64{l2}
		} else {
			forward[l1] = append(forward[l1], l2)
		}

		_, ok = backward[l2]
		if !ok {
			backward[l2] = []int64{l1}
		} else {
			backward[l2] = append(backward[l2], l1)
		}
	}
	remaining := lines[next+1:]
	for _, line := range remaining {
		split := strings.Split(line, ",")
		seenSet := make(map[int64]struct{})
		valueSet := make(map[int64]struct{})

		// Create a slice of int64
		arr := make([]int64, len(split))
		inOrder := true
		for n, s := range split {
			i, _ := strconv.ParseInt(s, 10, 64)
			arr[n] = i
			valueSet[i] = struct{}{}
			rules := forward[i]
			seenSet[i] = struct{}{}
			for _, r := range rules {
				if _, ok := seenSet[r]; ok {
					inOrder = false
					break
				}
			}
		}
		// Ignoring inorder this time.
		if inOrder {
			continue
		}

		// Put the list in order using rules
		orderedByRules := []int64{}
		seenSet = make(map[int64]struct{})
		for {
			var value int64 = -1
			// Find the first value in valueSet (there must be a better way)
			if len(valueSet) == 0 {
				break
			}

			for k, _ := range valueSet {
				value = k
				break
			}

			for {
				swapped := false
				// What must come before value
				rules := backward[value]
				// Check to see if any values in arr is in rules
				for _, r := range rules {
					if _, ok := valueSet[r]; ok {
						swapped = true
						value = r
						break
					}
				}
				// If nothing was swapped, we found the item with highest precedence
				if !swapped {
					break
				}
			}

			seenSet[value] = struct{}{}
			delete(valueSet, value)
			orderedByRules = append(orderedByRules, value)
		}

		midpoint := (len(orderedByRules) / 2)
		total += orderedByRules[midpoint]
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
