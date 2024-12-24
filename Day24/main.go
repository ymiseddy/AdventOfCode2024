package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 24"

const (
	OpAnd = "AND"
	OpOr  = "OR"
	OpXor = "XOR"
)

var operators = map[string]func(int, int) int{
	OpAnd: func(a, b int) int { return a & b },
	OpOr:  func(a, b int) int { return a | b },
	OpXor: func(a, b int) int { return a ^ b },
}

type Operation struct {
	LOperand    string
	ROperand    string
	Operation   string
	Result      string
	ResultValue int
	Completed   bool
}

func ParseInput(lines []string) (map[string]int, []*Operation) {

	operations := []*Operation{}
	initialStates := map[string]int{}
	lastLine := 0
	for n, line := range lines {
		if line == "" {
			lastLine = n
			break
		}
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			fmt.Printf("Invalid input at line number %d: %s\n", n, line)
			panic("Invalid line")
		}
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		initialStates[parts[0]] = val

	}

	parse := regexp.MustCompile(`(\w+) (AND|OR|XOR) (\w+) -> (\w+)`)
	// Targets are only used once - confirmed.
	for n, line := range lines[lastLine+1:] {
		match := parse.FindStringSubmatch(line)
		if match == nil {
			fmt.Printf("Invalid input at line number %d: %s\n", n+lastLine+1, line)
			panic("Invalid line")
		}
		left := match[1]
		right := match[3]
		operation := match[2]
		result := match[4]
		operations = append(operations, &Operation{
			LOperand:    left,
			ROperand:    right,
			Operation:   operation,
			Result:      result,
			ResultValue: 0,
			Completed:   false,
		})
	}
	return initialStates, operations
}

func CanEvaluate(states map[string]int, operation *Operation) bool {
	_, leftOk := states[operation.LOperand]
	_, rightOk := states[operation.ROperand]
	return leftOk && rightOk
}

func Evaluate(states map[string]int, operation *Operation) {
	fn := operators[operation.Operation]
	left := states[operation.LOperand]
	right := states[operation.ROperand]
	states[operation.Result] = fn(left, right)
	operation.ResultValue = states[operation.Result]
	operation.Completed = true

}

func Puzzle1(lines []string) int {
	total := 0
	states, operations := ParseInput(lines)
	zTstates := evaluateAll(operations, states)
	for n, zT := range zTstates {
		x := math.Pow(2, float64(n)) * float64(states[zT])
		total += int(x)
	}

	return total
}

func evaluateAll(operations []*Operation, states map[string]int) []string {
	zStates := []string{}
	for true {
		anyOperations := false
		for _, operation := range operations {
			if operation.Completed {
				continue
			}
			if !CanEvaluate(states, operation) {
				continue
			}
			Evaluate(states, operation)
			if !operation.Completed {
				panic("Operation not completed")
			}
			if operation.Result[0] == 'z' {
				zStates = append(zStates, operation.Result)
			}
			anyOperations = true
		}

		if !anyOperations {
			break
		}
	}
	slices.Sort(zStates)
	return zStates
}

func (o *Operation) String() string {
	return fmt.Sprintf("%s %s %s-> %s", o.LOperand, o.Operation, o.ROperand, o.Result)
}

func TracePath(state string, outputMap map[string]*Operation, indent string) {
	op, ok := outputMap[state]
	if !ok {
		return
	}
	fmt.Printf("%s%v\n", indent, op)
	indent += "  "
	TracePath(op.LOperand, outputMap, indent)
	TracePath(op.ROperand, outputMap, indent)

	op, ok = outputMap[op.Result]
}

func DoTrace(state string, outputMap map[string]*Operation) {
	fmt.Printf("Tracing %s\n", state)
	TracePath(state, outputMap, "")
	fmt.Println()
}

// right answer:  ckb,kbs,ksv,nbd,tqq,z06,z20,z39
func backTrack(operations []*Operation) []string {

	errorSet := map[string]struct{}{}

	checkResult := map[string]struct{}{}

	adds := map[string]int{}
	xycarrys := map[string]int{}

	for _, op := range operations {
		if (op.LOperand[0] == 'x' || op.LOperand[0] == 'y') &&
			(op.ROperand[0] == 'x' || op.ROperand[0] == 'y') {
			idx, err := strconv.Atoi(op.LOperand[1:])
			idx2, _ := strconv.Atoi(op.ROperand[1:])
			if err != nil {
				panic(err)
			}

			if idx != idx2 {
				panic("Invalid result")
			}

			if op.Operation == OpXor {
				adds[op.Result] = idx
			} else if op.Operation == OpAnd {
				xycarrys[op.Result] = idx
			}
		}
	}

	for _, op := range operations {
		if op.Result[0] == 'z' {
			if op.Operation != OpXor && op.Result != "z45" {
				errorSet[op.Result] = struct{}{}
			}
		}
		if (op.LOperand[0] == 'x' || op.LOperand[0] == 'y') &&
			(op.ROperand[0] == 'x' || op.ROperand[0] == 'y') {
			if op.Result[0] == 'z' {
				if op.Result == "z00" {
					continue
				}
				errorSet[op.Result] = struct{}{}
			}
			checkResult[op.Result] = struct{}{}
		} else if op.Operation == OpXor {
			if op.Result[0] == 'z' {
				continue
			}
			errorSet[op.Result] = struct{}{}
		} else if op.Operation == OpOr {
			_, lok := xycarrys[op.LOperand]
			_, rok := xycarrys[op.ROperand]
			if !lok && !rok {
				errorSet[op.ROperand] = struct{}{}
			}
		}

	}
	for _, op := range operations {
		if op.Operation == OpXor && op.Result[0] != 'z' {
			if (op.LOperand[0] == 'x' || op.LOperand[0] == 'y') &&
				(op.ROperand[0] == 'x' || op.ROperand[0] == 'y') {
				continue
			}
			left := op.LOperand
			right := op.ROperand
			_, okL := checkResult[left]
			_, okR := checkResult[right]
			if !okL || !okR {
				errorSet[op.Result] = struct{}{}
			}

		}
	}
	result := []string{}
	for k, _ := range errorSet {
		result = append(result, k)
	}
	slices.Sort(result)
	return result
}

func Puzzle2(lines []string) string {
	_, operations := ParseInput(lines)

	outputMap := map[string]*Operation{}
	for _, op := range operations {
		outputMap[op.Result] = op
	}
	resultSet := backTrack(operations)
	name := strings.Join(resultSet, ",")
	return name
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
