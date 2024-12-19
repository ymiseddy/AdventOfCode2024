package main

import (
	"fmt"
	"github.com/ymiseddy/AdventOfCode2024/priorityqueue"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 18"

type coord struct {
	x int
	y int
}

func ParseInput(lines []string) ([]coord, coord, int) {
	parts := strings.Split(lines[0], ",")
	steps, _ := strconv.Atoi(parts[0])
	width, _ := strconv.Atoi(parts[1])
	height, _ := strconv.Atoi(parts[2])
	mapSize := coord{width, height}
	coords := make([]coord, 0, 1024)
	for _, line := range lines[1:] {
		st := strings.Split(line, ",")
		x, _ := strconv.Atoi(st[0])
		y, _ := strconv.Atoi(st[1])
		coords = append(coords, coord{x, y})
	}
	return coords, mapSize, steps
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

func adjacent(c coord, width int, height int, rocks []coord) []coord {
	result := make([]coord, 0, 4)
outer:
	for _, dir := range cardinalDirections {
		x := c.x + dir[0]
		y := c.y + dir[1]
		if x < 0 || x >= width || y < 0 || y >= height {
			continue
		}
		// Does rocks contain this coord?
		for _, rock := range rocks {
			if rock.x == x && rock.y == y {
				continue outer
			}
		}
		result = append(result, coord{x, y})
	}
	return result
}

type coordStep struct {
	coord coord
	step  int
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

var debug = false

func ShowGridStep(grid [][]rune, pos coord) {
	if !debug {
		return
	}
	Clear()
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if pos.x == x && pos.y == y {
				fmt.Printf("@")
				continue
			}
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
	fmt.Println(pos)
	time.Sleep(time.Millisecond * 250)
}

var mapWidth = 70
var mapHeight = 70
var simSteps = 1024

func Puzzle1(lines []string) int {
	data, mapSize, steps := ParseInput(lines)
	mapWidth = mapSize.x
	mapHeight = mapSize.y
	simSteps = steps

	total := 0
	width := mapWidth + 1
	height := mapHeight + 1
	grid := make([][]rune, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			grid[i][j] = '.'
		}
	}

	for _, rock := range data[:simSteps] {
		grid[rock.y][rock.x] = '#'
	}

	path := AStarPath(grid)
	for _, step := range path {
		grid[step.y][step.x] = 'O'
	}
	total = len(path)
	return total
}

func adjacentPath(position coord, grid [][]rune) []coord {
	result := make([]coord, 0, 8)
	for _, dir := range cardinalDirections {
		x := position.x + dir[0]
		y := position.y + dir[1]
		if x < 0 || x > len(grid[0])-1 || y < 0 || y > len(grid)-1 {
			continue
		}
		if grid[y][x] == '.' {
			result = append(result, coord{x, y})
		}
	}
	return result
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func cost(position coord, step int, exit coord) int {
	// Manhattan distance
	distance := absInt(position.x-exit.x) + absInt(position.y-exit.y)
	return step + distance
}

func AStarPath(grid [][]rune) []coord {
	position := coord{0, 0}
	exit := coord{len(grid[0]) - 1, len(grid) - 1}
	queue := priorityqueue.New[coord]()
	stepMap := make(map[coord]int)
	stepMap[position] = 0
	queue.PushItem(position, cost(position, 0, exit))

	visited := make(map[coord]struct{})
	parent := make(map[coord]coord)
	parent[position] = coord{-1, -1}

	for queue.Len() > 0 {
		position := queue.PopItem()
		step := stepMap[position]
		ShowGridStep(grid, position)
		adjacentPositions := adjacentPath(position, grid)
		for _, adjacentPosition := range adjacentPositions {
			if adjacentPosition.x == len(grid[0])-1 && adjacentPosition.y == len(grid)-1 {
				// Reconstruct path.
				path := make([]coord, 0, 1024)
				path = append(path, adjacentPosition)
				for p := position; p.x != 0 || p.y != 0; p = parent[p] {
					path = append(path, p)
				}
				return path
			}
			if _, ok := visited[adjacentPosition]; ok {
				if queue.Contains(adjacentPosition) {
					cost := cost(adjacentPosition, step+1, exit)
					queue.UpdateGreaterItemPriority(adjacentPosition, cost)
					stepMap[adjacentPosition] = step + 1
				}
				continue
			}
			parent[adjacentPosition] = position
			adjacentCost := cost(adjacentPosition, step+1, exit)
			queue.PushItem(adjacentPosition, adjacentCost)
			visited[adjacentPosition] = struct{}{}
		}
	}

	return []coord{}
}

func BFSPath(grid [][]rune) []coord {
	position := coord{0, 0}
	queue := make([]coord, 0, 1024)
	queue = append(queue, position)
	visited := make(map[coord]struct{})
	parent := make(map[coord]coord)
	parent[position] = coord{-1, -1}

	for len(queue) > 0 {
		position = queue[0]
		//fmt.Printf("Position: %v\n", position)
		ShowGridStep(grid, position)
		queue = queue[1:]
		adjacentPositions := adjacentPath(position, grid)
		for _, adjacentPosition := range adjacentPositions {
			if adjacentPosition.x == len(grid[0])-1 && adjacentPosition.y == len(grid)-1 {
				// Reconstruct path.
				path := make([]coord, 0, 1024)
				path = append(path, adjacentPosition)
				for p := position; p.x != 0 || p.y != 0; p = parent[p] {
					path = append(path, p)
				}
				return path
			}
			if _, ok := visited[adjacentPosition]; ok {
				continue
			}
			parent[adjacentPosition] = position
			queue = append(queue, adjacentPosition)
			visited[adjacentPosition] = struct{}{}
		}
	}

	return []coord{}
}

func Puzzle2(lines []string) coord {
	data, mapSize, steps := ParseInput(lines)
	mapWidth = mapSize.x
	mapHeight = mapSize.y
	simSteps = steps

	width := mapWidth + 1
	height := mapHeight + 1
	grid := make([][]rune, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			grid[i][j] = '.'
		}
	}

	for _, rock := range data[:simSteps] {
		grid[rock.y][rock.x] = '#'
	}
	path := AStarPath(grid)
	for _, rock := range data[simSteps:] {

		grid[rock.y][rock.x] = '#'
		recomputePath := false

		// Optimization: only recompute path if rock is in clearPath
		for _, p := range path {
			if p.x == rock.x && p.y == rock.y {
				recomputePath = true
				break
			}
		}

		if !recomputePath {
			continue
		}
		path = AStarPath(grid)
		if len(path) == 0 {
			return rock
		}
	}

	return path[0]
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
