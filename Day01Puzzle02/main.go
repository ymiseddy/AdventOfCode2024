package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var title string = "Advent of Code 2024, Day 1, Puzzle 2"

func processFile(file *os.File) int64 {
	scanner := bufio.NewScanner(file)
	var listL []int64
	mapR := make(map[int64]int64)
	for scanner.Scan() {
		// Scan and parse the values from the line
		line := scanner.Text()
		fields := strings.Fields(line)
		numL, err := strconv.ParseInt(fields[0], 10, 64)
		if err != nil {
			panic(err)
		}
		numR, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			panic(err)
		}
		// Add values to the left list
		listL = append(listL, numL)

		// If the right value already exists, increment the value, otherwise set it to 1
		if _, ok := mapR[numR]; ok {
			mapR[numR]++
		} else {
			mapR[numR] = 1
		}
	}

	// Iterate ofer the left list
	var sum int64 = 0
	for _, numL := range listL {
		if count, ok := mapR[numL]; ok {
			sum += numL * count
		}
	}
	return sum
}

func main() {
	fmt.Println(title)
	code := processFile(os.Stdin)
	fmt.Printf("Result: %d\n", code)
}
