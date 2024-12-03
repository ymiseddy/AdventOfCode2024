package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "strings"
)

var title string = "Advent of Code 2024, Day 3"

func Puzzle1(lines []string) int64 {
	total := int64(0)
	for _, line := range lines {
		total += evaluateMul(line)
	}
	return total
}

func evaluateMul(line string) int64 {
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := mulRe.FindAllStringSubmatch(line, -1)
	total := int64(0)
	for _, match := range matches {
		x, _ := strconv.ParseInt(match[1], 10, 64)
		y, _ := strconv.ParseInt(match[2], 10, 64)
		total += x * y
	}
	return total
}

func Puzzle2(lines []string) int64 {
	total := int64(0)
	doEnabled := true
	for _, line := range lines {
		for {
			idx := 0
			if doEnabled {
				idx = strings.Index(line, "don't()")
				if idx == -1 {
					total += evaluateMul(line)

					fmt.Println("Line ended with do()")
					idx = len(line)
					break
				}
				doPart := line[:idx]
				line = line[idx:]
				fmt.Println("***** Do:", doPart)
				total += evaluateMul(doPart)
				doEnabled = false
			}

			if !doEnabled {
				idx = strings.Index(line, "do()")
				fmt.Println(line, idx)
				if idx == -1 {
					doEnabled = false
					fmt.Println("Line ended with don't()")
					break
				}
				origLine := line
				dontPart := line[:idx]
				fmt.Println("***** Don't:", dontPart)
				line = line[idx:]
				if dontPart+line != origLine {
					panic("You done goofed")
				}
				doEnabled = true
			}
		}
	}

	return total
}

func main() {
	fmt.Println(title)
	// Read all text from stdin
	lines := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, text)
	}

	res1 := Puzzle1(lines)
	fmt.Println("Puzzle 1 result: ", res1)
	res2 := Puzzle2(lines)
	fmt.Println("Puzzle 2 result: ", res2)
}
