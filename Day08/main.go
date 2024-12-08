package main

import (
	"fmt"
	"os"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 8"

type Coordinate struct {
	x int
	y int
}

type Frequency struct {
	x         int
	y         int
	frequency rune
}

func DisplayMap(runes [][]rune) {
	for _, line := range runes {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func Puzzle1(lines []string) int64 {
	runes := [][]rune{}
	for _, line := range lines {
		runes = append(runes, []rune(line))
	}
	maxY := len(runes) - 1
	maxX := len(runes[0]) - 1
	DisplayMap(runes)

	// Create a map of frequencies
	freqMap := make(map[rune][]Coordinate)
	for y, line := range runes {
		for x, c := range line {
			if c != '.' {
				// Does this frequency already exist?
				coord := Coordinate{x: x, y: y}
				if _, ok := freqMap[c]; ok {
					// Add the frequency to the map
					freqMap[c] = append(freqMap[c], coord)
				} else {
					// Create a new frequency
					freqMap[c] = []Coordinate{coord}
				}
			}
		}
	}

	antinodes := make(map[Coordinate]struct{})

	for _, freq := range freqMap {
		// For each pair of frequencies
		for _, freq1 := range freq {
			for _, freq2 := range freq {
				if freq1.x == freq2.x && freq1.y == freq2.y {
					// Skip if the frequencies are the same
					continue
				}

				// Compute line of sight between the two frequencies
				dx := freq1.x - freq2.x
				dy := freq1.y - freq2.y
				// Two antinodes are created on either side of the line of sight
				antinode1 := Coordinate{x: freq1.x + dx, y: freq1.y + dy}
				// Check if antinode is on the map
				if antinode1.x >= 0 && antinode1.x <= maxX && antinode1.y >= 0 && antinode1.y <= maxY {
					// Add the antinode to the map
					antinodes[antinode1] = struct{}{}
					runes[antinode1.y][antinode1.x] = '#'
					DisplayMap(runes)
				}
				antinode2 := Coordinate{x: freq2.x - dx, y: freq2.y - dy}
				if antinode2.x >= 0 && antinode2.x <= maxX && antinode2.y >= 0 && antinode2.y <= maxY {
					// Add the antinode to the map
					antinodes[antinode2] = struct{}{}
					runes[antinode2.y][antinode2.x] = '#'
				}
			}
		}
	}
	DisplayMap(runes)
	return int64(len(antinodes))
}

func Puzzle2(lines []string) int64 {
	runes := [][]rune{}
	for _, line := range lines {
		runes = append(runes, []rune(line))
	}
	maxY := len(runes) - 1
	maxX := len(runes[0]) - 1
	DisplayMap(runes)

	// Create a map of frequencies
	freqMap := make(map[rune][]Coordinate)
	for y, line := range runes {
		for x, c := range line {
			if c != '.' {
				// Does this frequency already exist?
				coord := Coordinate{x: x, y: y}
				if _, ok := freqMap[c]; ok {
					// Add the frequency to the map
					freqMap[c] = append(freqMap[c], coord)
				} else {
					// Create a new frequency
					freqMap[c] = []Coordinate{coord}
				}
			}
		}
	}

	antinodes := make(map[Coordinate]struct{})

	for _, freq := range freqMap {
		// For each pair of frequencies
		for _, freq1 := range freq {
			for _, freq2 := range freq {
				if freq1.x == freq2.x && freq1.y == freq2.y {
					// Skip if the frequencies are the same
					continue
				}

				// Compute line of sight between the two frequencies
				dx := freq1.x - freq2.x
				dy := freq1.y - freq2.y
				nx := freq1.x
				ny := freq1.y
				for {
					if nx < 0 || nx > maxX || ny < 0 || ny > maxY {
						break
					}
					antinode1 := Coordinate{x: nx, y: ny}
					antinodes[antinode1] = struct{}{}
					runes[antinode1.y][antinode1.x] = '#'
					// Increment at the end because in this scenario the
					// antenna locations are considered antinodes.
					nx += dx
					ny += dy
				}

				nx = freq2.x
				ny = freq2.y
				for {
					if nx < 0 || nx > maxX || ny < 0 || ny > maxY {
						break
					}
					antinode2 := Coordinate{x: nx, y: ny}
					antinodes[antinode2] = struct{}{}
					runes[antinode2.y][antinode2.x] = '#'
					// Decrement at the end because in this scenario the
					// antenna locations are considered antinodes.
					nx -= dx
					ny -= dy
				}
			}
		}
	}
	DisplayMap(runes)
	return int64(len(antinodes))
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
