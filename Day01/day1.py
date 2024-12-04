import numpy as np


def puzzle1(data: list[list[int]]) -> int:
    list_1 = np.array([x[0] for x in data])
    list_1.sort()
    list_2 = np.array([x[1] for x in data])
    list_2.sort()
    diff = np.abs(list_1 - list_2)
    return np.sum(diff)


def puzzle2(data: list[list[int]]) -> int:
    list_1 = np.array([x[0] for x in data])
    list_1.sort()
    # Dictionary of counts
    counts = {}
    for i in [x[1] for x in data]:
        counts[i] = counts.get(i, 0) + 1

    totals = [x * counts.get(x, 0) for x in list_1]
    return np.sum(totals)


if __name__ == "__main__":
    with open("input.txt") as f:
        lines = f.readlines()

    data = []
    for line in lines:
        line = line.strip()
        parts = line.split()
        parts = [int(p) for p in parts]
        data.append(parts)

    result1 = puzzle1(data)
    print(f"Result 1: {result1}")

    result1 = puzzle2(data)
    print(f"Result 2: {result1}")
