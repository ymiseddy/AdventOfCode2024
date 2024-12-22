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

func Puzzle1(lines []string) int {
	total := 0
	data := ParseInput(lines)

	for _, line := range data {
		pathLength := FindRouteLength(line, 'A', 3)
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

func GetPath(start rune, end rune, prefHorizontal bool) []rune {
	if start == end {
		return []rune{'A'}
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

	spaceCoord := coordinateMap[' ']
	if padGrid[spaceCoord.Y][spaceCoord.X] != ' ' {
		panic("Space is not valid")
	}

	startCoord := coordinateMap[start]
	endCoord := coordinateMap[end]
	deltaX := endCoord.X - startCoord.X
	deltaY := endCoord.Y - startCoord.Y
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
	sequence := []rune{}

	if startCoord.Y == spaceCoord.Y && endCoord.X == spaceCoord.X {
		// Can't move horizontally first.
		for i := 0; i < moveY; i++ {
			sequence = append(sequence, yChar)
		}
		for i := 0; i < moveX; i++ {
			sequence = append(sequence, xChar)
		}
	} else if startCoord.X == spaceCoord.X && endCoord.Y == spaceCoord.Y {
		// Can't move vertically first.
		for i := 0; i < moveX; i++ {
			sequence = append(sequence, xChar)
		}
		for i := 0; i < moveY; i++ {
			sequence = append(sequence, yChar)
		}
	} else {
		if prefHorizontal {
			for i := 0; i < moveX; i++ {
				sequence = append(sequence, xChar)
			}
			for i := 0; i < moveY; i++ {
				sequence = append(sequence, yChar)
			}
		} else {
			for i := 0; i < moveY; i++ {
				sequence = append(sequence, yChar)
			}
			for i := 0; i < moveX; i++ {
				sequence = append(sequence, xChar)
			}
		}
	}

	startCoord = coordinateMap[start]
	for _, x := range sequence {
		dirCoord := directionKeyToDirection[x]
		dx, dy := dirCoord.X+startCoord.X, dirCoord.Y+startCoord.Y
		startCoord = shared.Coord{X: dx, Y: dy}
		if dx == spaceCoord.X && dy == spaceCoord.Y {
			fmt.Printf("Start: %s End: %s dx:%d dy:%d: %s\n", string(start), string(end), dx, dy, string(sequence))
			panic("Direction is not valid")
		}
	}

	sequence = append(sequence, 'A')
	// fmt.Printf("GetPath: %s -> %s dx:%d dy:%d: %s\n", string(start), string(end), deltaX, deltaY, string(sequence))
	return sequence
}

var routeLengthMemo = map[memoKey]int{}

func FindRouteLength(characters []rune, start rune, depth int) int {
	//fmt.Printf("FindRouteLength depth %d: %s\n", depth, string(characters))
	key := memoKey{string(characters), start, depth}
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
		var seq []rune
		var seq2 []rune
		if current == char {
			seq = []rune{'A'}
			seq2 = nil
		} else {
			seq = GetPath(current, char, true)
			seq2 = GetPath(current, char, false)
			if string(seq) == string(seq2) {
				seq2 = nil
			}
		}
		length := FindRouteLength([]rune(seq), 'A', depth-1)
		if seq2 != nil {
			length2 := FindRouteLength([]rune(seq2), 'A', depth-1)
			if length2 < length {
				length = length2
			}
		}
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
		pathLength := FindRouteLength(line, 'A', 3)

		intLine, _ := strconv.Atoi(string(line[0 : len(line)-1]))
		fmt.Printf("Int line: %d * length: %d\n", intLine, pathLength)
		//fmt.Printf("Int line: %d * length: %d\n", intLine, pathLength)
		total += intLine * pathLength
	}
	return total
}

func main() {
	buildCoordMap()
	fmt.Println(title)
	// Read all text from stdin

	lines := shared.ReadLinesFromStream(os.Stdin)

	res1 := Puzzle1(lines)
	fmt.Println("Puzzle 1 result: ", res1)
	res2 := Puzzle2(lines)
	fmt.Println("Puzzle 2 result: ", res2)
}
