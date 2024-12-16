package main

import (
	"fmt"
	"os"
	//"slices"

	"github.com/ymiseddy/AdventOfCode2024/priorityqueue"
	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 16"

var debug = true

func ReadGrid(lines []string) [][]rune {
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

var cardinalDirections = [][]int{
	// North
	{0, -1},
	// East
	{1, 0},
	// South
	{0, 1},
	// West
	{-1, 0},
}
var directionIntToNameMap = map[int]string{
	0: "North",
	1: "East",
	2: "South",
	3: "West",
}

var directionToCharMap = map[int]rune{
	0: '^',
	1: '>',
	2: '<',
	3: 'v',
}

func Find(grid [][]rune, target rune) (int, int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == target {
				return x, y
			}
		}
	}
	panic("Target not found")
}

type Coord struct {
	X int
	Y int
}

type CoordWithDir struct {
	X         int
	Y         int
	Direction int
}

func (c CoordWithDir) Coord() Coord {
	return Coord{X: c.X, Y: c.Y}
}

func ClockWise(direction int) int {
	return (direction + 1) % 4
}

func CounterClockWise(direction int) int {
	return (direction + 3) % 4
}

type CoordCostMap struct {
	costMapWithDir map[CoordWithDir]int
	costMap        map[Coord]int
}

func NewCoordCostMap() *CoordCostMap {
	return &CoordCostMap{
		costMapWithDir: make(map[CoordWithDir]int),
		costMap:        make(map[Coord]int),
	}
}

func (ccm *CoordCostMap) GetCost(coord Coord) int {
	if cost, ok := ccm.costMap[coord]; ok {
		return cost
	}
	return 9223372036854775807
}

func (ccm *CoordCostMap) GetDirCost(coord CoordWithDir) int {
	if cost, ok := ccm.costMapWithDir[coord]; ok {
		return cost
	}
	return 9223372036854775807
}

func (ccm *CoordCostMap) SetDirCost(coord CoordWithDir, cost int) {
	ccm.costMapWithDir[coord] = cost
	if cost < ccm.GetCost(coord.Coord()) {
		ccm.costMap[coord.Coord()] = cost
	}
}

var coordCostMap = NewCoordCostMap()

func FindPath(grid [][]rune, startX, startY, direction int, cost int) int {
	pq := priorityqueue.New[CoordWithDir]()
	cost = 0

	coord := CoordWithDir{startX, startY, direction}
	coordCostMap.SetDirCost(coord, 0)
	pq.PushItem(coord, 0)

	minCost := 9223372036854775807
	for pq.Len() > 0 {
		coord := pq.PopItem()
		currentCost := coordCostMap.GetDirCost(coord)

		// If we can move forward, push that.
		nx, ny := coord.X+cardinalDirections[coord.Direction][0], coord.Y+cardinalDirections[coord.Direction][1]
		if grid[ny][nx] == 'E' {
			newCost := currentCost + 1
			newCoord := CoordWithDir{
				X:         nx,
				Y:         ny,
				Direction: coord.Direction,
			}
			if newCost <= coordCostMap.GetDirCost(newCoord) {
				if newCost < minCost {
					minCost = newCost
				}
				coordCostMap.SetDirCost(newCoord, newCost)
				pq.PushItem(newCoord, newCost)
			}
		}
		if grid[ny][nx] == '.' {
			newCost := currentCost + 1
			newCoord := CoordWithDir{
				X:         nx,
				Y:         ny,
				Direction: coord.Direction,
			}
			if newCost <= coordCostMap.GetDirCost(newCoord) {
				coordCostMap.SetDirCost(newCoord, newCost)
				pq.PushItem(newCoord, newCost)
			}
		}
		newCost := currentCost + 1000
		counterClockWiseDir := CounterClockWise(coord.Direction)
		newCoord := CoordWithDir{
			X:         coord.X,
			Y:         coord.Y,
			Direction: counterClockWiseDir,
		}
		if newCost < coordCostMap.GetDirCost(newCoord) {
			coordCostMap.SetDirCost(newCoord, newCost)
			pq.PushItem(newCoord, coordCostMap.GetDirCost(newCoord))
		}

		clockWiseDir := ClockWise(coord.Direction)
		newCoord = CoordWithDir{
			X:         coord.X,
			Y:         coord.Y,
			Direction: clockWiseDir,
		}
		if newCost <= coordCostMap.GetDirCost(newCoord) {
			coordCostMap.SetDirCost(newCoord, newCost)
			pq.PushItem(newCoord, coordCostMap.GetDirCost(newCoord))
		}
	}
	return minCost
}

func Puzzle1(lines []string) int {
	var total int = 0
	grid := ReadGrid(lines)
	shared.MaybeShowGrid(grid, debug)
	startX, startY := Find(grid, 'S')
	total = FindPath(grid, startX, startY, 1, 0)

	return total
}

func WalkPath(grid [][]rune, endX, endY int, direction int) int {

	walked := 1
	minCost := -1
	grid[endY][endX] = 'O'
	// Find the minimum cost out
	for _, dir := range cardinalDirections {
		nx, ny := endX+dir[0], endY+dir[1]

		if grid[ny][nx] == 'S' {
			grid[ny][nx] = 'O'
			fmt.Println("Start: ", walked)
			break
		}

		if grid[ny][nx] == '.' || grid[ny][nx] == 'S' {
			if minCost == -1 {
				minCost = coordCostMap.GetCost(Coord{nx, ny})
			} else if minCost > coordCostMap.GetCost(Coord{nx, ny}) {
				minCost = coordCostMap.GetCost(Coord{nx, ny})
			}
		}
	}

	for d, dir := range cardinalDirections {
		nx, ny := endX+dir[0], endY+dir[1]

		if coordCostMap.GetCost(Coord{nx, ny}) == minCost {
			walked += WalkPath(grid, nx, ny, d)
		}

		if d == direction && (coordCostMap.GetCost(Coord{nx, ny})-1000) == minCost {
			walked += WalkPath(grid, nx, ny, d)
		}
	}
	return walked
}

func Puzzle2(lines []string) int {
	var total int = 0
	grid := ReadGrid(lines)
	ex, ey := Find(grid, 'E')
	fmt.Printf("Found exit at %d, %d\n", ex, ey)
	//coord := Coord{ex, ey}

	dir := -1
	minCost := 9223372036854775807
	for x := 0; x < 4; x++ {
		coordX := CoordWithDir{ex, ey, x}
		cost := coordCostMap.GetDirCost(coordX)
		if cost < minCost {
			minCost = cost
			dir = x
			fmt.Printf("Direction %d has cost %d\n", x, cost)
		}
	}

	dir = (dir + 2) % 4

	_ = WalkPath(grid, ex, ey, dir)
	shared.MaybeShowGrid(grid, debug)

	total = 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 'O' {
				total += 1
			}
		}
	}
	/*
		for v, k := range coordCostMap.costMap {
			fmt.Printf("Cost %d for %v\n", k, v)
		}
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
