package main

import (
	"fmt"
	"github.com/ymiseddy/AdventOfCode2024/shared"
	"os"
)

var title string = "Advent of Code 2024, Day 4"
var xmasSearch string = "XMAS"
var masSearch string = "MAS"

func plotMapForward(x int, y int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y][x+i] = rune(searchString[i])
	}
}

func plotMapBackward(x int, y int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y][x-i] = rune(searchString[i])
	}
}

func plotMapUp(x int, y int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y-i][x] = rune(searchString[i])
	}
}

func plotMapDown(x int, y int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y+i][x] = rune(searchString[i])
	}
}

func plotDiagonalDownRight(x int, y int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y+i][x+i] = rune(searchString[i])
	}
}

func plotDiagonalDownLeft(x int, y int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y+i][x-i] = rune(searchString[i])
	}
}

func plotDiagonalUpRight(x int, y int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y-i][x+i] = rune(searchString[i])
	}
}

func plotDiagonalUpLeft(x int, y int, charmap [][]rune, searchString string) {
	for i := 0; i < len(searchString); i++ {
		charmap[y-i][x-i] = rune(searchString[i])
	}
}

func searchForward(x, y int, lines []string, searchString string) bool {
	tack := ""
	booFound := false
	if y == 10 && x == 6 {
		fmt.Println("Found it!")
		booFound = true
	}

	if x+len(searchString) > len(lines[y]) {
		if booFound {
			fmt.Println("Past the end.")
		}
		return false
	}

	for i := 0; i < len(searchString); i++ {
		tack += string(lines[y][x+i])
		if lines[y][i+x] != searchString[i] {
			return false
		}
	}
	//fmt.Println("Forward: ", tack)
	return true
}

func searchBackward(x, y int, lines []string, searchString string) bool {
	tack := ""
	if x-len(searchString)+1 < 0 {
		return false
	}

	for i := 0; i < len(searchString); i++ {
		tack += string(lines[y][x-i])
		if lines[y][x-i] != searchString[i] {
			return false
		}
	}
	//fmt.Println("Backward: ", tack)
	return true
}

func searchUp(x, y int, lines []string, searchString string) bool {
	tack := ""
	if y-len(searchString)+1 < 0 {
		return false
	}
	for i := 0; i < len(searchString); i++ {
		tack += string(lines[y-i][x])
		if lines[y-i][x] != searchString[i] {
			return false
		}
	}

	//fmt.Println("Up: ", tack)
	return true
}

func searchDown(x, y int, lines []string, searchString string) bool {
	tack := ""
	if y+len(searchString) > len(lines) {
		return false
	}
	for i := 0; i < len(searchString); i++ {
		tack += string(lines[y+i][x])
		if lines[y+i][x] != searchString[i] {
			return false
		}
	}
	//fmt.Println("Down: ", tack)
	return true
}

func searchDiagonalDownRight(x, y int, lines []string, searchString string) bool {
	tack := ""
	if y+len(searchString) > len(lines) {
		return false
	}
	if x+len(searchString) > len(lines[y]) {
		return false
	}

	for i := 0; i < len(searchString); i++ {
		tack += string(lines[y+i][x+i])
		if lines[y+i][x+i] != searchString[i] {
			return false
		}
	}
	//fmt.Println("Diagonal Down Right: ", tack)
	return true
}

func searchDiagonalDownLeft(x, y int, lines []string, searchString string) bool {
	tack := ""
	if y+len(searchString) > len(lines) {
		return false
	}
	if x-len(searchString)+1 < 0 {
		return false
	}

	for i := 0; i < len(searchString); i++ {
		tack += string(lines[y+i][x-i])
		if lines[y+i][x-i] != searchString[i] {
			return false
		}
	}
	return true
}

func searchDiagonalUpRight(x, y int, lines []string, searchString string) bool {
	tack := ""
	if y-len(searchString)+1 < 0 {
		return false
	}
	if x+len(searchString) > len(lines[y]) {
		return false
	}

	for i := 0; i < len(searchString); i++ {
		tack += string(lines[y-i][x+i])
		if lines[y-i][x+i] != searchString[i] {
			return false
		}
	}
	return true
}

func searchDiagonalUpLeft(x, y int, lines []string, searchString string) bool {
	tack := ""
	if y-len(searchString)+1 < 0 {
		return false
	}
	if x-len(searchString)+1 < 0 {
		return false
	}

	for i := 0; i < len(searchString); i++ {
		tack += string(lines[y-i][x-i])
		if lines[y-i][x-i] != searchString[i] {
			return false
		}
	}
	return true
}

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
	/*
		fmt.Println()
		for y := 0; y < len(charmap); y++ {
			fmt.Println(string(charmap[y]))
		}
		fmt.Println()
	*/
	return total
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
	/*
		fmt.Println()
		for y := 0; y < len(charmap); y++ {
			fmt.Println(string(charmap[y]))
		}
		fmt.Println()
	*/
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
