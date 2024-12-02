package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var title string = "Advent of Code 2024, Day 2, Puzzle 1"

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

func puzzle1(file *os.File) int64 {
	scanner := bufio.NewScanner(file)
	safe := int64(0)
outer:
	for scanner.Scan() {
		// Scan and parse the values from the line
		line := scanner.Text()
		fields := fieldsAsInts(line)
		originalDir := int64(0)
		for n, field := range fields {
			if n == 0 {
				continue
			}
			delta := fields[n-1] - field
			absDelta := max(delta, -delta)
			if delta == 0 {
				fmt.Println(fields, "zero")
				continue outer
			}
			dir := delta / absDelta
			if originalDir == 0 {
				originalDir = dir
			}
			if dir != originalDir {
				fmt.Println(fields, "direction change")
				continue outer
			}

			if absDelta < 1 || absDelta > 3 {
				fmt.Println(fields, "unsafe - out of range")
				continue outer
			}
		}
		safe += 1
		fmt.Println(fields, "safe")
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

func checkFields(fields []int64) bool {
	originalDir := int64(0)
	for n, field := range fields {
		if n == 0 {
			continue
		}
		delta := fields[n-1] - field
		absDelta := max(delta, -delta)
		if delta == 0 {
			return false
		}
		dir := delta / absDelta
		if originalDir == 0 {
			originalDir = dir
		}
		if dir != originalDir {
			return false
		}

		if absDelta < 1 || absDelta > 3 {
			return false
		}
	}
	return true
}

func puzzle2(file *os.File) int64 {
	number := 0
	scanner := bufio.NewScanner(file)
	safe := int64(0)
	for scanner.Scan() {
		number++
		// Scan and parse the values from the line
		line := scanner.Text()
		fields := fieldsAsInts(line)
		if checkFields(fields) {
			safe++
			continue
		}
		for n := 0; n < len(fields); n++ {
			newFields := slices.Concat(fields[:n], fields[n+1:])

			if checkFields(newFields) {
				fmt.Println("Safe after removing", n)
				safe++
				break
			}
		}
	}

	// Iterate ofer the left list
	return safe
}

func main() {
	//x := []int{20, 21, 24, 25, 27, 29, 27}
	// n := 2
	// fmt.Println(append(x[:n], x[n+1:]...))
	/*
		for n := 0; n < len(x); n++ {
			b := slices.Concat(x[:n], x[n+1:])
			fmt.Println(n, n+1, b)
		}
	*/

	fmt.Println(title)
	code := puzzle2(os.Stdin)
	fmt.Printf("Result: %d\n", code)

}
