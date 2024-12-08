package shared

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
