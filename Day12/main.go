package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var cardinalDirections = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

var cardinalAndDiagonalDirections = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

var title string = "Advent of Code 2024, Day 12"

func scanForUnexplored(runeArray [][]rune, x, y int) (int, int) {
	for my := y; my < len(runeArray); my++ {
		for mx := x; mx < len(runeArray[my]); mx++ {
			if isDigit(runeArray[my][mx]) {
				runeArray[my][mx] = '.'
			}
		}
	}
	for my := y; my < len(runeArray); my++ {
		for mx := x; mx < len(runeArray[my]); mx++ {
			if runeArray[my][mx] != '.' && !isDigit(runeArray[my][mx]) {
				return mx, my
			}
		}
	}
	return -1, -1
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func toDigit(ct int64) rune {
	if ct < 10 {
		return rune(ct + 48)
	}
	return rune(ct + 87)
}

func scanPlotArea(runeArray [][]rune, x, y int) (int64, int64, rune) {
	plotType := runeArray[y][x]
	plotTypeVisited := 100 + plotType
	runeArray[y][x] = plotTypeVisited
	area := int64(1)
	perimeter := int64(0)
	totalPerimeter := int64(0)

	// Look over cardinal directions
	for _, dir := range cardinalDirections {
		nx := x + dir[0]
		ny := y + dir[1]
		if nx < 0 || ny < 0 || ny >= len(runeArray) || nx >= len(runeArray[ny]) {
			perimeter++
			if x == 1 && y == 1 {
				//fmt.Printf("direction outside: %d, %d\n", dir[0], dir[1])
			}

			totalPerimeter++
			continue
		}

		if runeArray[ny][nx] != plotType &&
			runeArray[ny][nx] != plotTypeVisited &&
			!isDigit(runeArray[ny][nx]) {
			perimeter++
			totalPerimeter++
		}
		if runeArray[ny][nx] == plotType {
			narea, nperimiter, _ := scanPlotArea(runeArray, nx, ny)
			area += narea
			totalPerimeter += nperimiter
		}
	}
	runeArray[y][x] = toDigit(perimeter)
	return area, totalPerimeter, plotType
}

func Puzzle1(lines []string) int64 {
	var total int64 = 0
	var runeArray [][]rune
	for y := 0; y < len(lines); y++ {
		// trim line
		lines[y] = strings.TrimSpace(lines[y])
		runeArray = append(runeArray, []rune(lines[y]))
	}
	x := 0
	y := 0
	for true {
		x, y := scanForUnexplored(runeArray, x, y)
		if x == -1 {
			break
		}
		area, perimeter, _ := scanPlotArea(runeArray, x, y)
		total += area * perimeter
	}
	return total
}

func showMap(runeArray [][]rune, cx, cy int) {
	for my := 0; my < len(runeArray); my++ {
		fmt.Printf("%d: ", len(runeArray[my]))
		for mx := 0; mx < len(runeArray[my]); mx++ {
			if mx == cx && my == cy {
				fmt.Printf("@")
			} else {
				fmt.Printf("%c", runeArray[my][mx])
			}
		}
		fmt.Println()
	}
}

func scanPlotArea2(runeArray [][]rune, x, y int) (int64, int64, rune) {
	// Leaving in a bunch of debug code.
	// Realization for this puzzle: The number of sides == the number of corners.
	// Source: Some Guy (Probably Euler)

	plotType := runeArray[y][x]
	plotTypeVisited := 100 + plotType
	var fencePlot [][]rune = [][]rune{
		{'.', '.', '.'},
		{'.', '@', '.'},
		{'.', '.', '.'},
	}
	runeArray[y][x] = plotTypeVisited
	area := int64(1)

	fenceDirections := []bool{false, false, false, false,
		false, false, false, false}
	corners := int64(0)
	myCorners := int64(0)

	// Look for fences in cardinal and diagonal directions
	for n, dir := range cardinalAndDiagonalDirections {
		nx := x + dir[0]
		ny := y + dir[1]
		fx := 1 + dir[0]
		fy := 1 + dir[1]
		if nx < 0 || ny < 0 || ny >= len(runeArray) || nx >= len(runeArray[ny]) {
			fenceDirections[n] = true
			fencePlot[fy][fx] = 'X'
			continue
		}

		if runeArray[ny][nx] != plotType &&
			runeArray[ny][nx] != plotTypeVisited &&
			!isDigit(runeArray[ny][nx]) {
			fencePlot[fy][fx] = 'X'
			fenceDirections[n] = true
		}
	}

	//fmt.Printf("------------------------------------------\n")
	//fmt.Printf("Fence directions: %v\n", fenceDirections)
	//showMap(fencePlot, 1, 1)
	if fenceDirections[1] && fenceDirections[4] {
		//fmt.Printf("Top right corner\n")
		myCorners += 1
	}
	if fenceDirections[4] && fenceDirections[6] {
		//fmt.Printf("Bottom right corner\n")
		myCorners += 1
	}
	if fenceDirections[6] && fenceDirections[3] {
		//fmt.Printf("Bottom left corner\n")
		myCorners += 1
	}
	if fenceDirections[3] && fenceDirections[1] {
		//fmt.Printf("Top left corner\n")
		myCorners += 1
	}

	//	012
	//	3 4
	//	567

	//	XXX
	//	X@O
	//	XOX
	if fenceDirections[0] && !fenceDirections[1] && !fenceDirections[3] {
		//fmt.Printf("Isolated top left corner\n")
		myCorners += 1
	}

	if fenceDirections[2] && !fenceDirections[1] && !fenceDirections[4] {
		//fmt.Printf("Isolated top right corner\n")
		myCorners += 1
	}

	if fenceDirections[7] && !fenceDirections[6] && !fenceDirections[4] {
		//fmt.Printf("Isolated bottom right corner\n")
		myCorners += 1
	}
	if fenceDirections[5] && !fenceDirections[3] && !fenceDirections[6] {
		//fmt.Printf("Isolated bottom left corner\n")
		myCorners += 1
	}

	//fmt.Printf("------------------------------------------\n")

	runeArray[y][x] = toDigit(myCorners)
	//panic("Stop here.")

	// Look over cardinal directions
	for n, dir := range cardinalDirections {
		nx := x + dir[0]
		ny := y + dir[1]

		if nx < 0 || ny < 0 || ny >= len(runeArray) || nx >= len(runeArray[ny]) {
			continue
		}

		if runeArray[ny][nx] != plotType &&
			runeArray[ny][nx] != plotTypeVisited &&
			!isDigit(runeArray[ny][nx]) {
			fenceDirections[n] = true
		}
		if runeArray[ny][nx] == plotType {
			narea, ncorners, _ := scanPlotArea2(runeArray, nx, ny)
			area += narea
			corners += ncorners
		}
	}

	return area, corners + myCorners, plotType
}

func Puzzle2(lines []string) int64 {
	var total int64 = 0
	var runeArray [][]rune
	for y := 0; y < len(lines); y++ {
		// trim line
		lines[y] = strings.TrimSpace(lines[y])
		runeArray = append(runeArray, []rune(lines[y]))
	}
	x := 0
	y := 0
	for true {
		x, y := scanForUnexplored(runeArray, x, y)
		if x == -1 {
			break
		}
		area, sides, runeType := scanPlotArea2(runeArray, x, y)
		//showMap(runeArray, -1, -1)
		fmt.Printf("x: %d y: %d sides: %d area: %d runeType: %c\n", x, y, sides, area, runeType)

		total += area * sides
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
