package main

import (
	"fmt"
	"math"
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

func Puzzle1(lines []string) int {
	total := 0
	data := ParseInput(lines)

	for _, line := range data {
		pathLength := FindBestRouteLength(line, 'A', 3)
		intLine, _ := strconv.Atoi(string(line[0 : len(line)-1]))
		total += intLine * pathLength
		fmt.Printf("Int line: %d * length: %d\n", intLine, pathLength)
	}

	return total
}

type memoKey struct {
	characters     string
	start          rune
	depth          int
	prefHorizontal bool
}

var routeMemo = map[memoKey][]rune{}

var numericPad = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
}

var directionPad = [][]rune{
	{' ', '^', 'A'},
	{'<', 'V', '>'},
}

var directionKeys = []rune{
	'^', 'V', '<', '>',
}

var numericRuneToCoord = map[rune]shared.Coord{}
var directionalRuneToCoord = map[rune]shared.Coord{}

func buildCoordMap() {
	for y, row := range numericPad {
		for x, cell := range row {
			numericRuneToCoord[cell] = shared.Coord{X: x, Y: y}
		}
	}
	for y, row := range directionPad {
		for x, cell := range row {
			directionalRuneToCoord[cell] = shared.Coord{X: x, Y: y}
		}
	}
}

var directionKeyToDirection = map[rune]shared.Coord{
	'^': {X: 0, Y: -1},
	'>': {X: 1, Y: 0},
	'V': {X: 0, Y: 1},
	'<': {X: -1, Y: 0},
}

func GetPaths(start rune, end rune) [][]rune {
	if start == end {
		return [][]rune{{'A'}}
	}
	_, isStartDirectional := directionKeyToDirection[start]
	_, isEndDirectional := directionKeyToDirection[end]

	// Figure out which pad to use
	coordinateMap := numericRuneToCoord
	padGrid := numericPad
	if isStartDirectional || isEndDirectional {
		coordinateMap = directionalRuneToCoord
		padGrid = directionPad
	}

	// Sanity check that the space coordinate is valid.
	spaceCoord := coordinateMap[' ']
	if padGrid[spaceCoord.Y][spaceCoord.X] != ' ' {
		panic("Space is not valid")
	}

	currentCoord := coordinateMap[start]
	endCoord := coordinateMap[end]
	deltaX := endCoord.X - currentCoord.X
	deltaY := endCoord.Y - currentCoord.Y
	moveX := shared.Abs(deltaX)
	moveY := shared.Abs(deltaY)

	xChar := '>'
	if deltaX < 0 {
		xChar = '<'
	}

	yChar := 'V'
	if deltaY < 0 {
		yChar = '^'
	}

	sequences := [][]rune{}

	if currentCoord.Y == spaceCoord.Y && endCoord.X == spaceCoord.X {
		// Can't move horizontally first.
		sequence := []rune{}
		for i := 0; i < moveY; i++ {
			sequence = append(sequence, yChar)
		}
		for i := 0; i < moveX; i++ {
			sequence = append(sequence, xChar)
		}
		sequence = append(sequence, 'A')
		sequences = append(sequences, sequence)
	} else if currentCoord.X == spaceCoord.X && endCoord.Y == spaceCoord.Y {
		sequence := []rune{}
		// Can't move vertically first.
		for i := 0; i < moveX; i++ {
			sequence = append(sequence, xChar)
		}
		for i := 0; i < moveY; i++ {
			sequence = append(sequence, yChar)
		}
		sequence = append(sequence, 'A')
		sequences = append(sequences, sequence)
	} else {
		sequence := []rune{}
		for i := 0; i < moveX; i++ {
			sequence = append(sequence, xChar)
		}
		for i := 0; i < moveY; i++ {
			sequence = append(sequence, yChar)
		}
		sequence = append(sequence, 'A')
		sequences = append(sequences, sequence)

		sequence = []rune{}
		for i := 0; i < moveY; i++ {
			sequence = append(sequence, yChar)
		}
		for i := 0; i < moveX; i++ {
			sequence = append(sequence, xChar)
		}
		sequence = append(sequence, 'A')
		sequences = append(sequences, sequence)
	}

	currentCoord = coordinateMap[start]
	return sequences
}

var routeLengthMemo = map[routeLengthMemoKey]int{}

type routeLengthMemoKey struct {
	characters string
	start      rune
	depth      int
}

func FindBestRouteLength(characters []rune, start rune, depth int) int {
	key := routeLengthMemoKey{string(characters), start, depth}
	if val, ok := routeLengthMemo[key]; ok {
		return val
	}
	// fmt.Printf("Depth: %d\n", depth)
	count := 0
	current := start
	if depth == 0 {
		return len(characters)
	}

	for _, char := range characters {
		var seq [][]rune
		if current == char {
			seq = [][]rune{{'A'}}
		} else {
			seq = GetPaths(current, char)
		}

		// Max int
		minLength := math.MaxInt64
		chosenSeq := []rune{}
		for _, seq := range seq {
			length := FindBestRouteLength(seq, 'A', depth-1)
			if length < minLength {
				minLength = length
				chosenSeq = seq
			}
		}
		newKey := routeLengthMemoKey{string(chosenSeq), start, depth - 1}
		routeLengthMemo[newKey] = minLength
		count += minLength
		current = char
	}
	return count
}

func Puzzle2(lines []string) int {
	total := 0
	data := ParseInput(lines)
	for _, line := range data {
		pathLength := FindBestRouteLength(line, 'A', 26)
		intLine, _ := strconv.Atoi(string(line[0 : len(line)-1]))
		fmt.Printf("Int line: %d * length: %d\n", intLine, pathLength)
		//fmt.Printf("Int line: %d * length: %d\n", intLine, pathLength)
		total += intLine * pathLength
	}
	return total
}

func main() {
	buildCoordMap()
	lines := shared.ReadLinesFromStream(os.Stdin)
	// Read all text from stdin
	fmt.Println(title)

	res1 := Puzzle1(lines)
	fmt.Println("Puzzle 1 result: ", res1)
	res2 := Puzzle2(lines)
	fmt.Println("Puzzle 2 result: ", res2)
}
