package main

import (
	"fmt"
	"github.com/ymiseddy/AdventOfCode2024/shared"
	"os"
)

var title string = "Advent of Code 2024, Day 4"
var xmasSearch string = "XMAS"
var masSearch string = "MAS"

func plotMap(x int, y int, dx int, dy int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y+i*dy][x+i*dx] = rune(searchString[i])
	}
}

var curryPlotMap = func(dx int, dy int) func(x int, y int, charmap [][]rune, searchString string) {
	return func(x int, y int, charmap [][]rune, searchString string) {
		plotMap(x, y, dx, dy, charmap, searchString)
	}
}

var plotMapForward = curryPlotMap(1, 0)
var plotMapBackward = curryPlotMap(-1, 0)
var plotMapUp = curryPlotMap(0, -1)
var plotMapDown = curryPlotMap(0, 1)
var plotDiagonalDownRight = curryPlotMap(1, 1)
var plotDiagonalDownLeft = curryPlotMap(-1, 1)
var plotDiagonalUpRight = curryPlotMap(1, -1)
var plotDiagonalUpLeft = curryPlotMap(-1, -1)

func search(x, y, dx, dy int, lines []string, searchString string) bool {
	xFinal := x + len(searchString)*dx
	yFinal := y + len(searchString)*dy
	if xFinal > len(lines[y]) ||
		yFinal > len(lines) ||
		xFinal < -1 ||
		yFinal < -1 {
		return false
	}

	for i := 0; i < len(searchString); i++ {
		sx := x + dx*i
		sy := y + dy*i
		if lines[sy][sx] != searchString[i] {
			return false
		}
	}
	return true
}

func currySearch(dx int, dy int) func(x, y int, lines []string, searchString string) bool {
	return func(x, y int, lines []string, searchString string) bool {
		return search(x, y, dx, dy, lines, searchString)
	}
}

var searchForward = currySearch(1, 0)
var searchBackward = currySearch(-1, 0)
var searchUp = currySearch(0, -1)
var searchDown = currySearch(0, 1)
var searchDiagonalDownRight = currySearch(1, 1)
var searchDiagonalDownLeft = currySearch(-1, 1)
var searchDiagonalUpRight = currySearch(1, -1)
var searchDiagonalUpLeft = currySearch(-1, -1)

func Puzzle1(lines []string) int64 {
	total := int64(0)
	charmap := make([][]rune, len(lines))
	for y := 0; y < len(lines); y++ {
		charmap[y] = make([]rune, len(lines[y]))
		for x := 0; x < len(lines[y]); x++ {
			charmap[y][x] = '.'
		}
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] != xmasSearch[0] {
				continue
			}

			if searchForward(x, y, lines, xmasSearch) {
				plotMapForward(x, y, charmap, xmasSearch)
				total++
			}

			if searchBackward(x, y, lines, xmasSearch) {
				plotMapBackward(x, y, charmap, xmasSearch)
				total++
			}

			if searchUp(x, y, lines, xmasSearch) {
				plotMapUp(x, y, charmap, xmasSearch)
				total++
			}

			if searchDown(x, y, lines, xmasSearch) {
				plotMapDown(x, y, charmap, xmasSearch)
				total++
			}

			if searchDiagonalDownRight(x, y, lines, xmasSearch) {
				plotDiagonalDownRight(x, y, charmap, xmasSearch)
				total++
			}

			if searchDiagonalDownLeft(x, y, lines, xmasSearch) {
				plotDiagonalDownLeft(x, y, charmap, xmasSearch)
				total++
			}

			if searchDiagonalUpRight(x, y, lines, xmasSearch) {
				plotDiagonalUpRight(x, y, charmap, xmasSearch)
				total++
			}

			if searchDiagonalUpLeft(x, y, lines, xmasSearch) {
				plotDiagonalUpLeft(x, y, charmap, xmasSearch)
				total++
			}
		}
	}
	// drawPlot(charmap)
	return total
}

func drawPlot(charmap [][]rune) {
	fmt.Println()
	for y := 0; y < len(charmap); y++ {
		fmt.Println(string(charmap[y]))
	}
	fmt.Println()
}

func Puzzle2(lines []string) int64 {
	total := int64(0)
	charmap := make([][]rune, len(lines))
	for y := 0; y < len(lines); y++ {
		charmap[y] = make([]rune, len(lines[y]))
		for x := 0; x < len(lines[y]); x++ {
			charmap[y][x] = '.'
		}
	}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if y < 1 || x < 1 || x > len(lines[y])-2 || y > len(lines[0])-2 {
				continue
			}

			if lines[y][x] != 'A' {
				continue
			}
			left := x - 1
			right := x + 1
			top := y - 1
			bottom := y + 1

			crisFound := false
			crossFound := false
			if searchDiagonalDownRight(left, top, lines, masSearch) {
				crisFound = true
				plotDiagonalDownRight(left, top, charmap, masSearch)
			}

			if searchDiagonalUpLeft(right, bottom, lines, masSearch) {
				crisFound = true
				plotDiagonalUpLeft(right, bottom, charmap, masSearch)
			}

			if searchDiagonalDownLeft(right, top, lines, masSearch) {
				crossFound = true
				plotDiagonalDownLeft(right, top, charmap, masSearch)
			}
			if searchDiagonalUpRight(left, bottom, lines, masSearch) {
				crossFound = true
				plotDiagonalUpRight(left, bottom, charmap, masSearch)
			}

			if crisFound && crossFound {
				total++
			}
		}
	}
	//drawPlot(charmap)
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
