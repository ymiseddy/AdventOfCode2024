package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 14"

type Robot struct {
	x  int
	y  int
	dx int
	dy int
}

func readInputs(lines []string) []Robot {
	var robots []Robot
	for _, line := range lines {
		if line == "" {
			continue
		}
		re := regexp.MustCompile(`^p=([+\-]?\d+),[+\-]?(\d+) v=([+\-]?\d+),([+\-]?\d+)$`)
		match := re.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Sprintf("Invalid line: '%s'", line))
		}
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		dx, _ := strconv.Atoi(match[3])
		dy, _ := strconv.Atoi(match[4])

		robots = append(robots, Robot{
			x:  x,
			y:  y,
			dx: dx,
			dy: dy,
		})
	}
	return robots
}

func Move(robots []Robot, steps, width int, height int) {
	for n := 0; n < len(robots); n++ {
		robots[n].x = (robots[n].x + steps*(width+robots[n].dx)) % width
		robots[n].y = (robots[n].y + steps*(height+robots[n].dy)) % height
		if robots[n].y >= height || robots[n].x >= width {
			panic(fmt.Sprintf("Robot %v out of bounds\n", robots[n]))
		}
	}
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func DrawBots(robots []Robot, width, height int) {
	var runeMap [][]rune
	for y := 0; y < height; y++ {
		runeMap = append(runeMap, []rune{})
		for x := 0; x < width; x++ {
			runeMap[y] = append(runeMap[y], '.')
		}
	}

	// Add Bots
	for _, robot := range robots {
		if robot.x >= width || robot.y >= height {
			fmt.Printf("Width %d, Height %d\n", width, height)
			panic(fmt.Sprintf("Robot %v out of bounds\n", robot))
		}
		if runeMap[robot.y][robot.x] == '.' {
			runeMap[robot.y][robot.x] = '1'
		} else if isDigit(runeMap[robot.y][robot.x]) {
			runeMap[robot.y][robot.x] += 1
		}
	}

	// Print map
	for _, row := range runeMap {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func Puzzle1(lines []string) int64 {
	var total int64 = 0
	maxWidth := 101
	maxHeight := 103
	/*
		maxWidth := 11
		maxHeight := 7
	*/
	gutterX := (maxWidth - 1) / 2
	gutterY := (maxHeight - 1) / 2
	robots := readInputs(lines)
	Move(robots, 100, maxWidth, maxHeight)
	quadrants := []int{0, 0, 0, 0}
	fmt.Printf("Gutter: %d, %d\n", gutterX, gutterY)
	for _, robot := range robots {
		q := 0
		if robot.x == gutterX || robot.y == gutterY {
			fmt.Printf("x %d, y %d - Gutter\n", robot.x, robot.y)
			continue
		}
		if robot.x < gutterX && robot.y < gutterY {
			q = 0
		}
		if robot.x < gutterX && robot.y > gutterY {
			q = 1
		}
		if robot.x > gutterX && robot.y < gutterY {
			q = 2
		}
		if robot.x > gutterX && robot.y > gutterY {
			q = 3
		}
		quadrants[q]++
		fmt.Printf("x %d, y %d - Quadrant %d\n", robot.x, robot.y, q)
	}
	DrawBots(robots, maxWidth, maxHeight)
	fmt.Printf("Quadrants: %v\n", quadrants)

	total = 1
	for _, quadrant := range quadrants {
		total *= int64(quadrant)
	}

	return total
}

// For puzzle 2, I'm looking for how "clustered" the bots are to each other.
// Compute the centroid (mean x, mean y) and compute the mean distance from
// this centroid.  I stepped this down by 5 from 35 to 25 and it worked!
func Puzzle2(lines []string) int64 {
	var total int64 = 0
	maxWidth := 101
	maxHeight := 103
	robots := readInputs(lines)
	steps := int64(0)
	for true {
		Move(robots, 1, maxWidth, maxHeight)
		steps++

		// Compute the robot centroid
		var centroidX float64 = 0
		var centroidY float64 = 0
		for _, robot := range robots {
			centroidX += float64(robot.x)
			centroidY += float64(robot.y)
		}
		centroidX /= float64(len(robots))
		centroidY /= float64(len(robots))
		// Compute the mean distance from the centroid
		var meanDistance float64 = 0
		for _, robot := range robots {
			meanDistance += math.Sqrt(math.Pow(float64(robot.x)-centroidX, 2) + math.Pow(float64(robot.y)-centroidY, 2))
		}
		meanDistance /= float64(len(robots))
		if meanDistance < 25 {
			DrawBots(robots, maxWidth, maxHeight)
			fmt.Printf("Mean distance: %.2f\n", meanDistance)
			fmt.Printf("Steps: %d\n", steps)
			return int64(steps)
		}
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
