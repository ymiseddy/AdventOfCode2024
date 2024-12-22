package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day "

func ParseInput(lines []string) [][]rune {
	data := [][]rune{}
	for _, line := range lines {
		data = append(data, []rune(line))
	}
	return data
}

type runePair struct {
	a rune
	b rune
}

var shortestPaths = map[runePair]string{
	{'0', '1'}: "^<A",
	{'0', '2'}: "^A",
	{'0', '3'}: "^>A",
	{'0', '4'}: "^<^A",
	{'0', '5'}: "^^A",
	{'0', '6'}: "^^>A",
	{'0', '7'}: "^^^<A",
	{'0', '8'}: "^^^A",
	{'0', '9'}: "^^^>A",
	{'0', 'A'}: ">A",
	{'1', '0'}: ">VA",
	{'1', '2'}: ">A",
	{'1', '3'}: ">>A",
	{'1', '4'}: "^A",
	{'1', '5'}: "^>A",
	{'1', '6'}: "^>>A",
	{'1', '7'}: "^^A",
	{'1', '8'}: "^^>A",
	{'1', '9'}: "^^>>A",
	{'1', 'A'}: ">>VA",
	{'2', '0'}: "VA",
	{'2', '1'}: "<A",
	{'2', '3'}: ">A",
	{'2', '4'}: "<^A",
	{'2', '5'}: "^A",
	{'2', '6'}: "^>A",
	{'2', '7'}: "<^^A",
	{'2', '8'}: "^^A",
	{'2', '9'}: "^^>A",
	{'2', 'A'}: "V>A",
	{'3', '0'}: "<VA",
	{'3', '1'}: "<<A",
	{'3', '2'}: "<A",
	{'3', '4'}: "<<^A",
	{'3', '5'}: "<^A",
	{'3', '6'}: "^A",
	{'3', '7'}: "<<^^A",
	{'3', '8'}: "<^^A",
	{'3', '9'}: "^^A",
	{'3', 'A'}: "VA",
	{'4', '0'}: ">VVA",
	{'4', '1'}: "VA",
	{'4', '2'}: "V>A",
	{'4', '3'}: "V>>A",
	{'4', '5'}: ">A",
	{'4', '6'}: ">>A",
	{'4', '7'}: "^A",
	{'4', '8'}: "^>A",
	{'4', '9'}: "^>>A",
	{'4', 'A'}: ">>VVA",
	{'5', '0'}: "VVA",
	{'5', '1'}: "<VA",
	{'5', '2'}: "VA",
	{'5', '3'}: "V>A",
	{'5', '4'}: "<A",
	{'5', '6'}: ">A",
	{'5', '7'}: "<^A",
	{'5', '8'}: "^A",
	{'5', '9'}: "^>A",
	{'5', 'A'}: "VV>A",
	{'6', '0'}: "<VVA",
	{'6', '1'}: "<<VA",
	{'6', '2'}: "<VA",
	{'6', '3'}: "VA",
	{'6', '4'}: "<<A",
	{'6', '5'}: "<A",
	{'6', '7'}: "<<^A",
	{'6', '8'}: "<^A",
	{'6', '9'}: "^A",
	{'6', 'A'}: "VVA",
	{'7', '0'}: ">VVVA",
	{'7', '1'}: "VVA",
	{'7', '2'}: "VV>A",
	{'7', '3'}: "VV>>A",
	{'7', '4'}: "VA",
	{'7', '5'}: "V>A",
	{'7', '6'}: "V>>A",
	{'7', '8'}: ">A",
	{'7', '9'}: ">>A",
	{'7', 'A'}: ">>VVVA",
	{'8', '0'}: "VVVA",
	{'8', '1'}: "<VVA",
	{'8', '2'}: "VVA",
	{'8', '3'}: "VV>A",
	{'8', '4'}: "<VA",
	{'8', '5'}: "VA",
	{'8', '6'}: "V>A",
	{'8', '7'}: "<A",
	{'8', '9'}: ">A",
	{'8', 'A'}: "VVV>A",
	{'9', '0'}: "<VVVA",
	{'9', '1'}: "<<VVA",
	{'9', '2'}: "<VVA",
	{'9', '3'}: "VVA",
	{'9', '4'}: "<<VA",
	{'9', '5'}: "<VA",
	{'9', '6'}: "VA",
	{'9', '7'}: "<<A",
	{'9', '8'}: "<A",
	{'9', 'A'}: "VVVA",
	{'<', '>'}: ">>A",
	{'<', 'A'}: ">>^A",
	{'<', 'V'}: ">A",
	{'<', '^'}: ">^A",
	{'>', '<'}: "<<A",
	{'>', 'A'}: "^A",
	{'>', 'V'}: "<A",
	{'>', '^'}: "<^A",
	{'A', '0'}: "<A",
	{'A', '1'}: "^<<A",
	{'A', '2'}: "<^A",
	{'A', '3'}: "^A",
	{'A', '4'}: "^^<<A",
	{'A', '5'}: "<^^A",
	{'A', '6'}: "^^A",
	{'A', '7'}: "^^^<<A",
	{'A', '8'}: "<^^^A",
	{'A', '9'}: "^^^A",
	{'A', '<'}: "V<<A",
	{'A', '>'}: "VA",
	{'A', 'V'}: "<VA",
	{'A', '^'}: "<A",
	{'V', '<'}: "<A",
	{'V', '>'}: ">A",
	{'V', 'A'}: "^>A",
	{'V', '^'}: "^A",
	{'^', '<'}: "V<A",
	{'^', '>'}: "V>A",
	{'^', 'A'}: ">A",
	{'^', 'V'}: "VA",
}

func Puzzle1(lines []string) int {
	total := 0
	data := ParseInput(lines)

	for _, line := range data {
		path := FindRoute2(line, 'A', 3)
		pathLength := len(path)
		intLine, _ := strconv.Atoi(string(line[0 : len(line)-1]))
		total += intLine * pathLength
		fmt.Printf("Int line: %d * length: %d\n", intLine, pathLength)
	}

	return total
}

type memoKey struct {
	characters string
	start      rune
	depth      int
}

var routeMemo = map[memoKey][]rune{}

func FindRoute2(characters []rune, start rune, depth int) []rune {
	key := memoKey{string(characters), start, depth}
	if val, ok := routeMemo[key]; ok {
		return val
	}
	//fmt.Printf("FindRoute2: %s\n", string(characters))
	sequence := []rune{}
	current := start
	if depth == 0 {
		return characters
	}

	for _, char := range characters {
		if char == current {
			sequence = append(sequence, 'A')
			current = char
			continue
		}
		x := shortestPaths[runePair{current, char}]
		sequence = append(sequence, []rune(x)...)
		current = char
	}
	route := FindRoute2(sequence, 'A', depth-1)
	newKey := memoKey{string(route), start, depth - 1}
	routeMemo[newKey] = route
	return route
}

var routeLengthMemo = map[memoKey]int{}

func FindRouteLength(characters []rune, start rune, depth int) int {
	fmt.Printf("FindRouteLength depth %d: %s\n", depth, string(characters))
	key := memoKey{string(characters), start, depth}
	if val, ok := routeLengthMemo[key]; ok {
		return val
	}
	fmt.Printf("Depth: %d\n", depth)
	count := 0
	current := start
	if depth == 0 {
		return len(characters)
	}

	for _, char := range characters {
		var seq string
		if current == char {
			seq = "A"
		} else {
			seq = shortestPaths[runePair{current, char}]
		}
		length := FindRouteLength([]rune(seq), 'A', depth-1)
		newKey := memoKey{string(seq), start, depth - 1}
		routeLengthMemo[newKey] = length
		count += length
		current = char
	}
	return count
}

func Puzzle2(lines []string) int {
	total := 0
	data := ParseInput(lines)
	for _, line := range data {
		pathLength := FindRouteLength(line, 'A', 26)
		//pathLength := len(path)
		//fmt.Printf("Path: %s\n", string(path))

		intLine, _ := strconv.Atoi(string(line[0 : len(line)-1]))
		fmt.Printf("Int line: %d * length: %d\n", intLine, pathLength)
		//fmt.Printf("Int line: %d * length: %d\n", intLine, pathLength)
		total += intLine * pathLength
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
