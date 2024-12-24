package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/set"
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

func FindCliqueSizedN(source string, connectionMap map[string][]string, track []string, n int) [][]string {

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
	for connectionName := range connectionsMap {
		threeConnections := FindCliqueSizedN(connectionName, connectionMap, []string{}, 3)
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

func FindLargestClique(connectionMap map[string][]string, unitSet map[string]int) []string {
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
				if !slices.Contains(connectionMap[dest], prevDest) {
					continue outer
				}
			}
			newTrack := slices.Clone(track)
			newTrack = append(newTrack, dest)
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
	largestInterconnection := FindLargestClique(connectionMap, unitSet)
	slices.Sort(largestInterconnection)
	password = strings.Join(largestInterconnection, ",")
	return password
}

func buildConnectionMapSets(connections [][]string) map[string]*set.Set[string] {
	connectionMap := map[string]*set.Set[string]{}
	for _, connection := range connections {
		src, dest := connection[0], connection[1]
		if _, ok := connectionMap[src]; !ok {
			connectionMap[src] = set.NewSet[string]()
		}
		connectionMap[src].Add(dest)
		if _, ok := connectionMap[dest]; !ok {
			connectionMap[dest] = set.NewSet[string]()
		}
		connectionMap[dest].Add(src)
	}
	return connectionMap
}

func SetName(set *set.Set[string]) string {
	slice := set.ToSlice()
	slices.Sort(slice)
	return strings.Join(slice, ",")
}

func FindLargestClique_sets(connectionMap map[string]*set.Set[string], unitSet map[string]int) *set.Set[string] {

	var findConnection func(parents *set.Set[string], toExplore *set.Set[string]) *set.Set[string]
	memoized := map[string]*set.Set[string]{}

	findConnection = func(parents *set.Set[string], toExplore *set.Set[string]) *set.Set[string] {
		name := SetName(parents)
		if parents.Size() > 0 {
			if val, ok := memoized[name]; ok {
				return val
			}
		}
		longestResult := parents
		for v := range toExplore.All() {
			if parents.Contains(v) {
				continue
			}
			neighbors := connectionMap[v]
			if parents.Size() > 0 {
				if parents.Contains(v) {
					continue
				}
				// All parents must be in neighbors.
				if set.Intersection(parents, neighbors).Size() < parents.Size() {
					continue
				}
			}
			newParents := parents.With(v)
			neighbors = set.Difference(neighbors, parents)
			result := findConnection(newParents, neighbors)
			/*
				newName := SetName(newParents)
				memoized[newName] = result
			*/
			if result.Size() > longestResult.Size() {
				longestResult = result
			}
		}
		memoized[name] = longestResult
		return longestResult
	}

	units := set.NewSet[string]()
	for unitName := range unitSet {
		units.Add(unitName)
	}

	return findConnection(set.NewSet[string](), units)
}

func Puzzle2Sets(lines []string) string {
	password := ""
	connections, unitSet := ParseInput(lines)
	connectionMap := buildConnectionMapSets(connections)
	largestInterconnection := FindLargestClique_sets(connectionMap, unitSet)
	intersectionKeys := largestInterconnection.ToSlice()
	slices.Sort(intersectionKeys)
	password = strings.Join(intersectionKeys, ",")
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
	res3 := Puzzle2Sets(lines)
	fmt.Println("Puzzle 2 sets: ", res3)
}
