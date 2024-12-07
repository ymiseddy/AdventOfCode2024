package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 5"

func FindGuardCurrentPosition(lines []string) (int, int) {
	for y, line := range lines {
		// Find the position of '^' in the line
		x := strings.Index(line, "^")
		if x > -1 {
			// fmt.Println("Guard found at ", x, y)
			return x, y
		}
	}
	return -1, -1
}

func Puzzle1(lines []string) int64 {
	total := int64(0)
	outMap := make([][]rune, len(lines))
	for y, line := range lines {
		outMap[y] = make([]rune, len(line))
		for x, c := range line {
			outMap[y][x] = c
		}
	}
	x, y := FindGuardCurrentPosition(lines)
	xdir := 0
	ydir := -1
	total += 1
	outMap[y][x] = 'X'
	for {
		newx := x + xdir
		newy := y + ydir
		if newx < 0 || newx >= len(lines[0]) || newy < 0 || newy >= len(lines) {
			// We've gone off the edge of the map
			break
		}
		if lines[newy][newx] == '#' {
			// Rotate direction 90 degrees
			tmp := xdir
			xdir = -ydir
			ydir = tmp
		} else {
			// Move forward
			x = newx
			y = newy
			if outMap[y][x] != 'X' {
				total += 1
			}
			outMap[y][x] = 'X'
		}
	}

	// printMap(outMap)

	return total
}

var dirList []rune = []rune{'^', '>', 'v', '<'}

type obstacle struct {
	x int
	y int
}

func TracePathForLoop(x, y, dir int, outMap [][]rune) bool {
	xdir := 0
	ydir := 0
	if dir == 0 {
		xdir = 0
		ydir = -1
	} else if dir == 1 {
		xdir = 1
		ydir = 0
	} else if dir == 2 {
		xdir = 0
		ydir = 1
	} else if dir == 3 {
		xdir = -1
		ydir = 0
	}
	outMap[y][x] = dirList[dir]

	for {
		newx := x + xdir
		newy := y + ydir
		if newx < 0 || newx >= len(outMap[0]) || newy < 0 || newy >= len(outMap) {
			// We've gone off the edge of the map
			// fmt.Println("Hit edge of map - no loop")
			return false
		}
		if outMap[newy][newx] == '#' || outMap[newy][newx] == 'O' {
			// Rotate direction 90 degrees
			tmp := xdir
			xdir = -ydir
			ydir = tmp
			dir = (dir + 1) % 4
			continue
		}
		x = newx
		y = newy
		//Check if dirList contains outMap[y][x]
		if outMap[y][x] == dirList[dir] {
			// fmt.Println("Loop detected")
			return true
		}
		outMap[y][x] = dirList[dir]
	}
}

func makeMapCopy(lines []string) [][]rune {
	outMap := make([][]rune, len(lines))
	for y, line := range lines {
		outMap[y] = make([]rune, len(line))
		for x, c := range line {
			outMap[y][x] = c
		}
	}
	return outMap
}

func printMap(outMap [][]rune) {
	for _, line := range outMap {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func Puzzle2(lines []string) int64 {
	total := int64(0)
	startX, startY := FindGuardCurrentPosition(lines)

	// Brute force method
	for y, line := range lines {
		for x, c := range line {
			if c == '.' {
				outMap := makeMapCopy(lines)
				outMap[y][x] = 'O'
				loop := TracePathForLoop(startX, startY, 0, outMap)
				if loop {
					fmt.Println("Loop detected at ", x, y)
					printMap(outMap)
					total += 1
				}
			}
		}
	}
	return total
}

/*
func Puzzle2(lines []string) int64 {
	total := int64(0)
	obstacles := []obstacle{}
	outMap := makeMapCopy(lines)
	x, y := FindGuardCurrentPosition(lines)
	guardX, guardY := x, y

	// Current direction 0 up 1 right 2 down 3 left
	dir := 0
	xdir := 0
	ydir := -1
	outMap[y][x] = 'X'
	for {
		newx := x + xdir
		newy := y + ydir
		if newx < 0 || newx >= len(lines[0]) || newy < 0 || newy >= len(lines) {
			// We've gone off the edge of the map
			break
		}
		if lines[newy][newx] == '#' {
			// Rotate direction 90 degrees
			dir = (dir + 1) % 4
			tmp := xdir
			xdir = -ydir
			ydir = tmp
		} else {
			if lines[newx][newy] == '.' {
				// Check if placing obstacle would create a loop
				newMap := makeMapCopy(lines)
				newMap[guardY][guardX] = '.'
				newMap[y+ydir][x+xdir] = 'O'
				isLoop := TracePathForLoop(x, y, (dir+1)%4, newMap)
				if isLoop {
					ob := obstacle{x: x + xdir, y: y + ydir}
					if slices.Contains(obstacles, ob) {
						fmt.Println("Obstacle already exists")
					} else {
						total += 1
						obstacles = append(obstacles, ob)
					}
					fmt.Println(x+xdir, y+ydir)
					//fmt.Println("Could place obstacle at ", x+xdir, y+ydir)
					// printMap(newMap)
				}
			}
			// Move forward
			x = newx
			y = newy
			if outMap[y][x] != '.' {
				// total += 1
			}
			outMap[y][x] = dirList[dir]
		}
	}

	return total
}
*/

func main() {
	fmt.Println(title)
	// Read all text from stdin
	lines := shared.ReadLinesFromStream(os.Stdin)
	res1 := Puzzle1(lines)
	fmt.Println("Puzzle 1 result: ", res1)
	res2 := Puzzle2(lines)
	fmt.Println("Puzzle 2 result: ", res2)

}
