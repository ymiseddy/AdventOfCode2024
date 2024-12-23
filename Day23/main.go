package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 23"

func ParseInput(lines []string) ([][]string, map[string]int) {
	// Split on dash
	var endpointSet = map[string]int{}
	output := [][]string{}
	ct := 0
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic("Invalid line")
		}
		for _, part := range parts {
			if _, ok := endpointSet[part]; !ok {
				endpointSet[part] = ct
				ct++
			}
		}
		connection := []string{parts[0], parts[1]}
		output = append(output, connection)
	}
	return output, endpointSet
}

func FindNConnection(source string, connectionMap map[string][]string, track []string, n int) [][]string {

	var findSubResult func(track []string, n int) [][]string
	findSubResult = func(track []string, n int) [][]string {
		result := [][]string{}
		current := track[len(track)-1]
	outer:
		for _, dest := range connectionMap[current] {
			if slices.Contains(track, dest) {
				continue
			}
			for _, prevDest := range track {
				if !slices.Contains(connectionMap[dest], prevDest) {
					continue outer
				}
			}
			newTrack := append(track, dest)
			slices.Sort(newTrack)
			if len(newTrack) == n {
				result = append(result, newTrack)
			} else {
				result = append(result, findSubResult(newTrack, n)...)
			}
		}
		return result
	}
	return findSubResult([]string{source}, n)
}

func Puzzle1(lines []string) int {
	total := 0
	connections, connectionsMap := ParseInput(lines)
	connectionMap := buildConnectionMap(connections)

	allConnections := [][]string{}
	connectionSet := map[string]struct{}{}
	for connectionName, _ := range connectionsMap {
		threeConnections := FindNConnection(connectionName, connectionMap, []string{}, 3)
		for _, threeConnection := range threeConnections {
			slices.Sort(threeConnection)
			setName := strings.Join(threeConnection, "-")
			if _, ok := connectionSet[setName]; ok {
				continue
			}
			connectionSet[setName] = struct{}{}
			allConnections = append(allConnections, threeConnection)
		}
	}
outer:
	for _, threeConnection := range allConnections {
		for _, connectionName := range threeConnection {
			if connectionName[0] == 't' {
				total++
				continue outer
			}
		}
	}
	return total
}

func buildConnectionMap(connections [][]string) map[string][]string {
	connectionMap := map[string][]string{}
	for _, connection := range connections {
		src, dest := connection[0], connection[1]
		if _, ok := connectionMap[src]; !ok {
			connectionMap[src] = []string{}
		}
		connectionMap[src] = append(connectionMap[src], dest)
		if _, ok := connectionMap[dest]; !ok {
			connectionMap[dest] = []string{}
		}
		connectionMap[dest] = append(connectionMap[dest], src)
	}
	return connectionMap
}

func FindLargestInterconnection(connectionMap map[string][]string, unitSet map[string]int) []string {
	var findConnection func(track []string, connections int) []string

	memoized := map[string][]string{}
	findConnection = func(track []string, connections int) []string {
		current := track[len(track)-1]
		name := strings.Join(track, ",")
		if val, ok := memoized[name]; ok {
			// fmt.Printf("Found memoized: %s\n", name)
			return val
		}
		longest := track
	outer:
		for _, dest := range connectionMap[current] {
			if slices.Contains(track, dest) {
				continue
			}
			for _, prevDest := range track {
				if prevDest == dest {
					panic("Already in track - how are we getting here?")
				}
			}
			for _, prevDest := range track {
				if !slices.Contains(connectionMap[dest], prevDest) {
					continue outer
				}
			}
			newTrack := slices.Clone(track)
			newTrack = append(newTrack, dest)
			// Check newtrack for duplicates
			slices.Sort(newTrack)
			newTrackName := strings.Join(newTrack, ",")
			subResult := findConnection(newTrack, connections)
			memoized[newTrackName] = subResult
			name := strings.Join(newTrack, ",")
			memoized[name] = subResult
			if len(subResult) > len(longest) {
				longest = subResult
			}
		}
		// Check for duplicates in longest
		return longest
	}
	// Create a map of all nodes and their number of connections
	nodeConnections := map[string]int{}
	for unitName, _ := range unitSet {
		for _, dest := range connectionMap[unitName] {
			if _, ok := nodeConnections[dest]; !ok {
				nodeConnections[dest] = 0
			}
			nodeConnections[dest]++
		}
	}
	maxNodeConnections := 0
	for _, n := range nodeConnections {
		if n > maxNodeConnections {
			maxNodeConnections = n
		}
	}
	longestOverall := []string{}
	for node, n := range nodeConnections {
		if n != maxNodeConnections {
			continue
		}
		xx := findConnection([]string{node}, maxNodeConnections)
		if len(xx) > len(longestOverall) {
			longestOverall = xx
		}
	}

	return longestOverall
}

func Puzzle2(lines []string) string {
	password := ""
	connections, unitSet := ParseInput(lines)
	connectionMap := buildConnectionMap(connections)
	largestInterconnection := FindLargestInterconnection(connectionMap, unitSet)
	slices.Sort(largestInterconnection)
	password = strings.Join(largestInterconnection, ",")
	return password
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
