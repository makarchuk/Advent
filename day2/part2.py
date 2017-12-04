from itertools import permutations

data = []
with open('data') as f:
    for row in f:
        data.append([int(x) for x in row.split()])


def row_value(row):
    for a, b in permutations(row, 2):
        if not a % b:
            return a / b

print(sum([row_value(x) for x in data]))
