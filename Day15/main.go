package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 15"

func readGrid(lines []string) ([][]rune, string) {
	var grid [][]rune
	for l, line := range lines {
		if line == "" {
			// Join the rest of the lines
			rest := strings.Join(lines[l+1:], "")
			return grid, rest
		}
		grid = append(grid, []rune(line))
	}
	return nil, ""
}

func readGridWide(lines []string) ([][]rune, string) {
	var grid [][]rune
	for l, line := range lines {
		if line == "" {
			// Join the rest of the lines
			rest := strings.Join(lines[l+1:], "")
			return grid, rest
		}

		row := []rune{}
		for _, char := range line {
			if char == '#' {
				row = append(row, '#')
				row = append(row, '#')
			}
			if char == '@' {
				row = append(row, '@')
				row = append(row, '.')
			}
			if char == 'O' {
				row = append(row, '[')
				row = append(row, ']')
			}
			if char == '.' {
				row = append(row, '.')
				row = append(row, '.')
			}
		}
		grid = append(grid, row)
	}
	return nil, ""
}

var cardinalDirectionsMap = map[rune][]int{
	'^': {0, -1},
	'v': {0, 1},
	'>': {1, 0},
	'<': {-1, 0},
}

var debug = false

func showGrid(grid [][]rune) {
	if !debug {
		return
	}
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func findBot(grid [][]rune) (int, int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				return x, y
			}
		}
	}
	return -1, -1
}

func computeBoxCost(grid [][]rune) int64 {
	total := int64(0)
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'O' || cell == '[' {
				total += int64(y*100 + x)
			}
		}
	}
	return total
}

func maybeMove(move rune, grid [][]rune, botX, botY int) (int, int, bool) {
	dir := cardinalDirectionsMap[move]
	fmt.Printf("Moving %c %d, %d dir=%v\n", move, botX, botY, dir)
	item := grid[botY][botX]
	newX := botX + dir[0]
	newY := botY + dir[1]
	newC := grid[newY][newX]
	fmt.Printf("%c\n", newC)
	if newC == '@' {
		fmt.Printf("Hit bot at %d, %d\n", newX, newY)
		panic("This shouldn't happen")
	}
	if newC == '#' {
		fmt.Println("Hit wall")
		return botX, botY, false
	} else if newC == 'O' {
		_, _, res := maybeMove(move, grid, newX, newY)
		if !res {
			return botX, botY, false
		}
	}
	grid[botY][botX] = '.'
	grid[newY][newX] = item
	return newX, newY, true
}

func Puzzle1(lines []string) int64 {
	var total int64 = 0
	grid, moves := readGrid(lines)
	if grid == nil {
		panic("Invalid input")
	}

	fmt.Println(moves)
	botX, botY := findBot(grid)
	if botX == -1 || botY == -1 {
		panic("Bot not found")
	}
	fmt.Println("Bot found at", botX, botY)
	for _, move := range moves {
		botX, botY, _ = maybeMove(move, grid, botX, botY)
		fmt.Println()
	}

	showGrid(grid)
	total = computeBoxCost(grid)
	return total
}

func canMove(move rune, grid [][]rune, botX, botY int) bool {
	dir := cardinalDirectionsMap[move]
	newX := botX + dir[0]
	newY := botY + dir[1]
	newC := grid[newY][newX]
	if move == '<' || move == '>' {
		if newC == '.' {
			return true
		}
		if newC == '#' {
			return false
		}
		if newC == '[' || newC == ']' {
			return canMove(move, grid, newX, newY)
		}
	}

	if move == '^' || move == 'v' {
		if newC == '.' {
			return true
		}
		if newC == '#' {
			return false
		}
		if newC == '[' {
			return canMove(move, grid, newX, newY) && canMove(move, grid, newX+1, newY)
		}
		if newC == ']' {
			return canMove(move, grid, newX-1, newY) && canMove(move, grid, newX, newY)
		}
	}

	panic("Invalid move")
}

func makeMove(move rune, grid [][]rune, botX, botY int) (int, int) {
	dir := cardinalDirectionsMap[move]
	item := grid[botY][botX]
	newX := botX + dir[0]
	newY := botY + dir[1]
	newC := grid[newY][newX]
	if move == '<' || move == '>' {
		if newC == '.' {
			grid[newY][newX] = item
		}
		if newC == '[' || newC == ']' {
			makeMove(move, grid, newX, newY)
		}
	}
	if move == '^' || move == 'v' {
		if newC == '.' {
			grid[newY][newX] = item
		}
		if newC == '[' {
			makeMove(move, grid, newX, newY)
			makeMove(move, grid, newX+1, newY)
		}
		if newC == ']' {
			makeMove(move, grid, newX, newY)
			makeMove(move, grid, newX-1, newY)
		}
	}
	grid[botY][botX] = '.'
	grid[newY][newX] = item
	return newX, newY
}

func Puzzle2(lines []string) int64 {
	var total int64 = 0
	grid, moves := readGridWide(lines)
	if grid == nil {
		panic("Invalid input")
	}
	showGrid(grid)
	fmt.Println(moves)
	botX, botY := findBot(grid)
	if botX == -1 || botY == -1 {
		panic("Bot not found")
	}
	fmt.Println("Bot found at", botX, botY)
	for _, move := range moves {
		if canMove(move, grid, botX, botY) {
			botX, botY = makeMove(move, grid, botX, botY)
		}
		showGrid(grid)
		fmt.Println()
	}

	showGrid(grid)
	total = computeBoxCost(grid)
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
