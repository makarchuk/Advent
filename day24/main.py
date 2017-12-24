from copy import deepcopy


def find_bridges(end, ports, bridge):
    bridge = deepcopy(bridge)
    found = False
    for i, port in enumerate(ports):
        if port[0] == end:
            new_bridge = deepcopy(bridge)
            new_bridge.append(port)
            available_ports = ports[0:i] + ports[i + 1:]
            found = True
            for x in find_bridges(port[1], available_ports, new_bridge):
                yield x
        if port[1] == end:
            found = True
            new_bridge = deepcopy(bridge)
            new_bridge.append(port)
            available_ports = ports[0:i] + ports[i + 1:]
            for x in find_bridges(port[0], available_ports, new_bridge):
                yield x
    if not found:
        yield bridge


def part1(inp):
    bridges = list(find_bridges(0, inp, []))
    strongest = max(bridges, key=strength)
    print("STRONGEST: {}".format(strongest))


def part2(inp):
    bridges = list(find_bridges(0, inp, []))
    longest = len(max(bridges, key=len))
    all_longest = [b for b in bridges if len(b) == longest]
    strongest = max(all_longest, key=strength)
    print("STRONGEST: {}".format(strongest))


def strength(bridge):
    return sum((sum(x) for x in bridge))


def test_input():
    return [
        (0, 2),
        (2, 2),
        (2, 3),
        (3, 4),
        (3, 5),
        (0, 1),
        (10, 1),
        (9, 10)
    ]


def real_input():
    return [
        (31, 13),
        (34, 4),
        (49, 49),
        (23, 37),
        (47, 45),
        (32, 4),
        (12, 35),
        (37, 30),
        (41, 48),
        (0, 47),
        (32, 30),
        (12, 5),
        (37, 31),
        (7, 41),
        (10, 28),
        (35, 4),
        (28, 35),
        (20, 29),
        (32, 20),
        (31, 43),
        (48, 14),
        (10, 11),
        (27, 6),
        (9, 24),
        (8, 28),
        (45, 48),
        (8, 1),
        (16, 19),
        (45, 45),
        (0, 4),
        (29, 33),
        (2, 5),
        (33, 9),
        (11, 7),
        (32, 10),
        (44, 1),
        (40, 32),
        (2, 45),
        (16, 16),
        (1, 18),
        (38, 36),
        (34, 24),
        (39, 44),
        (32, 37),
        (26, 46),
        (25, 33),
        (9, 10),
        (0, 29),
        (38, 8),
        (33, 33),
        (49, 19),
        (18, 20),
        (49, 39),
        (18, 39),
        (26, 13),
        (19, 32)
    ]


if __name__ == '__main__':
    # part1(test_input())
    # part1(real_input())
    part2(real_input())
