package shared

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func MaybeShowGrid(grid [][]rune, debug bool) {
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

func ReadLinesFromStream(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func ReadIntsFromFile(fileName string) ([][]int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return ReadIntsFromStream(file)
}

func ConvertStringToInts(source []string) []int64 {
	result := make([]int64, len(source))
	for i, s := range source {
		result[i], _ = strconv.ParseInt(s, 10, 64)
	}
	return result
}

func ReadIntsFromStream(file *os.File) ([][]int64, error) {
	scanner := bufio.NewScanner(file)
	var lines [][]int64
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		ints := make([]int64, len(fields))
		for i, field := range fields {
			val, err := strconv.ParseInt(field, 10, 64)
			if err != nil {
				return nil, err
			}
			ints[i] = val
		}
		lines = append(lines, ints)
	}
	return lines, nil
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

type Coord struct {
	X int
	Y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDistance(c1 Coord, c2 Coord) int {
	return Abs(c1.X-c2.X) + Abs(c1.Y-c2.Y)
}

func EuclideanDistance(c1 Coord, c2 Coord) int {
	return int(math.Sqrt(float64(c1.X-c2.X)*float64(c1.X-c2.X) + float64(c1.Y-c2.Y)*float64(c1.Y-c2.Y)))
}

func (c Coord) Move(dir int) Coord {
	dx, dy := CardinalDirections[dir][0], CardinalDirections[dir][1]
	c.X += dx
	c.Y += dy
	return c
}

func (c Coord) Adjacencies() []Coord {
	var result []Coord
	for _, dir := range []int{0, 1, 2, 3} {
		result = append(result, c.Move(dir))
	}
	return result
}

var CardinalDirections = [][]int{
	// North
	{0, -1},
	// East
	{1, 0},
	// South
	{0, 1},
	// West
	{-1, 0},
}

func Clockwise(dir int) int {
	return (dir + 1) % 4
}

func CounterClockwise(dir int) int {
	return (dir + 3) % 4
}

func ShowGridStep(grid [][]rune, debug bool, positions []Coord, sleepTime int) {
	if !debug {
		return
	}
	positionMap := make(map[Coord]struct{})
	for _, pos := range positions {
		positionMap[pos] = struct{}{}
	}
	Clear()
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if _, ok := positionMap[Coord{x, y}]; ok {
				fmt.Printf("@")
				continue
			}
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
	if sleepTime > 0 {
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}
}
