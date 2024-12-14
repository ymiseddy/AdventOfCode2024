import sympy as sp
import sys
import re
import math


class PuzzleInput:
    def __init__(self, button_a, button_b, prize):
        self.button_a = button_a
        self.button_b = button_b
        self.prize = prize


# Read and parse the puzzle input
def read_puzzle_input(lines):
    button_a_re = re.compile(r"^Button A: X\+(\d+), Y\+(\d+)$")
    button_b_re = re.compile(r"^Button B: X\+(\d+), Y\+(\d+)$")
    prize_re = re.compile(r"^Prize: X=(\d+), Y=(\d+)$")
    puzzle_inputs = []
    puzzle_input = PuzzleInput([], [], [])
    for n, line in enumerate(lines):
        match = button_a_re.match(line)
        if match:
            strx = match.group(1)
            stry = match.group(2)
            x = int(strx)
            y = int(stry)
            puzzle_input.button_a = [x, y]
            continue
        match = button_b_re.match(line)
        if match:
            strx = match.group(1)
            stry = match.group(2)
            x = int(strx)
            y = int(stry)
            puzzle_input.button_b = [x, y]
            continue
        match = prize_re.match(line)
        if match:
            strx = match.group(1)
            stry = match.group(2)
            x = int(strx)
            y = int(stry)
            puzzle_input.prize = [x, y]
            if n != len(lines) - 1:
                continue
        if line == "" or n == len(lines) - 1:
            puzzle_inputs.append(puzzle_input)
            puzzle_input = PuzzleInput([], [], [])
            continue
        raise Exception(f"Invalid line: {line}")
    return puzzle_inputs


def find_min_prize(puzzle_input):
    a_cost = 3
    b_cost = 1

    a, b = sp.symbols("a b")
    eq1 = sp.Eq(
        a*puzzle_input.button_a[0] + b*puzzle_input.button_b[0],
        puzzle_input.prize[0])
    eq2 = sp.Eq(
        a*puzzle_input.button_a[1] + b*puzzle_input.button_b[1],
        puzzle_input.prize[1])
    eqs = [eq1, eq2]
    sol = sp.solve(eqs, [a, b])
    print(sol)
    # Get numeric solutions
    print(sp.N(sol[a]), sp.N(sol[b]))
    ra = sp.N(sol[a])
    rb = sp.N(sol[b])
    if math.floor(ra) == ra and math.floor(rb) == rb:
        cost = a_cost*ra + b_cost*rb
        print(f"cost: {cost}")
        return cost
    return 0


def puzzle1(lines):
    puzzle_inputs = read_puzzle_input(lines)
    total = 0
    for puzzle_input in puzzle_inputs:
        cost = find_min_prize(puzzle_input)
        total += cost
    return total


def puzzle2(lines):
    puzzle_inputs = read_puzzle_input(lines)
    total = 0
    for puzzle_input in puzzle_inputs:
        puzzle_input.prize[0] = 10000000000000 + puzzle_input.prize[0]
        puzzle_input.prize[1] = 10000000000000 + puzzle_input.prize[1]
        cost = find_min_prize(puzzle_input)
        total += cost
    return total


if __name__ == "__main__":
    # Read lines from stdin
    lines = sys.stdin.readlines()

    # Trim the trailing newline
    lines = [line.strip() for line in lines]

    res1 = puzzle1(lines)
    print("Puzzle 1 result: ", res1)
    res2 = puzzle2(lines)
    print("Puzzle 2 result: ", res2)
