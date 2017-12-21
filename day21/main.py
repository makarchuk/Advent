def parse(row):
    return [list(x) for x in row.split('/')]


def load_mapping(inp):
    mapping = []
    for line in inp:
        pattern_row, res_row = line.strip().split(' => ')
        pattern = parse(pattern_row)
        res = parse(res_row)
        mapping.append((pattern, res))
    return mapping


def rotate(pattern):
    rotated = []
    for i in range(len(pattern) - 1, -1, -1):
        rotated.append([pattern[x][i] for x in range(len(pattern))])
    return rotated


def reverse(pattern):
    reversed = []
    for i in range(len(pattern) - 1, -1, -1):
        reversed.append([pattern[i][x] for x in range(len(pattern))])
    return reversed


def ranges(divisors):
    prev = None
    for x in divisors:
        if prev is not None:
            yield prev, x
        prev = x


def transform(image, patterns):
    result = []
    if not len(image) % 2:
        size = 2
    else:
        size = 3
    patterns = [p for p in patterns if len(p[0]) == size]
    divisors = list(range(0, len(image), size))
    divisors.append(len(image))
    for i, (f, t) in enumerate(ranges(divisors)):
        transformed_row = []
        for f2, t2 in ranges(divisors):
            piece = [[x for x in row[f2:t2]] for row in image[f:t]]
            transformed_piece = _transform(piece, patterns)
            if not transformed_row:
                for x in range(len((transformed_piece))):
                    transformed_row.append([])
            for x in range(len(transformed_piece)):
                transformed_row[x].extend(transformed_piece[x])
        for x in transformed_row:
            result.append(x)
    return result


def _transform(piece, patterns):
    for pattern, transformed in patterns:
        for x in range(2):
            if x:
                p = reverse(pattern)
            else:
                p = pattern
            for x in range(4):
                p = rotate(p)
                if p == piece:
                    return transformed


def pprint(pattern, ch='>'):
    print(ch * 20)
    for x in pattern:
        print("".join(x))
    print(ch * 20)


def part1_test():
    test_input = [
        ".#.",
        "..#",
        "###"
    ]
    test_patterns = [
        '../.# => ##./#../...',
        '.#./..#/### => #..#/..../..../#..#'
    ]
    transform(test_input, load_mapping(test_patterns))


def part1():
    inp = [
        ".#.",
        "..#",
        "###",
    ]
    pprint(inp, '>')

    patterns = load_mapping(open('input'))
    # patterns = load_mapping(['##./#.#/#.. => .##./.#../.#.#/..#.'])
    for x in range(5):
        inp = transform(inp, patterns)
        pprint(inp, '<')
    print("Sum is {}".format(sum(sum(x == '#' for x in row) for row in inp)))


def part2():
    inp = [
        ".#.",
        "..#",
        "###",
    ]
    pprint(inp, '>')

    patterns = load_mapping(open('input'))
    # patterns = load_mapping(['##./#.#/#.. => .##./.#../.#.#/..#.'])
    for x in range(18):
        inp = transform(inp, patterns)
        pprint(inp, '<')
    print("Sum is {}".format(sum(sum(x == '#' for x in row) for row in inp)))


if __name__ == '__main__':
    part1_test()
    part1()
    part2()
