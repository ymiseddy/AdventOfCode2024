package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ymiseddy/AdventOfCode2024/shared"
)

var title string = "Advent of Code 2024, Day 17"

var debug = true

const (
	RegA = iota + 4
	RegB
	RegC
)

const (
	Adv = iota
	Bxl
	Bst
	Jnz
	Bxc
	Out
	Bdv
	Cdv
)

var opcodeMapToName = map[int]string{
	Adv: "Adv",
	Bxl: "Bxl",
	Bst: "Bst",
	Jnz: "Jnz",
	Bxc: "Bxc",
	Out: "Out",
	Bdv: "Bdv",
	Cdv: "Cdv",
}

type VirtualMachine struct {
	Memory     []int
	Ip         int
	Regs       []int
	Output     []string
	OutputInts []int
}

func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		Memory:     make([]int, 0, 1024),
		Ip:         0,
		Regs:       make([]int, 3),
		Output:     make([]string, 0, 1024),
		OutputInts: make([]int, 0, 1024),
	}
}

func (vm *VirtualMachine) Dissasemble() {
	x := 0
	for x < len(vm.Memory) {
		opcode := vm.Memory[x]
		operand := vm.Memory[x+1]
		if vm.Ip == x {
			fmt.Printf("*")
		} else {
			fmt.Printf(" ")
		}
		operandStr := fmt.Sprintf("%d", operand)
		if usesComboOperandMap[opcode] {
			if operand <= 3 {
				operandStr = fmt.Sprintf("%d (literal)", operand)
			} else {
				operandStr = fmt.Sprintf("Reg%c", 'A'-4+operand)
			}
		}
		fmt.Printf("%-4d %-10s %s\n", x, opcodeMapToName[opcode], operandStr)
		x += 2
	}
}

func (vm *VirtualMachine) String() string {
	return fmt.Sprintf("Memory: %v, Ip: %v, Regs: %v", vm.Memory, vm.Ip, vm.Regs)
}

func (vm *VirtualMachine) HasNextInstruction() bool {
	return vm.Ip < len(vm.Memory)
}

var usesComboOperandMap = map[int]bool{
	Adv: true,
	Bxl: false,
	Bst: true,
	Jnz: false,
	Bxc: false,
	Out: true,
	Bdv: true,
	Cdv: true,
}

func (vm *VirtualMachine) ExecInstruction() {
	opcode := vm.Memory[vm.Ip]
	operand := vm.Memory[vm.Ip+1]
	switch opcode {
	case Adv:
		comboOperand := vm.ComboOperand(operand)
		comboOperand = int(math.Pow(2, float64(comboOperand)))
		div := vm.Regs[RegA-4] / comboOperand
		vm.Regs[RegA-4] = div
	case Bxl:
		vm.Regs[RegB-4] = vm.Regs[RegB-4] ^ operand
	case Bst:
		comboOperand := vm.ComboOperand(operand)
		vm.Regs[RegB-4] = comboOperand % 8
	case Jnz:
		if vm.Regs[RegA-4] != 0 {
			vm.Ip = operand
			return
		}
	case Bxc:
		// Ignores operand
		vm.Regs[RegB-4] = vm.Regs[RegB-4] ^ vm.Regs[RegC-4]
	case Out:
		comboOperand := vm.ComboOperand(operand)
		vm.Output = append(vm.Output, fmt.Sprintf("%d", comboOperand%8))
		vm.OutputInts = append(vm.OutputInts, comboOperand%8)
	case Bdv:
		comboOperand := vm.ComboOperand(operand)
		comboOperand = int(math.Pow(2, float64(comboOperand)))
		vm.Regs[RegB-4] = vm.Regs[RegA-4] / comboOperand
	case Cdv:
		comboOperand := vm.ComboOperand(operand)
		comboOperand = int(math.Pow(2, float64(comboOperand)))
		vm.Regs[RegC-4] = vm.Regs[RegA-4] / comboOperand
	default:
		panic("Invalid opcode")
	}
	vm.Ip += 2
}

func (vm *VirtualMachine) ComboOperand(op int) int {
	if op <= 3 {
		return op
	}
	switch op {
	case RegA:
		return vm.Regs[0]
	case RegB:
		return vm.Regs[1]
	case RegC:
		return vm.Regs[2]
	default:
		panic("Invalid operand")
	}
}

func (vm *VirtualMachine) Run() {
	for vm.HasNextInstruction() {
		vm.ExecInstruction()
	}
}

func ReadInstructions(lines []string) *VirtualMachine {
	registerRe := regexp.MustCompile(`Register (A|B|C): (\d+)`)
	programRe := regexp.MustCompile(`Program: (.+)`)
	vm := NewVirtualMachine()
	for _, line := range lines {
		if line == "" {
			continue
		}
		registerMatch := registerRe.FindStringSubmatch(line)
		if registerMatch != nil {
			switch registerMatch[1] {
			case "A":
				vm.Regs[0], _ = strconv.Atoi(registerMatch[2])
			case "B":
				vm.Regs[1], _ = strconv.Atoi(registerMatch[2])
			case "C":
				vm.Regs[2], _ = strconv.Atoi(registerMatch[2])
			default:
				panic("Invalid register")
			}
			continue
		}
		programMatch := programRe.FindStringSubmatch(line)
		if programMatch != nil {
			// Split the program by commas
			program := strings.Split(programMatch[1], ",")
			for _, p := range program {
				// Convert the program to an integer
				i, _ := strconv.Atoi(p)
				vm.Memory = append(vm.Memory, i)
			}
			continue
		}
	}

	return vm
}

func Puzzle1(lines []string) string {
	vm := ReadInstructions(lines)
	for vm.HasNextInstruction() {
		vm.ExecInstruction()
	}

	output := strings.Join(vm.Output, ",")
	return output
}

func FindSolution(vm *VirtualMachine, a, d int) int {
	if d == len(vm.Memory) {
		return a
	}

	for x := 0; x < 8; x++ {
		nv := a*8 + x
		vm.Ip = 0
		vm.Regs[RegA-4] = nv
		vm.Regs[RegB-4] = 0
		vm.Regs[RegC-4] = 0
		vm.Output = make([]string, 0, 1024)
		vm.OutputInts = make([]int, 0, 1024)
		vm.Run()

		// Compare outputInts with memory at the end.
		len_outputs := len(vm.OutputInts)
		mem_start := len(vm.Memory) - len_outputs
		allMatch := true
		for i := 0; i < len(vm.OutputInts); i++ {
			if vm.OutputInts[i] != vm.Memory[mem_start+i] {
				allMatch = false
			}
		}
		if allMatch {
			res := FindSolution(vm, nv, d+1)
			if res != 0 {
				return res
			}
		}
	}

	return 0
}

func Puzzle2(lines []string) int {
	var total int = 0
	vm := ReadInstructions(lines)
	//fmt.Println("-------------------------------------")
	//for x := 0; x < 4097; x++ {
	// fmt.Printf("%v\n", vm)
	//vm.Dissasemble()
	res := FindSolution(vm, 0, 0)
	total = res
	//}
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
