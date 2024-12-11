package main

import (
	"fmt"
	"github.com/ymiseddy/AdventOfCode2024/shared"
	"os"
)

var title string = "Advent of Code 2024, Day 9"

func expand(line string) ([]int, int) {
	var result = make([]int, 0, 34)
	idCounter := 0
	for i, c := range line {
		digit := int(c - '0')
		if i%2 == 0 {
			for j := 0; j < digit; j++ {
				result = append(result, idCounter)
			}
			idCounter++
		} else {
			for j := 0; j < digit; j++ {
				result = append(result, -1)
			}
		}
	}
	return result, idCounter - 1
}

var display = false

func Display(expansion []int) {
	if !display {
		return
	}
	for _, v := range expansion {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()
}

func Puzzle1(lines []string) int64 {
	total := int64(0)
	expansion, _ := expand(lines[0])
	// Display(expansion)
	idxA, idxB := 0, len(expansion)-1
	for {
		if idxA >= idxB {
			break
		}
		if expansion[idxA] != -1 {
			idxA++
			continue
		}
		if expansion[idxB] == -1 {
			idxB--
			continue
		}
		expansion[idxA] = expansion[idxB]
		expansion[idxB] = -1
		idxA++
		idxB--
	}

	total = computeChecksum(expansion, total)

	return total
}

func computeChecksum(expansion []int, total int64) int64 {
	for n, v := range expansion {
		if v == -1 {
			continue
		}
		total += int64(n) * int64(v)
	}
	return total
}

func MaybeFit(fileId int, expansion []int) {

	// Traverse expansion in reverse order
	end := -1
	for i := len(expansion) - 1; i >= 0; i-- {
		if expansion[i] == fileId {
			end = i
			break
		}
	}
	if end == -1 {
		return
	}

	start := end
	for {
		if expansion[start] != fileId {
			start += 1
			break
		}
		if start == 0 {
			break
		}
		start--
	}
	sz := end - start + 1
	// fmt.Println(fileId, start, end, sz)

	// Find a spot large enought to fit the file
	startFound := false
	destStart := 0
	destEnd := -1

	for {
		if startFound {
			if destStart >= start {
				break
			}
			if expansion[destEnd] == -1 {
				destEnd++
				if destEnd-destStart >= sz {
					break
				}
				if destEnd >= len(expansion) {
					break
				}
				continue
			}
			destEnd--
			// We found a spot
			if destEnd-destStart >= sz {
				break
			}
			// We didn't find a spot
			startFound = false
			destEnd = -1
			destStart += 1
			continue
		} else {
			if destStart >= len(expansion) {
				break
			}
			if expansion[destStart] == -1 {
				startFound = true
				destEnd = destStart
				continue
			}
			destStart++
		}
	}

	if destEnd == -1 {
		return
	}
	if destEnd-destStart < sz {
		return
	}

	if destEnd == -1 {
		return
	}
	for x := 0; x < sz; x++ {
		expansion[destStart+x] = fileId
		expansion[start+x] = -1
	}
	fmt.Printf("Can fit %d at %d-%d\n", fileId, destStart, destEnd)
	Display(expansion)

}

func Puzzle2(lines []string) int64 {
	total := int64(0)
	expansion, maxId := expand(lines[0])
	Display(expansion)
	fmt.Println(maxId)
	for i := maxId; i >= 0; i-- {
		MaybeFit(i, expansion)
	}

	total = computeChecksum(expansion, total)
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
