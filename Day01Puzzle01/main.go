package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	// "strconv"
)

var title string = "Advent of Code 2024, Day 1, Puzzle 1"

func processFile(file *os.File) float64 {
	scanner := bufio.NewScanner(file)
	var listL []float64
	var listR []float64
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		numL, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			panic(err)
		}
		numR, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			panic(err)
		}
		listL = append(listL, numL)
		listR = append(listR, numR)
	}

	// Sort the lists
	sort.Float64s(listL)
	sort.Float64s(listR)

	// Calculate the sum of the differences
	sum := 0.0
	for i := 0; i < len(listL); i++ {
		sum += math.Abs(listL[i] - listR[i])
	}

	return sum
}

func main() {
	fmt.Println(title)
	code := processFile(os.Stdin)
	fmt.Printf("Result %.2f\n", code)
}
